// Package provider regroups all the basic CRUD functions to access the objects in the cluster
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_provider_namespace.go at 2019-06-12 21:55:14.865918493 +0300 EEST m=+0.001996219
package provider

import (
	"github.com/twuillemin/kuboxy/pkg/connector"
	"github.com/twuillemin/kuboxy/pkg/context"
	"github.com/twuillemin/kuboxy/pkg/event"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	rbacv1 "k8s.io/api/rbac/v1"
)

// GetServices returns all the Service. If an empty namespace is given, returns all the Service
func GetServices(contextName string, namespace string) ([]corev1.Service, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetServices(contextName, namespace); results != nil {
		return results, nil
	}

	return connector.GetServices(clientset, namespace)
}

// GetService returns the Service by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func GetService(contextName string, namespace string, name string) (*corev1.Service, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetServices(contextName, namespace); results != nil {
		for _, service := range results {
			if service.Name == name {
				return &service, nil
			}
		}
		return nil, nil
	}

	return connector.GetService(clientset, namespace, name)
}

// CreateService creates the Service with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func CreateService(contextName string, namespace string, service *corev1.Service) (*corev1.Service, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.CreateService(clientset, namespace, service)
}

// UpdateService updates the Service with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func UpdateService(contextName string, namespace string, service *corev1.Service) (*corev1.Service, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.UpdateService(clientset, namespace, service)
}

// DeleteService deletes the Service by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func DeleteService(contextName string, namespace string, name string) error {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return err
	}

	return connector.DeleteService(clientset, namespace, name)
}

// GetPods returns all the Pod. If an empty namespace is given, returns all the Pod
func GetPods(contextName string, namespace string) ([]corev1.Pod, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetPods(contextName, namespace); results != nil {
		return results, nil
	}

	return connector.GetPods(clientset, namespace)
}

// GetPod returns the Pod by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func GetPod(contextName string, namespace string, name string) (*corev1.Pod, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetPods(contextName, namespace); results != nil {
		for _, pod := range results {
			if pod.Name == name {
				return &pod, nil
			}
		}
		return nil, nil
	}

	return connector.GetPod(clientset, namespace, name)
}

// CreatePod creates the Pod with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func CreatePod(contextName string, namespace string, pod *corev1.Pod) (*corev1.Pod, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.CreatePod(clientset, namespace, pod)
}

// UpdatePod updates the Pod with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func UpdatePod(contextName string, namespace string, pod *corev1.Pod) (*corev1.Pod, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.UpdatePod(clientset, namespace, pod)
}

// DeletePod deletes the Pod by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func DeletePod(contextName string, namespace string, name string) error {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return err
	}

	return connector.DeletePod(clientset, namespace, name)
}

// GetPersistentVolumeClaims returns all the PersistentVolumeClaim. If an empty namespace is given, returns all the PersistentVolumeClaim
func GetPersistentVolumeClaims(contextName string, namespace string) ([]corev1.PersistentVolumeClaim, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetPersistentVolumeClaims(contextName, namespace); results != nil {
		return results, nil
	}

	return connector.GetPersistentVolumeClaims(clientset, namespace)
}

// GetPersistentVolumeClaim returns the PersistentVolumeClaim by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func GetPersistentVolumeClaim(contextName string, namespace string, name string) (*corev1.PersistentVolumeClaim, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetPersistentVolumeClaims(contextName, namespace); results != nil {
		for _, persistentVolumeClaim := range results {
			if persistentVolumeClaim.Name == name {
				return &persistentVolumeClaim, nil
			}
		}
		return nil, nil
	}

	return connector.GetPersistentVolumeClaim(clientset, namespace, name)
}

// CreatePersistentVolumeClaim creates the PersistentVolumeClaim with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func CreatePersistentVolumeClaim(contextName string, namespace string, persistentVolumeClaim *corev1.PersistentVolumeClaim) (*corev1.PersistentVolumeClaim, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.CreatePersistentVolumeClaim(clientset, namespace, persistentVolumeClaim)
}

// UpdatePersistentVolumeClaim updates the PersistentVolumeClaim with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func UpdatePersistentVolumeClaim(contextName string, namespace string, persistentVolumeClaim *corev1.PersistentVolumeClaim) (*corev1.PersistentVolumeClaim, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.UpdatePersistentVolumeClaim(clientset, namespace, persistentVolumeClaim)
}

// DeletePersistentVolumeClaim deletes the PersistentVolumeClaim by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func DeletePersistentVolumeClaim(contextName string, namespace string, name string) error {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return err
	}

	return connector.DeletePersistentVolumeClaim(clientset, namespace, name)
}

