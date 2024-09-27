/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

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
	v1alpha1 "github.com/apache/rocketmq-externals/rocketmq-knative/source/pkg/apis/sources/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RocketMQSourceLister helps list RocketMQSource.
type RocketMQSourceLister interface {
	// List lists all RocketMQSources in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.RocketMQSource, err error)
	// RocketMQSources returns an object that can list and get RocketMQSources.
	RocketMQSources(namespace string) RocketMQSourceNamespaceLister
	RocketMQSourceListerExpansion
}

// rocketmqSourceLister implements the RocketMQSourceLister interface.
type rocketmqsourceSourceLister struct {
	indexer cache.Indexer
}

// NewRocketMQSourceLister returns a new RocketMQSourceLister.
func NewRocketMQSourceLister(indexer cache.Indexer) RocketMQSourceLister {
	return &rocketmqsourceSourceLister{indexer: indexer}
}

// List lists all RocketMQSources in the indexer.
func (s *rocketmqsourceSourceLister) List(selector labels.Selector) (ret []*v1alpha1.RocketMQSource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RocketMQSource))
	})
	return ret, err
}

// RocketMQSources returns an object that can list and get RocketMQSources.
func (s *rocketmqsourceSourceLister) RocketMQSources(namespace string) RocketMQSourceNamespaceLister {
	return rocketmqSourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RocketMQSourceNamespaceLister helps list and get RocketMQSources.
type RocketMQSourceNamespaceLister interface {
	// List lists all RocketMQSources in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.RocketMQSource, err error)
	// Get retrieves the RocketMQSource from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.RocketMQSource, error)
	RocketMQSourceNamespaceListerExpansion
}

// rocketmqSourceNamespaceLister implements the RocketMQSourceNamespaceLister
// interface.
type rocketmqSourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RocketMQSources in the indexer for a given namespace.
func (s rocketmqSourceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RocketMQSource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RocketMQSource))
	})
	return ret, err
}

// Get retrieves the RocketMQSource from the indexer for a given namespace and name.
func (s rocketmqSourceNamespaceLister) Get(name string) (*v1alpha1.RocketMQSource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("rocketmqsource"), name)
	}
	return obj.(*v1alpha1.RocketMQSource), nil
}
