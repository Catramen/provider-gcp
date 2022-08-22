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

package featuremembership

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

// SetupHubFeatureMembership adds a controller that reconciles HubFeatureMembership managed
// resources.
func SetupHubFeatureMembership(mgr ctrl.Manager, o controller.Options) error {
	name := managed.ControllerName(v1alpha1.HubFeatureMembershipGroupKind)

	cps := []managed.ConnectionPublisher{managed.NewAPISecretPublisher(mgr.GetClient(), mgr.GetScheme())}

	r := managed.NewReconciler(mgr,
		resource.ManagedKind(v1alpha1.HubFeatureMembershipGroupVersionKind),
		managed.WithExternalConnecter(&Connector{kube: mgr.GetClient()}),
		managed.WithPollInterval(o.PollInterval),
		managed.WithLogger(o.Logger.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		managed.WithConnectionPublishers(cps...))

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		For(&v1alpha1.HubFeatureMembership{}).
		Complete(ratelimiter.NewReconciler(name, r, o.GlobalRateLimiter))
}

// Connector is the connector of hub feature.
type Connector struct {
	kube client.Client
}

// Connect to clients!
func (c *Connector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
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
		dcl.WithBasePath(gkeHubBase + "v1beta/"),
	}

	dclConfig := dcl.NewConfig(configOptions...)
	s := gkehub.NewClient(dclConfig)
	return &External{kube: c.kube, projectID: projectID, gkeHub: s}, nil
}

// External is the external resource manager.
type External struct {
	kube      client.Client
	gkeHub    *gkehub.Client
	projectID string
}

// Observe resources!
func (e *External) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) { // nolint:gocyclo
	cr, ok := mg.(*v1alpha1.HubFeatureMembership)
	if !ok {
		return managed.ExternalObservation{}, fmt.Errorf("not feature")
	}
	obj := &gkehub.FeatureMembership{
		Project:          dcl.String(e.projectID),
		Location:         dcl.String(cr.Spec.ForProvider.Location),
		Feature:          dcl.String(cr.Spec.ForProvider.Feature),
		Membership:       dcl.String(cr.Spec.ForProvider.Membership),
		Configmanagement: expandGkeHubFeatureMembershipConfigmanagement(cr),
	}
	_, err := e.gkeHub.GetFeatureMembership(ctx, obj)
	if err != nil {
		if dcl.IsNotFound(err) {
			return managed.ExternalObservation{}, nil
		}
		return managed.ExternalObservation{}, fmt.Errorf("failed to create observe membership %w", err)
	}
	cr.Status.SetConditions(xpv1.Available())
	return managed.ExternalObservation{
		ResourceExists:   true,
		ResourceUpToDate: true,
	}, nil
}

// Create resources!
func (e *External) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*v1alpha1.HubFeatureMembership)
	if !ok {
		return managed.ExternalCreation{}, fmt.Errorf("not feature")
	}
	obj := &gkehub.FeatureMembership{
		Project:          dcl.String(e.projectID),
		Location:         dcl.String(cr.Spec.ForProvider.Location),
		Feature:          dcl.String(cr.Spec.ForProvider.Feature),
		Membership:       dcl.String(cr.Spec.ForProvider.Membership),
		Configmanagement: expandGkeHubFeatureMembershipConfigmanagement(cr),
	}
	_, err := e.gkeHub.ApplyFeatureMembership(ctx, obj, CreateDirective...)
	if err != nil {
		return managed.ExternalCreation{}, fmt.Errorf("failed to create feature membership %w", err)
	}
	return managed.ExternalCreation{}, nil
}

// Update resource!
func (e *External) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	return managed.ExternalUpdate{}, fmt.Errorf("HubFeatureMembership Update is not supported yet")
}

// Delete resource!
func (e *External) Delete(ctx context.Context, mg resource.Managed) error {
	cr, ok := mg.(*v1alpha1.HubFeatureMembership)
	if !ok {
		return fmt.Errorf("not feature")
	}
	cr.SetConditions(xpv1.Deleting())

	obj := &gkehub.FeatureMembership{
		Project:          dcl.String(e.projectID),
		Location:         dcl.String(cr.Spec.ForProvider.Location),
		Feature:          dcl.String(cr.Spec.ForProvider.Feature),
		Membership:       dcl.String(cr.Spec.ForProvider.Membership),
		Configmanagement: expandGkeHubFeatureMembershipConfigmanagement(cr),
	}
	err := e.gkeHub.DeleteFeatureMembership(ctx, obj)
	if err != nil {
		if dcl.IsNotFound(err) {
			return nil
		}
		return fmt.Errorf("failed to create feature membership %w", err)
	}
	return nil
}

func expandGkeHubFeatureMembershipConfigmanagement(cr *v1alpha1.HubFeatureMembership) *gkehub.FeatureMembershipConfigmanagement {
	return &gkehub.FeatureMembershipConfigmanagement{
		ConfigSync: &gkehub.FeatureMembershipConfigmanagementConfigSync{
			Git: expandGkeHubFeatureMembershipConfigmanagementConfigSyncGit(cr),
		},
		Version: dcl.StringOrNil(cr.Spec.ForProvider.ConfigManagement.Version),
	}
}

func expandGkeHubFeatureMembershipConfigmanagementConfigSyncGit(cr *v1alpha1.HubFeatureMembership) *gkehub.FeatureMembershipConfigmanagementConfigSyncGit {
	return &gkehub.FeatureMembershipConfigmanagementConfigSyncGit{
		//GcpServiceAccountEmail: dcl.String(obj["gcp_service_account_email"].(string)),
		//HttpsProxy:             dcl.String(obj["https_proxy"].(string)),
		PolicyDir:  dcl.String(cr.Spec.ForProvider.ConfigManagement.Git.PolicyDir),
		SecretType: dcl.String(cr.Spec.ForProvider.ConfigManagement.Git.SecretType),
		SyncBranch: dcl.String(cr.Spec.ForProvider.ConfigManagement.Git.Branch),
		SyncRepo:   dcl.String(cr.Spec.ForProvider.ConfigManagement.Git.Repo),
		//SyncRev:                dcl.String(obj["sync_rev"].(string)),
		//SyncWaitSecs:           dcl.String(obj["sync_wait_secs"].(string)),
	}
}
