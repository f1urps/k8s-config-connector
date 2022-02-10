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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/dataflow/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDataflowJobs implements DataflowJobInterface
type FakeDataflowJobs struct {
	Fake *FakeDataflowV1beta1
	ns   string
}

var dataflowjobsResource = schema.GroupVersionResource{Group: "dataflow.cnrm.cloud.google.com", Version: "v1beta1", Resource: "dataflowjobs"}

var dataflowjobsKind = schema.GroupVersionKind{Group: "dataflow.cnrm.cloud.google.com", Version: "v1beta1", Kind: "DataflowJob"}

// Get takes name of the dataflowJob, and returns the corresponding dataflowJob object, and an error if there is any.
func (c *FakeDataflowJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.DataflowJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dataflowjobsResource, c.ns, name), &v1beta1.DataflowJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataflowJob), err
}

// List takes label and field selectors, and returns the list of DataflowJobs that match those selectors.
func (c *FakeDataflowJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.DataflowJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dataflowjobsResource, dataflowjobsKind, c.ns, opts), &v1beta1.DataflowJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.DataflowJobList{ListMeta: obj.(*v1beta1.DataflowJobList).ListMeta}
	for _, item := range obj.(*v1beta1.DataflowJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataflowJobs.
func (c *FakeDataflowJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dataflowjobsResource, c.ns, opts))

}

// Create takes the representation of a dataflowJob and creates it.  Returns the server's representation of the dataflowJob, and an error, if there is any.
func (c *FakeDataflowJobs) Create(ctx context.Context, dataflowJob *v1beta1.DataflowJob, opts v1.CreateOptions) (result *v1beta1.DataflowJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dataflowjobsResource, c.ns, dataflowJob), &v1beta1.DataflowJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataflowJob), err
}

// Update takes the representation of a dataflowJob and updates it. Returns the server's representation of the dataflowJob, and an error, if there is any.
func (c *FakeDataflowJobs) Update(ctx context.Context, dataflowJob *v1beta1.DataflowJob, opts v1.UpdateOptions) (result *v1beta1.DataflowJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dataflowjobsResource, c.ns, dataflowJob), &v1beta1.DataflowJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataflowJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataflowJobs) UpdateStatus(ctx context.Context, dataflowJob *v1beta1.DataflowJob, opts v1.UpdateOptions) (*v1beta1.DataflowJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dataflowjobsResource, "status", c.ns, dataflowJob), &v1beta1.DataflowJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataflowJob), err
}

// Delete takes name of the dataflowJob and deletes it. Returns an error if one occurs.
func (c *FakeDataflowJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dataflowjobsResource, c.ns, name), &v1beta1.DataflowJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataflowJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dataflowjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.DataflowJobList{})
	return err
}

// Patch applies the patch and returns the patched dataflowJob.
func (c *FakeDataflowJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.DataflowJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dataflowjobsResource, c.ns, name, pt, data, subresources...), &v1beta1.DataflowJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataflowJob), err
}