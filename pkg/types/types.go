package types

// ObjectType is the type of object that is managed by the Back End.
type ObjectType string

const (
	// Namespace references objects that are of type corev1.Namespace
	Namespace ObjectType = "Namespace"
	// Node references objects that are of type corev1.Node
	Node ObjectType = "Node"
	// ReplicationController references objects that are of type corev1.ReplicationController
	ReplicationController ObjectType = "ReplicationController"
	// PersistentVolume references objects that are of type corev1.PersistentVolume
	PersistentVolume ObjectType = "PersistentVolume"
	// ClusterRole references objects that are of type rbacv1.ClusterRole
	ClusterRole ObjectType = "ClusterRole"
	// ClusterRoleBinding references objects that are of type rbacv1.ClusterRoleBinding
	ClusterRoleBinding ObjectType = "ClusterRoleBinding"
	// Service references objects that are of type corev1.Service
	Service ObjectType = "Service"
	// Pod references objects that are of type corev1.Pod
	Pod ObjectType = "Pod"
	// PersistentVolumeClaim references objects that are of type corev1.PersistentVolumeClaim
	PersistentVolumeClaim ObjectType = "PersistentVolumeClaim"
	// ConfigMap references objects that are of type corev1.ConfigMap
	ConfigMap ObjectType = "ConfigMap"
	// Secret references objects that are of type corev1.Secret
	Secret ObjectType = "Secret"
	// ServiceAccount references objects that are of type corev1.ServiceAccount
	ServiceAccount ObjectType = "ServiceAccount"
	// Deployment references objects that are of type appsv1.Deployment
	Deployment ObjectType = "Deployment"
	// StatefulSet references objects that are of type appsv1.StatefulSet
	StatefulSet ObjectType = "StatefulSet"
	// DaemonSet references objects that are of type appsv1.DaemonSet
	DaemonSet ObjectType = "DaemonSet"
	// ReplicaSet references objects that are of type appsv1.ReplicaSet
	ReplicaSet ObjectType = "ReplicaSet"
	// NetworkPolicy references objects that are of type networkingv1.NetworkPolicy
	NetworkPolicy ObjectType = "NetworkPolicy"
	// Role references objects that are of type rbacv1.Role
	Role ObjectType = "Role"
	// RoleBinding references objects that are of type rbacv1.RoleBinding
	RoleBinding ObjectType = "RoleBinding"
	// Job references objects that are of type batchv1.Job
	Job ObjectType = "Job"
	// CronJob references objects that are of type batchv1beta1.CronJob
	CronJob ObjectType = "CronJob"
	// StorageClass references objects that are of type storagev1.StorageClass
	StorageClass ObjectType = "StorageClass"
	// NodeMetrics references objects that are of type metricsv1beta1.NodeMetrics
	NodeMetrics ObjectType = "NodeMetrics"
	// PodMetrics references objects that are of type metricsv1beta1.StorageClass
	PodMetrics ObjectType = "PodMetrics"
)

// ObjectDefinition gives all the properties of an object that can be used into the go template
type ObjectDefinition struct {
	// The type of the object
	Type ObjectType
	// The family of object: cluster, namespace, metrics, etc.
	Family ObjectFamily
	// Variable is the name the variable to be used in the code, for example "pod"
	Variable string
	// PluralVariable is the pluralized name of the variable to be used in the code, for example "pods"
	PluralVariable string
	// Name is the name of the object by itself, for example "Pod"
	Name string
	// Plural is the pluralized name of the object by itself, for example "Pods"
	Plural string
	// The Kubernetes name of the object, including its package, for example "core.Pod"
	FullName string
	// The name of the REST provider used with the kubernetes client, for example "CoreV1()" for pods
	RestProvider string
	// The name of the REST resources, for example "pods" for pods / CoreV1()
	RestResourceName string
}

// IsClusterFamily checks if the object is related to cluster (excepted metrics)
func (definition *ObjectDefinition) IsClusterFamily() bool {
	return definition.Family == ClusterFamily
}

// IsClusterMetricsFamily checks if the object is related to cluster metrics
func (definition *ObjectDefinition) IsClusterMetricsFamily() bool {
	return definition.Family == ClusterMetricsFamily
}

// IsNamespaceFamily checks if the object is related to namespace (excepted metrics)
func (definition *ObjectDefinition) IsNamespaceFamily() bool {
	return definition.Family == NamespaceFamily
}

// IsNamespaceMetricsFamily checks if the object is related to namespace metrics
func (definition *ObjectDefinition) IsNamespaceMetricsFamily() bool {
	return definition.Family == NamespaceMetricsFamily
}

// ObjectFamily is the family (or category) with which the object is related. The family
// is defined by having an influence on the object calling API
type ObjectFamily int

const (
	// ClusterFamily references the object related to cluster (excepted metrics)
	ClusterFamily ObjectFamily = iota
	// NamespaceFamily references the object related to namespace (excepted metrics)
	NamespaceFamily
	// ClusterMetricsFamily references the object related to cluster metrics
	ClusterMetricsFamily
	// NamespaceMetricsFamily references the object related to namespace metrics
	NamespaceMetricsFamily
)

