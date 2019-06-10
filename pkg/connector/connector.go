// Package connector regroups all the basic CRUD functions to access the objects in the cluster
package connector

//go:generate go run gen/gen_connector_cluster.go
//go:generate go run gen/gen_connector_cluster_metrics.go
//go:generate go run gen/gen_connector_namespace.go
//go:generate go run gen/gen_connector_namespace_metrics.go

import (
	core "k8s.io/api/core/v1"
)

// getValidNameSpace ensures that a valid name space is returned. If the given namespace is empty, then the
// default namespace is used.
func getValidNameSpace(namespace string) string {
	// Ensure there is a namespace
	if len(namespace) == 0 {
		return core.NamespaceDefault
	}
	return namespace
}
