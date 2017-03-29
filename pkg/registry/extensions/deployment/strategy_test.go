/*
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package deployment

import (
	"reflect"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/testapi"
	apitesting "k8s.io/kubernetes/pkg/api/testing"
	"k8s.io/kubernetes/pkg/apis/extensions"
)

func TestSelectableFieldLabelConversions(t *testing.T) {
	apitesting.TestSelectableFieldLabelConversionsOfKind(t,
		testapi.Extensions.GroupVersion().String(),
		"Deployment",
		DeploymentToSelectableFields(&extensions.Deployment{}),
		nil,
	)
}

func TestStatusUpdates(t *testing.T) {
	tests := []struct {
		old      runtime.Object
		obj      runtime.Object
		expected runtime.Object
	}{
		{
			old:      newDeployment(map[string]string{"test": "label"}, map[string]string{"test": "annotation"}),
			obj:      newDeployment(map[string]string{"test": "label", "sneaky": "label"}, map[string]string{"test": "annotation"}),
			expected: newDeployment(map[string]string{"test": "label"}, map[string]string{"test": "annotation"}),
		},
		{
			old:      newDeployment(map[string]string{"test": "label"}, map[string]string{"test": "annotation"}),
			obj:      newDeployment(map[string]string{"test": "label"}, map[string]string{"test": "annotation", "sneaky": "annotation"}),
			expected: newDeployment(map[string]string{"test": "label"}, map[string]string{"test": "annotation", "sneaky": "annotation"}),
		},
	}

	for _, test := range tests {
		deploymentStatusStrategy{}.PrepareForUpdate(genericapirequest.NewContext(), test.obj, test.old)
		if !reflect.DeepEqual(test.expected, test.obj) {
			t.Errorf("Unexpected object mismatch! Expected:\n%#v\ngot:\n%#v", test.expected, test.obj)
		}
	}
}

func TestInvalidMaxUnavailable(t *testing.T) {
	deployment := newDeploymentWithStategyAndReplica(10, extensions.DeploymentStrategy{
		Type: extensions.RollingUpdateDeploymentStrategyType,
		RollingUpdate: &extensions.RollingUpdateDeployment{
			MaxSurge:       func(i int) intstr.IntOrString { x := intstr.FromInt(i); return x }(3),
			MaxUnavailable: func(i int) intstr.IntOrString { x := intstr.FromInt(i); return x }(20),
		},
	})
	deploymentStrategy{}.PrepareForCreate(genericapirequest.NewContext(), deployment)
	if deployment.Spec.Strategy.RollingUpdate.MaxUnavailable.IntValue() != 10 {
		t.Errorf("Unexpected maxUnavalible! Expected:\n%#v\ngot:\n%#v", 10, deployment.Spec.Strategy.RollingUpdate.MaxUnavailable.IntValue())
	}
	oldDeployment := newDeploymentWithStategyAndReplica(10, extensions.DeploymentStrategy{
		Type: extensions.RollingUpdateDeploymentStrategyType,
		RollingUpdate: &extensions.RollingUpdateDeployment{
			MaxSurge:       func(i int) intstr.IntOrString { x := intstr.FromInt(i); return x }(3),
			MaxUnavailable: func(i int) intstr.IntOrString { x := intstr.FromInt(i); return x }(3),
		},
	})
	deploymentStrategy{}.PrepareForUpdate(genericapirequest.NewContext(), deployment, oldDeployment)
	if deployment.Spec.Strategy.RollingUpdate.MaxUnavailable.IntValue() != 10 {
		t.Errorf("Unexpected maxUnavalible! Expected:\n%#v\ngot:\n%#v", 10, deployment.Spec.Strategy.RollingUpdate.MaxUnavailable.IntValue())
	}
}

func newDeployment(labels, annotations map[string]string) *extensions.Deployment {
	return &extensions.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "test",
			Labels:      labels,
			Annotations: annotations,
		},
		Spec: extensions.DeploymentSpec{
			Replicas: 1,
			Strategy: extensions.DeploymentStrategy{
				Type: extensions.RecreateDeploymentStrategyType,
			},
			Template: api.PodTemplateSpec{
				Spec: api.PodSpec{
					Containers: []api.Container{
						{
							Name:  "test",
							Image: "test",
						},
					},
				},
			},
		},
	}
}

func newDeploymentWithStategyAndReplica(replicas int32, strategy extensions.DeploymentStrategy) *extensions.Deployment {
	return &extensions.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
		Spec: extensions.DeploymentSpec{
			Replicas: replicas,
			Strategy: strategy,
			Template: api.PodTemplateSpec{
				Spec: api.PodSpec{
					Containers: []api.Container{
						{
							Name:  "test",
							Image: "test",
						},
					},
				},
			},
		},
	}
}
