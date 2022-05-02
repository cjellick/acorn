package appdefinition

import (
	"strconv"
	"strings"

	v1 "github.com/acorn-io/acorn/pkg/apis/acorn.io/v1"
	"github.com/acorn-io/acorn/pkg/labels"
	"github.com/acorn-io/acorn/pkg/system"
	"github.com/acorn-io/baaah/pkg/meta"
	"github.com/acorn-io/baaah/pkg/typed"
	name2 "github.com/rancher/wrangler/pkg/name"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func addAlias(aliases []*corev1.Service, aliasServiceName, aliasTarget string, svc *corev1.Service) []*corev1.Service {
	for _, existing := range aliases {
		if existing.Name == aliasServiceName {
			for _, newPort := range svc.Spec.Ports {
				found := false
				for _, existingPort := range existing.Spec.Ports {
					if existingPort.Port == newPort.Port {
						found = true
						break
					}
				}
				if !found {
					existing.Spec.Ports = append(existing.Spec.Ports, newPort)
				}
			}
			return aliases
		}
	}

	newSvc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      aliasServiceName,
			Namespace: svc.Namespace,
			Labels:    svc.Labels,
		},
		Spec: svc.Spec,
	}
	newSvc.Spec.Selector = map[string]string{
		labels.AcornAlias + aliasTarget: "true",
		labels.AcornAppNamespace:        svc.Spec.Selector[labels.AcornAppNamespace],
		labels.AcornAppName:             svc.Spec.Selector[labels.AcornAppName],
		labels.AcornManaged:             "true",
	}

	return append(aliases, newSvc)
}

func toServices(appInstance *v1.AppInstance) (result []meta.Object) {
	var aliases []*corev1.Service

	for _, entry := range typed.Sorted(appInstance.Status.AppSpec.Containers) {
		service := toService(appInstance, entry.Key, entry.Value)
		if service != nil {
			for _, alias := range entry.Value.Aliases {
				aliases = addAlias(aliases, alias.Name, alias.Name, service)
			}
			result = append(result, service)
		}
		publishService := toPublishService(appInstance, entry.Key, entry.Value)
		if publishService != nil {
			if len(entry.Value.Aliases) > 0 {
				alias := entry.Value.Aliases[0]
				aliases = addAlias(aliases, PublishServiceName(appInstance, alias.Name), alias.Name, publishService)
			} else {
				result = append(result, publishService)
			}
		}
	}

	for _, entry := range typed.Sorted(appInstance.Status.AppSpec.Acorns) {
		service := toService(appInstance, entry.Key, v1.Container{Ports: entry.Value.Ports})
		if service != nil {
			service, ptrService := toAcornService(appInstance, service)
			result = append(result, service, ptrService)
		}
		publishService := toPublishService(appInstance, entry.Key, v1.Container{Ports: entry.Value.Ports})
		if publishService != nil {
			publishService, _ = toAcornService(appInstance, publishService)
			result = append(result, publishService)
		}
	}

	for _, alias := range aliases {
		result = append(result, alias)
	}

	return result
}

func toAcornLabels(l map[string]string) map[string]string {
	result := map[string]string{}
	for k, v := range l {
		if k == labels.AcornContainerName {
			k = labels.AcornAcornName
		}
		result[k] = v
	}
	return result
}

func toAcornService(app *v1.AppInstance, svc *corev1.Service) (*corev1.Service, *corev1.Service) {
	systemName := name2.SafeConcatName(svc.Name, svc.Namespace, string(app.UID[:12]))
	ptrSvc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      svc.Name,
			Namespace: svc.Namespace,
			Labels:    toAcornLabels(svc.Labels),
		},
		Spec: corev1.ServiceSpec{
			Type:         corev1.ServiceTypeExternalName,
			ExternalName: systemName + "." + system.Namespace + "." + system.ClusterDomain,
		},
	}
	svc.Name = systemName
	svc.Namespace = system.Namespace
	svc.Labels = toAcornLabels(svc.Labels)
	svc.Spec.Selector = toAcornLabels(svc.Spec.Selector)
	svc.Spec.InternalTrafficPolicy = &[]corev1.ServiceInternalTrafficPolicyType{corev1.ServiceInternalTrafficPolicyLocal}[0]
	return svc, ptrSvc
}

func toServicePort(port v1.Port) corev1.ServicePort {
	servicePort := corev1.ServicePort{
		Name:     strconv.Itoa(int(port.Port)),
		Protocol: corev1.ProtocolTCP,
		Port:     port.Port,
		TargetPort: intstr.IntOrString{
			IntVal: port.ContainerPort,
		},
	}
	switch port.Protocol {
	case v1.ProtocolTCP:
	case v1.ProtocolUDP:
		servicePort.Protocol = corev1.ProtocolUDP
	case v1.ProtocolHTTP:
		fallthrough
	case v1.ProtocolHTTPS:
		str := strings.ToUpper(string(port.Protocol))
		servicePort.AppProtocol = &str
	}
	return servicePort
}

func isLayer4Port(port v1.Port) bool {
	switch port.Protocol {
	case v1.ProtocolUDP:
		return true
	case v1.ProtocolTCP:
		return true
	case v1.ProtocolHTTPS:
		return true
	default:
		return false
	}
}

func PublishServiceName(appInstance *v1.AppInstance, containerName string) string {
	// UID based name is to avoid name conflict. For example if somebody had to containers, foo and foo-publish.
	return name2.SafeConcatName(containerName, "publish", string(appInstance.UID)[:12])
}

func toPublishService(appInstance *v1.AppInstance, containerName string, container v1.Container) *corev1.Service {
	var (
		ports []corev1.ServicePort
	)
	if appInstance.Spec.Stop != nil && *appInstance.Spec.Stop {
		// remove all publishes
		return nil
	}
	for _, port := range container.Ports {
		if port.Publish && isLayer4Port(port) {
			ports = append(ports, toServicePort(port))
		}
	}

	for _, entry := range typed.Sorted(container.Sidecars) {
		for _, port := range entry.Value.Ports {
			if port.Publish && isLayer4Port(port) {
				ports = append(ports, toServicePort(port))
			}
		}
	}

	if len(ports) == 0 {
		return nil
	}

	return &corev1.Service{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      PublishServiceName(appInstance, containerName),
			Namespace: appInstance.Status.Namespace,
			Labels:    containerLabels(appInstance, containerName),
		},
		Spec: corev1.ServiceSpec{
			Ports:    ports,
			Selector: containerLabels(appInstance, containerName),
			Type:     corev1.ServiceTypeLoadBalancer,
		},
	}
}

func toService(appInstance *v1.AppInstance, name string, container v1.Container) *corev1.Service {
	var (
		ports []corev1.ServicePort
	)
	for _, port := range container.Ports {
		ports = append(ports, toServicePort(port))
	}
	for _, entry := range typed.Sorted(container.Sidecars) {
		for _, port := range entry.Value.Ports {
			ports = append(ports, toServicePort(port))
		}
	}

	if len(ports) == 0 {
		return nil
	}

	return &corev1.Service{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: appInstance.Status.Namespace,
			Labels:    containerLabels(appInstance, name),
		},
		Spec: corev1.ServiceSpec{
			Ports:    ports,
			Selector: containerLabels(appInstance, name),
			Type:     corev1.ServiceTypeClusterIP,
		},
	}
}
