/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "harmonycloud.cn/stellaris/pkg/apis/multicluster/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MultiClusterResourceBindingLister helps list MultiClusterResourceBindings.
// All objects returned here must be treated as read-only.
type MultiClusterResourceBindingLister interface {
	// List lists all MultiClusterResourceBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MultiClusterResourceBinding, err error)
	// MultiClusterResourceBindings returns an object that can list and get MultiClusterResourceBindings.
	MultiClusterResourceBindings(namespace string) MultiClusterResourceBindingNamespaceLister
	MultiClusterResourceBindingListerExpansion
}

// multiClusterResourceBindingLister implements the MultiClusterResourceBindingLister interface.
type multiClusterResourceBindingLister struct {
	indexer cache.Indexer
}

// NewMultiClusterResourceBindingLister returns a new MultiClusterResourceBindingLister.
func NewMultiClusterResourceBindingLister(indexer cache.Indexer) MultiClusterResourceBindingLister {
	return &multiClusterResourceBindingLister{indexer: indexer}
}

// List lists all MultiClusterResourceBindings in the indexer.
func (s *multiClusterResourceBindingLister) List(selector labels.Selector) (ret []*v1alpha1.MultiClusterResourceBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MultiClusterResourceBinding))
	})
	return ret, err
}

// MultiClusterResourceBindings returns an object that can list and get MultiClusterResourceBindings.
func (s *multiClusterResourceBindingLister) MultiClusterResourceBindings(namespace string) MultiClusterResourceBindingNamespaceLister {
	return multiClusterResourceBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MultiClusterResourceBindingNamespaceLister helps list and get MultiClusterResourceBindings.
// All objects returned here must be treated as read-only.
type MultiClusterResourceBindingNamespaceLister interface {
	// List lists all MultiClusterResourceBindings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MultiClusterResourceBinding, err error)
	// Get retrieves the MultiClusterResourceBinding from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.MultiClusterResourceBinding, error)
	MultiClusterResourceBindingNamespaceListerExpansion
}

// multiClusterResourceBindingNamespaceLister implements the MultiClusterResourceBindingNamespaceLister
// interface.
type multiClusterResourceBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MultiClusterResourceBindings in the indexer for a given namespace.
func (s multiClusterResourceBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MultiClusterResourceBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MultiClusterResourceBinding))
	})
	return ret, err
}

// Get retrieves the MultiClusterResourceBinding from the indexer for a given namespace and name.
func (s multiClusterResourceBindingNamespaceLister) Get(name string) (*v1alpha1.MultiClusterResourceBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("multiclusterresourcebinding"), name)
	}
	return obj.(*v1alpha1.MultiClusterResourceBinding), nil
}
