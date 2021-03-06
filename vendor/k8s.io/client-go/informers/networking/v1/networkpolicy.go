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

package v1

import (
	time "time"

	networking_v1 "k8s.io/api/networking/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	kubernetes "k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/listers/networking/v1"
	cache "k8s.io/client-go/tools/cache"
)

// NetworkPolicyInformer provides access to a shared informer and lister for
// NetworkPolicies.
type NetworkPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.NetworkPolicyLister
}

type networkPolicyInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewNetworkPolicyInformer constructs a new informer for NetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNetworkPolicyInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.NetworkingV1().NetworkPolicies(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.NetworkingV1().NetworkPolicies(namespace).Watch(options)
			},
		},
		&networking_v1.NetworkPolicy{},
		resyncPeriod,
		indexers,
	)
}

func defaultNetworkPolicyInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewNetworkPolicyInformer(client, meta_v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *networkPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networking_v1.NetworkPolicy{}, defaultNetworkPolicyInformer)
}

func (f *networkPolicyInformer) Lister() v1.NetworkPolicyLister {
	return v1.NewNetworkPolicyLister(f.Informer().GetIndexer())
}
