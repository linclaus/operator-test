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
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	webappv1 "github.com/linclaus/operator-test/api/v1"
)

// GuestbookReconciler reconciles a Guestbook object
type GuestbookReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=webapp.mycrd.com,resources=guestbooks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=webapp.mycrd.com,resources=guestbooks/status,verbs=get;update;patch

func (r *GuestbookReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	_ = r.Log.WithValues("guestbook", req.NamespacedName)

	// your logic here
	gb := &webappv1.Guestbook{}
	if err := r.Get(ctx, req.NamespacedName, gb); err != nil {
		fmt.Errorf("couldn't find object:%s", req.String())
	} else {
		//打印Detail和Created
		r.Log.V(1).Info("Successfully get detail", "Detail", gb.Spec.Detail)
		r.Log.V(1).Info("", "Created", gb.Status.Created)
	}
	fmt.Printf("guestbook:%s\n", gb)
	// 2. Change Created
	if !gb.Status.Created {
		gb.Status.Created = true
		r.Update(ctx, gb)
	}

	return ctrl.Result{}, nil
}

func (r *GuestbookReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&webappv1.Guestbook{}).
		Complete(r)
}
