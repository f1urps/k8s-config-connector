// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/apigee/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApigeeNATAddresses implements ApigeeNATAddressInterface
type FakeApigeeNATAddresses struct {
	Fake *FakeApigeeV1alpha1
	ns   string
}

var apigeenataddressesResource = schema.GroupVersionResource{Group: "apigee.cnrm.cloud.google.com", Version: "v1alpha1", Resource: "apigeenataddresses"}

var apigeenataddressesKind = schema.GroupVersionKind{Group: "apigee.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ApigeeNATAddress"}

// Get takes name of the apigeeNATAddress, and returns the corresponding apigeeNATAddress object, and an error if there is any.
func (c *FakeApigeeNATAddresses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApigeeNATAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apigeenataddressesResource, c.ns, name), &v1alpha1.ApigeeNATAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApigeeNATAddress), err
}

// List takes label and field selectors, and returns the list of ApigeeNATAddresses that match those selectors.
func (c *FakeApigeeNATAddresses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApigeeNATAddressList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apigeenataddressesResource, apigeenataddressesKind, c.ns, opts), &v1alpha1.ApigeeNATAddressList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApigeeNATAddressList{ListMeta: obj.(*v1alpha1.ApigeeNATAddressList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApigeeNATAddressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apigeeNATAddresses.
func (c *FakeApigeeNATAddresses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apigeenataddressesResource, c.ns, opts))

}

// Create takes the representation of a apigeeNATAddress and creates it.  Returns the server's representation of the apigeeNATAddress, and an error, if there is any.
func (c *FakeApigeeNATAddresses) Create(ctx context.Context, apigeeNATAddress *v1alpha1.ApigeeNATAddress, opts v1.CreateOptions) (result *v1alpha1.ApigeeNATAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apigeenataddressesResource, c.ns, apigeeNATAddress), &v1alpha1.ApigeeNATAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApigeeNATAddress), err
}

// Update takes the representation of a apigeeNATAddress and updates it. Returns the server's representation of the apigeeNATAddress, and an error, if there is any.
func (c *FakeApigeeNATAddresses) Update(ctx context.Context, apigeeNATAddress *v1alpha1.ApigeeNATAddress, opts v1.UpdateOptions) (result *v1alpha1.ApigeeNATAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apigeenataddressesResource, c.ns, apigeeNATAddress), &v1alpha1.ApigeeNATAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApigeeNATAddress), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApigeeNATAddresses) UpdateStatus(ctx context.Context, apigeeNATAddress *v1alpha1.ApigeeNATAddress, opts v1.UpdateOptions) (*v1alpha1.ApigeeNATAddress, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apigeenataddressesResource, "status", c.ns, apigeeNATAddress), &v1alpha1.ApigeeNATAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApigeeNATAddress), err
}

// Delete takes name of the apigeeNATAddress and deletes it. Returns an error if one occurs.
func (c *FakeApigeeNATAddresses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(apigeenataddressesResource, c.ns, name, opts), &v1alpha1.ApigeeNATAddress{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApigeeNATAddresses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apigeenataddressesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApigeeNATAddressList{})
	return err
}

// Patch applies the patch and returns the patched apigeeNATAddress.
func (c *FakeApigeeNATAddresses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApigeeNATAddress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apigeenataddressesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApigeeNATAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApigeeNATAddress), err
}
