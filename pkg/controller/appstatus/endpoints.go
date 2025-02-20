package appstatus

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/acorn-io/baaah/pkg/typed"
	v1 "github.com/acorn-io/runtime/pkg/apis/internal.acorn.io/v1"
	"github.com/acorn-io/runtime/pkg/labels"
	"github.com/acorn-io/runtime/pkg/publish"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	klabels "k8s.io/apimachinery/pkg/labels"
	kclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func serviceEndpoints(ctx context.Context, c kclient.Client, app *v1.AppInstance) (endpoints []v1.Endpoint, _ error) {
	serviceList := &corev1.ServiceList{}
	err := c.List(ctx, serviceList, &kclient.ListOptions{
		Namespace: app.Status.Namespace,
		LabelSelector: klabels.SelectorFromSet(map[string]string{
			labels.AcornManaged:        "true",
			labels.AcornServicePublish: "true",
		}),
	})
	if err != nil {
		return nil, err
	}

	for _, service := range serviceList.Items {
		containerName := service.Labels[labels.AcornContainerName]
		if containerName == "" {
			continue
		}

		for _, port := range service.Spec.Ports {
			var protocol v1.Protocol

			switch port.Protocol {
			case corev1.ProtocolTCP:
				protocol = v1.ProtocolTCP
			case corev1.ProtocolUDP:
				protocol = v1.ProtocolUDP
			default:
				continue
			}

			for _, ingress := range service.Status.LoadBalancer.Ingress {
				if ingress.Hostname != "" {
					endpoints = append(endpoints, v1.Endpoint{
						Target:     containerName,
						TargetPort: port.TargetPort.IntVal,
						Address:    fmt.Sprintf("%s:%d", ingress.Hostname, port.Port),
						Protocol:   protocol,
					})
				} else if ingress.IP != "" {
					endpoints = append(endpoints, v1.Endpoint{
						Target:     containerName,
						TargetPort: port.TargetPort.IntVal,
						Address:    fmt.Sprintf("%s:%d", ingress.IP, port.Port),
						Protocol:   protocol,
					})
				}
			}

			if len(service.Status.LoadBalancer.Ingress) == 0 {
				endpoints = append(endpoints, v1.Endpoint{
					Target:     containerName,
					TargetPort: port.TargetPort.IntVal,
					Address:    fmt.Sprintf("<Pending Ingress>:%d", port.Port),
					Protocol:   protocol,
					Pending:    true,
				})
			}
		}
	}

	return
}

func ingressEndpoints(ctx context.Context, c kclient.Client, app *v1.AppInstance) (endpoints []v1.Endpoint, _ error) {
	ingressList := &networkingv1.IngressList{}
	err := c.List(ctx, ingressList, &kclient.ListOptions{
		Namespace: app.Status.Namespace,
		LabelSelector: klabels.SelectorFromSet(map[string]string{
			labels.AcornManaged: "true",
		}),
	})
	if err != nil {
		return nil, err
	}

	for _, ingress := range ingressList.Items {
		targetStr := ingress.Annotations[labels.AcornTargets]
		if targetStr == "" {
			return nil, err
		}

		targets := map[string]publish.Target{}
		if err := json.Unmarshal([]byte(targetStr), &targets); err != nil {
			return nil, err
		}

		for _, entry := range typed.Sorted(targets) {
			hostname, target := entry.Key, entry.Value
			hostnameOverride := ingress.Annotations[labels.AcornPublishURL]
			if hostnameOverride != "" {
				hostname = hostnameOverride
			}

			endpoints = append(endpoints, v1.Endpoint{
				Target:     target.Service,
				TargetPort: target.Port,
				Address:    hostname,
				Protocol:   v1.ProtocolHTTP,
				Pending:    len(ingress.Status.LoadBalancer.Ingress) == 0,
			})
		}
	}

	return
}

func (a *appStatusRenderer) readEndpoints() error {
	// reset state
	a.app.Status.AppStatus.Endpoints = nil

	ingressEndpoints, err := ingressEndpoints(a.ctx, a.c, a.app)
	if err != nil {
		return err
	}

	serviceEndpoints, err := serviceEndpoints(a.ctx, a.c, a.app)
	if err != nil {
		return err
	}

	eps := append(ingressEndpoints, serviceEndpoints...)

	ingressTLSHosts, err := ingressTLSHosts(a.ctx, a.c, a.app)
	if err != nil {
		return err
	}

	for i, ep := range eps {
		if ep.Protocol == v1.ProtocolHTTP {
			ep.PublishProtocol = v1.PublishProtocolHTTP
			if _, ok := ingressTLSHosts[strings.Split(ep.Address, ":")[0]]; ok {
				ep.PublishProtocol = v1.PublishProtocolHTTPS
			}
		} else {
			ep.PublishProtocol = v1.PublishProtocol(ep.Protocol)
		}
		eps[i] = ep
	}

	sort.Slice(eps, func(i, j int) bool {
		return eps[i].Address < eps[j].Address
	})

	a.app.Status.AppStatus.Endpoints = eps
	return nil
}

func ingressTLSHosts(ctx context.Context, client kclient.Client, app *v1.AppInstance) (map[string]interface{}, error) {
	ingresses := &networkingv1.IngressList{}
	err := client.List(ctx, ingresses, &kclient.ListOptions{
		Namespace: app.Status.Namespace,
		LabelSelector: klabels.SelectorFromSet(map[string]string{
			labels.AcornManaged: "true",
			labels.AcornAppName: app.Name,
		}),
	})
	if err != nil {
		return nil, err
	}

	ingressTLSHosts := map[string]interface{}{}
	for _, ingress := range ingresses.Items {
		if ingress.Spec.TLS != nil {
			for _, tls := range ingress.Spec.TLS {
				for _, host := range tls.Hosts {
					ingressTLSHosts[host] = nil
				}
			}
		}
	}

	return ingressTLSHosts, nil
}
