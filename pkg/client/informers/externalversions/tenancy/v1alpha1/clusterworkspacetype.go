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
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v3"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"

	tenancyv1alpha1 "github.com/kcp-dev/kcp/pkg/apis/tenancy/v1alpha1"
	scopedclientset "github.com/kcp-dev/kcp/pkg/client/clientset/versioned"
	clientset "github.com/kcp-dev/kcp/pkg/client/clientset/versioned/cluster"
	"github.com/kcp-dev/kcp/pkg/client/informers/externalversions/internalinterfaces"
	tenancyv1alpha1listers "github.com/kcp-dev/kcp/pkg/client/listers/tenancy/v1alpha1"
)

// ClusterWorkspaceTypeClusterInformer provides access to a shared informer and lister for
// ClusterWorkspaceTypes.
type ClusterWorkspaceTypeClusterInformer interface {
	Cluster(logicalcluster.Name) ClusterWorkspaceTypeInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() tenancyv1alpha1listers.ClusterWorkspaceTypeClusterLister
}

type clusterWorkspaceTypeClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterWorkspaceTypeClusterInformer constructs a new informer for ClusterWorkspaceType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterWorkspaceTypeClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredClusterWorkspaceTypeClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterWorkspaceTypeClusterInformer constructs a new informer for ClusterWorkspaceType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterWorkspaceTypeClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TenancyV1alpha1().ClusterWorkspaceTypes().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TenancyV1alpha1().ClusterWorkspaceTypes().Watch(context.TODO(), options)
			},
		},
		&tenancyv1alpha1.ClusterWorkspaceType{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterWorkspaceTypeClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredClusterWorkspaceTypeClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName: kcpcache.ClusterIndexFunc,
	},
		f.tweakListOptions,
	)
}

func (f *clusterWorkspaceTypeClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&tenancyv1alpha1.ClusterWorkspaceType{}, f.defaultInformer)
}

func (f *clusterWorkspaceTypeClusterInformer) Lister() tenancyv1alpha1listers.ClusterWorkspaceTypeClusterLister {
	return tenancyv1alpha1listers.NewClusterWorkspaceTypeClusterLister(f.Informer().GetIndexer())
}

// ClusterWorkspaceTypeInformer provides access to a shared informer and lister for
// ClusterWorkspaceTypes.
type ClusterWorkspaceTypeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() tenancyv1alpha1listers.ClusterWorkspaceTypeLister
}

func (f *clusterWorkspaceTypeClusterInformer) Cluster(clusterName logicalcluster.Name) ClusterWorkspaceTypeInformer {
	return &clusterWorkspaceTypeInformer{
		informer: f.Informer().Cluster(clusterName),
		lister:   f.Lister().Cluster(clusterName),
	}
}

type clusterWorkspaceTypeInformer struct {
	informer cache.SharedIndexInformer
	lister   tenancyv1alpha1listers.ClusterWorkspaceTypeLister
}

func (f *clusterWorkspaceTypeInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *clusterWorkspaceTypeInformer) Lister() tenancyv1alpha1listers.ClusterWorkspaceTypeLister {
	return f.lister
}

type clusterWorkspaceTypeScopedInformer struct {
	factory          internalinterfaces.SharedScopedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

func (f *clusterWorkspaceTypeScopedInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&tenancyv1alpha1.ClusterWorkspaceType{}, f.defaultInformer)
}

func (f *clusterWorkspaceTypeScopedInformer) Lister() tenancyv1alpha1listers.ClusterWorkspaceTypeLister {
	return tenancyv1alpha1listers.NewClusterWorkspaceTypeLister(f.Informer().GetIndexer())
}

// NewClusterWorkspaceTypeInformer constructs a new informer for ClusterWorkspaceType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterWorkspaceTypeInformer(client scopedclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterWorkspaceTypeInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterWorkspaceTypeInformer constructs a new informer for ClusterWorkspaceType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterWorkspaceTypeInformer(client scopedclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TenancyV1alpha1().ClusterWorkspaceTypes().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TenancyV1alpha1().ClusterWorkspaceTypes().Watch(context.TODO(), options)
			},
		},
		&tenancyv1alpha1.ClusterWorkspaceType{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterWorkspaceTypeScopedInformer) defaultInformer(client scopedclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterWorkspaceTypeInformer(client, resyncPeriod, cache.Indexers{}, f.tweakListOptions)
}
