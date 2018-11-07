package v3

import (
	"github.com/rancher/norman/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GlobalDNS struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	//ObjectMeta.Name is the FQDN
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlobalDNSSpec   `json:"spec,omitempty"`
	Status GlobalDNSStatus `json:"status,omitempty"`
}

type GlobalDNSSpec struct {
	RootDomain   string   `json:"rootDomain" norman:"required"`
	ProviderName string   `json:"providerName,omitempty" norman:"required"`
	ClusterNames []string `json:"clusterNames,omitempty" norman:"type=array[reference[cluster]]"`
}

type GlobalDNSStatus struct {
	Endpoints []string `json:"endpoints,omitempty"`
}

type MultiClusterApp struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MultiClusterAppSpec   `json:"spec,omitempty"`
	Status MultiClusterAppStatus `json:"status,omitempty"`
}

type MultiClusterAppSpec struct {
	ReleaseName      string              `json:"releaseName,omitempty"`
	ReleaseNamespace string              `json:"releaseNamespace,omitempty"`
	Answers          map[string]string   `json:"answers,omitempty"`
	HelmChart        *HelmChartReference `json:"helmChart,omitempty"`
	Template         *TemplateReference  `json:"template,omitempty"`
	Targets          []Target            `json:"targets,omitempty" norman:"required"`
}

type MultiClusterAppStatus struct {
	HealthState string `json:"healthState,omitempty"`
}

type HelmChartReference struct {
	RepositoryURL string `json:"repositoryUrl,omitempty" norman:"required"`
	Reference     string `json:"reference,omitempty" norman:"required"`
	Version       string `json:"version,omitempty"`
}

type TemplateReference struct {
	Name    string `json:"name,omitempty" norman:"required,type=reference[template]"`
	Version string `json:"version,omitempty" norman:"required,type=reference[templateVersion]"`
}

type Target struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TargetSpec   `json:"spec,omitempty" norman:"required"`
	Status TargetStatus `json:"status,omitempty" norman:"required"`
}

type TargetSpec struct {
	ClusterName   string            `json:"clusterName,omitempty" norman:"type=reference[cluster]"`
	ClusterConfig *ClusterConfig    `json:"clusterConfig,omitempty"`
	Answers       map[string]string `json:"answers,omitempty"`
}

type TargetStatus struct {
	ChartReleaseName string `json:"chartReleaseName,omitempty"`
	HealthState      string `json:"healthState,omitempty"`
}

type ClusterConfig struct {
	Server                   string `json:"server,omitempty" norman:"required"`
	CertificateAuthorityPath string `json:"certificateAuthorityPath,omitempty"`
	CertificateAuthorityData []byte `json:"certificateAuthorityData,omitempty"`
	ClientCertificatePath    string `json:"clientCertificatePath,omitempty"`
	ClientCertificateData    []byte `json:"clientCertificateData,omitempty"`
	ClientKeyPath            string `json:"clientKeyPath,omitempty"`
	ClientKeyData            []byte `json:"clientKeyData,omitempty"`
	TokenFile                string `json:"tokenFile,omitempty"`
}
