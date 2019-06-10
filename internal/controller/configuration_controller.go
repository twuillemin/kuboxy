package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/twuillemin/kuboxy/pkg/context"
)

func registerConfigurationController(e *echo.Echo) {

	// Declare the routes
	e.GET("api/v1/configuration/", getConfiguration)

	// The users
	e.GET("api/v1/configuration/users/", getConfigurationUsers)
	e.POST("/api/v1/configuration/users/:name/username-password", createConfigurationUserUserNamePassword)
	e.PUT("/api/v1/configuration/users/:name/username-password", updateConfigurationUserUserNamePassword)
	e.POST("/api/v1/configuration/users/:name/certificate-file", createConfigurationUserFile)
	e.PUT("/api/v1/configuration/users/:name/certificate-file", updateConfigurationUserFile)
	e.POST("/api/v1/configuration/users/:name/certificate-embedded", createConfigurationUserEmbedded)
	e.PUT("/api/v1/configuration/users/:name/certificate-embedded", updateConfigurationUserEmbedded)

	// The cluster
	e.GET("api/v1/configuration/clusters/", getConfigurationClusters)
	e.POST("/api/v1/configuration/clusters/:name/insecure", createConfigurationClusterInsecure)
	e.PUT("/api/v1/configuration/clusters/:name/insecure", updateConfigurationClusterInsecure)
	e.POST("/api/v1/configuration/clusters/:name/certificate-file", createConfigurationClusterFile)
	e.PUT("/api/v1/configuration/clusters/:name/certificate-file", updateConfigurationClusterFile)
	e.POST("/api/v1/configuration/clusters/:name/certificate-embedded", createConfigurationClusterEmbedded)
	e.PUT("/api/v1/configuration/clusters/:name/certificate-embedded", updateConfigurationClusterEmbedded)

	// The context
	e.GET("api/v1/configuration/contexts/", getConfigurationContexts)
	e.POST("/api/v1/configuration/contexts/:name", createConfigurationContext)
	e.PUT("/api/v1/configuration/contexts/:name", updateConfigurationContext)
}

// getConfiguration generates a JSON representation of all the configuration
// @Summary Retrieve the configuration
// @Description get the configuration
// @ID get-configuration
// @Tags Configuration
// @Produce application/json
// @Success 200 {object} context.KubeConfig
// @Router /api/v1/configuration/ [get]
func getConfiguration(e echo.Context) error {

	conf, err := context.GetKubeConfig()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return e.JSON(http.StatusOK, conf)
}

// getConfigurationUsers generates a JSON representation of the users
// @Summary Retrieve the users
// @Description get the users
// @ID get-configuration-users
// @Tags Configuration
// @Produce application/json
// @Success 200 {array} context.NamedUser
// @Router /api/v1/configuration/users/ [get]
func getConfigurationUsers(e echo.Context) error {

	users, err := context.GetKubeUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return e.JSON(http.StatusOK, users)
}

// createConfigurationUserUserNamePassword creates a new user in the configuration
// @Summary Create a new user
// @Description Create a new user by giving its name in the configuration and the username and the
// @Description password to connect to the server
// @ID post-configuration-user-username-password
// @Tags Configuration
// @Accept json
// @Produce application/json
// @Param name path string true "the name of the user in the configuration"
// @Param body body context.ParamCredentialsUserNamePassword true "the credentials"
// @Success 200 {object} context.NamedUser
// @Success 400 {object} echo.HTTPError
// @Success 409 {object} echo.HTTPError
// @Router /api/v1/configuration/users/{name}/username-password [post]
func createConfigurationUserUserNamePassword(e echo.Context) error {

	name := e.Param("name")

	// Ensure that the user does not already exist
	user, err := context.GetKubeUser(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if user != nil {
		return echo.NewHTTPError(http.StatusConflict, fmt.Errorf("the user %s already exist", name))
	}

	// Parse the credentials
	credentials := new(context.ParamCredentialsUserNamePassword)
	if err = e.Bind(credentials); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// Add the user
	err = context.SetUserWithUserNamePassword(name, *credentials)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// Read the newly created object
	user, err = context.GetKubeUser(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if user == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("unable to retrieve the user %s from configuration after creation", name))
	}

	return e.JSON(http.StatusOK, user)
}

