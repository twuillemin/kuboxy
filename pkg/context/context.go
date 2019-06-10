package context

import (
	"fmt"
	"github.com/twuillemin/kuboxy/internal/configuration"
	"gopkg.in/yaml.v2"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
	"os"
	"path"
	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

// The known contextNames (bool is not used)
var contextNames = make(map[string]bool)

// The relations between all the known configuration and their connection
var clientsets = make(map[string]*kubernetes.Clientset)

// The relations between all the known configuration and their connection
var versionedClientsets = make(map[string]*metrics.Clientset)

// Internal copy of the context configuration file name
var contextConfigurationFileName = ""

// LoadContexts loads all the possible configuration from a Kubectl config file. If there are no configuration, try to default
// to the InCluster mode
func LoadContexts(applicationConfiguration configuration.ApplicationConfiguration) error {

	// Keep the name of context configuration file (could be something like ~/.kube/config) as it could be used multiple
	// times
	contextConfigurationFileName = applicationConfiguration.KubeContextConfigurationFile

	// Ensure there is a context file (even empty)
	if err := ensureContextConfigFile(); err != nil {
		return err
	}

	config, err := GetKubeConfig()
	if err != nil {
		return err
	}

	// Otherwise declare the configuration found (but do not make the client set)
	for i := 0; i < len(config.Contexts); i++ {
		contextNames[config.Contexts[i].Name] = true
	}

	return nil
}

// ensureContextConfigFile ensures that the given context configuration file exist and is readable. If needed create
// an empty one
func ensureContextConfigFile() (err error) {

	// If the context file exists, just return
	if _, err = os.Stat(contextConfigurationFileName); err == nil {
		return nil
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("unable to access to the context configuration file due to: %v", err.Error())
	}

	// Get the directory for the configuration files
	directory := path.Dir(contextConfigurationFileName)

	// If the context file does not exist, create it
	if _, err = os.Stat(directory); err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("unable to access to the directory for context configuration file due to: %v", err.Error())
		}
		if err = os.MkdirAll(directory, 0700); err != nil {
			return fmt.Errorf("unable to create to the directory for context configuration file due to: %v", err.Error())
		}
	}

	// At this point the file does not exist, so just create an empty one
	newContextFile, err := os.Create(contextConfigurationFileName)
	if err != nil {
		return fmt.Errorf("unable to create a new context configuration file due to: %v", err.Error())
	}

	defer func() {
		err = newContextFile.Close()
	}()

	if err = os.Chmod(contextConfigurationFileName, 0600); err != nil {
		return fmt.Errorf("unable to define access rights to the new context configuration file due to: %v", err.Error())
	}

	// Create a default configuration file content
	emptyConfig := KubeConfig{
		APIVersion:  "v1",
		Kind:        "Config",
		Preferences: make(map[string]string),
	}
	data, err := yaml.Marshal(&emptyConfig)
	if err != nil {
		return fmt.Errorf("unable to create default context content due to: %v", err.Error())
	}

	// Write the default content to the config file
	if _, err = newContextFile.Write(data); err != nil {
		return fmt.Errorf("unable to write the default context content due to: %v", err.Error())
	}

	return nil
}

// GetContextNames gives the list of all contextNames known by the application
func GetContextNames() []string {
	result := make([]string, 0, len(contextNames))
	for contextName := range contextNames {
		result = append(result, contextName)
	}
	return result
}

// GetClientset gives a clientset for the given contextName
func GetClientset(contextName string) (*kubernetes.Clientset, error) {

	// Try to get it from the cache
	existingClientset := clientsets[contextName]
	if existingClientset != nil {
		return existingClientset, nil
	}

	// Check if the contextName exists in the list of possible contextNames
	if _, ok := contextNames[contextName]; !ok {
		return nil, &NotFoundError{contextName}
	}

	// Update the configuration
	err := UseContext(contextName)
	if err != nil {
		return nil, err
	}

	// Try to build an empty config, that should default to in cluster mode
	config, err := clientcmd.BuildConfigFromFlags("", contextConfigurationFileName)
	if err != nil {
		return nil, err
	}

	// Set High QPS and Burst because we may query intensively when doing the report
	config.QPS = 1e6
	config.Burst = 1e6

	// Try to connect with the update config
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	// Keep the client set
	clientsets[contextName] = clientset

	return clientset, nil
}

// GetMetrics gives a clientset for the given contextName
func GetMetrics(contextName string) (*metrics.Clientset, error) {

	// Try to get it from the cache
	existingClientset := versionedClientsets[contextName]
	if existingClientset != nil {
		return existingClientset, nil
	}

	// Check if the contextName exists in the list of possible contextNames
	if _, ok := contextNames[contextName]; !ok {
		return nil, &NotFoundError{contextName}
	}

	// Update the configuration
	err := UseContext(contextName)
	if err != nil {
		return nil, err
	}

	// Try to build an empty config, that should default to in cluster mode
	config, err := clientcmd.BuildConfigFromFlags("", contextConfigurationFileName)
	if err != nil {
		return nil, err
	}

	// Set High QPS and Burst because we may query intensively when doing the report
	config.QPS = 1e6
	config.Burst = 1e6

	// Try to connect with the update config
	versioned, err := metrics.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	// Keep the client set
	versionedClientsets[contextName] = versioned

	return versioned, nil
}