// GetConfigMaps returns all the ConfigMap. If an empty namespace is given, returns all the ConfigMap
func GetConfigMaps(contextName string, namespace string) ([]corev1.ConfigMap, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetConfigMaps(contextName, namespace); results != nil {
		return results, nil
	}

	return connector.GetConfigMaps(clientset, namespace)
}

// GetConfigMap returns the ConfigMap by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func GetConfigMap(contextName string, namespace string, name string) (*corev1.ConfigMap, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetConfigMaps(contextName, namespace); results != nil {
		for _, configMap := range results {
			if configMap.Name == name {
				return &configMap, nil
			}
		}
		return nil, nil
	}

	return connector.GetConfigMap(clientset, namespace, name)
}

// CreateConfigMap creates the ConfigMap with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func CreateConfigMap(contextName string, namespace string, configMap *corev1.ConfigMap) (*corev1.ConfigMap, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.CreateConfigMap(clientset, namespace, configMap)
}

// UpdateConfigMap updates the ConfigMap with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func UpdateConfigMap(contextName string, namespace string, configMap *corev1.ConfigMap) (*corev1.ConfigMap, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.UpdateConfigMap(clientset, namespace, configMap)
}

// DeleteConfigMap deletes the ConfigMap by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func DeleteConfigMap(contextName string, namespace string, name string) error {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return err
	}

	return connector.DeleteConfigMap(clientset, namespace, name)
}

// GetReplicationControllers returns all the ReplicationController. If an empty namespace is given, returns all the ReplicationController
func GetReplicationControllers(contextName string, namespace string) ([]corev1.ReplicationController, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetReplicationControllers(contextName, namespace); results != nil {
		return results, nil
	}

	return connector.GetReplicationControllers(clientset, namespace)
}

// GetReplicationController returns the ReplicationController by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func GetReplicationController(contextName string, namespace string, name string) (*corev1.ReplicationController, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetReplicationControllers(contextName, namespace); results != nil {
		for _, replicationController := range results {
			if replicationController.Name == name {
				return &replicationController, nil
			}
		}
		return nil, nil
	}

	return connector.GetReplicationController(clientset, namespace, name)
}

// CreateReplicationController creates the ReplicationController with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func CreateReplicationController(contextName string, namespace string, replicationController *corev1.ReplicationController) (*corev1.ReplicationController, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.CreateReplicationController(clientset, namespace, replicationController)
}

// UpdateReplicationController updates the ReplicationController with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func UpdateReplicationController(contextName string, namespace string, replicationController *corev1.ReplicationController) (*corev1.ReplicationController, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.UpdateReplicationController(clientset, namespace, replicationController)
}

// DeleteReplicationController deletes the ReplicationController by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func DeleteReplicationController(contextName string, namespace string, name string) error {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return err
	}

	return connector.DeleteReplicationController(clientset, namespace, name)
}

// GetSecrets returns all the Secret. If an empty namespace is given, returns all the Secret
func GetSecrets(contextName string, namespace string) ([]corev1.Secret, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetSecrets(contextName, namespace); results != nil {
		return results, nil
	}

	return connector.GetSecrets(clientset, namespace)
}

// GetSecret returns the Secret by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func GetSecret(contextName string, namespace string, name string) (*corev1.Secret, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetSecrets(contextName, namespace); results != nil {
		for _, secret := range results {
			if secret.Name == name {
				return &secret, nil
			}
		}
		return nil, nil
	}

	return connector.GetSecret(clientset, namespace, name)
}

// CreateSecret creates the Secret with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func CreateSecret(contextName string, namespace string, secret *corev1.Secret) (*corev1.Secret, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.CreateSecret(clientset, namespace, secret)
}

// UpdateSecret updates the Secret with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func UpdateSecret(contextName string, namespace string, secret *corev1.Secret) (*corev1.Secret, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.UpdateSecret(clientset, namespace, secret)
}

// DeleteSecret deletes the Secret by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func DeleteSecret(contextName string, namespace string, name string) error {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return err
	}

	return connector.DeleteSecret(clientset, namespace, name)
}

// GetServiceAccounts returns all the ServiceAccount. If an empty namespace is given, returns all the ServiceAccount
func GetServiceAccounts(contextName string, namespace string) ([]corev1.ServiceAccount, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetServiceAccounts(contextName, namespace); results != nil {
		return results, nil
	}

	return connector.GetServiceAccounts(clientset, namespace)
}

