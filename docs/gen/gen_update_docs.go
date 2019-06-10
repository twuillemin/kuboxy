// The following directive is necessary to make the package coherent:

// + build ignore

// This program updates the docs.go file with the definition from external packages. It can be invoked by running
// go generate

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	storagev1 "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	metricsv1beta1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

func main() {

	newLines := generateAdditionalDocumentation()

	originalLines := readDocs()

	file, err := os.Create("docs.go")
	dieIfError(err)
	defer func() {
		dieIfError(file.Close())
	}()

	w := bufio.NewWriter(file)
	for _, originalLine := range originalLines {
		_, err = fmt.Fprintln(w, originalLine)
		dieIfError(err)
		if originalLine == "    \"definitions\": {" {
			for _, newLine := range newLines {
				_, err = fmt.Fprintf(w, "        %s", newLine)
				dieIfError(err)
			}
		}
	}

	dieIfError(w.Flush())
}

func readDocs() []string {

	file, err := os.Open("docs.go")
	dieIfError(err)
	defer func() {
		dieIfError(file.Close())
	}()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	dieIfError(scanner.Err())

	return lines
}

func dieIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func generateAdditionalDocumentation() []string {
	results := make([]string, 0, 50)
	missings := make(map[string]bool)

	// Head objects (having a direct endpoint)

	results = addNewObject(results, missings, appsv1.Deployment{}, "appsv1", true)
	results = addNewObject(results, missings, appsv1.StatefulSet{}, "appsv1", true)
	results = addNewObject(results, missings, appsv1.DaemonSet{}, "appsv1", true)
	results = addNewObject(results, missings, appsv1.ReplicaSet{}, "appsv1", true)

	results = addNewObject(results, missings, batchv1.Job{}, "batchv1", true)

	results = addNewObject(results, missings, batchv1beta1.CronJob{}, "batchv1beta1", true)

	results = addNewObject(results, missings, corev1.ConfigMap{}, "corev1", true)
	results = addNewObject(results, missings, corev1.Namespace{}, "corev1", true)
	results = addNewObject(results, missings, corev1.Node{}, "corev1", true)
	results = addNewObject(results, missings, corev1.PersistentVolume{}, "corev1", true)
	results = addNewObject(results, missings, corev1.PersistentVolumeClaim{}, "corev1", true)
	results = addNewObject(results, missings, corev1.Pod{}, "corev1", true)
	results = addNewObject(results, missings, corev1.ReplicationController{}, "corev1", true)
	results = addNewObject(results, missings, corev1.Secret{}, "corev1", true)
	results = addNewObject(results, missings, corev1.Service{}, "corev1", true)
	results = addNewObject(results, missings, corev1.ServiceAccount{}, "corev1", true)

	results = addNewObject(results, missings, metricsv1beta1.PodMetrics{}, "metricsv1beta1", true)
	results = addNewObject(results, missings, metricsv1beta1.NodeMetrics{}, "metricsv1beta1", true)

	results = addNewObject(results, missings, networkingv1.NetworkPolicy{}, "networkingv1", true)

	results = addNewObject(results, missings, rbacv1.ClusterRole{}, "rbacv1", true)
	results = addNewObject(results, missings, rbacv1.ClusterRoleBinding{}, "rbacv1", true)
	results = addNewObject(results, missings, rbacv1.Role{}, "rbacv1", true)
	results = addNewObject(results, missings, rbacv1.RoleBinding{}, "rbacv1", true)

	results = addNewObject(results, missings, storagev1.StorageClass{}, "storagev1", true)

	// Non head objects (not having a direct endpoint)

	results = addNewObject(results, missings, appsv1.DeploymentCondition{}, "appsv1", false)
	results = addNewObject(results, missings, appsv1.DeploymentSpec{}, "appsv1", false)
	results = addNewObject(results, missings, appsv1.DeploymentStatus{}, "appsv1", false)
	results = addNewObject(results, missings, appsv1.DeploymentStrategy{}, "appsv1", false)
	results = addNewObject(results, missings, appsv1.RollingUpdateDeployment{}, "appsv1", false)
	results = addNewObject(results, missings, appsv1.DaemonSetSpec{}, "appsv1", false)
	results = addNewObject(results, missings, appsv1.DaemonSetStatus{}, "appsv1", false)
	results = addNewObject(results, missings, appsv1.DaemonSetCondition{}, "appsv1", false)
	results = addNewObject(results, missings, appsv1.DaemonSetUpdateStrategy{}, "appsv1", false)
	results = addNewObject(results, missings, appsv1.ReplicaSetSpec{}, "appsv1", false)
	results = addNewObject(results, missings, appsv1.ReplicaSetStatus{}, "appsv1", false)
	results = addNewObject(results, missings, appsv1.ReplicaSetCondition{}, "appsv1", false)
	results = addNewObject(results, missings, appsv1.StatefulSetStatus{}, "appsv1", false)
	results = addNewObject(results, missings, appsv1.StatefulSetSpec{}, "appsv1", false)
	results = addNewObject(results, missings, appsv1.StatefulSetCondition{}, "appsv1", false)
	results = addNewObject(results, missings, appsv1.StatefulSetUpdateStrategy{}, "appsv1", false)
	results = addNewObject(results, missings, appsv1.RollingUpdateDaemonSet{}, "appsv1", false)
	results = addNewObject(results, missings, appsv1.RollingUpdateStatefulSetStrategy{}, "appsv1", false)

	results = addNewObject(results, missings, batchv1.JobSpec{}, "batchv1", false)
	results = addNewObject(results, missings, batchv1.JobStatus{}, "batchv1", false)
	results = addNewObject(results, missings, batchv1.JobCondition{}, "batchv1", false)

	results = addNewObject(results, missings, batchv1beta1.CronJobSpec{}, "batchv1beta1", false)
	results = addNewObject(results, missings, batchv1beta1.CronJobStatus{}, "batchv1beta1", false)
	results = addNewObject(results, missings, batchv1beta1.JobTemplateSpec{}, "batchv1beta1", false)

	results = addNewObject(results, missings, corev1.Affinity{}, "corev1", false)
	results = addNewObject(results, missings, corev1.AttachedVolume{}, "corev1", false)
	results = addNewObject(results, missings, corev1.Capabilities{}, "corev1", false)
	results = addNewObject(results, missings, corev1.ClientIPConfig{}, "corev1", false)
	results = addNewObject(results, missings, corev1.ConfigMapEnvSource{}, "corev1", false)
	results = addNewObject(results, missings, corev1.ConfigMapKeySelector{}, "corev1", false)
	results = addNewObject(results, missings, corev1.ConfigMapNodeConfigSource{}, "corev1", false)
	results = addNewObject(results, missings, corev1.Container{}, "corev1", false)
	results = addNewObject(results, missings, corev1.ContainerImage{}, "corev1", false)
	results = addNewObject(results, missings, corev1.ContainerPort{}, "corev1", false)
	results = addNewObject(results, missings, corev1.ContainerState{}, "corev1", false)
	results = addNewObject(results, missings, corev1.ContainerStateRunning{}, "corev1", false)
	results = addNewObject(results, missings, corev1.ContainerStateTerminated{}, "corev1", false)
	results = addNewObject(results, missings, corev1.ContainerStateWaiting{}, "corev1", false)
	results = addNewObject(results, missings, corev1.ContainerStatus{}, "corev1", false)
	results = addNewObject(results, missings, corev1.DaemonEndpoint{}, "corev1", false)
	results = addNewObject(results, missings, corev1.EnvFromSource{}, "corev1", false)
	results = addNewObject(results, missings, corev1.EnvVar{}, "corev1", false)
	results = addNewObject(results, missings, corev1.EnvVarSource{}, "corev1", false)
	results = addNewObject(results, missings, corev1.ExecAction{}, "corev1", false)
	results = addNewObject(results, missings, corev1.HTTPGetAction{}, "corev1", false)
	results = addNewObject(results, missings, corev1.HTTPHeader{}, "corev1", false)
	results = addNewObject(results, missings, corev1.Handler{}, "corev1", false)
	results = addNewObject(results, missings, corev1.HostAlias{}, "corev1", false)
	results = addNewObject(results, missings, corev1.Lifecycle{}, "corev1", false)
	results = addNewObject(results, missings, corev1.LoadBalancerIngress{}, "corev1", false)
	results = addNewObject(results, missings, corev1.LoadBalancerStatus{}, "corev1", false)
	results = addNewObject(results, missings, corev1.LocalObjectReference{}, "corev1", false)
	results = addNewObject(results, missings, corev1.NamespaceSpec{}, "corev1", false)
	results = addNewObject(results, missings, corev1.NamespaceStatus{}, "corev1", false)
	results = addNewObject(results, missings, corev1.NodeAddress{}, "corev1", false)
	results = addNewObject(results, missings, corev1.NodeAffinity{}, "corev1", false)
	results = addNewObject(results, missings, corev1.NodeCondition{}, "corev1", false)
	results = addNewObject(results, missings, corev1.NodeConfigSource{}, "corev1", false)
	results = addNewObject(results, missings, corev1.NodeConfigStatus{}, "corev1", false)
	results = addNewObject(results, missings, corev1.NodeDaemonEndpoints{}, "corev1", false)
	results = addNewObject(results, missings, corev1.NodeSelector{}, "corev1", false)
	results = addNewObject(results, missings, corev1.NodeSelectorRequirement{}, "corev1", false)
	results = addNewObject(results, missings, corev1.NodeSelectorTerm{}, "corev1", false)
	results = addNewObject(results, missings, corev1.NodeSpec{}, "corev1", false)
	results = addNewObject(results, missings, corev1.NodeStatus{}, "corev1", false)
	results = addNewObject(results, missings, corev1.NodeSystemInfo{}, "corev1", false)
	results = addNewObject(results, missings, corev1.ObjectFieldSelector{}, "corev1", false)
	results = addNewObject(results, missings, corev1.ObjectReference{}, "corev1", false)
	results = addNewObject(results, missings, corev1.PersistentVolumeClaimCondition{}, "corev1", false)
	results = addNewObject(results, missings, corev1.PersistentVolumeClaimSpec{}, "corev1", false)
	results = addNewObject(results, missings, corev1.PersistentVolumeClaimStatus{}, "corev1", false)
	results = addNewObject(results, missings, corev1.PersistentVolumeSpec{}, "corev1", false)
	results = addNewObject(results, missings, corev1.PersistentVolumeStatus{}, "corev1", false)
	results = addNewObject(results, missings, corev1.PodAffinity{}, "corev1", false)
	results = addNewObject(results, missings, corev1.PodAffinityTerm{}, "corev1", false)
	results = addNewObject(results, missings, corev1.PodAntiAffinity{}, "corev1", false)
	results = addNewObject(results, missings, corev1.PodCondition{}, "corev1", false)
	results = addNewObject(results, missings, corev1.PodDNSConfig{}, "corev1", false)
	results = addNewObject(results, missings, corev1.PodDNSConfigOption{}, "corev1", false)
	results = addNewObject(results, missings, corev1.PodReadinessGate{}, "corev1", false)
	results = addNewObject(results, missings, corev1.PodSecurityContext{}, "corev1", false)
	results = addNewObject(results, missings, corev1.PodSpec{}, "corev1", false)
	results = addNewObject(results, missings, corev1.PodStatus{}, "corev1", false)
	results = addNewObject(results, missings, corev1.PodTemplateSpec{}, "corev1", false)
	results = addNewObject(results, missings, corev1.PreferredSchedulingTerm{}, "corev1", false)
	results = addNewObject(results, missings, corev1.Probe{}, "corev1", false)
	results = addNewObject(results, missings, corev1.ReplicationControllerSpec{}, "corev1", false)
	results = addNewObject(results, missings, corev1.ReplicationControllerStatus{}, "corev1", false)
	results = addNewObject(results, missings, corev1.ReplicationControllerCondition{}, "corev1", false)
	results = addNewObject(results, missings, corev1.ResourceFieldSelector{}, "corev1", false)
	results = addNewObject(results, missings, corev1.ResourceRequirements{}, "corev1", false)
	results = addNewObject(results, missings, corev1.SELinuxOptions{}, "corev1", false)
	results = addNewObject(results, missings, corev1.SecretEnvSource{}, "corev1", false)
	results = addNewObject(results, missings, corev1.SecretKeySelector{}, "corev1", false)
	results = addNewObject(results, missings, corev1.SecurityContext{}, "corev1", false)
	results = addNewObject(results, missings, corev1.ServicePort{}, "corev1", false)
	results = addNewObject(results, missings, corev1.ServiceSpec{}, "corev1", false)
	results = addNewObject(results, missings, corev1.ServiceStatus{}, "corev1", false)
	results = addNewObject(results, missings, corev1.SessionAffinityConfig{}, "corev1", false)
	results = addNewObject(results, missings, corev1.Sysctl{}, "corev1", false)
	results = addNewObject(results, missings, corev1.TCPSocketAction{}, "corev1", false)
	results = addNewObject(results, missings, corev1.Taint{}, "corev1", false)
	results = addNewObject(results, missings, corev1.Toleration{}, "corev1", false)
	results = addNewObject(results, missings, corev1.TopologySelectorTerm{}, "corev1", false)
	results = addNewObject(results, missings, corev1.TopologySelectorLabelRequirement{}, "corev1", false)
	results = addNewObject(results, missings, corev1.TypedLocalObjectReference{}, "corev1", false)
	results = addNewObject(results, missings, corev1.Volume{}, "corev1", false)
	results = addNewObject(results, missings, corev1.VolumeDevice{}, "corev1", false)
	results = addNewObject(results, missings, corev1.VolumeMount{}, "corev1", false)
	results = addNewObject(results, missings, corev1.VolumeNodeAffinity{}, "corev1", false)
	results = addNewObject(results, missings, corev1.WeightedPodAffinityTerm{}, "corev1", false)

	results = addNewObject(results, missings, intstr.IntOrString{}, "intstr", false)

	results = addNewObject(results, missings, metricsv1beta1.ContainerMetrics{}, "metricsv1beta1", false)

	results = addNewObject(results, missings, metav1.Duration{}, "metav1", false)
	results = addNewObject(results, missings, metav1.Initializer{}, "metav1", false)
	results = addNewObject(results, missings, metav1.Initializers{}, "metav1", false)
	results = addNewObject(results, missings, metav1.LabelSelector{}, "metav1", false)
	results = addNewObject(results, missings, metav1.LabelSelectorRequirement{}, "metav1", false)
	results = addNewObject(results, missings, metav1.ListMeta{}, "metav1", false)
	results = addNewObject(results, missings, metav1.ObjectMeta{}, "metav1", false)
	results = addNewObject(results, missings, metav1.OwnerReference{}, "metav1", false)
	results = addNewObject(results, missings, metav1.Status{}, "metav1", false)
	results = addNewObject(results, missings, metav1.StatusCause{}, "metav1", false)
	results = addNewObject(results, missings, metav1.StatusDetails{}, "metav1", false)

	results = addNewObject(results, missings, networkingv1.IPBlock{}, "networkingv1", false)
	results = addNewObject(results, missings, networkingv1.NetworkPolicySpec{}, "networkingv1", false)
	results = addNewObject(results, missings, networkingv1.NetworkPolicyEgressRule{}, "networkingv1", false)
	results = addNewObject(results, missings, networkingv1.NetworkPolicyIngressRule{}, "networkingv1", false)
	results = addNewObject(results, missings, networkingv1.NetworkPolicyPeer{}, "networkingv1", false)
	results = addNewObject(results, missings, networkingv1.NetworkPolicyPort{}, "networkingv1", false)

	results = addNewObject(results, missings, rbacv1.PolicyRule{}, "rbacv1", false)
	results = addNewObject(results, missings, rbacv1.RoleRef{}, "rbacv1", false)
	results = addNewObject(results, missings, rbacv1.Subject{}, "rbacv1", false)
	results = addNewObject(results, missings, rbacv1.AggregationRule{}, "rbacv1", false)

	results = addNewObject(results, missings, resource.Quantity{}, "resource", false)

	// Add the hard coded objects
	results = addMapOfStrings(results)

	missingNames := make([]string, 0, len(missings))
	for k := range missings {
		missingNames = append(missingNames, k)
	}

	sort.Slice(missingNames, func(i, j int) bool {
		return strings.Compare(missingNames[i], missingNames[j]) < 0
	})
	for _, missing := range missingNames {

		generated := false
		expected := fmt.Sprintf("\"%s\": {\n", missing)
		for _, generatedLine := range results {
			if generatedLine == expected {
				generated = true
			}
		}

		if generated {
			fmt.Printf("%v\n", missing)
		} else {
			fmt.Printf(">>>>>>%v\n", missing)
		}

	}

	return results
}

