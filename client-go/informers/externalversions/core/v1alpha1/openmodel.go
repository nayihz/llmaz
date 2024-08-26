/*
Copyright 2024.

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
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	corev1alpha1 "github.com/inftyai/llmaz/api/core/v1alpha1"
	versioned "github.com/inftyai/llmaz/client-go/clientset/versioned"
	internalinterfaces "github.com/inftyai/llmaz/client-go/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/inftyai/llmaz/client-go/listers/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// OpenModelInformer provides access to a shared informer and lister for
// OpenModels.
type OpenModelInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.OpenModelLister
}

type openModelInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewOpenModelInformer constructs a new informer for OpenModel type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewOpenModelInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredOpenModelInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredOpenModelInformer constructs a new informer for OpenModel type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredOpenModelInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LlmazV1alpha1().OpenModels(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LlmazV1alpha1().OpenModels(namespace).Watch(context.TODO(), options)
			},
		},
		&corev1alpha1.OpenModel{},
		resyncPeriod,
		indexers,
	)
}

func (f *openModelInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredOpenModelInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *openModelInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1alpha1.OpenModel{}, f.defaultInformer)
}

func (f *openModelInformer) Lister() v1alpha1.OpenModelLister {
	return v1alpha1.NewOpenModelLister(f.Informer().GetIndexer())
}
