/*
Copyright 2022 The Katalyst Authors.

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

	v1alpha1 "github.com/kubewharf/katalyst-api/pkg/apis/config/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKatalystCustomConfigs implements KatalystCustomConfigInterface
type FakeKatalystCustomConfigs struct {
	Fake *FakeConfigV1alpha1
	ns   string
}

var katalystcustomconfigsResource = schema.GroupVersionResource{Group: "config.katalyst.kubewharf.io", Version: "v1alpha1", Resource: "katalystcustomconfigs"}

var katalystcustomconfigsKind = schema.GroupVersionKind{Group: "config.katalyst.kubewharf.io", Version: "v1alpha1", Kind: "KatalystCustomConfig"}

// Get takes name of the katalystCustomConfig, and returns the corresponding katalystCustomConfig object, and an error if there is any.
func (c *FakeKatalystCustomConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KatalystCustomConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(katalystcustomconfigsResource, c.ns, name), &v1alpha1.KatalystCustomConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KatalystCustomConfig), err
}

// List takes label and field selectors, and returns the list of KatalystCustomConfigs that match those selectors.
func (c *FakeKatalystCustomConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KatalystCustomConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(katalystcustomconfigsResource, katalystcustomconfigsKind, c.ns, opts), &v1alpha1.KatalystCustomConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KatalystCustomConfigList{ListMeta: obj.(*v1alpha1.KatalystCustomConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.KatalystCustomConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested katalystCustomConfigs.
func (c *FakeKatalystCustomConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(katalystcustomconfigsResource, c.ns, opts))

}

// Create takes the representation of a katalystCustomConfig and creates it.  Returns the server's representation of the katalystCustomConfig, and an error, if there is any.
func (c *FakeKatalystCustomConfigs) Create(ctx context.Context, katalystCustomConfig *v1alpha1.KatalystCustomConfig, opts v1.CreateOptions) (result *v1alpha1.KatalystCustomConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(katalystcustomconfigsResource, c.ns, katalystCustomConfig), &v1alpha1.KatalystCustomConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KatalystCustomConfig), err
}

// Update takes the representation of a katalystCustomConfig and updates it. Returns the server's representation of the katalystCustomConfig, and an error, if there is any.
func (c *FakeKatalystCustomConfigs) Update(ctx context.Context, katalystCustomConfig *v1alpha1.KatalystCustomConfig, opts v1.UpdateOptions) (result *v1alpha1.KatalystCustomConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(katalystcustomconfigsResource, c.ns, katalystCustomConfig), &v1alpha1.KatalystCustomConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KatalystCustomConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKatalystCustomConfigs) UpdateStatus(ctx context.Context, katalystCustomConfig *v1alpha1.KatalystCustomConfig, opts v1.UpdateOptions) (*v1alpha1.KatalystCustomConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(katalystcustomconfigsResource, "status", c.ns, katalystCustomConfig), &v1alpha1.KatalystCustomConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KatalystCustomConfig), err
}

// Delete takes name of the katalystCustomConfig and deletes it. Returns an error if one occurs.
func (c *FakeKatalystCustomConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(katalystcustomconfigsResource, c.ns, name, opts), &v1alpha1.KatalystCustomConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKatalystCustomConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(katalystcustomconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.KatalystCustomConfigList{})
	return err
}

// Patch applies the patch and returns the patched katalystCustomConfig.
func (c *FakeKatalystCustomConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KatalystCustomConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(katalystcustomconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.KatalystCustomConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KatalystCustomConfig), err
}
