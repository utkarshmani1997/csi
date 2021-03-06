/*
Copyright 2019 The OpenEBS Authors

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
	time "time"

	corev1alpha1 "github.com/openebs/csi/pkg/apis/openebs.io/core/v1alpha1"
	internalclientset "github.com/openebs/csi/pkg/generated/clientset/core/internalclientset"
	internalinterfaces "github.com/openebs/csi/pkg/generated/informer/core/externalversions/internalinterfaces"
	v1alpha1 "github.com/openebs/csi/pkg/generated/lister/core/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CSIVolumeInformer provides access to a shared informer and lister for
// CSIVolumes.
type CSIVolumeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CSIVolumeLister
}

type cSIVolumeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCSIVolumeInformer constructs a new informer for CSIVolume type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCSIVolumeInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCSIVolumeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCSIVolumeInformer constructs a new informer for CSIVolume type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCSIVolumeInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OpenebsV1alpha1().CSIVolumes(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OpenebsV1alpha1().CSIVolumes(namespace).Watch(options)
			},
		},
		&corev1alpha1.CSIVolume{},
		resyncPeriod,
		indexers,
	)
}

func (f *cSIVolumeInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCSIVolumeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cSIVolumeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1alpha1.CSIVolume{}, f.defaultInformer)
}

func (f *cSIVolumeInformer) Lister() v1alpha1.CSIVolumeLister {
	return v1alpha1.NewCSIVolumeLister(f.Informer().GetIndexer())
}