// updateConfigurationUserUserNamePassword updates an existing user in the configuration
// @Summary Update an existing user
// @Description Update an existing user by giving its name in the configuration and the username and the
// @Description password to connect to the server
// @ID put-configuration-user-username-password
// @Tags Configuration
// @Accept json
// @Produce application/json
// @Param name path string true "the name of the user in the configuration"
// @Param body body context.ParamCredentialsUserNamePassword true "the credentials"
// @Success 200 {object} context.NamedUser
// @Success 404 {object} echo.HTTPError
// @Success 409 {object} echo.HTTPError
// @Router /api/v1/configuration/users/{name}/username-password [put]
func updateConfigurationUserUserNamePassword(e echo.Context) error {

	name := e.Param("name")

	// Ensure that the user does not already exist
	user, err := context.GetKubeUser(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if user == nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("the user %s does not exist", name))
	}

	// Parse the credentials
	credentials := new(context.ParamCredentialsUserNamePassword)
	if err = e.Bind(credentials); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// Add the user
	err = context.SetUserWithUserNamePassword(name, *credentials)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// Read the newly created object
	user, err = context.GetKubeUser(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if user == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("unable to retrieve the user %s from configuration after update", name))
	}

	return e.JSON(http.StatusOK, user)
}

// createConfigurationUserFile creates a new user in the configuration with the certificate given as local files
// @Summary Create a new user with the certificates given as local files
// @Description Create a new user by giving its name in the configuration and its certificates as local files
// @ID post-configuration-user-file
// @Tags Configuration
// @Accept json
// @Produce application/json
// @Param name path string true "the name of the user in the configuration"
// @Param body body context.ParamCredentialsCertificateFile true "the credentials"
// @Success 200 {object} context.NamedUser
// @Success 400 {object} echo.HTTPError
// @Success 409 {object} echo.HTTPError
// @Router /api/v1/configuration/users/{name}/certificate-file [post]
func createConfigurationUserFile(e echo.Context) error {

	name := e.Param("name")

	// Ensure that the user does not already exist
	user, err := context.GetKubeUser(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if user != nil {
		return echo.NewHTTPError(http.StatusConflict, fmt.Errorf("the user %s already exist", name))
	}

	// Parse the credentials
	credentials := new(context.ParamCredentialsCertificateFile)
	if err = e.Bind(credentials); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// Add the user
	err = context.SetUserWithCertificateFile(name, *credentials)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// Read the newly created object
	user, err = context.GetKubeUser(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if user == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("unable to retrieve the user %s from configuration after creation", name))
	}

	return e.JSON(http.StatusOK, user)
}

// updateConfigurationUserFile updates an existing user in the configuration with the certificate given as local files
// @Summary Update an existing user with the certificates given as local files
// @Description Update an existing user by giving its name in the configuration and its certificates as local files
// @ID put-configuration-user-file
// @Tags Configuration
// @Accept json
// @Produce application/json
// @Param name path string true "the name of the user in the configuration"
// @Param body body context.ParamCredentialsCertificateFile true "the credentials"
// @Success 200 {object} context.NamedUser
// @Success 404 {object} echo.HTTPError
// @Success 409 {object} echo.HTTPError
// @Router /api/v1/configuration/users/{name}/certificate-file [put]
func updateConfigurationUserFile(e echo.Context) error {

	name := e.Param("name")

	// Ensure that the user does not already exist
	user, err := context.GetKubeUser(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if user == nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("the user %s does not exist", name))
	}

	// Parse the credentials
	credentials := new(context.ParamCredentialsCertificateFile)
	if err = e.Bind(credentials); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// Add the user
	err = context.SetUserWithCertificateFile(name, *credentials)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// Read the newly created object
	user, err = context.GetKubeUser(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if user == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("unable to retrieve the user %s from configuration after update", name))
	}

	return e.JSON(http.StatusOK, user)
}

// createConfigurationUserEmbedded creates a new user in the configuration with the certificate embedded
// @Summary Create a new user with the certificates embedded
// @Description Create a new user by giving its name in the configuration and its certificates embedded
// @ID post-configuration-user-embedded
// @Tags Configuration
// @Accept json
// @Produce application/json
// @Param name path string true "the name of the user in the configuration"
// @Param body body context.ParamCredentialsCertificateEmbedded true "the credentials"
// @Success 200 {object} context.NamedUser
// @Success 400 {object} echo.HTTPError
// @Success 409 {object} echo.HTTPError
// @Router /api/v1/configuration/users/{name}/certificate-embedded [post]
func createConfigurationUserEmbedded(e echo.Context) error {

	name := e.Param("name")

	// Ensure that the user does not already exist
	user, err := context.GetKubeUser(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if user != nil {
		return echo.NewHTTPError(http.StatusConflict, fmt.Errorf("the user %s already exist", name))
	}

	// Parse the credentials
	credentials := new(context.ParamCredentialsCertificateEmbedded)
	if err = e.Bind(credentials); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// Add the user
	err = context.SetUserWithCertificateFileEmbedded(name, *credentials)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// Read the newly created object
	user, err = context.GetKubeUser(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if user == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("unable to retrieve the user %s from configuration after creation", name))
	}

	return e.JSON(http.StatusOK, user)
}

