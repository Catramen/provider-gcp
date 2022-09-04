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

package hubmembership

import (
	"context"
	"fmt"

	dcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	gkehub "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/beta"
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/controller"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"google.golang.org/api/option"
	htransport "google.golang.org/api/transport/http"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane-contrib/provider-gcp/apis/anthos/v1alpha1"

	gcp "github.com/crossplane-contrib/provider-gcp/pkg/clients"
)

var (
	// CreateDirective restricts Apply to creating resources for Create
	CreateDirective = []dcl.ApplyOption{
		dcl.WithLifecycleParam(dcl.BlockAcquire),
		dcl.WithLifecycleParam(dcl.BlockDestruction),
		dcl.WithLifecycleParam(dcl.BlockModification),
	}

	// UpdateDirective restricts Apply to modifying resources for Update
	UpdateDirective = []dcl.ApplyOption{
		dcl.WithLifecycleParam(dcl.BlockCreation),
		dcl.WithLifecycleParam(dcl.BlockDestruction),
	}
)

// const gkeHubBase = "https://gkehub.googleapis.com/"
const gkeHubBase = "https://autopush-gkehub.sandbox.googleapis.com/"

// SetupHubMembership adds a controller that reconciles HubMembership managed
// resources.
func SetupHubMembership(mgr ctrl.Manager, o controller.Options) error {
	name := managed.ControllerName(v1alpha1.HubMembershipGroupKind)

	cps := []managed.ConnectionPublisher{managed.NewAPISecretPublisher(mgr.GetClient(), mgr.GetScheme())}

	r := managed.NewReconciler(mgr,
		resource.ManagedKind(v1alpha1.HubMembershipGroupVersionKind),
		managed.WithExternalConnecter(&MembershipConnector{kube: mgr.GetClient()}),
		managed.WithPollInterval(o.PollInterval),
		managed.WithLogger(o.Logger.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		managed.WithConnectionPublishers(cps...))

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		For(&v1alpha1.HubMembership{}).
		Complete(ratelimiter.NewReconciler(name, r, o.GlobalRateLimiter))
}

// MembershipConnector is the connector of hub membership.
type MembershipConnector struct {
	kube client.Client
}

// Connect to clients!
func (c *MembershipConnector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	projectID, opts, err := gcp.GetConnectionInfo(ctx, c.kube, mg)
	if err != nil {
		return nil, err
	}
	opts = append(opts, option.WithAudiences(gkeHubBase))
	httpClient, _, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}

	configOptions := []dcl.ConfigOption{
		dcl.WithHTTPClient(httpClient),
		dcl.WithBasePath(gkeHubBase + "v1/"),
	}

	dclConfig := dcl.NewConfig(configOptions...)
	s := gkehub.NewClient(dclConfig)
	return &MembershipExternal{kube: c.kube, projectID: projectID, gkeHub: s}, nil
}

// MembershipExternal is the external resource manager.
type MembershipExternal struct {
	kube      client.Client
	gkeHub    *gkehub.Client
	projectID string
}

// Observe resources!
func (e *MembershipExternal) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) { // nolint:gocyclo
	cr, ok := mg.(*v1alpha1.HubMembership)
	if !ok {
		return managed.ExternalObservation{}, fmt.Errorf("not membership")
	}
	obj := &gkehub.Membership{
		Project:  dcl.String(e.projectID),
		Name:     dcl.String(cr.Spec.ForProvider.MembershipID),
		Location: dcl.String("global"),
	}
	resp, err := e.gkeHub.GetMembership(ctx, obj)
	if err != nil {
		if dcl.IsNotFound(err) {
			return managed.ExternalObservation{}, nil
		}
		return managed.ExternalObservation{}, fmt.Errorf("failed to observe membership  %w", err)
	}
	state := string(*resp.State.Code)
	cr.Status.AtProvider.State = state
	switch state {
	case "READY":
		cr.Status.SetConditions(xpv1.Available())
	case "ERROR":
		cr.Status.SetConditions(xpv1.Unavailable())
	case "CREATING":
		cr.Status.SetConditions(xpv1.Creating())
	case "DELETING":
		cr.Status.SetConditions(xpv1.Deleting())
	default:
		cr.Status.SetConditions(xpv1.Creating())
	}

	return managed.ExternalObservation{
		ResourceExists:   true,
		ResourceUpToDate: true,
	}, nil
}

// Create resources!
func (e *MembershipExternal) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*v1alpha1.HubMembership)
	if !ok {
		return managed.ExternalCreation{}, fmt.Errorf("not membership")
	}
	if cr.Spec.ForProvider.GKEClusterID == "" {
		return managed.ExternalCreation{}, fmt.Errorf("No cluster ID")
	}
	obj := &gkehub.Membership{
		Project: dcl.String(e.projectID),
		Name:    dcl.String(cr.Spec.ForProvider.MembershipID),
		Endpoint: &gkehub.MembershipEndpoint{
			GkeCluster: &gkehub.MembershipEndpointGkeCluster{
				ResourceLink: dcl.String(fmt.Sprintf("//container.googleapis.com/%s", cr.Spec.ForProvider.GKEClusterID)),
			},
		},
		Location: dcl.String("global"),
	}
	_, err := e.gkeHub.ApplyMembership(ctx, obj, CreateDirective...)
	if err != nil {
		return managed.ExternalCreation{}, fmt.Errorf("failed to create membership  %w", err)
	}
	return managed.ExternalCreation{}, nil
}

// Update resource!
func (e *MembershipExternal) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	return managed.ExternalUpdate{}, fmt.Errorf("HubMembership Update is not supported yet")
}

// Delete resource!
func (e *MembershipExternal) Delete(ctx context.Context, mg resource.Managed) error {
	cr, ok := mg.(*v1alpha1.HubMembership)
	if !ok {
		return fmt.Errorf("not membership")
	}
	cr.SetConditions(xpv1.Deleting())
	obj := &gkehub.Membership{
		Project:  dcl.String(e.projectID),
		Name:     dcl.String(cr.Spec.ForProvider.MembershipID),
		Location: dcl.String("global"),
	}
	err := e.gkeHub.DeleteMembership(ctx, obj)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil
		}
		return fmt.Errorf("failed to delete membership  %w", err)
	}
	return nil
}
