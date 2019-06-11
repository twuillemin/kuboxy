// Package connector regroups all the basic CRUD functions to access the objects in the cluster
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_connector_cluster.go at 2019-02-17 23:50:39.514439735 +0200 EET m=+0.001370208
package connector

import (
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GetNamespaces returns all the Namespace.
func GetNamespaces(clientset *kubernetes.Clientset) ([]corev1.Namespace, error) {

	client := clientset.CoreV1().Namespaces()
	namespace, err := client.List(metav1.ListOptions{})
	if err != nil {
		return []corev1.Namespace{}, err
	}
	return namespace.Items, nil
}

// GetNamespace returns the Namespace by its name.
func GetNamespace(clientset *kubernetes.Clientset, name string) (*corev1.Namespace, error) {

	client := clientset.CoreV1().Namespaces()
	return client.Get(name, metav1.GetOptions{})
}

// CreateNamespace creates the Namespace with the given model.
func CreateNamespace(clientset *kubernetes.Clientset, namespace *corev1.Namespace) (*corev1.Namespace, error) {

	client := clientset.CoreV1().Namespaces()
	return client.Create(namespace)
}

// UpdateNamespace updates the Namespace with the given model.
func UpdateNamespace(clientset *kubernetes.Clientset, namespace *corev1.Namespace) (*corev1.Namespace, error) {

	client := clientset.CoreV1().Namespaces()
	return client.Update(namespace)
}

// DeleteNamespace deletes the Namespace by its name.
func DeleteNamespace(clientset *kubernetes.Clientset, name string) error {

	client := clientset.CoreV1().Namespaces()

	deletePolicy := metav1.DeletePropagationForeground

	return client.Delete(name, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
}

// GetNodes returns all the Node.
func GetNodes(clientset *kubernetes.Clientset) ([]corev1.Node, error) {

	client := clientset.CoreV1().Nodes()
	node, err := client.List(metav1.ListOptions{})
	if err != nil {
		return []corev1.Node{}, err
	}
	return node.Items, nil
}

// GetNode returns the Node by its name.
func GetNode(clientset *kubernetes.Clientset, name string) (*corev1.Node, error) {

	client := clientset.CoreV1().Nodes()
	return client.Get(name, metav1.GetOptions{})
}

// CreateNode creates the Node with the given model.
func CreateNode(clientset *kubernetes.Clientset, node *corev1.Node) (*corev1.Node, error) {

	client := clientset.CoreV1().Nodes()
	return client.Create(node)
}

// UpdateNode updates the Node with the given model.
func UpdateNode(clientset *kubernetes.Clientset, node *corev1.Node) (*corev1.Node, error) {

	client := clientset.CoreV1().Nodes()
	return client.Update(node)
}

// DeleteNode deletes the Node by its name.
func DeleteNode(clientset *kubernetes.Clientset, name string) error {

	client := clientset.CoreV1().Nodes()

	deletePolicy := metav1.DeletePropagationForeground

	return client.Delete(name, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
}

// GetPersistentVolumes returns all the PersistentVolume.
func GetPersistentVolumes(clientset *kubernetes.Clientset) ([]corev1.PersistentVolume, error) {

	client := clientset.CoreV1().PersistentVolumes()
	persistentVolume, err := client.List(metav1.ListOptions{})
	if err != nil {
		return []corev1.PersistentVolume{}, err
	}
	return persistentVolume.Items, nil
}

// GetPersistentVolume returns the PersistentVolume by its name.
func GetPersistentVolume(clientset *kubernetes.Clientset, name string) (*corev1.PersistentVolume, error) {

	client := clientset.CoreV1().PersistentVolumes()
	return client.Get(name, metav1.GetOptions{})
}

// CreatePersistentVolume creates the PersistentVolume with the given model.
func CreatePersistentVolume(clientset *kubernetes.Clientset, persistentVolume *corev1.PersistentVolume) (*corev1.PersistentVolume, error) {

	client := clientset.CoreV1().PersistentVolumes()
	return client.Create(persistentVolume)
}

// UpdatePersistentVolume updates the PersistentVolume with the given model.
func UpdatePersistentVolume(clientset *kubernetes.Clientset, persistentVolume *corev1.PersistentVolume) (*corev1.PersistentVolume, error) {

	client := clientset.CoreV1().PersistentVolumes()
	return client.Update(persistentVolume)
}

// DeletePersistentVolume deletes the PersistentVolume by its name.
func DeletePersistentVolume(clientset *kubernetes.Clientset, name string) error {

	client := clientset.CoreV1().PersistentVolumes()

	deletePolicy := metav1.DeletePropagationForeground

	return client.Delete(name, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
}

// GetClusterRoles returns all the ClusterRole.
func GetClusterRoles(clientset *kubernetes.Clientset) ([]rbacv1.ClusterRole, error) {

	client := clientset.RbacV1().ClusterRoles()
	clusterRole, err := client.List(metav1.ListOptions{})
	if err != nil {
		return []rbacv1.ClusterRole{}, err
	}
	return clusterRole.Items, nil
}

// GetClusterRole returns the ClusterRole by its name.
func GetClusterRole(clientset *kubernetes.Clientset, name string) (*rbacv1.ClusterRole, error) {

	client := clientset.RbacV1().ClusterRoles()
	return client.Get(name, metav1.GetOptions{})
}

// CreateClusterRole creates the ClusterRole with the given model.
func CreateClusterRole(clientset *kubernetes.Clientset, clusterRole *rbacv1.ClusterRole) (*rbacv1.ClusterRole, error) {

	client := clientset.RbacV1().ClusterRoles()
	return client.Create(clusterRole)
}

// UpdateClusterRole updates the ClusterRole with the given model.
func UpdateClusterRole(clientset *kubernetes.Clientset, clusterRole *rbacv1.ClusterRole) (*rbacv1.ClusterRole, error) {

	client := clientset.RbacV1().ClusterRoles()
	return client.Update(clusterRole)
}

// DeleteClusterRole deletes the ClusterRole by its name.
func DeleteClusterRole(clientset *kubernetes.Clientset, name string) error {

	client := clientset.RbacV1().ClusterRoles()

	deletePolicy := metav1.DeletePropagationForeground

	return client.Delete(name, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
}

// GetClusterRoleBindings returns all the ClusterRoleBinding.
func GetClusterRoleBindings(clientset *kubernetes.Clientset) ([]rbacv1.ClusterRoleBinding, error) {

	client := clientset.RbacV1().ClusterRoleBindings()
	clusterRoleBinding, err := client.List(metav1.ListOptions{})
	if err != nil {
		return []rbacv1.ClusterRoleBinding{}, err
	}
	return clusterRoleBinding.Items, nil
}

// GetClusterRoleBinding returns the ClusterRoleBinding by its name.
func GetClusterRoleBinding(clientset *kubernetes.Clientset, name string) (*rbacv1.ClusterRoleBinding, error) {

	client := clientset.RbacV1().ClusterRoleBindings()
	return client.Get(name, metav1.GetOptions{})
}

// CreateClusterRoleBinding creates the ClusterRoleBinding with the given model.
func CreateClusterRoleBinding(clientset *kubernetes.Clientset, clusterRoleBinding *rbacv1.ClusterRoleBinding) (*rbacv1.ClusterRoleBinding, error) {

	client := clientset.RbacV1().ClusterRoleBindings()
	return client.Create(clusterRoleBinding)
}

// UpdateClusterRoleBinding updates the ClusterRoleBinding with the given model.
func UpdateClusterRoleBinding(clientset *kubernetes.Clientset, clusterRoleBinding *rbacv1.ClusterRoleBinding) (*rbacv1.ClusterRoleBinding, error) {

	client := clientset.RbacV1().ClusterRoleBindings()
	return client.Update(clusterRoleBinding)
}

// DeleteClusterRoleBinding deletes the ClusterRoleBinding by its name.
func DeleteClusterRoleBinding(clientset *kubernetes.Clientset, name string) error {

	client := clientset.RbacV1().ClusterRoleBindings()

	deletePolicy := metav1.DeletePropagationForeground

	return client.Delete(name, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
}

// GetStorageClasses returns all the StorageClass.
func GetStorageClasses(clientset *kubernetes.Clientset) ([]storagev1.StorageClass, error) {

	client := clientset.StorageV1().StorageClasses()
	storageClass, err := client.List(metav1.ListOptions{})
	if err != nil {
		return []storagev1.StorageClass{}, err
	}
	return storageClass.Items, nil
}

// GetStorageClass returns the StorageClass by its name.
func GetStorageClass(clientset *kubernetes.Clientset, name string) (*storagev1.StorageClass, error) {

	client := clientset.StorageV1().StorageClasses()
	return client.Get(name, metav1.GetOptions{})
}

// CreateStorageClass creates the StorageClass with the given model.
func CreateStorageClass(clientset *kubernetes.Clientset, storageClass *storagev1.StorageClass) (*storagev1.StorageClass, error) {

	client := clientset.StorageV1().StorageClasses()
	return client.Create(storageClass)
}

// UpdateStorageClass updates the StorageClass with the given model.
func UpdateStorageClass(clientset *kubernetes.Clientset, storageClass *storagev1.StorageClass) (*storagev1.StorageClass, error) {

	client := clientset.StorageV1().StorageClasses()
	return client.Update(storageClass)
}

// DeleteStorageClass deletes the StorageClass by its name.
func DeleteStorageClass(clientset *kubernetes.Clientset, name string) error {

	client := clientset.StorageV1().StorageClasses()

	deletePolicy := metav1.DeletePropagationForeground

	return client.Delete(name, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
}