// updateConfigurationUserEmbedded updates an existing user in the configuration with the certificate embedded
// @Summary Update an existing user with the certificates embedded
// @Description Update an existing user by giving its name in the configuration and its certificates embedded
// @ID put-configuration-user-embedded
// @Tags Configuration
// @Accept json
// @Produce application/json
// @Param name path string true "the name of the user in the configuration"
// @Param body body context.ParamCredentialsCertificateEmbedded true "the credentials"
// @Success 200 {object} context.NamedUser
// @Success 404 {object} echo.HTTPError
// @Success 409 {object} echo.HTTPError
// @Router /api/v1/configuration/users/{name}/certificate-embedded [put]
func updateConfigurationUserEmbedded(e echo.Context) error {

	name := e.Param("name")

	// Ensure that the user does not already exist
	user, err := context.GetKubeUser(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if user == nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("the user %s does not exist", name))
	}

	// Parse the credentials
	credentials := new(context.ParamCredentialsCertificateEmbedded)
	if err = e.Bind(credentials); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// Add the user
	err = context.SetUserWithCertificateFileEmbedded(name, *credentials)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// Read the newly created object
	user, err = context.GetKubeUser(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if user == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("unable to retrieve the user %s from configuration after update", name))
	}

	return e.JSON(http.StatusOK, user)
}

// getConfigurationClusters generates a JSON representation of the clusters
// @Summary Retrieve the clusters
// @Description get the clusters
// @ID get-configuration-clusters
// @Tags Configuration
// @Produce application/json
// @Success 200 {array} context.NamedCluster
// @Router /api/v1/configuration/clusters/ [get]
func getConfigurationClusters(e echo.Context) error {

	clusters, err := context.GetKubeClusters()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return e.JSON(http.StatusOK, clusters)
}

// createConfigurationClusterInsecure creates a new cluster in the configuration
// @Summary Create a new cluster
// @Description Create a new cluster for which the TLS certificate is not verified
// @ID post-configuration-cluster-insecure
// @Tags Configuration
// @Accept json
// @Produce application/json
// @Param name path string true "the name of the cluster in the configuration"
// @Param body body context.ParamClusterInsecure true "the definition of the cluster"
// @Success 200 {object} context.NamedCluster
// @Success 400 {object} echo.HTTPError
// @Success 409 {object} echo.HTTPError
// @Router /api/v1/configuration/clusters/{name}/insecure [post]
func createConfigurationClusterInsecure(e echo.Context) error {

	name := e.Param("name")

	// Ensure that the cluster does not already exist
	cluster, err := context.GetKubeCluster(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if cluster != nil {
		return echo.NewHTTPError(http.StatusConflict, fmt.Errorf("the cluster %s already exist", name))
	}

	// Parse the cluster information
	clusterInsecure := new(context.ParamClusterInsecure)
	if err = e.Bind(clusterInsecure); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// Add the cluster
	err = context.SetClusterInsecure(name, *clusterInsecure)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// Read the newly created object
	cluster, err = context.GetKubeCluster(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if cluster == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("unable to retrieve the cluster %s from configuration after creation", name))
	}

	return e.JSON(http.StatusOK, cluster)
}

// updateConfigurationClusterInsecure updates an existing cluster in the configuration
// @Summary Update an existing cluster
// @Description Update an existing cluster for which the TLS certificate is not verified
// @ID put-configuration-cluster-insecure
// @Tags Configuration
// @Accept json
// @Produce application/json
// @Param name path string true "the name of the cluster in the configuration"
// @Param body body context.ParamClusterInsecure true "the definition of the cluster"
// @Success 200 {object} context.NamedCluster
// @Success 404 {object} echo.HTTPError
// @Success 409 {object} echo.HTTPError
// @Router /api/v1/configuration/clusters/{name}/insecure [put]
func updateConfigurationClusterInsecure(e echo.Context) error {

	name := e.Param("name")

	// Ensure that the cluster does not already exist
	cluster, err := context.GetKubeCluster(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if cluster == nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("the cluster %s does not exist", name))
	}

	// Parse the cluster information
	clusterInsecure := new(context.ParamClusterInsecure)
	if err = e.Bind(clusterInsecure); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// Add the cluster
	err = context.SetClusterInsecure(name, *clusterInsecure)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// Read the newly created object
	cluster, err = context.GetKubeCluster(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if cluster == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("unable to retrieve the cluster %s from configuration after update", name))
	}

	return e.JSON(http.StatusOK, cluster)
}

