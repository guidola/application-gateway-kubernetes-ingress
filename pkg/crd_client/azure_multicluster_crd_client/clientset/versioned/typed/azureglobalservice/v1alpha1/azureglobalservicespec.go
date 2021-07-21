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

	v1alpha1 "github.com/Azure/application-gateway-kubernetes-ingress/pkg/apis/azureglobalservice/v1alpha1"
	scheme "github.com/Azure/application-gateway-kubernetes-ingress/pkg/crd_client/azure_multicluster_crd_client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AzureGlobalServiceSpecsGetter has a method to return a AzureGlobalServiceSpecInterface.
// A group's client should implement this interface.
type AzureGlobalServiceSpecsGetter interface {
	AzureGlobalServiceSpecs(namespace string) AzureGlobalServiceSpecInterface
}

// AzureGlobalServiceSpecInterface has methods to work with AzureGlobalServiceSpec resources.
type AzureGlobalServiceSpecInterface interface {
	Create(ctx context.Context, azureGlobalServiceSpec *v1alpha1.AzureGlobalServiceSpec, opts v1.CreateOptions) (*v1alpha1.AzureGlobalServiceSpec, error)
	Update(ctx context.Context, azureGlobalServiceSpec *v1alpha1.AzureGlobalServiceSpec, opts v1.UpdateOptions) (*v1alpha1.AzureGlobalServiceSpec, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.AzureGlobalServiceSpec, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AzureGlobalServiceSpecList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AzureGlobalServiceSpec, err error)
	AzureGlobalServiceSpecExpansion
}

// azureGlobalServiceSpecs implements AzureGlobalServiceSpecInterface
type azureGlobalServiceSpecs struct {
	client rest.Interface
	ns     string
}

// newAzureGlobalServiceSpecs returns a AzureGlobalServiceSpecs
func newAzureGlobalServiceSpecs(c *AzureglobalservicesV1alpha1Client, namespace string) *azureGlobalServiceSpecs {
	return &azureGlobalServiceSpecs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the azureGlobalServiceSpec, and returns the corresponding azureGlobalServiceSpec object, and an error if there is any.
func (c *azureGlobalServiceSpecs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AzureGlobalServiceSpec, err error) {
	result = &v1alpha1.AzureGlobalServiceSpec{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("azureglobalservicespecs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzureGlobalServiceSpecs that match those selectors.
func (c *azureGlobalServiceSpecs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AzureGlobalServiceSpecList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzureGlobalServiceSpecList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("azureglobalservicespecs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azureGlobalServiceSpecs.
func (c *azureGlobalServiceSpecs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("azureglobalservicespecs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a azureGlobalServiceSpec and creates it.  Returns the server's representation of the azureGlobalServiceSpec, and an error, if there is any.
func (c *azureGlobalServiceSpecs) Create(ctx context.Context, azureGlobalServiceSpec *v1alpha1.AzureGlobalServiceSpec, opts v1.CreateOptions) (result *v1alpha1.AzureGlobalServiceSpec, err error) {
	result = &v1alpha1.AzureGlobalServiceSpec{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("azureglobalservicespecs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(azureGlobalServiceSpec).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a azureGlobalServiceSpec and updates it. Returns the server's representation of the azureGlobalServiceSpec, and an error, if there is any.
func (c *azureGlobalServiceSpecs) Update(ctx context.Context, azureGlobalServiceSpec *v1alpha1.AzureGlobalServiceSpec, opts v1.UpdateOptions) (result *v1alpha1.AzureGlobalServiceSpec, err error) {
	result = &v1alpha1.AzureGlobalServiceSpec{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("azureglobalservicespecs").
		Name(azureGlobalServiceSpec.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(azureGlobalServiceSpec).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the azureGlobalServiceSpec and deletes it. Returns an error if one occurs.
func (c *azureGlobalServiceSpecs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("azureglobalservicespecs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azureGlobalServiceSpecs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("azureglobalservicespecs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched azureGlobalServiceSpec.
func (c *azureGlobalServiceSpecs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AzureGlobalServiceSpec, err error) {
	result = &v1alpha1.AzureGlobalServiceSpec{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("azureglobalservicespecs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
