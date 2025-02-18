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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/storage/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStorageHMACKeys implements StorageHMACKeyInterface
type FakeStorageHMACKeys struct {
	Fake *FakeStorageV1alpha1
	ns   string
}

var storagehmackeysResource = schema.GroupVersionResource{Group: "storage.cnrm.cloud.google.com", Version: "v1alpha1", Resource: "storagehmackeys"}

var storagehmackeysKind = schema.GroupVersionKind{Group: "storage.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "StorageHMACKey"}

// Get takes name of the storageHMACKey, and returns the corresponding storageHMACKey object, and an error if there is any.
func (c *FakeStorageHMACKeys) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.StorageHMACKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(storagehmackeysResource, c.ns, name), &v1alpha1.StorageHMACKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageHMACKey), err
}

// List takes label and field selectors, and returns the list of StorageHMACKeys that match those selectors.
func (c *FakeStorageHMACKeys) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.StorageHMACKeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(storagehmackeysResource, storagehmackeysKind, c.ns, opts), &v1alpha1.StorageHMACKeyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StorageHMACKeyList{ListMeta: obj.(*v1alpha1.StorageHMACKeyList).ListMeta}
	for _, item := range obj.(*v1alpha1.StorageHMACKeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageHMACKeys.
func (c *FakeStorageHMACKeys) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(storagehmackeysResource, c.ns, opts))

}

// Create takes the representation of a storageHMACKey and creates it.  Returns the server's representation of the storageHMACKey, and an error, if there is any.
func (c *FakeStorageHMACKeys) Create(ctx context.Context, storageHMACKey *v1alpha1.StorageHMACKey, opts v1.CreateOptions) (result *v1alpha1.StorageHMACKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(storagehmackeysResource, c.ns, storageHMACKey), &v1alpha1.StorageHMACKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageHMACKey), err
}

// Update takes the representation of a storageHMACKey and updates it. Returns the server's representation of the storageHMACKey, and an error, if there is any.
func (c *FakeStorageHMACKeys) Update(ctx context.Context, storageHMACKey *v1alpha1.StorageHMACKey, opts v1.UpdateOptions) (result *v1alpha1.StorageHMACKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(storagehmackeysResource, c.ns, storageHMACKey), &v1alpha1.StorageHMACKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageHMACKey), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStorageHMACKeys) UpdateStatus(ctx context.Context, storageHMACKey *v1alpha1.StorageHMACKey, opts v1.UpdateOptions) (*v1alpha1.StorageHMACKey, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(storagehmackeysResource, "status", c.ns, storageHMACKey), &v1alpha1.StorageHMACKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageHMACKey), err
}

// Delete takes name of the storageHMACKey and deletes it. Returns an error if one occurs.
func (c *FakeStorageHMACKeys) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(storagehmackeysResource, c.ns, name, opts), &v1alpha1.StorageHMACKey{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageHMACKeys) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(storagehmackeysResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.StorageHMACKeyList{})
	return err
}

// Patch applies the patch and returns the patched storageHMACKey.
func (c *FakeStorageHMACKeys) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.StorageHMACKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(storagehmackeysResource, c.ns, name, pt, data, subresources...), &v1alpha1.StorageHMACKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageHMACKey), err
}
