/*
Copyright 2019 The Crossplane Authors.
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

package anthos

import (
	"context"
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/controller"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane-contrib/provider-gcp/apis/anthos/v1alpha1"
)

//	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// SetupAttachedCluster adds a controller that reconciles AttachedCluster managed
// resources.
func SetupAttachedCluster(mgr ctrl.Manager, o controller.Options) error {
	name := managed.ControllerName(v1alpha1.AttachedClusterGroupKind)

	cps := []managed.ConnectionPublisher{managed.NewAPISecretPublisher(mgr.GetClient(), mgr.GetScheme())}

	r := managed.NewReconciler(mgr,
		resource.ManagedKind(v1alpha1.AttachedClusterGroupVersionKind),
		managed.WithExternalConnecter(&attachedClusterConnector{kube: mgr.GetClient()}),
		managed.WithPollInterval(o.PollInterval),
		managed.WithLogger(o.Logger.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		managed.WithConnectionPublishers(cps...))

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		For(&v1alpha1.AttachedCluster{}).
		Complete(ratelimiter.NewReconciler(name, r, o.GlobalRateLimiter))
}

type attachedClusterConnector struct {
	kube client.Client
}

func (c *attachedClusterConnector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	return &attachedClusterExternal{}, nil
}

type attachedClusterExternal struct {
}

func (e *attachedClusterExternal) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) { // nolint:gocyclo
	return managed.ExternalObservation{}, nil
}

func (e *attachedClusterExternal) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	return managed.ExternalCreation{}, nil
}

func (e *attachedClusterExternal) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	return managed.ExternalUpdate{}, fmt.Errorf("attachedCluster Update is not supported yet")
}

func (e *attachedClusterExternal) Delete(ctx context.Context, mg resource.Managed) error {
	return nil
}