// GetServiceAccount returns the ServiceAccount by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func GetServiceAccount(contextName string, namespace string, name string) (*corev1.ServiceAccount, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetServiceAccounts(contextName, namespace); results != nil {
		for _, serviceAccount := range results {
			if serviceAccount.Name == name {
				return &serviceAccount, nil
			}
		}
		return nil, nil
	}

	return connector.GetServiceAccount(clientset, namespace, name)
}

// CreateServiceAccount creates the ServiceAccount with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func CreateServiceAccount(contextName string, namespace string, serviceAccount *corev1.ServiceAccount) (*corev1.ServiceAccount, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.CreateServiceAccount(clientset, namespace, serviceAccount)
}

// UpdateServiceAccount updates the ServiceAccount with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func UpdateServiceAccount(contextName string, namespace string, serviceAccount *corev1.ServiceAccount) (*corev1.ServiceAccount, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.UpdateServiceAccount(clientset, namespace, serviceAccount)
}

// DeleteServiceAccount deletes the ServiceAccount by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func DeleteServiceAccount(contextName string, namespace string, name string) error {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return err
	}

	return connector.DeleteServiceAccount(clientset, namespace, name)
}

// GetDeployments returns all the Deployment. If an empty namespace is given, returns all the Deployment
func GetDeployments(contextName string, namespace string) ([]appsv1.Deployment, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetDeployments(contextName, namespace); results != nil {
		return results, nil
	}

	return connector.GetDeployments(clientset, namespace)
}

// GetDeployment returns the Deployment by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func GetDeployment(contextName string, namespace string, name string) (*appsv1.Deployment, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetDeployments(contextName, namespace); results != nil {
		for _, deployment := range results {
			if deployment.Name == name {
				return &deployment, nil
			}
		}
		return nil, nil
	}

	return connector.GetDeployment(clientset, namespace, name)
}

// CreateDeployment creates the Deployment with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func CreateDeployment(contextName string, namespace string, deployment *appsv1.Deployment) (*appsv1.Deployment, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.CreateDeployment(clientset, namespace, deployment)
}

// UpdateDeployment updates the Deployment with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func UpdateDeployment(contextName string, namespace string, deployment *appsv1.Deployment) (*appsv1.Deployment, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.UpdateDeployment(clientset, namespace, deployment)
}

// DeleteDeployment deletes the Deployment by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func DeleteDeployment(contextName string, namespace string, name string) error {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return err
	}

	return connector.DeleteDeployment(clientset, namespace, name)
}

// GetStatefulSets returns all the StatefulSet. If an empty namespace is given, returns all the StatefulSet
func GetStatefulSets(contextName string, namespace string) ([]appsv1.StatefulSet, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetStatefulSets(contextName, namespace); results != nil {
		return results, nil
	}

	return connector.GetStatefulSets(clientset, namespace)
}

// GetStatefulSet returns the StatefulSet by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func GetStatefulSet(contextName string, namespace string, name string) (*appsv1.StatefulSet, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetStatefulSets(contextName, namespace); results != nil {
		for _, statefulSet := range results {
			if statefulSet.Name == name {
				return &statefulSet, nil
			}
		}
		return nil, nil
	}

	return connector.GetStatefulSet(clientset, namespace, name)
}

// CreateStatefulSet creates the StatefulSet with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func CreateStatefulSet(contextName string, namespace string, statefulSet *appsv1.StatefulSet) (*appsv1.StatefulSet, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.CreateStatefulSet(clientset, namespace, statefulSet)
}

// UpdateStatefulSet updates the StatefulSet with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func UpdateStatefulSet(contextName string, namespace string, statefulSet *appsv1.StatefulSet) (*appsv1.StatefulSet, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.UpdateStatefulSet(clientset, namespace, statefulSet)
}

// DeleteStatefulSet deletes the StatefulSet by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func DeleteStatefulSet(contextName string, namespace string, name string) error {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return err
	}

	return connector.DeleteStatefulSet(clientset, namespace, name)
}

// GetDaemonSets returns all the DaemonSet. If an empty namespace is given, returns all the DaemonSet
func GetDaemonSets(contextName string, namespace string) ([]appsv1.DaemonSet, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetDaemonSets(contextName, namespace); results != nil {
		return results, nil
	}

	return connector.GetDaemonSets(clientset, namespace)
}

// GetDaemonSet returns the DaemonSet by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func GetDaemonSet(contextName string, namespace string, name string) (*appsv1.DaemonSet, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetDaemonSets(contextName, namespace); results != nil {
		for _, daemonSet := range results {
			if daemonSet.Name == name {
				return &daemonSet, nil
			}
		}
		return nil, nil
	}

	return connector.GetDaemonSet(clientset, namespace, name)
}

