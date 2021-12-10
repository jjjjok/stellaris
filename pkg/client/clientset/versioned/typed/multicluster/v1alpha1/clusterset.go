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
	"context"
	"time"

	v1alpha1 "harmonycloud.cn/stellaris/pkg/apis/multicluster/v1alpha1"
	scheme "harmonycloud.cn/stellaris/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterSetsGetter has a method to return a ClusterSetInterface.
// A group's client should implement this interface.
type ClusterSetsGetter interface {
	ClusterSets() ClusterSetInterface
}

// ClusterSetInterface has methods to work with ClusterSet resources.
type ClusterSetInterface interface {
	Create(ctx context.Context, clusterSet *v1alpha1.ClusterSet, opts v1.CreateOptions) (*v1alpha1.ClusterSet, error)
	Update(ctx context.Context, clusterSet *v1alpha1.ClusterSet, opts v1.UpdateOptions) (*v1alpha1.ClusterSet, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClusterSet, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClusterSetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterSet, err error)
	ClusterSetExpansion
}

// clusterSets implements ClusterSetInterface
type clusterSets struct {
	client rest.Interface
}

// newClusterSets returns a ClusterSets
func newClusterSets(c *MulticlusterV1alpha1Client) *clusterSets {
	return &clusterSets{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterSet, and returns the corresponding clusterSet object, and an error if there is any.
func (c *clusterSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterSet, err error) {
	result = &v1alpha1.ClusterSet{}
	err = c.client.Get().
		Resource("clustersets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterSets that match those selectors.
func (c *clusterSets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterSetList{}
	err = c.client.Get().
		Resource("clustersets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterSets.
func (c *clusterSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clustersets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterSet and creates it.  Returns the server's representation of the clusterSet, and an error, if there is any.
func (c *clusterSets) Create(ctx context.Context, clusterSet *v1alpha1.ClusterSet, opts v1.CreateOptions) (result *v1alpha1.ClusterSet, err error) {
	result = &v1alpha1.ClusterSet{}
	err = c.client.Post().
		Resource("clustersets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterSet).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterSet and updates it. Returns the server's representation of the clusterSet, and an error, if there is any.
func (c *clusterSets) Update(ctx context.Context, clusterSet *v1alpha1.ClusterSet, opts v1.UpdateOptions) (result *v1alpha1.ClusterSet, err error) {
	result = &v1alpha1.ClusterSet{}
	err = c.client.Put().
		Resource("clustersets").
		Name(clusterSet.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterSet).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterSet and deletes it. Returns an error if one occurs.
func (c *clusterSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clustersets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clustersets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterSet.
func (c *clusterSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterSet, err error) {
	result = &v1alpha1.ClusterSet{}
	err = c.client.Patch(pt).
		Resource("clustersets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
