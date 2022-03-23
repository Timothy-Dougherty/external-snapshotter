/*
Copyright 2022 The Kubernetes Authors.

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

	volumesnapshotv1 "github.com/kubernetes-csi/external-snapshotter/client/v5/apis/volumesnapshot/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVolumeSnapshotClasses implements VolumeSnapshotClassInterface
type FakeVolumeSnapshotClasses struct {
	Fake *FakeSnapshotV1
}

var volumesnapshotclassesResource = schema.GroupVersionResource{Group: "snapshot.storage.k8s.io", Version: "v1", Resource: "volumesnapshotclasses"}

var volumesnapshotclassesKind = schema.GroupVersionKind{Group: "snapshot.storage.k8s.io", Version: "v1", Kind: "VolumeSnapshotClass"}

// Get takes name of the volumeSnapshotClass, and returns the corresponding volumeSnapshotClass object, and an error if there is any.
func (c *FakeVolumeSnapshotClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *volumesnapshotv1.VolumeSnapshotClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(volumesnapshotclassesResource, name), &volumesnapshotv1.VolumeSnapshotClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*volumesnapshotv1.VolumeSnapshotClass), err
}

// List takes label and field selectors, and returns the list of VolumeSnapshotClasses that match those selectors.
func (c *FakeVolumeSnapshotClasses) List(ctx context.Context, opts v1.ListOptions) (result *volumesnapshotv1.VolumeSnapshotClassList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(volumesnapshotclassesResource, volumesnapshotclassesKind, opts), &volumesnapshotv1.VolumeSnapshotClassList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &volumesnapshotv1.VolumeSnapshotClassList{ListMeta: obj.(*volumesnapshotv1.VolumeSnapshotClassList).ListMeta}
	for _, item := range obj.(*volumesnapshotv1.VolumeSnapshotClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested volumeSnapshotClasses.
func (c *FakeVolumeSnapshotClasses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(volumesnapshotclassesResource, opts))
}

// Create takes the representation of a volumeSnapshotClass and creates it.  Returns the server's representation of the volumeSnapshotClass, and an error, if there is any.
func (c *FakeVolumeSnapshotClasses) Create(ctx context.Context, volumeSnapshotClass *volumesnapshotv1.VolumeSnapshotClass, opts v1.CreateOptions) (result *volumesnapshotv1.VolumeSnapshotClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(volumesnapshotclassesResource, volumeSnapshotClass), &volumesnapshotv1.VolumeSnapshotClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*volumesnapshotv1.VolumeSnapshotClass), err
}

// Update takes the representation of a volumeSnapshotClass and updates it. Returns the server's representation of the volumeSnapshotClass, and an error, if there is any.
func (c *FakeVolumeSnapshotClasses) Update(ctx context.Context, volumeSnapshotClass *volumesnapshotv1.VolumeSnapshotClass, opts v1.UpdateOptions) (result *volumesnapshotv1.VolumeSnapshotClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(volumesnapshotclassesResource, volumeSnapshotClass), &volumesnapshotv1.VolumeSnapshotClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*volumesnapshotv1.VolumeSnapshotClass), err
}

// Delete takes name of the volumeSnapshotClass and deletes it. Returns an error if one occurs.
func (c *FakeVolumeSnapshotClasses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(volumesnapshotclassesResource, name, opts), &volumesnapshotv1.VolumeSnapshotClass{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVolumeSnapshotClasses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(volumesnapshotclassesResource, listOpts)

	_, err := c.Fake.Invokes(action, &volumesnapshotv1.VolumeSnapshotClassList{})
	return err
}

// Patch applies the patch and returns the patched volumeSnapshotClass.
func (c *FakeVolumeSnapshotClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *volumesnapshotv1.VolumeSnapshotClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(volumesnapshotclassesResource, name, pt, data, subresources...), &volumesnapshotv1.VolumeSnapshotClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*volumesnapshotv1.VolumeSnapshotClass), err
}
