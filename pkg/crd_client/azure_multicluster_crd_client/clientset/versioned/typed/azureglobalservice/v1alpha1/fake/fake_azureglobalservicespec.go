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

package fake

import (
	"context"

	v1alpha1 "github.com/Azure/application-gateway-kubernetes-ingress/pkg/apis/azureglobalservice/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAzureGlobalServiceSpecs implements AzureGlobalServiceSpecInterface
type FakeAzureGlobalServiceSpecs struct {
	Fake *FakeAzureglobalservicesV1alpha1
	ns   string
}

var azureglobalservicespecsResource = schema.GroupVersionResource{Group: "azureglobalservices.networking.aks.io", Version: "v1alpha1", Resource: "azureglobalservicespecs"}

var azureglobalservicespecsKind = schema.GroupVersionKind{Group: "azureglobalservices.networking.aks.io", Version: "v1alpha1", Kind: "AzureGlobalServiceSpec"}

// Get takes name of the azureGlobalServiceSpec, and returns the corresponding azureGlobalServiceSpec object, and an error if there is any.
func (c *FakeAzureGlobalServiceSpecs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AzureGlobalServiceSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(azureglobalservicespecsResource, c.ns, name), &v1alpha1.AzureGlobalServiceSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureGlobalServiceSpec), err
}

// List takes label and field selectors, and returns the list of AzureGlobalServiceSpecs that match those selectors.
func (c *FakeAzureGlobalServiceSpecs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AzureGlobalServiceSpecList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(azureglobalservicespecsResource, azureglobalservicespecsKind, c.ns, opts), &v1alpha1.AzureGlobalServiceSpecList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureGlobalServiceSpecList), err
}

// Watch returns a watch.Interface that watches the requested azureGlobalServiceSpecs.
func (c *FakeAzureGlobalServiceSpecs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(azureglobalservicespecsResource, c.ns, opts))

}

// Create takes the representation of a azureGlobalServiceSpec and creates it.  Returns the server's representation of the azureGlobalServiceSpec, and an error, if there is any.
func (c *FakeAzureGlobalServiceSpecs) Create(ctx context.Context, azureGlobalServiceSpec *v1alpha1.AzureGlobalServiceSpec, opts v1.CreateOptions) (result *v1alpha1.AzureGlobalServiceSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(azureglobalservicespecsResource, c.ns, azureGlobalServiceSpec), &v1alpha1.AzureGlobalServiceSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureGlobalServiceSpec), err
}

// Update takes the representation of a azureGlobalServiceSpec and updates it. Returns the server's representation of the azureGlobalServiceSpec, and an error, if there is any.
func (c *FakeAzureGlobalServiceSpecs) Update(ctx context.Context, azureGlobalServiceSpec *v1alpha1.AzureGlobalServiceSpec, opts v1.UpdateOptions) (result *v1alpha1.AzureGlobalServiceSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(azureglobalservicespecsResource, c.ns, azureGlobalServiceSpec), &v1alpha1.AzureGlobalServiceSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureGlobalServiceSpec), err
}

// Delete takes name of the azureGlobalServiceSpec and deletes it. Returns an error if one occurs.
func (c *FakeAzureGlobalServiceSpecs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(azureglobalservicespecsResource, c.ns, name), &v1alpha1.AzureGlobalServiceSpec{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzureGlobalServiceSpecs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(azureglobalservicespecsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzureGlobalServiceSpecList{})
	return err
}

// Patch applies the patch and returns the patched azureGlobalServiceSpec.
func (c *FakeAzureGlobalServiceSpecs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AzureGlobalServiceSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(azureglobalservicespecsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AzureGlobalServiceSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureGlobalServiceSpec), err
}
