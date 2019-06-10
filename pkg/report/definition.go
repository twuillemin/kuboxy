package report

import (
	"k8s.io/apimachinery/pkg/types"
)

// ClusterStateReport is the complete summary (so without all the details) for a cluster
type ClusterStateReport struct {
	Namespaces                   []NamespaceReport             `json:"namespaces"`
	Nodes                        []NodeReport                  `json:"nodes"`
	Services                     []ServiceReport               `json:"services"`
	Pods                         []PodReport                   `json:"pods"`
	PersistentVolumeReports      []PersistentVolumeReport      `json:"persistentVolumes"`
	PersistentVolumeClaimReports []PersistentVolumeClaimReport `json:"persistentVolumeClaims"`
	ConfigMapReports             []ConfigMapReport             `json:"configMaps"`
	SecretReports                []SecretReport                `json:"secrets"`
	Deployments                  []DeploymentReport            `json:"deployments"`
}

// NamespaceReport is the summary of the namespaces of the cluster
type NamespaceReport struct {
	ID            types.UID         `json:"id"`
	NameSpace     string            `json:"namespace"`
	Name          string            `json:"name"`
	Labels        map[string]string `json:"labels"`
	ServiceIds    []types.UID       `json:"serviceIds"`
	PodIds        []types.UID       `json:"podIds"`
	DeploymentIds []types.UID       `json:"deploymentsIds"`
}

// NodeReport is the summary of the nodes of the cluster
type NodeReport struct {
	ID              types.UID         `json:"id"`
	NameSpace       string            `json:"namespace"`
	Name            string            `json:"name"`
	Labels          map[string]string `json:"labels"`
	Condition       string            `json:"condition"`
	AllowedCPUCores int               `json:"allowedCPUCores"`
	AllowedRAM      int               `json:"allowedRAM"`
	UsageCPUCores   *float64          `json:"usageCPUCores"`
	UsageRAM        *int              `json:"usageRAM"`
	PodIds          []types.UID       `json:"podIds"`
}

// ServiceReport is the summary of the services of the cluster
type ServiceReport struct {
	ID          types.UID           `json:"id"`
	NameSpace   string              `json:"namespace"`
	Name        string              `json:"name"`
	ClusterIP   string              `json:"clusterIP"`
	ExternalIPs []string            `json:"externalIP"`
	Ports       []ServicePortReport `json:"ports"`
	Labels      map[string]string   `json:"labels"`
	Selector    map[string]string   `json:"selector"`
	PodIds      []types.UID         `json:"podIds"`
}

// ServicePortReport is the summary of the port used by a service
type ServicePortReport struct {
	Protocol         string  `json:"protocol"`
	InternalPort     int32   `json:"internalPort"`
	InternalPortName *string `json:"internalPortName"`
	NodePort         *int32  `json:"nodePort"`
}

// ContainerReport is the summary of the containers of a pod
type ContainerReport struct {
	ID        string          `json:"id"`
	NameSpace string          `json:"namespace"`
	Name      string          `json:"name"`
	Image     string          `json:"image"`
	Ports     []PodPortReport `json:"ports"`
	Ready     bool            `json:"ready"`
}

// PodPortReport is the summary of the ports of a pod
type PodPortReport struct {
	Protocol      string  `json:"protocol"`
	ContainerPort int32   `json:"containerPort"`
	Name          *string `json:"name"`
}

// PodReport is the summary of the pods in a cluster
type PodReport struct {
	ID            types.UID         `json:"id"`
	NameSpace     string            `json:"namespace"`
	Name          string            `json:"name"`
	Labels        map[string]string `json:"labels"`
	Condition     string            `json:"condition"`
	IP            string            `json:"ip"`
	UsageCPUCores *float64          `json:"usageCPUCores"`
	UsageRAM      *int              `json:"usageRAM"`
	Containers    []ContainerReport `json:"containers"`
	Volumes       []VolumeReport    `json:"volumes"`
}

