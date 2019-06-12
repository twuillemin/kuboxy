// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates event_namespace.go. It can be invoked by running
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

	objects := make([]types.ObjectDefinition, 0, len(types.NamespaceObjectDefinitions)+len(types.NamespaceMetricsObjectDefinitions))
	for _, obj := range types.NamespaceObjectDefinitions {
		objects = append(objects, obj)
	}
	for _, obj := range types.NamespaceMetricsObjectDefinitions {
		objects = append(objects, obj)
	}

	f, err := os.Create("event_namespace.go")
	die(err)
	defer f.Close()

	builderTemplate.Execute(
		f,
		struct {
			Timestamp         time.Time
			ObjectDefinitions []types.ObjectDefinition
		}{
			Timestamp:         time.Now(),
			ObjectDefinitions: objects,
		})
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var builderTemplate = template.Must(template.New("").Parse(`// Package event regroups all the definitions and functions to receive events from a cluster 
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_event_namespace.go at {{ .Timestamp }}
package event

import (
	"github.com/twuillemin/kuboxy/pkg/context"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metricsv1beta1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
)
{{ range .ObjectDefinitions }}
// Get{{ .Plural }} returns the list of all {{ .Plural }} known by the EventReceiver. The returned list
// is a copy and could be freely modified bt the caller
func Get{{ .Plural }}(contextName string, namespace string) []{{ .FullName }} {

	ctxReceiver, ok := contextReceivers[contextName]
	if !ok {
		return nil
	}

	nsReceiver, ok := ctxReceiver.namespaceReceivers[namespace]
	if !ok {
		return nil
	}

	receiver := nsReceiver.{{ .Variable }}EventReceiver
	if receiver == nil {
		return nil
	}

	return receiver.get{{ .Plural }}()
}

// Add{{ .Name }}EventClient adds a new client that will received events
func Add{{ .Name }}EventClient(contextName string, namespace string, client chan {{ .Name }}Event) error {

	// Get the receiver or create it
	ctxReceiver, ok := contextReceivers[contextName]
	if !ok {

		clientset, err := context.GetClientset(contextName)
		if err != nil {
			return err
		}

		metrics, err := context.GetMetrics(contextName)
		if err != nil {
			return err
		}

		ctxReceiver = &contextReceiver{
			clientset: clientset,
			metrics:   metrics,
		}
		
		contextReceivers[contextName] = ctxReceiver
	}

	nsReceiver, ok := ctxReceiver.namespaceReceivers[namespace]
	if !ok {
		nsReceiver = &namespaceReceiver{}
		ctxReceiver.namespaceReceivers[namespace] = nsReceiver
	}

	// If the event are not received, create a new event receiver
	receiver := nsReceiver.{{ .Variable }}EventReceiver
	if receiver == nil {
		{{ if or .IsClusterFamily .IsNamespaceFamily }}
		receiver = new{{ .Name }}EventReceiver(ctxReceiver.clientset, namespace)
		{{ else }}
		receiver = new{{ .Name }}EventReceiver(ctxReceiver.metrics, namespace)
		{{ end }}
		nsReceiver.{{ .Variable }}EventReceiver = receiver
	}

	// Add the client
	receiver.addClient(client)

	return nil
}

// Remove{{ .Name }}EventClient removes a client from receiving events
func Remove{{ .Name }}EventClient(contextName string, namespace string, client chan {{ .Name }}Event) {

	// Get the context receiver
	ctxReceiver, ok := contextReceivers[contextName]
	if !ok {
		return
	}

	// Get the namespace receiver
	nsReceiver, ok := ctxReceiver.namespaceReceivers[namespace]
	if !ok {
		return
	}

	// Get the receiver
	receiver := nsReceiver.{{ .Variable }}EventReceiver
	if receiver == nil {
		return
	}

	// Remove the client
	receiver.removeClient(client)
	
	// If no more client, stop receiving event
	if len(receiver.clients) == 0 {
		receiver.stop()
		nsReceiver.{{ .Variable }}EventReceiver = nil
	}
}
{{ end }}`))
