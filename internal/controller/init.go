package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"github.com/twuillemin/kuboxy/docs"
)

//go:generate go run gen/gen_events_controller_all.go
//go:generate go run gen/gen_events_controller_cluster.go
//go:generate go run gen/gen_events_controller_namespace.go
//go:generate go run gen/gen_labels_controller.go
//go:generate go run gen/gen_objects_controller_cluster.go
//go:generate go run gen/gen_objects_controller_cluster_metrics.go
//go:generate go run gen/gen_objects_controller_namespace.go
//go:generate go run gen/gen_objects_controller_namespace_metrics.go

// @title Kubernetes Proxy
// @version 0.3.0
// @description A single proxy for multiple Kubernetes clusters

// @contact.name Thomas Wuillemin
// @contact.url http://www.wuillemin.net/
// @contact.email thomas.wuillemin@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// RegisterControllers registers all the controller of the application within the given configuration
func RegisterControllers(e *echo.Echo) {

	// Add swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Programmatically set swagger info
	docs.SwaggerInfo.Title = "Kuboxy Service"
	docs.SwaggerInfo.Description = "The Kubernetes Proxy."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "api/v1"

	// For generating swagger info in the project
	// "go get github.com/swaggo/swag/cmd/swag" for generating swag executable
	// "swag init" in the root of the project

	registerConfigurationController(e)
	registerObjectClusterControllers(e)
	registerObjectClusterMetricsControllers(e)
	registerObjectNamespaceControllers(e)
	registerObjectNamespaceMetricsControllers(e)
	registerLabelsController(e)
	registerSummaryControllers(e)
	registerSearchControllers(e)
}

// RegisterEventWebSocketController register the controller for the websocket dedicated to events
func RegisterEventWebSocketController(ews *echo.Echo) {
	ews.GET("events", getEventsByWebSocket)
}