// CreateDaemonSet creates the DaemonSet with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func CreateDaemonSet(contextName string, namespace string, daemonSet *appsv1.DaemonSet) (*appsv1.DaemonSet, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.CreateDaemonSet(clientset, namespace, daemonSet)
}

// UpdateDaemonSet updates the DaemonSet with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func UpdateDaemonSet(contextName string, namespace string, daemonSet *appsv1.DaemonSet) (*appsv1.DaemonSet, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.UpdateDaemonSet(clientset, namespace, daemonSet)
}

// DeleteDaemonSet deletes the DaemonSet by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func DeleteDaemonSet(contextName string, namespace string, name string) error {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return err
	}

	return connector.DeleteDaemonSet(clientset, namespace, name)
}

// GetReplicaSets returns all the ReplicaSet. If an empty namespace is given, returns all the ReplicaSet
func GetReplicaSets(contextName string, namespace string) ([]appsv1.ReplicaSet, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetReplicaSets(contextName, namespace); results != nil {
		return results, nil
	}

	return connector.GetReplicaSets(clientset, namespace)
}

// GetReplicaSet returns the ReplicaSet by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func GetReplicaSet(contextName string, namespace string, name string) (*appsv1.ReplicaSet, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetReplicaSets(contextName, namespace); results != nil {
		for _, replicaSet := range results {
			if replicaSet.Name == name {
				return &replicaSet, nil
			}
		}
		return nil, nil
	}

	return connector.GetReplicaSet(clientset, namespace, name)
}

// CreateReplicaSet creates the ReplicaSet with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func CreateReplicaSet(contextName string, namespace string, replicaSet *appsv1.ReplicaSet) (*appsv1.ReplicaSet, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.CreateReplicaSet(clientset, namespace, replicaSet)
}

// UpdateReplicaSet updates the ReplicaSet with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func UpdateReplicaSet(contextName string, namespace string, replicaSet *appsv1.ReplicaSet) (*appsv1.ReplicaSet, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.UpdateReplicaSet(clientset, namespace, replicaSet)
}

// DeleteReplicaSet deletes the ReplicaSet by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func DeleteReplicaSet(contextName string, namespace string, name string) error {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return err
	}

	return connector.DeleteReplicaSet(clientset, namespace, name)
}

// GetNetworkPolicies returns all the NetworkPolicy. If an empty namespace is given, returns all the NetworkPolicy
func GetNetworkPolicies(contextName string, namespace string) ([]networkingv1.NetworkPolicy, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetNetworkPolicies(contextName, namespace); results != nil {
		return results, nil
	}

	return connector.GetNetworkPolicies(clientset, namespace)
}

// GetNetworkPolicy returns the NetworkPolicy by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func GetNetworkPolicy(contextName string, namespace string, name string) (*networkingv1.NetworkPolicy, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetNetworkPolicies(contextName, namespace); results != nil {
		for _, networkPolicy := range results {
			if networkPolicy.Name == name {
				return &networkPolicy, nil
			}
		}
		return nil, nil
	}

	return connector.GetNetworkPolicy(clientset, namespace, name)
}

// CreateNetworkPolicy creates the NetworkPolicy with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func CreateNetworkPolicy(contextName string, namespace string, networkPolicy *networkingv1.NetworkPolicy) (*networkingv1.NetworkPolicy, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.CreateNetworkPolicy(clientset, namespace, networkPolicy)
}

// UpdateNetworkPolicy updates the NetworkPolicy with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func UpdateNetworkPolicy(contextName string, namespace string, networkPolicy *networkingv1.NetworkPolicy) (*networkingv1.NetworkPolicy, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.UpdateNetworkPolicy(clientset, namespace, networkPolicy)
}

// DeleteNetworkPolicy deletes the NetworkPolicy by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func DeleteNetworkPolicy(contextName string, namespace string, name string) error {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return err
	}

	return connector.DeleteNetworkPolicy(clientset, namespace, name)
}

// GetRoles returns all the Role. If an empty namespace is given, returns all the Role
func GetRoles(contextName string, namespace string) ([]rbacv1.Role, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetRoles(contextName, namespace); results != nil {
		return results, nil
	}

	return connector.GetRoles(clientset, namespace)
}

// GetRole returns the Role by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func GetRole(contextName string, namespace string, name string) (*rbacv1.Role, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetRoles(contextName, namespace); results != nil {
		for _, role := range results {
			if role.Name == name {
				return &role, nil
			}
		}
		return nil, nil
	}

	return connector.GetRole(clientset, namespace, name)
}

// CreateRole creates the Role with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func CreateRole(contextName string, namespace string, role *rbacv1.Role) (*rbacv1.Role, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.CreateRole(clientset, namespace, role)
}

