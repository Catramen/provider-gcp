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
	"errors"
	"fmt"
	"io"
	"strings"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/controller"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/yaml"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane-contrib/provider-gcp/apis/anthos/v1alpha1"
	"github.com/crossplane-contrib/provider-gcp/pkg/clients/anthos"
)

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
	s, err := anthos.NewService(ctx, c.kube, mg)
	if err != nil {
		return nil, err
	}
	return &attachedClusterExternal{s: s}, nil
}

type attachedClusterExternal struct {
	s *anthos.Service
}

func (e *attachedClusterExternal) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) { // nolint:gocyclo
	cr, ok := mg.(*v1alpha1.AttachedCluster)
	if !ok {
		return managed.ExternalObservation{}, fmt.Errorf("not attached cluster")
	}
	if err := e.s.GetAttachedCluster(ctx, cr); err != nil {
		if errors.Is(err, anthos.ErrNotExist) {
			return managed.ExternalObservation{}, nil
		}
		return managed.ExternalObservation{}, fmt.Errorf("cannot create get cluster %v", err)
	}
	switch cr.Status.AtProvider.State {
	case "PROVISIONING":
		cr.Status.SetConditions(xpv1.Creating())
	case "RUNNING":
		cr.Status.SetConditions(xpv1.Available())
	case "ERROR":
		cr.Status.SetConditions(xpv1.Unavailable())
	}

	return managed.ExternalObservation{
		ResourceExists:   true,
		ResourceUpToDate: true,
	}, nil
}

func getResources(manifest string) ([]*unstructured.Unstructured, error) {
	d := yaml.NewYAMLOrJSONDecoder(strings.NewReader(manifest), 4096)
	var resources []*unstructured.Unstructured
	for {
		u := &unstructured.Unstructured{}
		if err := d.Decode(u); err != nil {
			if err != io.EOF {
				return nil, fmt.Errorf("could not decode object: %s", err)
			}
			break
		}
		if len(u.Object) == 0 {
			// We skip empty resources such as those generated by empty blocks between
			// resource separators.
			continue
		}
		resources = append(resources, u)
	}
	return resources, nil
}

func (e *attachedClusterExternal) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*v1alpha1.AttachedCluster)
	if !ok {
		return managed.ExternalCreation{}, fmt.Errorf("not attached cluster")
	}
	if e.s.Kube == nil {
		fmt.Println("kubeclient not built")
		return managed.ExternalCreation{}, nil
	}
	m, err := e.s.GetInstallManifest(ctx, cr.Spec.ForProvider.Location, anthos.InstallManifest{
		AttachedClusterID: cr.Name,
		PlatformVersion:   cr.Spec.ForProvider.PlatformVersion,
	})
	if err != nil {
		return managed.ExternalCreation{}, fmt.Errorf("cannot get install manifest %v", err)
	}
	resources, err := getResources(m.Manifest)
	if err != nil {
		return managed.ExternalCreation{}, err
	}

	c := resource.ClientApplicator{
		Client:     e.s.Kube,
		Applicator: resource.NewAPIPatchingApplicator(e.s.Kube),
	}
	for _, r := range resources {
		if err := c.Apply(ctx, r); err != nil {
			return managed.ExternalCreation{}, fmt.Errorf("failed to apply manifest: %v", err)
		}
	}

	if err := e.s.CreateAttachedCluster(ctx, cr); err != nil {
		return managed.ExternalCreation{}, fmt.Errorf("cannot create attached cluster %v", err)
	}

	return managed.ExternalCreation{}, nil
}

func (e *attachedClusterExternal) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	return managed.ExternalUpdate{}, fmt.Errorf("attachedCluster Update is not supported yet")
}

func (e *attachedClusterExternal) Delete(ctx context.Context, mg resource.Managed) error {
	cr, ok := mg.(*v1alpha1.AttachedCluster)
	if !ok {
		return fmt.Errorf("not attached cluster")
	}
	cr.SetConditions(xpv1.Deleting())
	return e.s.DeleteAttachedCluster(ctx, cr)
}