// ObjectDefinitions has all the definitions for the objects used by the API
var ObjectDefinitions = []ObjectDefinition{
	{Namespace, ClusterFamily, "namespace", "namespaces", "Namespace", "Namespaces", "corev1.Namespace", "CoreV1()", "namespaces"},
	{Node, ClusterFamily, "node", "nodes", "Node", "Nodes", "corev1.Node", "CoreV1()", "nodes"},
	{PersistentVolume, ClusterFamily, "persistentVolume", "persistentVolumes", "PersistentVolume", "PersistentVolumes", "corev1.PersistentVolume", "CoreV1()", "persistentvolumes"},
	{ClusterRole, ClusterFamily, "clusterRole", "clusterRoles", "ClusterRole", "ClusterRoles", "rbacv1.ClusterRole", "RbacV1()", "clusterroles"},
	{ClusterRoleBinding, ClusterFamily, "clusterRoleBinding", "clusterRoleBindings", "ClusterRoleBinding", "ClusterRoleBindings", "rbacv1.ClusterRoleBinding", "RbacV1()", "clusterrolebindings"},
	{StorageClass, ClusterFamily, "storageClass", "storageClasses", "StorageClass", "StorageClasses", "storagev1.StorageClass", "StorageV1()", "storageclasses"},
	{Service, NamespaceFamily, "service", "services", "Service", "Services", "corev1.Service", "CoreV1()", "services"},
	{Pod, NamespaceFamily, "pod", "pods", "Pod", "Pods", "corev1.Pod", "CoreV1()", "pods"},
	{PersistentVolumeClaim, NamespaceFamily, "persistentVolumeClaim", "persistentVolumeClaims", "PersistentVolumeClaim", "PersistentVolumeClaims", "corev1.PersistentVolumeClaim", "CoreV1()", "persistentvolumeclaims"},
	{ConfigMap, NamespaceFamily, "configMap", "configMaps", "ConfigMap", "ConfigMaps", "corev1.ConfigMap", "CoreV1()", "configmaps"},
	{ReplicationController, NamespaceFamily, "replicationController", "ReplicationControllers", "ReplicationController", "ReplicationControllers", "corev1.ReplicationController", "CoreV1()", "replicationcontrollers"},
	{Secret, NamespaceFamily, "secret", "secrets", "Secret", "Secrets", "corev1.Secret", "CoreV1()", "secrets"},
	{ServiceAccount, NamespaceFamily, "serviceAccount", "serviceAccounts", "ServiceAccount", "ServiceAccounts", "corev1.ServiceAccount", "CoreV1()", "serviceaccounts"},
	{Deployment, NamespaceFamily, "deployment", "deployments", "Deployment", "Deployments", "appsv1.Deployment", "AppsV1()", "deployments"},
	{StatefulSet, NamespaceFamily, "statefulSet", "statefulSets", "StatefulSet", "StatefulSets", "appsv1.StatefulSet", "AppsV1()", "statefulsets"},
	{DaemonSet, NamespaceFamily, "daemonSet", "daemonSets", "DaemonSet", "DaemonSets", "appsv1.DaemonSet", "AppsV1()", "daemonsets"},
	{ReplicaSet, NamespaceFamily, "replicaSet", "replicaSets", "ReplicaSet", "ReplicaSets", "appsv1.ReplicaSet", "AppsV1()", "replicasets"},
	{NetworkPolicy, NamespaceFamily, "networkPolicy", "networkPolicies", "NetworkPolicy", "NetworkPolicies", "networkingv1.NetworkPolicy", "NetworkingV1()", "networkpolicies"},
	{Role, NamespaceFamily, "role", "roles", "Role", "Roles", "rbacv1.Role", "RbacV1()", "roles"},
	{RoleBinding, NamespaceFamily, "roleBinding", "roleBindings", "RoleBinding", "RoleBindings", "rbacv1.RoleBinding", "RbacV1()", "rolebindings"},
	{Job, NamespaceFamily, "job", "jobs", "Job", "Jobs", "batchv1.Job", "BatchV1()", "jobs"},
	{CronJob, NamespaceFamily, "cronJob", "cronJobs", "CronJob", "CronJobs", "batchv1beta1.CronJob", "BatchV1beta1()", "cronjobs"},
	{NodeMetrics, ClusterMetricsFamily, "nodeMetrics", "nodeMetricses", "NodeMetrics", "NodeMetricses", "metricsv1beta1.NodeMetrics", "MetricsV1beta1()", "nodemetricses"},
	{PodMetrics, NamespaceMetricsFamily, "podMetrics", "podMetricses", "PodMetrics", "PodMetricses", "metricsv1beta1.PodMetrics", "MetricsV1beta1()", "podmetricses"},
}

// ClusterObjectDefinitions has all the definitions for the objects used by the API that are defined at the cluster level
var ClusterObjectDefinitions = func() []ObjectDefinition {
	results := make([]ObjectDefinition, 0, 5)
	for _, obj := range ObjectDefinitions {
		if obj.Family == ClusterFamily {
			results = append(results, obj)
		}
	}
	return results
}()

// NamespaceObjectDefinitions has all the definitions for the objects used by the API that are defined at the namespace level
var NamespaceObjectDefinitions = func() []ObjectDefinition {
	results := make([]ObjectDefinition, 0, 10)
	for _, obj := range ObjectDefinitions {
		if obj.Family == NamespaceFamily {
			results = append(results, obj)
		}
	}
	return results
}()

// ClusterMetricsObjectDefinitions has all the definitions for the objects used by the API that are defined at the metrics level
var ClusterMetricsObjectDefinitions = func() []ObjectDefinition {
	results := make([]ObjectDefinition, 0, 10)
	for _, obj := range ObjectDefinitions {
		if obj.Family == ClusterMetricsFamily {
			results = append(results, obj)
		}
	}
	return results
}()

// NamespaceMetricsObjectDefinitions has all the definitions for the objects used by the API that are defined at the metrics level
var NamespaceMetricsObjectDefinitions = func() []ObjectDefinition {
	results := make([]ObjectDefinition, 0, 10)
	for _, obj := range ObjectDefinitions {
		if obj.Family == NamespaceMetricsFamily {
			results = append(results, obj)
		}
	}
	return results
}()