func addNewObject(lines []string, missings map[string]bool, v interface{}, prefix string, headObject bool) []string {

	objResults, objMissing := iterateFields(v, prefix, headObject)
	lines = append(lines, objResults...)

	for _, missing := range objMissing {
		missings[missing] = true
	}
	return lines
}

func addMapOfStrings(lines []string) []string {

	lines = append(lines, fmt.Sprintf("\"MapOfStrings\": {\n"))
	lines = append(lines, fmt.Sprintf("    \"type\": \"object\",\n"))
	lines = append(lines, fmt.Sprintf("    \"additionalProperties\": {\n"))
	lines = append(lines, fmt.Sprintf("        \"type\": \"array\",\n"))
	lines = append(lines, fmt.Sprintf("        \"items\": {\n"))
	lines = append(lines, fmt.Sprintf("            \"type\": \"string\"\n"))
	lines = append(lines, fmt.Sprintf("        }\n"))
	lines = append(lines, fmt.Sprintf("    }\n"))
	lines = append(lines, fmt.Sprintf("},\n"))

	return lines
}

func iterateFields(v interface{}, prefix string, headObject bool) ([]string, []string) {

	results := make([]string, 0, 50)
	missing := make([]string, 0, 10)

	valueOf := reflect.ValueOf(v)
	typeOf := reflect.TypeOf(v)

	if headObject {
		results = append(results, fmt.Sprintf("\"%s\": {\n", typeOf.Name()))
	} else {
		results = append(results, fmt.Sprintf("\"%s.%s\": {\n", prefix, typeOf.Name()))
	}
	results = append(results, fmt.Sprintf("    \"type\": \"object\",\n"))
	results = append(results, fmt.Sprintf("    \"properties\": {\n"))

	for i := 0; i < valueOf.NumField(); i++ {

		value := valueOf.Field(i)
		fieldType := typeOf.Field(i)
		tag := getNameFromTag(fieldType)

		if len(tag) > 0 {
			results = append(results, fmt.Sprintf("        \"%v\": {\n", tag))

			if value.Type().String() == "*v1.Time" || value.Type().String() == "v1.Time" {
				results = append(results, fmt.Sprintf("            \"type\": \"string\"\n"))
			} else if value.Type().String() == "map[string]string" {
				results = append(results, fmt.Sprintf("            \"type\": \"object\",\n"))
				results = append(results, fmt.Sprintf("            \"additionalProperties\": {\n"))
				results = append(results, fmt.Sprintf("                \"type\": \"string\",\n"))
				results = append(results, fmt.Sprintf("            }\n"))
			} else if value.Type().String() == "types.UID" {
				results = append(results, fmt.Sprintf("            \"type\": \"string\"\n"))
			} else {
				if value.Kind() == reflect.Struct {
					targetType, _ := guessTarget(value, prefix)
					results = append(results, fmt.Sprintf("            \"type\": \"object\",\n"))
					results = append(results, fmt.Sprintf("            \"$ref\": \"#/definitions/%v\",\n", targetType))
					missing = append(missing, targetType)
				} else if value.Kind() == reflect.Ptr {
					targetType, isComplex := guessTarget(value, prefix)
					if isComplex {
						results = append(results, fmt.Sprintf("            \"$ref\": \"#/definitions/%v\",\n", targetType))
						missing = append(missing, targetType)
					} else {
						results = append(results, fmt.Sprintf("            \"type\": \"%v\"\n", targetType))
					}
				} else if value.Kind() == reflect.Slice {
					targetType, isComplex := guessTarget(value, prefix)
					results = append(results, fmt.Sprintf("            \"type\": \"array\",\n"))
					results = append(results, fmt.Sprintf("            \"items\": {\n"))
					if isComplex {
						results = append(results, fmt.Sprintf("                \"$ref\": \"#/definitions/%v\",\n", targetType))
						missing = append(missing, targetType)
					} else {
						results = append(results, fmt.Sprintf("                \"type\": \"%v\",\n", targetType))
					}
					results = append(results, fmt.Sprintf("            }\n"))
				} else {
					targetType, isComplex := guessTarget(value, prefix)
					if isComplex {
						results = append(results, fmt.Sprintf("            \"type\": \"%v\"\n", value.Kind()))
					} else {
						results = append(results, fmt.Sprintf("            \"type\": \"%v\"\n", targetType))
					}
				}
			}

			if i < valueOf.NumField()-1 {
				results = append(results, fmt.Sprintf("        },\n"))

			} else {
				results = append(results, fmt.Sprintf("        }\n"))
			}
		}

	}
	results = append(results, fmt.Sprintf("    }\n"))
	results = append(results, fmt.Sprintf("},\n"))

	return results, missing
}

