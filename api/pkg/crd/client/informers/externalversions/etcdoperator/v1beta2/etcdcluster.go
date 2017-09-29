/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package v1beta2

import (
	versioned "github.com/kubermatic/kubermatic/api/pkg/crd/client/clientset/versioned"
	internalinterfaces "github.com/kubermatic/kubermatic/api/pkg/crd/client/informers/externalversions/internalinterfaces"
	v1beta2 "github.com/kubermatic/kubermatic/api/pkg/crd/client/listers/etcdoperator/v1beta2"
	etcdoperator_v1beta2 "github.com/kubermatic/kubermatic/api/pkg/crd/etcdoperator/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// EtcdClusterInformer provides access to a shared informer and lister for
// EtcdClusters.
type EtcdClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta2.EtcdClusterLister
}

type etcdClusterInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewEtcdClusterInformer constructs a new informer for EtcdCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEtcdClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.EtcdoperatorV1beta2().EtcdClusters(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.EtcdoperatorV1beta2().EtcdClusters(namespace).Watch(options)
			},
		},
		&etcdoperator_v1beta2.EtcdCluster{},
		resyncPeriod,
		indexers,
	)
}

func defaultEtcdClusterInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewEtcdClusterInformer(client, v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *etcdClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&etcdoperator_v1beta2.EtcdCluster{}, defaultEtcdClusterInformer)
}

func (f *etcdClusterInformer) Lister() v1beta2.EtcdClusterLister {
	return v1beta2.NewEtcdClusterLister(f.Informer().GetIndexer())
}