// UpdateRole updates the Role with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func UpdateRole(contextName string, namespace string, role *rbacv1.Role) (*rbacv1.Role, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.UpdateRole(clientset, namespace, role)
}

// DeleteRole deletes the Role by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func DeleteRole(contextName string, namespace string, name string) error {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return err
	}

	return connector.DeleteRole(clientset, namespace, name)
}

// GetRoleBindings returns all the RoleBinding. If an empty namespace is given, returns all the RoleBinding
func GetRoleBindings(contextName string, namespace string) ([]rbacv1.RoleBinding, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetRoleBindings(contextName, namespace); results != nil {
		return results, nil
	}

	return connector.GetRoleBindings(clientset, namespace)
}

// GetRoleBinding returns the RoleBinding by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func GetRoleBinding(contextName string, namespace string, name string) (*rbacv1.RoleBinding, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetRoleBindings(contextName, namespace); results != nil {
		for _, roleBinding := range results {
			if roleBinding.Name == name {
				return &roleBinding, nil
			}
		}
		return nil, nil
	}

	return connector.GetRoleBinding(clientset, namespace, name)
}

// CreateRoleBinding creates the RoleBinding with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func CreateRoleBinding(contextName string, namespace string, roleBinding *rbacv1.RoleBinding) (*rbacv1.RoleBinding, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.CreateRoleBinding(clientset, namespace, roleBinding)
}

// UpdateRoleBinding updates the RoleBinding with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func UpdateRoleBinding(contextName string, namespace string, roleBinding *rbacv1.RoleBinding) (*rbacv1.RoleBinding, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.UpdateRoleBinding(clientset, namespace, roleBinding)
}

// DeleteRoleBinding deletes the RoleBinding by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func DeleteRoleBinding(contextName string, namespace string, name string) error {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return err
	}

	return connector.DeleteRoleBinding(clientset, namespace, name)
}

// GetJobs returns all the Job. If an empty namespace is given, returns all the Job
func GetJobs(contextName string, namespace string) ([]batchv1.Job, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetJobs(contextName, namespace); results != nil {
		return results, nil
	}

	return connector.GetJobs(clientset, namespace)
}

// GetJob returns the Job by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func GetJob(contextName string, namespace string, name string) (*batchv1.Job, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetJobs(contextName, namespace); results != nil {
		for _, job := range results {
			if job.Name == name {
				return &job, nil
			}
		}
		return nil, nil
	}

	return connector.GetJob(clientset, namespace, name)
}

// CreateJob creates the Job with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func CreateJob(contextName string, namespace string, job *batchv1.Job) (*batchv1.Job, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.CreateJob(clientset, namespace, job)
}

// UpdateJob updates the Job with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func UpdateJob(contextName string, namespace string, job *batchv1.Job) (*batchv1.Job, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.UpdateJob(clientset, namespace, job)
}

// DeleteJob deletes the Job by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func DeleteJob(contextName string, namespace string, name string) error {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return err
	}

	return connector.DeleteJob(clientset, namespace, name)
}

// GetCronJobs returns all the CronJob. If an empty namespace is given, returns all the CronJob
func GetCronJobs(contextName string, namespace string) ([]batchv1beta1.CronJob, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetCronJobs(contextName, namespace); results != nil {
		return results, nil
	}

	return connector.GetCronJobs(clientset, namespace)
}

// GetCronJob returns the CronJob by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func GetCronJob(contextName string, namespace string, name string) (*batchv1beta1.CronJob, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetCronJobs(contextName, namespace); results != nil {
		for _, cronJob := range results {
			if cronJob.Name == name {
				return &cronJob, nil
			}
		}
		return nil, nil
	}

	return connector.GetCronJob(clientset, namespace, name)
}

// CreateCronJob creates the CronJob with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func CreateCronJob(contextName string, namespace string, cronJob *batchv1beta1.CronJob) (*batchv1beta1.CronJob, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.CreateCronJob(clientset, namespace, cronJob)
}

// UpdateCronJob updates the CronJob with the given model. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func UpdateCronJob(contextName string, namespace string, cronJob *batchv1beta1.CronJob) (*batchv1beta1.CronJob, error) {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return nil, err
	}

	return connector.UpdateCronJob(clientset, namespace, cronJob)
}

// DeleteCronJob deletes the CronJob by its name. An optional namespace can be given, if none is given
// the operation takes place in the default name space.
func DeleteCronJob(contextName string, namespace string, name string) error {

	clientset, err := context.GetClientset(contextName)
	if err != nil {
		return err
	}

	return connector.DeleteCronJob(clientset, namespace, name)
}
