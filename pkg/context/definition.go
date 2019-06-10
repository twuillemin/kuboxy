package context

import "fmt"

// NamedContext is a Kubectl configuration, the association of a user and a cluster. This struct holds only a name and the
// actual definition structure
type NamedContext struct {
	Name              string            `yaml:"name" json:"name"`
	DefinitionContext DefinitionContext `yaml:"context,omitempty" json:"context,omitempty"`
}

// DefinitionContext is the actual definition of a Kubectl configuration
type DefinitionContext struct {
	User      string `yaml:"user,omitempty" json:"user"`
	Cluster   string `yaml:"cluster,omitempty" json:"cluster"`
	Namespace string `yaml:"namespace,omitempty" json:"namespace,omitempty"`
}

// NamedCluster is a Kubernetes cluster. This struct holds only a name and the
//// actual definition structure
type NamedCluster struct {
	Name              string            `yaml:"name" json:"name"`
	DefinitionCluster DefinitionCluster `yaml:"cluster,omitempty" json:"cluster,omitempty"`
}

// DefinitionCluster is the actual definition of a Kubernetes cluster
type DefinitionCluster struct {
	Server                   string `yaml:"server,omitempty" json:"server,omitempty"`
	InsecureSkipTLSVerify    bool   `yaml:"insecure-skip-tls-verify,omitempty" json:"insecure-skip-tls-verify,omitempty"`
	CertificateAuthority     string `yaml:"certificate-authority,omitempty" json:"certificate-authority,omitempty"`
	CertificateAuthorityData string `yaml:"certificate-authority-data,omitempty" json:"certificate-authority-data,omitempty"`
}

// NamedUser is a Kubernetes user. This struct holds only a name and the
//// actual definition structure
type NamedUser struct {
	Name           string         `yaml:"name" json:"name"`
	DefinitionUser DefinitionUser `yaml:"user,omitempty" json:"user,omitempty"`
}

// DefinitionUser is the actual definition of a Kubernetes user
type DefinitionUser struct {
	UserName              string `yaml:"username,omitempty" json:"username,omitempty"`
	Password              string `yaml:"password,omitempty" json:"password,omitempty"`
	ClientCertificate     string `yaml:"client-certificate,omitempty" json:"client-certificate,omitempty"`
	ClientKey             string `yaml:"client-key,omitempty" json:"client-key,omitempty"`
	ClientCertificateData string `yaml:"client-certificate-data,omitempty" json:"client-certificate-data,omitempty"`
	ClientKeyData         string `yaml:"client-key-data,omitempty" json:"client-key-data,omitempty"`
}

// KubeConfig is the complete configuration of a kubectl config file
type KubeConfig struct {
	APIVersion     string            `yaml:"apiVersion" json:"apiVersion"`
	Kind           string            `yaml:"kind" json:"kind"`
	Preferences    map[string]string `yaml:"preferences" json:"preferences"`
	Clusters       []NamedCluster    `yaml:"clusters,omitempty" json:"clusters,omitempty"`
	Contexts       []NamedContext    `yaml:"contexts,omitempty" json:"contexts,omitempty"`
	CurrentContext string            `yaml:"current-context,omitempty" json:"current-context,omitempty"`
	Users          []NamedUser       `yaml:"users,omitempty" json:"users,omitempty"`
}

// ParamCredentialsUserNamePassword  is the definition of a credential with a username and a password
type ParamCredentialsUserNamePassword struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

// ParamCredentialsCertificateFile  is the definition of a credential with a certificate kept as a local file
type ParamCredentialsCertificateFile struct {
	ClientCertificate string `json:"clientCertificate"`
	ClientKey         string `json:"clientKey"`
}

// ParamCredentialsCertificateEmbedded  is the definition of a credential with a certificate embedded
type ParamCredentialsCertificateEmbedded struct {
	ClientCertificateData string `json:"clientCertificateData"`
	ClientKeyData         string `json:"clientKeyData"`
}

// ParamClusterInsecure  is the definition of a cluster for which the TLS certificate is not verified
type ParamClusterInsecure struct {
	Server string `json:"server"`
}

// ParamClusterCertificateFile  is the definition of a cluster for which the TLS certificate is verified with a local file
type ParamClusterCertificateFile struct {
	Server               string `json:"server"`
	CertificateAuthority string `json:"certificateAuthority"`
}

// ParamClusterCertificateEmbedded  is the definition of a cluster for which the TLS certificate is verified with a local file
type ParamClusterCertificateEmbedded struct {
	Server                   string `json:"server"`
	CertificateAuthorityData string `json:"certificateAuthorityData"`
}

// ParamContext  is the definition of a context
type ParamContext struct {
	User      string `json:"user"`
	Cluster   string `json:"cluster"`
	Namespace string `json:"namespace"`
}

// NotFoundError is a trivial implementation of error.
type NotFoundError struct {
	contextName string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("the context \"%s\" does not exist", e.contextName)
}

// ContextName returns the name of the context that was not found
func (e *NotFoundError) ContextName() string {
	return e.contextName
}
