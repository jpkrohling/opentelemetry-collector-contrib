// Copyright 2019 OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kube

import (
	"sort"

	"go.uber.org/zap"
	api_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sprocessor/observability"
)

// OwnerProvider allows to dynamically assign constructor
type OwnerProvider func(
	logger *zap.Logger,
	client *kubernetes.Clientset,
	labelSelector labels.Selector,
	fieldSelector fields.Selector,
	namespace string,
) (OwnerAPI, error)

// ObjectOwner keeps single entry
type ObjectOwner struct {
	UID       types.UID
	ownerUIDs []types.UID
	namespace string
	kind      string
	name      string
}

// OwnerAPI describes functions that could allow retrieving owner info
type OwnerAPI interface {
	GetOwners(pod *api_v1.Pod) []*ObjectOwner
	GetServices(pod *api_v1.Pod) []string
	Start()
	Stop()
}

// OwnerCache is a simple structure which aids querying for owners
type OwnerCache struct {
	objectOwners map[string]*ObjectOwner
	podServices  map[string][]string

	clientset *kubernetes.Clientset
	logger    *zap.Logger

	stopCh    chan struct{}
	informers []cache.SharedIndexInformer
}

// Start runs the informers
func (op *OwnerCache) Start() {
	op.logger.Info("Staring K8S resource informers", zap.Int("#infomers", len(op.informers)))
	for _, informer := range op.informers {
		go informer.Run(op.stopCh)
	}
}

// Stop shutdowns the informers
func (op *OwnerCache) Stop() {
	close(op.stopCh)
}

func newOwnerProvider(
	logger *zap.Logger,
	clientset *kubernetes.Clientset,
	labelSelector labels.Selector,
	fieldSelector fields.Selector,
	namespace string) (OwnerAPI, error) {
	ownerCache := OwnerCache{}
	ownerCache.objectOwners = map[string]*ObjectOwner{}
	ownerCache.podServices = map[string][]string{}
	ownerCache.clientset = clientset
	ownerCache.logger = logger

	factory := informers.NewSharedInformerFactoryWithOptions(clientset, watchSyncPeriod,
		informers.WithNamespace(namespace),
		informers.WithTweakListOptions(func(opts *meta_v1.ListOptions) {
			opts.LabelSelector = labelSelector.String()
			opts.FieldSelector = fieldSelector.String()
		}))

	ownerCache.addOwnerInformer("ReplicaSet",
		factory.Extensions().V1beta1().ReplicaSets().Informer(),
		ownerCache.cacheObject,
		ownerCache.deleteObject)

	ownerCache.addOwnerInformer("Deployment",
		factory.Extensions().V1beta1().Deployments().Informer(),
		ownerCache.cacheObject,
		ownerCache.deleteObject)

	ownerCache.addOwnerInformer("StatefulSet",
		factory.Apps().V1().StatefulSets().Informer(),
		ownerCache.cacheObject,
		ownerCache.deleteObject)

	ownerCache.addOwnerInformer("Endpoint",
		factory.Core().V1().Endpoints().Informer(),
		ownerCache.cacheEndpoint,
		ownerCache.deleteEndpoint)

	return &ownerCache, nil
}

func (op *OwnerCache) addOwnerInformer(
	kind string,
	informer cache.SharedIndexInformer,
	cacheFunc func(kind string, obj interface{}),
	deleteFunc func(obj interface{})) {
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			observability.RecordOtherAdded()
			cacheFunc(kind, obj)
		},
		UpdateFunc: func(_, obj interface{}) {
			observability.RecordOtherUpdated()
			cacheFunc(kind, obj)
		},
		DeleteFunc: func(obj interface{}) {
			observability.RecordOtherDeleted()
			deleteFunc(obj)
		},
	})

	op.informers = append(op.informers, informer)
}

func (op *OwnerCache) deleteObject(obj interface{}) {
	delete(op.objectOwners, string(obj.(meta_v1.Object).GetUID()))
}

func (op *OwnerCache) cacheObject(kind string, obj interface{}) {
	meta := obj.(meta_v1.Object)

	oo := ObjectOwner{
		UID:       meta.GetUID(),
		namespace: meta.GetNamespace(),
		ownerUIDs: []types.UID{},
		kind:      kind,
		name:      meta.GetName(),
	}
	for _, or := range meta.GetOwnerReferences() {
		oo.ownerUIDs = append(oo.ownerUIDs, or.UID)
	}

	op.objectOwners[string(oo.UID)] = &oo
}

func (op *OwnerCache) addEndpointToPod(pod string, endpoint string) {
	services := op.podServices[pod]

	for _, it := range services {
		if it == endpoint {
			return
		}
	}

	services = append(services, endpoint)
	sort.Strings(services)
	op.podServices[pod] = services
}

func (op *OwnerCache) deleteEndpointFromPod(pod string, endpoint string) {
	services := op.podServices[pod]
	newServices := []string{}

	for _, it := range services {
		if it != endpoint {
			newServices = append(newServices, it)
		}
	}

	op.podServices[pod] = newServices
}

func (op *OwnerCache) genericEndpointOp(obj interface{}, endpointFunc func(pod string, endpoint string)) {
	ep := obj.(*api_v1.Endpoints)

	for _, it := range ep.Subsets {
		for _, addr := range it.Addresses {
			if addr.TargetRef != nil && addr.TargetRef.Kind == "Pod" {
				endpointFunc(addr.TargetRef.Name, ep.Name)
			}
		}
		for _, addr := range it.NotReadyAddresses {
			if addr.TargetRef != nil && addr.TargetRef.Kind == "Pod" {
				endpointFunc(addr.TargetRef.Name, ep.Name)
			}
		}
	}
}

func (op *OwnerCache) deleteEndpoint(obj interface{}) {
	op.genericEndpointOp(obj, op.deleteEndpointFromPod)
}

func (op *OwnerCache) cacheEndpoint(kind string, obj interface{}) {
	op.genericEndpointOp(obj, op.addEndpointToPod)
}

// GetServices returns a slice with matched services - in case no services are found, it returns an empty slice
func (op *OwnerCache) GetServices(pod *api_v1.Pod) []string {
	if oo, found := op.podServices[pod.Name]; found {
		return oo
	}
	return []string{}
}

// GetOwners goes through the cached data and assigns relevant metadata for pod
func (op *OwnerCache) GetOwners(pod *api_v1.Pod) []*ObjectOwner {
	objectOwners := []*ObjectOwner{}

	visited := map[types.UID]bool{}
	queue := []types.UID{}

	for _, or := range pod.OwnerReferences {
		if _, uidVisited := visited[or.UID]; !uidVisited {
			queue = append(queue, or.UID)
			visited[or.UID] = true
		}
	}

	for len(queue) > 0 {
		uid := queue[0]
		queue = queue[1:]
		if oo, found := op.objectOwners[string(uid)]; found {
			objectOwners = append(objectOwners, oo)

			for _, ownerUID := range oo.ownerUIDs {
				if _, uidVisited := visited[ownerUID]; !uidVisited {
					queue = append(queue, ownerUID)
					visited[ownerUID] = true
				}
			}
		}
	}

	return objectOwners
}
