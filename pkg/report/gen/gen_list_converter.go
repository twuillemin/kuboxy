// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates list_converter.go. It can be invoked by running
// go generate

package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

func main() {

	type ObjectType struct {
		Variable string
		Name     string
		FullName string
	}

	objectTypes := []ObjectType{
		{"service", "Service", "core.Service"},
		{"pod", "Pod", "core.Pod"},
		{"deployment", "Deployment", "apps.Deployment"},
	}

	f, err := os.Create("list_converter.go")
	die(err)
	defer f.Close()

	builderTemplate.Execute(
		f,
		struct {
			Timestamp   time.Time
			ObjectTypes []ObjectType
		}{
			Timestamp:   time.Now(),
			ObjectTypes: objectTypes,
		})
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var builderTemplate = template.Must(template.New("").Parse(`// Package report generates a report from the state of a cluster
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_list_converter.go at {{ .Timestamp }}
package report

import (
	apps "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

{{ range .ObjectTypes }}
// {{ .Variable }}sTo{{ .Name }}Ids convert a list of {{ .Name }}s to a list of their ids
func {{ .Variable }}sTo{{ .Name }}Ids(objects []{{ .FullName }}) []types.UID {
	
	if objects == nil {
		return make([]types.UID, 0, 0)
	}
	
	ids := make([]types.UID, 0, len(objects))
	for _, provider := range objects {
		ids = append(ids, provider.UID)
	}

	return ids
}
{{ end }}`))