// createConfigurationClusterFile creates a new cluster in the configuration with its certificate given as a local file
// @Summary Create a new cluster with its certificate given as a local file
// @Description Create a new cluster for which the TLS certificate is given as a local file
// @ID post-configuration-cluster-file
// @Tags Configuration
// @Accept json
// @Produce application/json
// @Param name path string true "the name of the cluster in the configuration"
// @Param body body context.ParamClusterCertificateFile true "the definition of the cluster"
// @Success 200 {object} context.NamedCluster
// @Success 400 {object} echo.HTTPError
// @Success 409 {object} echo.HTTPError
// @Router /api/v1/configuration/clusters/{name}/certificate-file [post]
func createConfigurationClusterFile(e echo.Context) error {

	name := e.Param("name")

	// Ensure that the cluster does not already exist
	cluster, err := context.GetKubeCluster(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if cluster != nil {
		return echo.NewHTTPError(http.StatusConflict, fmt.Errorf("the cluster %s already exist", name))
	}

	// Parse the cluster information
	clusterFile := new(context.ParamClusterCertificateFile)
	if err = e.Bind(clusterFile); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// Add the cluster
	err = context.SetClusterCertificateFile(name, *clusterFile)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// Read the newly created object
	cluster, err = context.GetKubeCluster(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if cluster == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("unable to retrieve the cluster %s from configuration after creation", name))
	}

	return e.JSON(http.StatusOK, cluster)
}

// updateConfigurationClusterFile updates an existing cluster in the configuration with its certificate given as a local file
// @Summary Update an existing cluster with its certificate given as a local file
// @Description Update an existing cluster for which the TLS certificate is given as a local file
// @ID put-configuration-cluster-file
// @Tags Configuration
// @Accept json
// @Produce application/json
// @Param name path string true "the name of the cluster in the configuration"
// @Param body body context.ParamClusterCertificateFile true "the definition of the cluster"
// @Success 200 {object} context.NamedCluster
// @Success 404 {object} echo.HTTPError
// @Success 409 {object} echo.HTTPError
// @Router /api/v1/configuration/clusters/{name}/certificate-file [put]
func updateConfigurationClusterFile(e echo.Context) error {

	name := e.Param("name")

	// Ensure that the cluster does not already exist
	cluster, err := context.GetKubeCluster(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if cluster == nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("the cluster %s does not exist", name))
	}

	// Parse the cluster information
	clusterFile := new(context.ParamClusterCertificateFile)
	if err = e.Bind(clusterFile); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// Add the cluster
	err = context.SetClusterCertificateFile(name, *clusterFile)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// Read the newly created object
	cluster, err = context.GetKubeCluster(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if cluster == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("unable to retrieve the cluster %s from configuration after update", name))
	}

	return e.JSON(http.StatusOK, cluster)
}

// createConfigurationClusterEmbedded creates a new cluster in the configuration with its certificate embedded
// @Summary Create a new cluster with its certificates embedded
// @Description Create a new cluster for which the TLS certificate is given as embedded (raw text)
// @ID post-configuration-cluster-embedded
// @Tags Configuration
// @Accept json
// @Produce application/json
// @Param name path string true "the name of the cluster in the configuration"
// @Param body body context.ParamClusterCertificateEmbedded true "the definition of the cluster"
// @Success 200 {object} context.NamedCluster
// @Success 400 {object} echo.HTTPError
// @Success 409 {object} echo.HTTPError
// @Router /api/v1/configuration/clusters/{name}/certificate-embedded [post]
func createConfigurationClusterEmbedded(e echo.Context) error {

	name := e.Param("name")

	// Ensure that the cluster does not already exist
	cluster, err := context.GetKubeCluster(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if cluster != nil {
		return echo.NewHTTPError(http.StatusConflict, fmt.Errorf("the cluster %s already exist", name))
	}

	// Parse the cluster information
	clusterEmbedded := new(context.ParamClusterCertificateEmbedded)
	if err = e.Bind(clusterEmbedded); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// Add the cluster
	err = context.SetClusterCertificateEmbedded(name, *clusterEmbedded)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// Read the newly created object
	cluster, err = context.GetKubeCluster(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if cluster == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("unable to retrieve the cluster %s from configuration after creation", name))
	}

	return e.JSON(http.StatusOK, cluster)
}

