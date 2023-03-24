package expr

import (
	"context"
	v1 "github.com/acorn-io/acorn/pkg/apis/internal.acorn.io/v1"
	"github.com/acorn-io/acorn/pkg/labels"
	"github.com/acorn-io/acorn/pkg/scheme"
	"github.com/acorn-io/baaah/pkg/router/tester"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	"testing"
)

func TestResolve(t *testing.T) {
	type args struct {
		objects   []kclient.Object
		expr      string
		namespace string
	}

	type want struct {
		kind      string
		name      string
		namespace string
		err       assert.ErrorAssertionFunc
	}

	objects := []kclient.Object{
		&corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: "project-name",
			},
		},
		&corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: "app-namespace",
				Labels: map[string]string{
					labels.AcornAppName:      "app-name",
					labels.AcornAppNamespace: "project-name",
				},
			},
		},
		&corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: "child-app-random-ns",
				Labels: map[string]string{
					labels.AcornAppName:      "child-app-random",
					labels.AcornAppNamespace: "project-name",
				},
			},
		},
		&v1.AppInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "app-name",
				Namespace: "project-name",
			},
			Status: v1.AppInstanceStatus{
				Namespace: "app-namespace",
			},
		},
		&v1.AppInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "child-app-random",
				Namespace: "project-name",
				Labels: map[string]string{
					labels.AcornAcornName:       "child-app",
					labels.AcornParentAcornName: "app-name",
				},
			},
			Status: v1.AppInstanceStatus{
				Namespace: "child-app-random-ns",
			},
		},
		&v1.ServiceInstance{
			TypeMeta: metav1.TypeMeta{},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "child-app-default-svc",
				Namespace: "child-app-random-ns",
			},
			Spec: v1.ServiceInstanceSpec{
				Default: true,
			},
		},
		&v1.ServiceInstance{
			TypeMeta: metav1.TypeMeta{},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "grandchild-app-default-svc",
				Namespace: "grandchild-app-random-ns",
			},
			Spec: v1.ServiceInstanceSpec{
				Default: true,
			},
		},
		&v1.ServiceInstance{
			TypeMeta: metav1.TypeMeta{},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "app-name-default-svc",
				Namespace: "app-namespace",
			},
			Spec: v1.ServiceInstanceSpec{
				Default: true,
				Secrets: []string{"secret"},
			},
		},
		&v1.AppInstance{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "grandchild-app-random",
				Namespace: "project-name",
				Labels: map[string]string{
					labels.AcornAcornName:       "grandchild-app",
					labels.AcornParentAcornName: "child-app-random",
				},
			},
			Status: v1.AppInstanceStatus{
				Namespace: "grandchild-app-random-ns",
			},
		},
		&corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "secret",
				Namespace: "app-namespace",
			},
		},
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "direct lookup of app",
			args: args{
				objects:   objects,
				expr:      "app-name",
				namespace: "project-name",
			},
			want: want{
				kind:      "ServiceInstance",
				name:      "app-name-default-svc",
				namespace: "app-namespace",
			},
		},
		{
			name: "lookup of app defined in acorn",
			args: args{
				objects:   objects,
				expr:      "child-app",
				namespace: "app-namespace",
			},
			want: want{
				kind:      "ServiceInstance",
				name:      "child-app-default-svc",
				namespace: "child-app-random-ns",
			},
		},
		{
			name: "lookup of app nested",
			args: args{
				objects:   objects,
				expr:      "child-app.grandchild-app",
				namespace: "app-namespace",
			},
			want: want{
				kind:      "ServiceInstance",
				name:      "grandchild-app-default-svc",
				namespace: "grandchild-app-random-ns",
			},
		},
		{
			name: "external lookup of service",
			args: args{
				objects:   objects,
				expr:      "app-name.app-name-default-svc",
				namespace: "project-name",
			},
			want: want{
				kind:      "ServiceInstance",
				name:      "app-name-default-svc",
				namespace: "app-namespace",
			},
		},
		{
			name: "external lookup of secret",
			args: args{
				objects:   objects,
				expr:      "app-name.app-name-default-svc.secret",
				namespace: "project-name",
			},
			want: want{
				kind:      "Secret",
				name:      "secret",
				namespace: "app-namespace",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &tester.Client{
				Objects:   tt.args.objects,
				SchemeObj: scheme.Scheme,
			}
			got, err := Resolve(context.Background(), req, tt.args.namespace, tt.args.expr)
			if tt.want.err == nil {
				assert.NoError(t, err)
			} else {
				tt.want.err(t, err)
			}
			if err == nil {
				gvk, err := apiutil.GVKForObject(got, scheme.Scheme)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tt.want.kind, gvk.Kind)
				assert.Equal(t, tt.want.name, got.GetName())
				assert.Equal(t, tt.want.namespace, got.GetNamespace())
			}
		})
	}
}
