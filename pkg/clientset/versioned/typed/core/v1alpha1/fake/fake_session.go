/*
Copyright 2020 The Tilt Dev Authors

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/tilt-dev/tilt/pkg/apis/core/v1alpha1"
)

// FakeSessions implements SessionInterface
type FakeSessions struct {
	Fake *FakeTiltV1alpha1
}

var sessionsResource = schema.GroupVersionResource{Group: "tilt.dev", Version: "v1alpha1", Resource: "sessions"}

var sessionsKind = schema.GroupVersionKind{Group: "tilt.dev", Version: "v1alpha1", Kind: "Session"}

// Get takes name of the session, and returns the corresponding session object, and an error if there is any.
func (c *FakeSessions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Session, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(sessionsResource, name), &v1alpha1.Session{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Session), err
}

// List takes label and field selectors, and returns the list of Sessions that match those selectors.
func (c *FakeSessions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SessionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(sessionsResource, sessionsKind, opts), &v1alpha1.SessionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SessionList{ListMeta: obj.(*v1alpha1.SessionList).ListMeta}
	for _, item := range obj.(*v1alpha1.SessionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sessions.
func (c *FakeSessions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(sessionsResource, opts))
}

// Create takes the representation of a session and creates it.  Returns the server's representation of the session, and an error, if there is any.
func (c *FakeSessions) Create(ctx context.Context, session *v1alpha1.Session, opts v1.CreateOptions) (result *v1alpha1.Session, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(sessionsResource, session), &v1alpha1.Session{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Session), err
}

// Update takes the representation of a session and updates it. Returns the server's representation of the session, and an error, if there is any.
func (c *FakeSessions) Update(ctx context.Context, session *v1alpha1.Session, opts v1.UpdateOptions) (result *v1alpha1.Session, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(sessionsResource, session), &v1alpha1.Session{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Session), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSessions) UpdateStatus(ctx context.Context, session *v1alpha1.Session, opts v1.UpdateOptions) (*v1alpha1.Session, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(sessionsResource, "status", session), &v1alpha1.Session{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Session), err
}

// Delete takes name of the session and deletes it. Returns an error if one occurs.
func (c *FakeSessions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(sessionsResource, name), &v1alpha1.Session{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSessions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(sessionsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SessionList{})
	return err
}

// Patch applies the patch and returns the patched session.
func (c *FakeSessions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Session, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(sessionsResource, name, pt, data, subresources...), &v1alpha1.Session{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Session), err
}