// VolumeReport is the summary of the volumes in a cluster
type VolumeReport struct {
	Name                      string      `json:"name"`
	StorageType               StorageType `json:"type"`
	PersistentVolumeClaimName *string     `json:"persistentVolumeClaimName"`
}

// PersistentVolumeReport is the summary of the persistent volumes in a cluster
type PersistentVolumeReport struct {
	ID          types.UID   `json:"id"`
	NameSpace   string      `json:"namespace"`
	Name        string      `json:"name"`
	StorageType StorageType `json:"type"`
	Capacity    *int        `json:"capacity"`
}

// PersistentVolumeClaimReport is the summary of the persistent volume claims in a cluster
type PersistentVolumeClaimReport struct {
	ID         types.UID `json:"id"`
	NameSpace  string    `json:"namespace"`
	Name       string    `json:"name"`
	VolumeName string    `json:"volumeName"`
}

// ConfigMapReport is the summary of the config maps in a cluster
type ConfigMapReport struct {
	ID        types.UID         `json:"id"`
	NameSpace string            `json:"namespace"`
	Name      string            `json:"name"`
	Data      map[string]string `json:"data"`
}

// SecretReport is the summary of the secrets in a cluster
type SecretReport struct {
	ID        types.UID         `json:"id"`
	NameSpace string            `json:"namespace"`
	Name      string            `json:"name"`
	Data      map[string][]byte `json:"data"`
}

// DeploymentReport is the summary of the deployments in a cluster
type DeploymentReport struct {
	ID        types.UID         `json:"id"`
	NameSpace string            `json:"namespace"`
	Name      string            `json:"name"`
	Labels    map[string]string `json:"labels"`
	Selector  map[string]string `json:"selector"`
	Volumes   []VolumeReport    `json:"volumes"`
}

// StorageType is a list of all possible storage for volumes and persistent volumes
type StorageType string

// These are valid address type of node.
const (
	/*
	 * Suitable for Volume and PersistentVolume
	 */

	// AWSElasticBlockStore is storage in Amazon
	AWSElasticBlockStore StorageType = "AWSElasticBlockStore"
	AzureDisk            StorageType = "AzureDisk"
	AzureFile            StorageType = "AzureFile"
	CephFS               StorageType = "CephFS"
	Cinder               StorageType = "Cinder"
	FC                   StorageType = "FC"
	Flocker              StorageType = "Flocker"
	FlexVolume           StorageType = "FlexVolume"
	GCEPersistentDisk    StorageType = "GCEPersistentDisk"
	Glusterfs            StorageType = "Glusterfs"
	HostPath             StorageType = "HostPath"
	ISCSI                StorageType = "ISCSI"
	NFS                  StorageType = "NFS"
	Quobyte              StorageType = "Quobyte"
	RBD                  StorageType = "RBD"
	ScaleIO              StorageType = "ScaleIO"
	StorageOS            StorageType = "StorageOS"
	VsphereVolume        StorageType = "VsphereVolume"

	/*
	 * Suitable only for PersistentVolume
	 */

	// CSI is storage in CSI
	CSI   StorageType = "CSI"
	Local StorageType = "Local"

	/*
	 * Suitable only for Volume
	 */

	// ConfigMap is storage in ConfigMap
	ConfigMap             StorageType = "ConfigMap"
	DownwardAPI           StorageType = "DownwardAPI"
	EmptyDir              StorageType = "EmptyDir"
	PersistentVolumeClaim StorageType = "PersistentVolumeClaim"
	PhotonPersistentDisk  StorageType = "PhotonPersistentDisk"
	PortworxVolume        StorageType = "PortworxVolume"
	Projected             StorageType = "Projected"
	Secret                StorageType = "Secret"

	// Unknown is an unknown storage, not good...
	Unknown StorageType = "Unknown"
)
