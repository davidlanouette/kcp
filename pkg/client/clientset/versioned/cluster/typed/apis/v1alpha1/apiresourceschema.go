//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	apisv1alpha1 "github.com/kcp-dev/kcp/pkg/apis/apis/v1alpha1"
	apisv1alpha1client "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/typed/apis/v1alpha1"
)

// APIResourceSchemasClusterGetter has a method to return a APIResourceSchemaClusterInterface.
// A group's cluster client should implement this interface.
type APIResourceSchemasClusterGetter interface {
	APIResourceSchemas() APIResourceSchemaClusterInterface
}

// APIResourceSchemaClusterInterface can operate on APIResourceSchemas across all clusters,
// or scope down to one cluster and return a apisv1alpha1client.APIResourceSchemaInterface.
type APIResourceSchemaClusterInterface interface {
	Cluster(logicalcluster.Path) apisv1alpha1client.APIResourceSchemaInterface
	List(ctx context.Context, opts metav1.ListOptions) (*apisv1alpha1.APIResourceSchemaList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type aPIResourceSchemasClusterInterface struct {
	clientCache kcpclient.Cache[*apisv1alpha1client.ApisV1alpha1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *aPIResourceSchemasClusterInterface) Cluster(clusterPath logicalcluster.Path) apisv1alpha1client.APIResourceSchemaInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).APIResourceSchemas()
}

// List returns the entire collection of all APIResourceSchemas across all clusters.
func (c *aPIResourceSchemasClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*apisv1alpha1.APIResourceSchemaList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).APIResourceSchemas().List(ctx, opts)
}

// Watch begins to watch all APIResourceSchemas across all clusters.
func (c *aPIResourceSchemasClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).APIResourceSchemas().Watch(ctx, opts)
}