// updateConfigurationClusterEmbedded updates an existing cluster in the configuration with its certificate embedded
// @Summary Update an existing cluster with its certificate embedded
// @Description Update an existing cluster for which the TLS certificate is given embedded (raw text)
// @ID put-configuration-cluster-embedded
// @Tags Configuration
// @Accept json
// @Produce application/json
// @Param name path string true "the name of the cluster in the configuration"
// @Param body body context.ParamClusterCertificateEmbedded true "the definition of the cluster"
// @Success 200 {object} context.NamedCluster
// @Success 404 {object} echo.HTTPError
// @Success 409 {object} echo.HTTPError
// @Router /api/v1/configuration/clusters/{name}/certificate-embedded [put]
func updateConfigurationClusterEmbedded(e echo.Context) error {

	name := e.Param("name")

	// Ensure that the cluster does not already exist
	cluster, err := context.GetKubeCluster(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if cluster == nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("the cluster %s does not exist", name))
	}

	// Parse the cluster information
	clusterEmbedded := new(context.ParamClusterCertificateEmbedded)
	if err = e.Bind(clusterEmbedded); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// Add the cluster
	err = context.SetClusterCertificateEmbedded(name, *clusterEmbedded)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// Read the newly created object
	cluster, err = context.GetKubeCluster(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if cluster == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("unable to retrieve the cluster %s from configuration after update", name))
	}

	return e.JSON(http.StatusOK, cluster)
}

// getConfigurationContexts generates a JSON representation of the contexts
// @Summary Retrieve the contexts
// @Description get the contexts
// @ID get-configuration-contexts
// @Tags Configuration
// @Produce application/json
// @Success 200 {array} context.NamedContext
// @Router /api/v1/configuration/contexts/ [get]
func getConfigurationContexts(e echo.Context) error {

	contexts, err := context.GetKubeContexts()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return e.JSON(http.StatusOK, contexts)
}

// createConfigurationContext creates a new context in the configuration
// @Summary Create a new context
// @Description Create a new context: user and cluster, with an optional namespace
// @ID post-configuration-context
// @Tags Configuration
// @Accept json
// @Produce application/json
// @Param name path string true "the name of the context in the configuration"
// @Param body body context.ParamContext true "the definition of the context"
// @Success 200 {object} context.NamedContext
// @Success 400 {object} echo.HTTPError
// @Success 409 {object} echo.HTTPError
// @Router /api/v1/configuration/contexts/{name} [post]
func createConfigurationContext(e echo.Context) error {

	name := e.Param("name")

	// Ensure that the context does not already exist
	kubeContext, err := context.GetKubeContext(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if kubeContext != nil {
		return echo.NewHTTPError(http.StatusConflict, fmt.Errorf("the context %s already exist", name))
	}

	// Parse the cluster information
	contextParam := new(context.ParamContext)
	if err = e.Bind(contextParam); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// Add the context
	err = context.SetContext(name, *contextParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// Read the newly created object
	kubeContext, err = context.GetKubeContext(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if kubeContext == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("unable to retrieve the context %s from configuration after creation", name))
	}

	return e.JSON(http.StatusOK, kubeContext)
}

// updateConfigurationContext updates an existing context in the configuration
// @Summary Update an existing context
// @Description Update an existing cluster: user and cluster, with an optional namespace
// @ID put-configuration-context
// @Tags Configuration
// @Accept json
// @Produce application/json
// @Param name path string true "the name of the context in the configuration"
// @Param body body context.ParamContext true "the definition of the context"
// @Success 200 {object} context.NamedCluster
// @Success 404 {object} echo.HTTPError
// @Success 409 {object} echo.HTTPError
// @Router /api/v1/configuration/contexts/{name} [put]
func updateConfigurationContext(e echo.Context) error {

	name := e.Param("name")

	// Ensure that the cluster does not already exist
	kubeContext, err := context.GetKubeContext(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if kubeContext == nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("the context %s does not exist", name))
	}

	// Parse the cluster information
	contextParam := new(context.ParamContext)
	if err = e.Bind(contextParam); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// Add the cluster
	err = context.SetContext(name, *contextParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// Read the newly created object
	kubeContext, err = context.GetKubeContext(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if kubeContext == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("unable to retrieve the context %s from configuration after update", name))
	}

	return e.JSON(http.StatusOK, kubeContext)
}
