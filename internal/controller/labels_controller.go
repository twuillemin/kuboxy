// Package controller regroups all the HTTP controllers of the application
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_labels_controller.go at 2019-02-25 19:47:13.7623547 +0200 EET m=+0.000943943
package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/twuillemin/kuboxy/pkg/provider"
)

func registerLabelsController(e *echo.Echo) {

	// Declare the routes
	e.GET("api/v1/labels/:contextName/:namespace", getLabels)
}

// getLabels generates a JSON representation of all the labels and their values in the given configuration
// @Summary Get all the labels and their values
// @Description Get all the labels and their values
// @ID get-labels
// @Tags Labels
// @Produce application/json
// @Param contextName path string true "the name of the context"
// @Param namespace path string true "the name of the namespace"
// @Success 200 {object} MapOfStrings
// @Failure 404 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /api/v1/labels/{contextName}/{namespace} [get]
func getLabels(e echo.Context) error {

	queryContextName := e.Param("contextName")
	queryNamespace := e.Param("namespace")

	labels := make(map[string]map[string]bool)

	// Get the state of the cluster
	namespaces, err := provider.GetNamespaces(queryContextName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, namespace := range namespaces {
		for k, v := range namespace.ObjectMeta.Labels {
			values, ok := labels[k]
			if !ok {
				values = make(map[string]bool)
			}
			values[v] = true
			labels[k] = values
		}
	}

	// Get the state of the cluster
	nodes, err := provider.GetNodes(queryContextName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, node := range nodes {
		for k, v := range node.ObjectMeta.Labels {
			values, ok := labels[k]
			if !ok {
				values = make(map[string]bool)
			}
			values[v] = true
			labels[k] = values
		}
	}

	// Get the state of the cluster
	persistentVolumes, err := provider.GetPersistentVolumes(queryContextName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, persistentVolume := range persistentVolumes {
		for k, v := range persistentVolume.ObjectMeta.Labels {
			values, ok := labels[k]
			if !ok {
				values = make(map[string]bool)
			}
			values[v] = true
			labels[k] = values
		}
	}

	// Get the state of the cluster
	clusterRoles, err := provider.GetClusterRoles(queryContextName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, clusterRole := range clusterRoles {
		for k, v := range clusterRole.ObjectMeta.Labels {
			values, ok := labels[k]
			if !ok {
				values = make(map[string]bool)
			}
			values[v] = true
			labels[k] = values
		}
	}

	// Get the state of the cluster
	clusterRoleBindings, err := provider.GetClusterRoleBindings(queryContextName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, clusterRoleBinding := range clusterRoleBindings {
		for k, v := range clusterRoleBinding.ObjectMeta.Labels {
			values, ok := labels[k]
			if !ok {
				values = make(map[string]bool)
			}
			values[v] = true
			labels[k] = values
		}
	}

	// Get the state of the cluster
	storageClasses, err := provider.GetStorageClasses(queryContextName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, storageClass := range storageClasses {
		for k, v := range storageClass.ObjectMeta.Labels {
			values, ok := labels[k]
			if !ok {
				values = make(map[string]bool)
			}
			values[v] = true
			labels[k] = values
		}
	}

	// Get the state of the cluster
	services, err := provider.GetServices(queryContextName, queryNamespace)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, service := range services {
		for k, v := range service.ObjectMeta.Labels {
			values, ok := labels[k]
			if !ok {
				values = make(map[string]bool)
			}
			values[v] = true
			labels[k] = values
		}
	}

	// Get the state of the cluster
	pods, err := provider.GetPods(queryContextName, queryNamespace)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, pod := range pods {
		for k, v := range pod.ObjectMeta.Labels {
			values, ok := labels[k]
			if !ok {
				values = make(map[string]bool)
			}
			values[v] = true
			labels[k] = values
		}
	}

	// Get the state of the cluster
	persistentVolumeClaims, err := provider.GetPersistentVolumeClaims(queryContextName, queryNamespace)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, persistentVolumeClaim := range persistentVolumeClaims {
		for k, v := range persistentVolumeClaim.ObjectMeta.Labels {
			values, ok := labels[k]
			if !ok {
				values = make(map[string]bool)
			}
			values[v] = true
			labels[k] = values
		}
	}

	// Get the state of the cluster
	configMaps, err := provider.GetConfigMaps(queryContextName, queryNamespace)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, configMap := range configMaps {
		for k, v := range configMap.ObjectMeta.Labels {
			values, ok := labels[k]
			if !ok {
				values = make(map[string]bool)
			}
			values[v] = true
			labels[k] = values
		}
	}

	// Get the state of the cluster
	ReplicationControllers, err := provider.GetReplicationControllers(queryContextName, queryNamespace)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, replicationController := range ReplicationControllers {
		for k, v := range replicationController.ObjectMeta.Labels {
			values, ok := labels[k]
			if !ok {
				values = make(map[string]bool)
			}
			values[v] = true
			labels[k] = values
		}
	}

	// Get the state of the cluster
	secrets, err := provider.GetSecrets(queryContextName, queryNamespace)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, secret := range secrets {
		for k, v := range secret.ObjectMeta.Labels {
			values, ok := labels[k]
			if !ok {
				values = make(map[string]bool)
			}
			values[v] = true
			labels[k] = values
		}
	}

	// Get the state of the cluster
	serviceAccounts, err := provider.GetServiceAccounts(queryContextName, queryNamespace)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, serviceAccount := range serviceAccounts {
		for k, v := range serviceAccount.ObjectMeta.Labels {
			values, ok := labels[k]
			if !ok {
				values = make(map[string]bool)
			}
			values[v] = true
			labels[k] = values
		}
	}

	// Get the state of the cluster
	deployments, err := provider.GetDeployments(queryContextName, queryNamespace)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, deployment := range deployments {
		for k, v := range deployment.ObjectMeta.Labels {
			values, ok := labels[k]
			if !ok {
				values = make(map[string]bool)
			}
			values[v] = true
			labels[k] = values
		}
	}

	// Get the state of the cluster
	statefulSets, err := provider.GetStatefulSets(queryContextName, queryNamespace)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, statefulSet := range statefulSets {
		for k, v := range statefulSet.ObjectMeta.Labels {
			values, ok := labels[k]
			if !ok {
				values = make(map[string]bool)
			}
			values[v] = true
			labels[k] = values
		}
	}

	// Get the state of the cluster
	daemonSets, err := provider.GetDaemonSets(queryContextName, queryNamespace)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, daemonSet := range daemonSets {
		for k, v := range daemonSet.ObjectMeta.Labels {
			values, ok := labels[k]
			if !ok {
				values = make(map[string]bool)
			}
			values[v] = true
			labels[k] = values
		}
	}

	// Get the state of the cluster
	replicaSets, err := provider.GetReplicaSets(queryContextName, queryNamespace)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, replicaSet := range replicaSets {
		for k, v := range replicaSet.ObjectMeta.Labels {
			values, ok := labels[k]
			if !ok {
				values = make(map[string]bool)
			}
			values[v] = true
			labels[k] = values
		}
	}

	// Get the state of the cluster
	networkPolicies, err := provider.GetNetworkPolicies(queryContextName, queryNamespace)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, networkPolicy := range networkPolicies {
		for k, v := range networkPolicy.ObjectMeta.Labels {
			values, ok := labels[k]
			if !ok {
				values = make(map[string]bool)
			}
			values[v] = true
			labels[k] = values
		}
	}

	// Get the state of the cluster
	roles, err := provider.GetRoles(queryContextName, queryNamespace)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, role := range roles {
		for k, v := range role.ObjectMeta.Labels {
			values, ok := labels[k]
			if !ok {
				values = make(map[string]bool)
			}
			values[v] = true
			labels[k] = values
		}
	}

	// Get the state of the cluster
	roleBindings, err := provider.GetRoleBindings(queryContextName, queryNamespace)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, roleBinding := range roleBindings {
		for k, v := range roleBinding.ObjectMeta.Labels {
			values, ok := labels[k]
			if !ok {
				values = make(map[string]bool)
			}
			values[v] = true
			labels[k] = values
		}
	}

	// Get the state of the cluster
	jobs, err := provider.GetJobs(queryContextName, queryNamespace)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, job := range jobs {
		for k, v := range job.ObjectMeta.Labels {
			values, ok := labels[k]
			if !ok {
				values = make(map[string]bool)
			}
			values[v] = true
			labels[k] = values
		}
	}

	// Get the state of the cluster
	cronJobs, err := provider.GetCronJobs(queryContextName, queryNamespace)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	for _, cronJob := range cronJobs {
		for k, v := range cronJob.ObjectMeta.Labels {
			values, ok := labels[k]
			if !ok {
				values = make(map[string]bool)
			}
			values[v] = true
			labels[k] = values
		}
	}

	results := make(map[string][]string)
	for tagName, tagValues := range labels {

		values := make([]string, 0, len(tagValues))
		for tagValue := range tagValues {
			values = append(values, tagValue)
		}

		results[tagName] = values
	}

	return e.JSON(http.StatusOK, results)
}