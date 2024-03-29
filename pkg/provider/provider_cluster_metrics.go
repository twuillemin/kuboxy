// Package provider regroups all the basic CRUD functions to access the objects in the cluster
//
// Code generated by go generate; DO NOT EDIT.
//
// This file was generated by gen_provider_cluster_metrics.go at 2019-06-17 12:12:21.012007208 +0300 EEST m=+0.000977640
package provider

import (
	"github.com/twuillemin/kuboxy/pkg/connector"
	"github.com/twuillemin/kuboxy/pkg/context"
	"github.com/twuillemin/kuboxy/pkg/event"
	metricsv1beta1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

// GetNodeMetricses returns all the NodeMetrics.
func GetNodeMetricses(contextName string) ([]metricsv1beta1.NodeMetrics, error) {

	metrics, err := context.GetMetrics(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetNodeMetricses(contextName); results != nil {
		return results, nil
	}

	return connector.GetNodeMetricses(metrics)
}

// GetNodeMetrics returns the NodeMetrics by its name.
func GetNodeMetrics(contextName string, name string) (*metricsv1beta1.NodeMetrics, error) {

	metrics, err := context.GetMetrics(contextName)
	if err != nil {
		return nil, err
	}

	if results := event.GetNodeMetricses(contextName); results != nil {
		for _, nodeMetrics := range results {
			if nodeMetrics.Name == name {
				return &nodeMetrics, nil
			}
		}
		return nil, nil
	}

	return connector.GetNodeMetrics(metrics, name)
}
