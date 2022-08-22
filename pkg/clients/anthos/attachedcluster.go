/*
Copyright 2021 The Crossplane Authors.
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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
	htransport "google.golang.org/api/transport/http"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane-contrib/provider-gcp/apis/anthos/v1alpha1"
	"github.com/crossplane-contrib/provider-gcp/apis/v1beta1"
	gcp "github.com/crossplane-contrib/provider-gcp/pkg/clients"
)

// ErrNotExist is the rror when the cluster does not exits
var ErrNotExist = errors.New("AttachedCluster does not exists")

// Service is the attached cluster service.
type Service struct {
	projectID string
	opts      []option.ClientOption
	Kube      client.Client
}

// NewService returns a new servcie.
func NewService(ctx context.Context, kube client.Client, mg resource.Managed) (*Service, error) {
	projectID, opts, err := gcp.GetConnectionInfo(ctx, kube, mg)
	if err != nil {
		return nil, err
	}

	pc := &v1beta1.ProviderConfig{}
	t := resource.NewProviderConfigUsageTracker(kube, &v1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, err
	}
	if err := kube.Get(ctx, types.NamespacedName{Name: mg.GetProviderConfigReference().Name}, pc); err != nil {
		return nil, err
	}

	var userClusterKube client.Client
	cred := pc.Spec.AnthosCredentials
	if cred != nil {
		var rc *rest.Config
		kc, err := resource.CommonCredentialExtractor(ctx, cred.Source, kube, cred.CommonCredentialSelectors)
		if err != nil {
			return nil, err
		}
		if rc, err = NewRESTConfig(kc); err != nil {
			return nil, err
		}
		userClusterKube, err = NewKubeClient(rc)
		if err != nil {
			return nil, err
		}
	}
	opts = append(opts, option.WithAudiences("https://autopush-gkemulticloud.sandbox.googleapis.com/"))
	return &Service{
		projectID: projectID,
		opts:      opts,
		Kube:      userClusterKube,
	}, nil
}

// GetInstallManifest gets the install manifests.
func (s *Service) GetInstallManifest(ctx context.Context, location string, i InstallManifest) (*InstallManifest, error) {
	return getInstallManifest(ctx, GCPOpts{
		Project: s.projectID,
		Region:  location,
	}, i, s.opts...)
}

// CreateAttachedCluster creates the attached clsuter.
func (s *Service) CreateAttachedCluster(ctx context.Context, a *v1alpha1.AttachedCluster) error {
	return createAttachedCluster(ctx, GCPOpts{s.projectID, a.Spec.ForProvider.Location},
		AttachedCluster{
			Name: a.Name,
			Authority: Authority{
				IssuerURL: a.Spec.ForProvider.Authority.IssuerURL,
			},
			Fleet: Fleet{
				Project: a.Spec.ForProvider.Fleet.Project,
			},
			PlatformVersion: a.Spec.ForProvider.PlatformVersion,
		}, s.opts...,
	)
}

// GetAttachedCluster gets the attached cluster.
func (s *Service) GetAttachedCluster(ctx context.Context, a *v1alpha1.AttachedCluster) error {
	at, err := getAttachedCluster(ctx, a.Name, GCPOpts{s.projectID, a.Spec.ForProvider.Location}, s.opts...)
	if err != nil {
		return err
	}
	a.Status.AtProvider.State = at.State
	a.Status.AtProvider.MembershipID = fmt.Sprintf("projects/%s/locations/%s/memberships/%s", a.Spec.ForProvider.Fleet.Project, a.Spec.ForProvider.Location, a.Name)
	return nil
}

// DeleteAttachedCluster deletes the attached cluster.
func (s *Service) DeleteAttachedCluster(ctx context.Context, a *v1alpha1.AttachedCluster) error {
	return deleteAttachedCluster(ctx, a.Name, GCPOpts{s.projectID, a.Spec.ForProvider.Location}, s.opts...)
}

// NewRESTConfig returns a rest config given a secret with connection information.
func NewRESTConfig(kubeconfig []byte) (*rest.Config, error) {
	ac, err := clientcmd.Load(kubeconfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load kubeconfig")
	}
	return restConfigFromAPIConfig(ac)
}

// NewKubeClient returns a kubernetes client given a secret with connection
// information.
func NewKubeClient(config *rest.Config) (client.Client, error) {
	kc, err := client.New(config, client.Options{})
	if err != nil {
		return nil, errors.Wrap(err, "cannot create Kubernetes client")
	}

	return kc, nil
}

func restConfigFromAPIConfig(c *api.Config) (*rest.Config, error) {
	if c.CurrentContext == "" {
		return nil, errors.New("currentContext not set in kubeconfig")
	}
	ctx := c.Contexts[c.CurrentContext]
	cluster := c.Clusters[ctx.Cluster]
	if cluster == nil {
		return nil, errors.Errorf("cluster for currentContext (%s) not found", c.CurrentContext)
	}
	user := c.AuthInfos[ctx.AuthInfo]
	if user == nil {
		// We don't require a user because it's possible user
		// authorization configuration will be loaded from a separate
		// set of identity credentials (e.g. Google Application Creds).
		user = &api.AuthInfo{}
	}
	return &rest.Config{
		Host:            cluster.Server,
		Username:        user.Username,
		Password:        user.Password,
		BearerToken:     user.Token,
		BearerTokenFile: user.TokenFile,
		Impersonate: rest.ImpersonationConfig{
			UserName: user.Impersonate,
			Groups:   user.ImpersonateGroups,
			Extra:    user.ImpersonateUserExtra,
		},
		AuthProvider: user.AuthProvider,
		ExecProvider: user.Exec,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure:   cluster.InsecureSkipTLSVerify,
			ServerName: cluster.TLSServerName,
			CertData:   user.ClientCertificateData,
			KeyData:    user.ClientKeyData,
			CAData:     cluster.CertificateAuthorityData,
		},
	}, nil
}

// AttachedCluster is the attached cluster resource.
type AttachedCluster struct {
	Name            string    `json:"name,omitempty"`
	Authority       Authority `json:"authority,omitempty"`
	Fleet           Fleet     `json:"fleet,omitempty"`
	PlatformVersion string    `json:"platform_version,omitempty"`
	State           string    `json:"state,omitempty"`
	UID             string    `json:"uid,omitempty"`
	CreateTime      string    `json:"create_time,omitempty"`
	Etag            string    `json:"etag,omitempty"`
}

// InstallManifest is the install manifest for attached cluster.
type InstallManifest struct {
	AttachedClusterID string `json:"attached_cluster_id,omitempty"`
	PlatformVersion   string `json:"platform_version,omitempty"`
	Manifest          string `json:"manifest,omitempty"`
}

// Authority is the authority setting.
type Authority struct {
	IssuerURL string `json:"issuer_url,omitempty"`
}

// Fleet is the fleet setting.
type Fleet struct {
	Project string `json:"project,omitempty"`
}

// GCPOpts sets the gcp parameters.
type GCPOpts struct {
	Project string
	Region  string
}

// GetInstallManifest just get the install manifest.
func getInstallManifest(ctx context.Context, gcp GCPOpts, i InstallManifest, opts ...option.ClientOption) (*InstallManifest, error) {
	c, _, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://autopush-gkemulticloud.sandbox.googleapis.com/v1/projects/%s/locations/%s/generateAttachedClusterInstallManifest?attached_cluster_id=%s&platform_version=%s", gcp.Project, gcp.Region, i.AttachedClusterID, i.PlatformVersion)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Gfe-Ssl", "yes")

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Printf("got response %s", string(r))
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("get manifest faield %s", string(r))
	}
	a := &InstallManifest{}
	if err := json.Unmarshal(r, a); err != nil {
		return nil, err
	}
	return a, nil
}

// CreateAttachedCluster creates an attached cluster
func createAttachedCluster(ctx context.Context, gcp GCPOpts, attached AttachedCluster, opts ...option.ClientOption) error {
	c, _, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://autopush-gkemulticloud.sandbox.googleapis.com/v1/projects/%s/locations/%s/attachedClusters?attached_cluster_id=%s", gcp.Project, gcp.Region, attached.Name)
	body, err := json.Marshal(attached)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(string(body)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Gfe-Ssl", "yes")

	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("got response %s", string(r))
	if resp.StatusCode != 200 {
		return fmt.Errorf("creation faield %s", string(r))
	}
	return nil
}

// DeleteAttachedCluster deletes an attached cluster.
func deleteAttachedCluster(ctx context.Context, name string, gcp GCPOpts, opts ...option.ClientOption) error {
	c, _, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://autopush-gkemulticloud.sandbox.googleapis.com/v1/projects/%s/locations/%s/attachedClusters/%s", gcp.Project, gcp.Region, name)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Gfe-Ssl", "yes")

	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("got response %s", string(r))
	if resp.StatusCode != 200 {
		return fmt.Errorf("deletion faield %s", string(r))
	}
	return nil
}

// GetAttachedCluster deletes an attached cluster.
func getAttachedCluster(ctx context.Context, name string, gcp GCPOpts, opts ...option.ClientOption) (*AttachedCluster, error) {
	c, _, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://autopush-gkemulticloud.sandbox.googleapis.com/v1/projects/%s/locations/%s/attachedClusters/%s", gcp.Project, gcp.Region, name)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Gfe-Ssl", "yes")

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Printf("got response %s", string(r))
	if resp.StatusCode == 404 {
		return nil, ErrNotExist
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("get faield %s", string(r))
	}
	a := &AttachedCluster{}
	if err := json.Unmarshal(r, a); err != nil {
		return nil, err
	}
	return a, nil
}
