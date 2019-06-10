// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates event_search_namespace.go. It can be invoked by running
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

	f, err := os.Create("search_namespace.go")
	die(err)
	defer f.Close()

	builderTemplate.Execute(
		f,
		struct {
			Timestamp         time.Time
			ObjectDefinitions []types.ObjectDefinition
		}{
			Timestamp:         time.Now(),
			ObjectDefinitions: types.NamespaceObjectDefinitions,
		})
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var builderTemplate = template.Must(template.New("").Parse(`// Package search regroups the functions to search any object in the cluster 
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_search_namespace.go at {{ .Timestamp }}
package search

import (
	"github.com/twuillemin/kuboxy/pkg/provider"
	"github.com/twuillemin/kuboxy/pkg/types"
)

func searchNamespaceObjects(contextName string, searchParameter *preparedParameter, results map[types.ObjectType][]interface{}) error {
{{ range .ObjectDefinitions }}
	if _, ok := searchParameter.objectTypes[types.{{ .Name }}]; ok {
		{{ .PluralVariable }}, err := provider.Get{{ .Plural }}(contextName, "")
		if err != nil {
			return err
		}
		{{ .Variable }}Results := make([]interface{},0,0)
		for _, {{ .Variable }} := range {{ .PluralVariable }} {
			if isValidNamespaceObject({{ .Variable }}.ObjectMeta, searchParameter) {
				{{ .Variable }}Results = append({{ .Variable }}Results, {{ .Variable }} )
			}
		}
		results[types.{{ .Name }}]={{ .Variable }}Results
	}
{{ end }}
	return nil
}
`))
