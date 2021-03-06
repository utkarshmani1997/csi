/*
Copyright 2019 The OpenEBS Authors

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
	v1alpha1 "github.com/openebs/csi/pkg/apis/openebs.io/maya/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCStorVolumeClaims implements CStorVolumeClaimInterface
type FakeCStorVolumeClaims struct {
	Fake *FakeOpenebsV1alpha1
	ns   string
}

var cstorvolumeclaimsResource = schema.GroupVersionResource{Group: "openebs.io", Version: "v1alpha1", Resource: "cstorvolumeclaims"}

var cstorvolumeclaimsKind = schema.GroupVersionKind{Group: "openebs.io", Version: "v1alpha1", Kind: "CStorVolumeClaim"}

// Get takes name of the cStorVolumeClaim, and returns the corresponding cStorVolumeClaim object, and an error if there is any.
func (c *FakeCStorVolumeClaims) Get(name string, options v1.GetOptions) (result *v1alpha1.CStorVolumeClaim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cstorvolumeclaimsResource, c.ns, name), &v1alpha1.CStorVolumeClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CStorVolumeClaim), err
}

// List takes label and field selectors, and returns the list of CStorVolumeClaims that match those selectors.
func (c *FakeCStorVolumeClaims) List(opts v1.ListOptions) (result *v1alpha1.CStorVolumeClaimList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cstorvolumeclaimsResource, cstorvolumeclaimsKind, c.ns, opts), &v1alpha1.CStorVolumeClaimList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CStorVolumeClaimList{ListMeta: obj.(*v1alpha1.CStorVolumeClaimList).ListMeta}
	for _, item := range obj.(*v1alpha1.CStorVolumeClaimList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cStorVolumeClaims.
func (c *FakeCStorVolumeClaims) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cstorvolumeclaimsResource, c.ns, opts))

}

// Create takes the representation of a cStorVolumeClaim and creates it.  Returns the server's representation of the cStorVolumeClaim, and an error, if there is any.
func (c *FakeCStorVolumeClaims) Create(cStorVolumeClaim *v1alpha1.CStorVolumeClaim) (result *v1alpha1.CStorVolumeClaim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cstorvolumeclaimsResource, c.ns, cStorVolumeClaim), &v1alpha1.CStorVolumeClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CStorVolumeClaim), err
}

// Update takes the representation of a cStorVolumeClaim and updates it. Returns the server's representation of the cStorVolumeClaim, and an error, if there is any.
func (c *FakeCStorVolumeClaims) Update(cStorVolumeClaim *v1alpha1.CStorVolumeClaim) (result *v1alpha1.CStorVolumeClaim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cstorvolumeclaimsResource, c.ns, cStorVolumeClaim), &v1alpha1.CStorVolumeClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CStorVolumeClaim), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCStorVolumeClaims) UpdateStatus(cStorVolumeClaim *v1alpha1.CStorVolumeClaim) (*v1alpha1.CStorVolumeClaim, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cstorvolumeclaimsResource, "status", c.ns, cStorVolumeClaim), &v1alpha1.CStorVolumeClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CStorVolumeClaim), err
}

// Delete takes name of the cStorVolumeClaim and deletes it. Returns an error if one occurs.
func (c *FakeCStorVolumeClaims) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cstorvolumeclaimsResource, c.ns, name), &v1alpha1.CStorVolumeClaim{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCStorVolumeClaims) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cstorvolumeclaimsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CStorVolumeClaimList{})
	return err
}

// Patch applies the patch and returns the patched cStorVolumeClaim.
func (c *FakeCStorVolumeClaims) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CStorVolumeClaim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cstorvolumeclaimsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CStorVolumeClaim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CStorVolumeClaim), err
}
