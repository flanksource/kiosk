/*
Copyright 2020 DevSpace Technologies Inc.

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
	v1alpha1 "github.com/kiosk-sh/kiosk/pkg/apis/config/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTemplates implements TemplateInterface
type FakeTemplates struct {
	Fake *FakeConfigV1alpha1
}

var templatesResource = schema.GroupVersionResource{Group: "config.kiosk.sh", Version: "v1alpha1", Resource: "templates"}

var templatesKind = schema.GroupVersionKind{Group: "config.kiosk.sh", Version: "v1alpha1", Kind: "Template"}

// Get takes name of the template, and returns the corresponding template object, and an error if there is any.
func (c *FakeTemplates) Get(name string, options v1.GetOptions) (result *v1alpha1.Template, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(templatesResource, name), &v1alpha1.Template{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Template), err
}

// List takes label and field selectors, and returns the list of Templates that match those selectors.
func (c *FakeTemplates) List(opts v1.ListOptions) (result *v1alpha1.TemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(templatesResource, templatesKind, opts), &v1alpha1.TemplateList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TemplateList{ListMeta: obj.(*v1alpha1.TemplateList).ListMeta}
	for _, item := range obj.(*v1alpha1.TemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested templates.
func (c *FakeTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(templatesResource, opts))
}

// Create takes the representation of a template and creates it.  Returns the server's representation of the template, and an error, if there is any.
func (c *FakeTemplates) Create(template *v1alpha1.Template) (result *v1alpha1.Template, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(templatesResource, template), &v1alpha1.Template{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Template), err
}

// Update takes the representation of a template and updates it. Returns the server's representation of the template, and an error, if there is any.
func (c *FakeTemplates) Update(template *v1alpha1.Template) (result *v1alpha1.Template, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(templatesResource, template), &v1alpha1.Template{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Template), err
}

// Delete takes name of the template and deletes it. Returns an error if one occurs.
func (c *FakeTemplates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(templatesResource, name), &v1alpha1.Template{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(templatesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.TemplateList{})
	return err
}

// Patch applies the patch and returns the patched template.
func (c *FakeTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Template, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(templatesResource, name, pt, data, subresources...), &v1alpha1.Template{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Template), err
}
