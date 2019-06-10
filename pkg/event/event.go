package event

//go:generate go run gen/gen_event_definition.go
//go:generate go run gen/gen_event_cluster.go
//go:generate go run gen/gen_event_namespace.go

import (
	"k8s.io/client-go/kubernetes"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

// Type is the type of event: Create, Update or Delete
type Type int

const (
	// Create is used when an object is created
	Create Type = iota
	// Update is used when an object is updated
	Update
	// Delete is used when an object is deleted
	Delete
)

type contextReceiver struct {
	clientset                       *kubernetes.Clientset
	metrics                         *metrics.Clientset
	namespaceEventReceiver          *namespaceEventReceiver
	nodeEventReceiver               *nodeEventReceiver
	persistentVolumeEventReceiver   *persistentVolumeEventReceiver
	clusterRoleEventReceiver        *clusterRoleEventReceiver
	clusterRoleBindingEventReceiver *clusterRoleBindingEventReceiver
	storageClassEventReceiver       *storageClassEventReceiver
	nodeMetricsEventReceiver        *nodeMetricsEventReceiver
	namespaceReceivers              map[string]*namespaceReceiver
}

type namespaceReceiver struct {
	serviceEventReceiver               *serviceEventReceiver
	podEventReceiver                   *podEventReceiver
	persistentVolumeClaimEventReceiver *persistentVolumeClaimEventReceiver
	configMapEventReceiver             *configMapEventReceiver
	secretEventReceiver                *secretEventReceiver
	serviceAccountEventReceiver        *serviceAccountEventReceiver
	replicationControllerEventReceiver *replicationControllerEventReceiver
	deploymentEventReceiver            *deploymentEventReceiver
	statefulSetEventReceiver           *statefulSetEventReceiver
	daemonSetEventReceiver             *daemonSetEventReceiver
	replicaSetEventReceiver            *replicaSetEventReceiver
	networkPolicyEventReceiver         *networkPolicyEventReceiver
	roleEventReceiver                  *roleEventReceiver
	roleBindingEventReceiver           *roleBindingEventReceiver
	jobEventReceiver                   *jobEventReceiver
	cronJobEventReceiver               *cronJobEventReceiver
	podMetricsEventReceiver            *podMetricsEventReceiver
}

// The list of all contextReceivers by context name
var contextReceivers = make(map[string]*contextReceiver)
