package context

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// GetKubeConfig reads the configuration file
func GetKubeConfig() (*KubeConfig, error) {

	if len(contextConfigurationFileName) == 0 {
		return nil, errors.New("the context was not initialized before use")
	}

	// Open the config file
	var config KubeConfig
	source, err := ioutil.ReadFile(contextConfigurationFileName)
	if err != nil {
		return nil, err
	}

	// Read the contextNames
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		return nil, err
	}

	return &config, err
}

// GetKubeUsers returns the users configured
func GetKubeUsers() ([]NamedUser, error) {

	config, err := GetKubeConfig()
	if err != nil {
		return nil, err
	}

	return config.Users, nil
}

// GetKubeUser return a single user if present, nil otherwise
func GetKubeUser(name string) (*NamedUser, error) {

	config, err := GetKubeConfig()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(config.Users); i++ {
		if config.Users[i].Name == name {
			return &(config.Users[i]), nil
		}
	}

	return nil, nil
}

// GetKubeClusters returns the clusters configured
func GetKubeClusters() ([]NamedCluster, error) {

	config, err := GetKubeConfig()
	if err != nil {
		return nil, err
	}

	return config.Clusters, nil
}

// GetKubeCluster return a single cluster if present, nil otherwise
func GetKubeCluster(name string) (*NamedCluster, error) {

	config, err := GetKubeConfig()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(config.Clusters); i++ {
		if config.Clusters[i].Name == name {
			return &(config.Clusters[i]), nil
		}
	}

	return nil, nil
}

// GetKubeContexts returns the clusters configured
func GetKubeContexts() ([]NamedContext, error) {

	config, err := GetKubeConfig()
	if err != nil {
		return nil, err
	}

	return config.Contexts, nil
}

// GetKubeContext return a single context if present, nil otherwise
func GetKubeContext(name string) (*NamedContext, error) {

	config, err := GetKubeConfig()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(config.Contexts); i++ {
		if config.Contexts[i].Name == name {
			return &(config.Contexts[i]), nil
		}
	}

	return nil, nil
}
