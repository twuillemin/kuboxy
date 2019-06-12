// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates provider_cluster_metrics.go. It can be invoked by running
// go generate

package main

import (
	"log"
	"os"
	"text/template"
	"time"

	"github.com/twuillemin/kuboxy/pkg/types"
)

func main() {

	f, err := os.Create("provider_cluster_metrics.go")
	die(err)
	defer f.Close()

	builderTemplate.Execute(
		f,
		struct {
			Timestamp         time.Time
			ObjectDefinitions []types.ObjectDefinition
		}{
			Timestamp:         time.Now(),
			ObjectDefinitions: types.ClusterMetricsObjectDefinitions,
		})
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var builderTemplate = template.Must(template.New("").Parse(`// Package provider regroups all the basic CRUD functions to access the objects in the cluster
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_provider_cluster_metrics.go at {{ .Timestamp }}
package provider

import (
	"github.com/twuillemin/kuboxy/pkg/connector"
	"github.com/twuillemin/kuboxy/pkg/context"
	"github.com/twuillemin/kuboxy/pkg/event"
	metricsv1beta1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
)
{{ range .ObjectDefinitions }}
// Get{{ .Plural }} returns all the {{ .Name }}.
func Get{{ .Plural }}(contextName string) ([]{{ .FullName }}, error) {

	metrics, err := context.GetMetrics(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.Get{{ .Plural }}(contextName); results != nil {
		return results, nil
	}

	return connector.Get{{ .Plural }}(metrics)
}

// Get{{ .Name }} returns the {{ .Name }} by its name.
func Get{{ .Name }}(contextName string, name string) (*{{ .FullName }}, error) {

	metrics, err := context.GetMetrics(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.Get{{ .Plural }}(contextName); results != nil {
		for _, {{ .Variable }} := range results {
			if {{ .Variable }}.Name == name {
				return &{{ .Variable }}, nil
			}
		}
		return nil, nil
	}

	return connector.Get{{ .Name }}(metrics, name)
}
{{ end }}`))
