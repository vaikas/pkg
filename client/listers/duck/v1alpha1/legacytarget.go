/*
Copyright 2018 The Knative Authors

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
	v1alpha1 "github.com/knative/pkg/apis/duck/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LegacyTargetLister helps list LegacyTargets.
type LegacyTargetLister interface {
	// List lists all LegacyTargets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.LegacyTarget, err error)
	// LegacyTargets returns an object that can list and get LegacyTargets.
	LegacyTargets(namespace string) LegacyTargetNamespaceLister
	LegacyTargetListerExpansion
}

// legacyTargetLister implements the LegacyTargetLister interface.
type legacyTargetLister struct {
	indexer cache.Indexer
}

// NewLegacyTargetLister returns a new LegacyTargetLister.
func NewLegacyTargetLister(indexer cache.Indexer) LegacyTargetLister {
	return &legacyTargetLister{indexer: indexer}
}

// List lists all LegacyTargets in the indexer.
func (s *legacyTargetLister) List(selector labels.Selector) (ret []*v1alpha1.LegacyTarget, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LegacyTarget))
	})
	return ret, err
}

// LegacyTargets returns an object that can list and get LegacyTargets.
func (s *legacyTargetLister) LegacyTargets(namespace string) LegacyTargetNamespaceLister {
	return legacyTargetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LegacyTargetNamespaceLister helps list and get LegacyTargets.
type LegacyTargetNamespaceLister interface {
	// List lists all LegacyTargets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.LegacyTarget, err error)
	// Get retrieves the LegacyTarget from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.LegacyTarget, error)
	LegacyTargetNamespaceListerExpansion
}

// legacyTargetNamespaceLister implements the LegacyTargetNamespaceLister
// interface.
type legacyTargetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LegacyTargets in the indexer for a given namespace.
func (s legacyTargetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LegacyTarget, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LegacyTarget))
	})
	return ret, err
}

// Get retrieves the LegacyTarget from the indexer for a given namespace and name.
func (s legacyTargetNamespaceLister) Get(name string) (*v1alpha1.LegacyTarget, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("legacytarget"), name)
	}
	return obj.(*v1alpha1.LegacyTarget), nil
}
