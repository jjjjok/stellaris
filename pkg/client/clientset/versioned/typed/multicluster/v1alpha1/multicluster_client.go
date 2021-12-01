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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "harmonycloud.cn/multi-cluster-manager/pkg/apis/multicluster/v1alpha1"
	"harmonycloud.cn/multi-cluster-manager/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type MulticlusterV1alpha1Interface interface {
	RESTClient() rest.Interface
	AggregatedResourcesGetter
	ClustersGetter
	ClusterResourcesGetter
	ClusterSetsGetter
	MultiClusterResourcesGetter
	MultiClusterResourceAggregatePoliciesGetter
	MultiClusterResourceAggregateRulesGetter
	NamespaceMappingsGetter
	ResourceAggregatePoliciesGetter
}

// MulticlusterV1alpha1Client is used to interact with features provided by the multicluster.harmonycloud.cn group.
type MulticlusterV1alpha1Client struct {
	restClient rest.Interface
}

func (c *MulticlusterV1alpha1Client) AggregatedResources(namespace string) AggregatedResourceInterface {
	return newAggregatedResources(c, namespace)
}

func (c *MulticlusterV1alpha1Client) Clusters() ClusterInterface {
	return newClusters(c)
}

func (c *MulticlusterV1alpha1Client) ClusterResources(namespace string) ClusterResourceInterface {
	return newClusterResources(c, namespace)
}

func (c *MulticlusterV1alpha1Client) ClusterSets() ClusterSetInterface {
	return newClusterSets(c)
}

func (c *MulticlusterV1alpha1Client) MultiClusterResources(namespace string) MultiClusterResourceInterface {
	return newMultiClusterResources(c, namespace)
}

func (c *MulticlusterV1alpha1Client) MultiClusterResourceAggregatePolicies(namespace string) MultiClusterResourceAggregatePolicyInterface {
	return newMultiClusterResourceAggregatePolicies(c, namespace)
}

func (c *MulticlusterV1alpha1Client) MultiClusterResourceAggregateRules(namespace string) MultiClusterResourceAggregateRuleInterface {
	return newMultiClusterResourceAggregateRules(c, namespace)
}

func (c *MulticlusterV1alpha1Client) NamespaceMappings(namespace string) NamespaceMappingInterface {
	return newNamespaceMappings(c, namespace)
}

func (c *MulticlusterV1alpha1Client) ResourceAggregatePolicies(namespace string) ResourceAggregatePolicyInterface {
	return newResourceAggregatePolicies(c, namespace)
}

// NewForConfig creates a new MulticlusterV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*MulticlusterV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &MulticlusterV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new MulticlusterV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *MulticlusterV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new MulticlusterV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *MulticlusterV1alpha1Client {
	return &MulticlusterV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *MulticlusterV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
