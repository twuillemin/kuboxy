package context

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

// UseContext defines the current default context
func UseContext(contextName string) error {

	config, err := GetKubeConfig()
	if err != nil {
		return err
	}

	// Check if the context exist
	found := false
	for _, context := range config.Contexts {
		if context.Name == contextName {
			found = true
			break
		}
	}
	if !found {
		return errors.New("the context to be used as default does not exist")
	}

	// Update the configuration
	config.CurrentContext = contextName

	return writeConfigFile(config)
}

// SetUserWithUserNamePassword adds or updates a user with the given credentials
func SetUserWithUserNamePassword(userName string, credentials ParamCredentialsUserNamePassword) error {

	return createOrUpdateUser(
		userName,
		NamedUser{
			Name: userName,
			DefinitionUser: DefinitionUser{
				UserName: credentials.UserName,
				Password: credentials.Password,
			},
		})
}

// SetUserWithCertificateFile adds or updates a user with the credentials given as local certificate files
func SetUserWithCertificateFile(userName string, credentials ParamCredentialsCertificateFile) error {

	return createOrUpdateUser(
		userName,
		NamedUser{
			Name: userName,
			DefinitionUser: DefinitionUser{
				ClientCertificate: credentials.ClientCertificate,
				ClientKey:         credentials.ClientKey,
			},
		})
}

// SetUserWithCertificateFileEmbedded adds or updates a user with the credentials given as embedded certificate files
func SetUserWithCertificateFileEmbedded(userName string, credentials ParamCredentialsCertificateEmbedded) error {

	return createOrUpdateUser(
		userName,
		NamedUser{
			Name: userName,
			DefinitionUser: DefinitionUser{
				ClientCertificateData: credentials.ClientCertificateData,
				ClientKeyData:         credentials.ClientKeyData,
			},
		})
}

// createOrUpdateUser creates or updates the given userName entry with the given user object
func createOrUpdateUser(userName string, user NamedUser) error {

	config, err := GetKubeConfig()
	if err != nil {
		return err
	}

	// Check if the user exist exist
	userIndex := -1
	for i, u := range config.Users {
		if u.Name == userName {
			userIndex = i
			break
		}
	}

	// Add or update to the users
	if userIndex < 0 {
		config.Users = append(config.Users, user)
	} else {
		config.Users[userIndex] = user
	}

	return writeConfigFile(config)
}

// SetClusterInsecure adds or updates a cluster without validating the server certificate
func SetClusterInsecure(clusterName string, cluster ParamClusterInsecure) error {

	return createOrUpdateCluster(
		clusterName,
		NamedCluster{
			Name: clusterName,
			DefinitionCluster: DefinitionCluster{
				Server:                cluster.Server,
				InsecureSkipTLSVerify: true,
			},
		})
}

// SetClusterCertificateFile adds or updates a cluster without validating the server certificate given as a local file
func SetClusterCertificateFile(clusterName string, cluster ParamClusterCertificateFile) error {

	return createOrUpdateCluster(
		clusterName,
		NamedCluster{
			Name: clusterName,
			DefinitionCluster: DefinitionCluster{
				Server:                cluster.Server,
				CertificateAuthority:  cluster.CertificateAuthority,
				InsecureSkipTLSVerify: false,
			},
		})
}

// SetClusterCertificateEmbedded adds or updates a cluster without validating the server certificate given as raw text
func SetClusterCertificateEmbedded(clusterName string, cluster ParamClusterCertificateEmbedded) error {

	return createOrUpdateCluster(
		clusterName,
		NamedCluster{
			Name: clusterName,
			DefinitionCluster: DefinitionCluster{
				Server:                   cluster.Server,
				CertificateAuthorityData: cluster.CertificateAuthorityData,
				InsecureSkipTLSVerify:    false,
			},
		})
}

// createOrUpdateUser creates or updates the given userName entry with the given user object
func createOrUpdateCluster(clusterName string, cluster NamedCluster) error {

	config, err := GetKubeConfig()
	if err != nil {
		return err
	}

	// Check if the user exist exist
	clusterIndex := -1
	for i, c := range config.Clusters {
		if c.Name == clusterName {
			clusterIndex = i
			break
		}
	}

	// Add or update to the clusters
	if clusterIndex < 0 {
		config.Clusters = append(config.Clusters, cluster)
	} else {
		config.Clusters[clusterIndex] = cluster
	}

	return writeConfigFile(config)
}

// SetContext adds or updates a context
func SetContext(contextName string, context ParamContext) error {

	return createOrUpdateContext(
		contextName,
		NamedContext{
			Name: contextName,
			DefinitionContext: DefinitionContext{
				Cluster:   context.Cluster,
				User:      context.User,
				Namespace: context.Namespace,
			},
		})
}

// createOrUpdateUser creates or updates the given userName entry with the given user object
func createOrUpdateContext(clusterName string, context NamedContext) error {

	config, err := GetKubeConfig()
	if err != nil {
		return err
	}

	// Check if the user exist exist
	contextIndex := -1
	for i, c := range config.Contexts {
		if c.Name == clusterName {
			contextIndex = i
			break
		}
	}

	// Add or update to the contexts
	if contextIndex < 0 {
		config.Contexts = append(config.Contexts, context)
	} else {
		config.Contexts[contextIndex] = context
	}

	return writeConfigFile(config)
}

// writeConfigFile writes the given configuration into the context configuration file
func writeConfigFile(config *KubeConfig) (err error) {

	file, err := os.OpenFile(contextConfigurationFileName, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("unable to open the context configuration file due to: %v", err.Error())
	}

	defer func() {
		err = file.Close()
	}()

	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("unable to marshall the given configuration due to: %v", err.Error())
	}

	// Write the content to the config file
	if _, err = file.Write(data); err != nil {
		return fmt.Errorf("unable to write the configuration into the file due to: %v", err.Error())
	}

	return nil
}
