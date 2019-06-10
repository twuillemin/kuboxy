package report

//go:generate go run gen/gen_list_converter.go

import (
	"fmt"
	"strings"
	"time"

	"github.com/twuillemin/kuboxy/pkg/provider"
	"k8s.io/apimachinery/pkg/types"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	metricsv1beta1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

// BuildReport builds a complete status report from a cluster status
func BuildReport(contextName string) (*ClusterStateReport, error) {

	start := time.Now()

	namespaces, err := provider.GetNamespaces(contextName)
	if err != nil {
		return nil, err
	}

	nodes, err := provider.GetNodes(contextName)
	if err != nil {
		return nil, err
	}

	nodeMetricses, err := provider.GetNodeMetricses(contextName)
	if err != nil {
		return nil, err
	}

	persistentVolumes, err := provider.GetPersistentVolumes(contextName)
	if err != nil {
		return nil, err
	}

	services := make(map[string][]corev1.Service)
	pods := make(map[string][]corev1.Pod)
	deployments := make(map[string][]appsv1.Deployment)
	podMetricses := make(map[string][]metricsv1beta1.PodMetrics)
	persistentVolumeClaims := make(map[string][]corev1.PersistentVolumeClaim)
	configMaps := make(map[string][]corev1.ConfigMap)
	secrets := make(map[string][]corev1.Secret)

	// Add the information for each namespace
	for _, namespace := range namespaces {
		namespaceServices, e := provider.GetServices(contextName, namespace.Name)
		if e != nil {
			return nil, e
		}
		namespacePods, e := provider.GetPods(contextName, namespace.Name)
		if e != nil {
			return nil, e
		}
		namespaceDeployments, e := provider.GetDeployments(contextName, namespace.Name)
		if e != nil {
			return nil, e
		}
		namespacePodMetricses, e := provider.GetPodMetricses(contextName, namespace.Name)
		if e != nil {
			return nil, e
		}
		namespacePersistentVolumeClaims, e := provider.GetPersistentVolumeClaims(contextName, namespace.Name)
		if e != nil {
			return nil, e
		}
		namespaceConfigMaps, e := provider.GetConfigMaps(contextName, namespace.Name)
		if e != nil {
			return nil, e
		}
		namespaceSecrets, e := provider.GetSecrets(contextName, namespace.Name)
		if e != nil {
			return nil, e
		}
		services[namespace.Name] = namespaceServices
		pods[namespace.Name] = namespacePods
		deployments[namespace.Name] = namespaceDeployments
		podMetricses[namespace.Name] = namespacePodMetricses
		persistentVolumeClaims[namespace.Name] = namespacePersistentVolumeClaims
		configMaps[namespace.Name] = namespaceConfigMaps
		secrets[namespace.Name] = namespaceSecrets
	}
	elapsed := time.Since(start)
	fmt.Printf("elapsed BuildReport/load data: %v\n", elapsed)

	return &ClusterStateReport{
		buildNamespacesReport(namespaces, services, pods, deployments),
		buildNodesReport(nodes, pods, nodeMetricses),
		buildServicesReport(services, pods),
		buildPodsReport(pods, podMetricses),
		buildPersistentVolumeReport(persistentVolumes),
		buildPersistentVolumeClaimReport(persistentVolumeClaims),
		buildConfigMapReport(configMaps),
		buildSecretReport(secrets),
		buildDeploymentsReport(deployments),
	}, nil
}

func buildNamespacesReport(
	namespaces []corev1.Namespace,
	services map[string][]corev1.Service,
	pods map[string][]corev1.Pod,
	deployments map[string][]appsv1.Deployment) []NamespaceReport {

	results := make([]NamespaceReport, 0, len(namespaces))

	// Add the information for each namespace
	for _, namespace := range namespaces {

		serviceIds := servicesToServiceIds(services[namespace.Name])
		podIds := podsToPodIds(pods[namespace.Name])
		deploymentsIds := deploymentsToDeploymentIds(deployments[namespace.Name])

		results = append(
			results,
			NamespaceReport{
				namespace.UID,
				namespace.Namespace,
				namespace.Name,
				namespace.Labels,
				serviceIds,
				podIds,
				deploymentsIds,
			})
	}

	return results
}

func buildNodesReport(
	nodes []corev1.Node,
	pods map[string][]corev1.Pod,
	nodeMetricses []metricsv1beta1.NodeMetrics) []NodeReport {

	results := make([]NodeReport, 0, len(pods))

	// Add the information for each node
	for _, node := range nodes {

		conditions := make([]string, 0, len(node.Status.Conditions))
		for _, condition := range node.Status.Conditions {
			if strings.Compare(string(condition.Status), "False") != 0 {
				conditions = append(conditions, string(condition.Type))
			}
		}

		podIds := make([]types.UID, 0)
		for _, namespacePods := range pods {
			for _, pod := range namespacePods {
				if pod.Spec.NodeName == node.Name {
					podIds = append(podIds, pod.UID)
				}
			}
		}

		allowedCPUCores, _ := node.Status.Capacity.Cpu().AsInt64()
		allowedRAM, _ := node.Status.Capacity.Memory().AsInt64()

		// Use pointer instead of numeric values for keeping the result as nil if no metrics is found
		var usageCPUCores *float64
		var usageRAM *int

		for _, nodeMetrics := range nodeMetricses {
			if nodeMetrics.Name == node.Name {
				cpuUsage := float64(nodeMetrics.Usage.Cpu().MilliValue()) / 1000
				usageCPUCores = &cpuUsage

				if i, ok := nodeMetrics.Usage.Memory().AsInt64(); ok {
					ramUsage := int(i)
					usageRAM = &(ramUsage)
				}
			}
		}

		results = append(
			results,
			NodeReport{
				node.UID,
				node.Namespace,
				node.Name,
				node.Labels,
				strings.Join(conditions, ", "),
				int(allowedCPUCores),
				int(allowedRAM),
				usageCPUCores,
				usageRAM,
				podIds,
			})
	}

	return results
}

func buildServicesReport(
	services map[string][]corev1.Service,
	pods map[string][]corev1.Pod) []ServiceReport {

	results := make([]ServiceReport, 0)

	// Add the information for each service
	for namespace, namespaceServices := range services {

		for _, service := range namespaceServices {

			servicePods := searchPodsForSelector(pods, namespace, service.Spec.Selector)
			podIds := make([]types.UID, 0, len(servicePods))
			for _, pod := range servicePods {
				podIds = append(podIds, pod.UID)
			}

			ports := make([]ServicePortReport, 0, len(service.Spec.Ports))
			for _, port := range service.Spec.Ports {

				// Use pointers instead of string/int values for keeping the result as nil if no information is found
				var portName *string
				var nodePort *int32

				if port.TargetPort.Type == intstr.String {
					portName = &(port.TargetPort.StrVal)
				}

				if port.NodePort != 0 {
					nodePort = &(port.NodePort)
				}

				ports = append(
					ports,
					ServicePortReport{
						string(port.Protocol),
						port.Port,
						portName,
						nodePort,
					})
			}

			results = append(
				results,
				ServiceReport{
					service.UID,
					service.Namespace,
					service.Name,
					service.Spec.ClusterIP,
					service.Spec.ExternalIPs,
					ports,
					service.Labels,
					service.Spec.Selector,
					podIds,
				})
		}
	}

	return results
}

func buildPodsReport(
	pods map[string][]corev1.Pod,
	podMetricses map[string][]metricsv1beta1.PodMetrics) []PodReport {

	results := make([]PodReport, 0)

	// Add the information for each pod
	for namespace, namespacePods := range pods {

		namespacePodMetricses := podMetricses[namespace]

		for _, pod := range namespacePods {

			containers := make([]ContainerReport, 0, len(pod.Spec.Containers))

			// Make a report for the containers
			for _, container := range pod.Spec.Containers {

				ports := make([]PodPortReport, 0, len(container.Ports))

				for _, port := range container.Ports {

					var portName *string
					if len(port.Name) > 0 {
						portName = &(port.Name)
					}

					ports = append(
						ports,
						PodPortReport{
							string(port.Protocol),
							port.ContainerPort,
							portName,
						})
				}

				id := "unknown id"
				ready := false

				for _, containerStatus := range pod.Status.ContainerStatuses {
					if container.Name == containerStatus.Name {
						id = containerStatus.ContainerID
						ready = containerStatus.Ready
					}
				}

				containers = append(
					containers,
					ContainerReport{
						id,
						pod.Namespace,
						container.Name,
						container.Image,
						ports,
						ready,
					})
			}

			conditions := make([]string, 0, len(pod.Status.Conditions))
			for _, condition := range pod.Status.Conditions {
				if strings.Compare(string(condition.Status), "False") != 0 {
					conditions = append(conditions, string(condition.Type))
				}
			}

			// Use pointers instead of float/int values for keeping the result as nil if no information is found
			var usageCPUCores *float64
			var usageRAM *int

			// The metrics are given by container, so sum them
			for _, podMetrics := range namespacePodMetricses {

				if podMetrics.Name != pod.Name {
					totalCPU := 0.0
					totalRAM := 0
					for _, container := range podMetrics.Containers {
						totalCPU += float64(container.Usage.Cpu().MilliValue()) / 1000
						if i, ok := container.Usage.Memory().AsInt64(); ok {
							totalRAM += int(i)
						}
					}
					usageCPUCores = &totalCPU
					usageRAM = &(totalRAM)
				}
			}

			results = append(
				results,
				PodReport{
					pod.UID,
					pod.Namespace,
					pod.Name,
					pod.Labels,
					strings.Join(conditions, ", "),
					pod.Status.PodIP,
					usageCPUCores,
					usageRAM,
					containers,
					buildVolumes(pod.Spec.Volumes),
				})
		}
	}

	return results
}

func buildPersistentVolumeReport(persistentVolumes []corev1.PersistentVolume) []PersistentVolumeReport {

	results := make([]PersistentVolumeReport, 0, len(persistentVolumes))

	// Add the information for each pvc
	for _, persistentVolume := range persistentVolumes {

		// Use a pointer instead of an int to preserve missing information
		var storageCapacity *int

		if capacity, capacityOk := persistentVolume.Spec.Capacity[corev1.ResourceStorage]; capacityOk {
			if i, ok := capacity.AsInt64(); ok {
				value := int(i)
				storageCapacity = &value
			}
		}

		results = append(
			results,
			PersistentVolumeReport{
				persistentVolume.UID,
				persistentVolume.Namespace,
				persistentVolume.Name,
				getStorageTypeFromPersistentVolume(persistentVolume),
				storageCapacity,
			})
	}

	return results
}

func buildPersistentVolumeClaimReport(persistentVolumeClaims map[string][]corev1.PersistentVolumeClaim) []PersistentVolumeClaimReport {

	results := make([]PersistentVolumeClaimReport, 0)

	// Add the information for each pvc
	for _, namespacePersistentVolumeClaims := range persistentVolumeClaims {

		for _, persistentVolumeClaim := range namespacePersistentVolumeClaims {

			results = append(
				results,
				PersistentVolumeClaimReport{
					persistentVolumeClaim.UID,
					persistentVolumeClaim.Namespace,
					persistentVolumeClaim.Name,
					persistentVolumeClaim.Spec.String(),
				})
		}
	}

	return results
}

func buildConfigMapReport(configMaps map[string][]corev1.ConfigMap) []ConfigMapReport {

	results := make([]ConfigMapReport, 0)

	// Add the information for each config map
	for _, namespaceConfigMaps := range configMaps {
		for _, configMap := range namespaceConfigMaps {

			results = append(
				results,
				ConfigMapReport{
					configMap.UID,
					configMap.Namespace,
					configMap.Name,
					configMap.Data,
				})
		}
	}

	return results
}

func buildSecretReport(secrets map[string][]corev1.Secret) []SecretReport {

	results := make([]SecretReport, 0)

	// Add the information for each secret
	for _, namespaceSecrets := range secrets {
		for _, secret := range namespaceSecrets {

			results = append(
				results,
				SecretReport{
					secret.UID,
					secret.Namespace,
					secret.Name,
					secret.Data,
				})
		}
	}

	return results
}

func buildDeploymentsReport(deployments map[string][]appsv1.Deployment) []DeploymentReport {

	results := make([]DeploymentReport, 0)

	// Add the information for each deployment
	for _, namespaceDeployments := range deployments {
		for _, deployment := range namespaceDeployments {

			results = append(
				results,
				DeploymentReport{
					deployment.UID,
					deployment.Namespace,
					deployment.Name,
					deployment.Labels,
					deployment.Spec.Selector.MatchLabels,
					buildVolumes(deployment.Spec.Template.Spec.Volumes),
				})
		}
	}

	return results
}

func buildVolumes(volumes []corev1.Volume) []VolumeReport {

	results := make([]VolumeReport, 0, len(volumes))

	// Add the information for each deployment
	for _, volume := range volumes {

		storageType := getStorageTypeFromVolume(volume)

		// Use a pointer instead of string to preserve missing information
		var pvcName *string

		if storageType == PersistentVolumeClaim {
			name := volume.PersistentVolumeClaim.ClaimName
			pvcName = &name
		}

		results = append(
			results,
			VolumeReport{
				volume.Name,
				getStorageTypeFromVolume(volume),
				pvcName,
			})
	}

	return results
}

func searchPodsForSelector(
	pods map[string][]corev1.Pod,
	namespace string,
	selector map[string]string) []corev1.Pod {

	results := make([]corev1.Pod, 0)

	namespacePods := pods[namespace]
	if namespacePods == nil {
		return results
	}

	for _, pod := range namespacePods {
		podLabels := pod.Labels
		match := false
		for k, v := range selector {
			if podLabels[k] == v {
				match = true
				break
			}
		}

		if match {
			results = append(results, pod)
		}
	}

	return results
}
