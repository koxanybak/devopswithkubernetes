/*


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

package controllers

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	kbatch "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	batchv1 "stable.dwk/api/v1"
)

// DummySiteReconciler reconciles a DummySite object
type DummySiteReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

func constructPodForDummySite(ds *batchv1.DummySite, r *DummySiteReconciler) (*kbatch.Pod, error) {
	name := fmt.Sprintf(ds.Name)
	pod := &kbatch.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Labels:			map[string]string{"app": name},
			Annotations: 	make(map[string]string),
			Name:			name + "-pod",
			Namespace:		ds.Namespace,
		},
		Spec: kbatch.PodSpec{
			Containers: []kbatch.Container{
				kbatch.Container{
					Name: name,
					Image: "koxanybak/dummysite",
					Env: []kbatch.EnvVar{
						kbatch.EnvVar{Name: "WEBSITE_URL", Value: ds.Spec.WebsiteURL},
					},
				},
			},
			RestartPolicy: "OnFailure",
		},
	}

	if err := ctrl.SetControllerReference(ds, pod, r.Scheme); err != nil {
		return nil, err
	}

	return pod, nil
}

func constructSvcForDummySite(ds *batchv1.DummySite, r *DummySiteReconciler) (*kbatch.Service, error) {
	name := fmt.Sprintf(ds.Name)
	svc := &kbatch.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels:			make(map[string]string),
			Annotations: 	make(map[string]string),
			Name:			name + "-svc",
			Namespace:		ds.Namespace,
		},
		Spec: kbatch.ServiceSpec{
			Selector: map[string]string{"app": name},
			Type: kbatch.ServiceTypeClusterIP,
			Ports: []kbatch.ServicePort{
				kbatch.ServicePort{
					Port: 80,
					TargetPort: intstr.FromInt(8000),
					Protocol: kbatch.ProtocolTCP,
				},
			},
		},
	}

	if err := ctrl.SetControllerReference(ds, svc, r.Scheme); err != nil {
		return nil, err
	}

	return svc, nil
}

// +kubebuilder:rbac:groups=batch.stable.dwk,resources=dummysites,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=batch.stable.dwk,resources=dummysites/status,verbs=get;update;patch

func (r *DummySiteReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("dummysite", req.NamespacedName)

	// your logic here
	var ds batchv1.DummySite
	if err := r.Get(ctx, req.NamespacedName, &ds); err != nil {
		log.Error(err, "unable to fetch")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Construct
	pod, err := constructPodForDummySite(&ds, r)
	if err != nil {
		log.Error(err, "unable to construct pod from ds")
		return ctrl.Result{}, err
	}
	/* svc, err := constructSvcForDummySite(&ds, r)
	if err != nil {
		log.Error(err, "unable to construct svc from ds")
		return ctrl.Result{}, err
	} */

	// Apply
	if err := r.Create(ctx, pod); err != nil {
		log.Error(err, "unable to create Pod for Countdown")
	}
	/* if err := r.Create(ctx, svc); err != nil {
		log.Error(err, "unable to create Service for Countdown", "svc")
	} */

	return ctrl.Result{}, nil
}

func (r *DummySiteReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&batchv1.DummySite{}).
		Complete(r)
}