func getNameFromTag(field reflect.StructField) string {
	allTags := field.Tag.Get("json")
	tags := strings.Split(allTags, ",")
	for _, tag := range tags {
		if tag != "omitempty" && tag != "inline" {
			return tag
		}
	}
	return ""
}

func guessTarget(value reflect.Value, prefix string) (string, bool) {

	typeStr := value.Type().String()

	// special cases
	if typeStr == "v1.ObjectMeta" {
		return "metav1.ObjectMeta", true
	}

	if typeStr == "v1.OwnerReference" {
		return "metav1.OwnerReference", true
	}

	if typeStr[0] == '*' {
		targetType := strings.Replace(typeStr[1:], "v1", prefix, -1)
		return convertTargetType(targetType)
	}
	if typeStr[0] == '[' {
		targetType := strings.Replace(typeStr[2:], "v1", prefix, -1)
		return convertTargetType(targetType)
	}

	targetType := strings.Replace(typeStr, "v1", prefix, -1)
	return convertTargetType(targetType)
}

func convertTargetType(targetType string) (string, bool) {
	if targetType == "int64" || targetType == "int32" {
		return "integer", false
	}
	if targetType == "bool" {
		return "boolean", false
	}
	if targetType == "string" ||
		targetType == "corev1.FinalizerName" ||
		targetType == "corev1.PersistentVolumeAccessMode" ||
		targetType == "corev1.PersistentVolumeMode" ||
		targetType == "corev1.UniqueVolumeName" ||
		targetType == "corev1.Capability" ||
		targetType == "corev1.MountPropagationMode" ||
		targetType == "corev1.ProcMountType" ||
		targetType == "networkingv1.PolicyType" ||
		targetType == "networkingv1.Protocol" ||
		targetType == "storagev1.VolumeBindingMode" ||
		targetType == "storagev1.PersistentVolumeReclaimPolicy" {
		return "string", false
	}

	// Interpackage references
	if strings.HasSuffix(targetType, ".LabelSelector") {
		targetType = "metav1.LabelSelector"
	}
	if strings.HasSuffix(targetType, ".LabelSelectorRequirement") {
		targetType = "metav1.LabelSelectorRequirement"
	}
	if strings.HasSuffix(targetType, ".PodTemplateSpec") {
		targetType = "corev1.PodTemplateSpec"
	}
	if strings.HasSuffix(targetType, ".ObjectReference") {
		targetType = "corev1.ObjectReference"
	}
	if strings.HasSuffix(targetType, ".JobSpec") {
		targetType = "batchv1.JobSpec"
	}
	if strings.HasSuffix(targetType, ".TopologySelectorTerm") {
		targetType = "corev1.TopologySelectorTerm"
	}
	if strings.HasSuffix(targetType, ".TopologySelectorLabelRequirement") {
		targetType = "corev1.TopologySelectorLabelRequirement"
	}
	if strings.HasSuffix(targetType, ".Duration") {
		targetType = "metav1.Duration"
	}
	if strings.HasSuffix(targetType, ".ResourceList") {
		targetType = "corev1.ResourceList"
	}
	if strings.HasSuffix(targetType, ".ContainerMetrics") {
		targetType = "metricsv1beta1.ContainerMetrics"
	}

	// sub-object referencing head object
	if strings.HasSuffix(targetType, ".PersistentVolumeClaim") {
		targetType = "PersistentVolumeClaim"
	}

	// stupid concatenation
	targetType = strings.ReplaceAll(targetType, "batchv1beta1beta1", "batchv1beta1")

	return targetType, true
}
