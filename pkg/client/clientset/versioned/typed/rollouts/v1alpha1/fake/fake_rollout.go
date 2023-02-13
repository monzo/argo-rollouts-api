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

	v1alpha1 "github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRollouts implements RolloutInterface
type FakeRollouts struct {
	Fake *FakeArgoprojV1alpha1
	ns   string
}

var rolloutsResource = schema.GroupVersionResource{Group: "argoproj.io", Version: "v1alpha1", Resource: "rollouts"}

var rolloutsKind = schema.GroupVersionKind{Group: "argoproj.io", Version: "v1alpha1", Kind: "Rollout"}

// Get takes name of the rollout, and returns the corresponding rollout object, and an error if there is any.
func (c *FakeRollouts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Rollout, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rolloutsResource, c.ns, name), &v1alpha1.Rollout{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Rollout), err
}

// List takes label and field selectors, and returns the list of Rollouts that match those selectors.
func (c *FakeRollouts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RolloutList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rolloutsResource, rolloutsKind, c.ns, opts), &v1alpha1.RolloutList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RolloutList{ListMeta: obj.(*v1alpha1.RolloutList).ListMeta}
	for _, item := range obj.(*v1alpha1.RolloutList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rollouts.
func (c *FakeRollouts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rolloutsResource, c.ns, opts))

}

// Create takes the representation of a rollout and creates it.  Returns the server's representation of the rollout, and an error, if there is any.
func (c *FakeRollouts) Create(ctx context.Context, rollout *v1alpha1.Rollout, opts v1.CreateOptions) (result *v1alpha1.Rollout, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rolloutsResource, c.ns, rollout), &v1alpha1.Rollout{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Rollout), err
}

// Update takes the representation of a rollout and updates it. Returns the server's representation of the rollout, and an error, if there is any.
func (c *FakeRollouts) Update(ctx context.Context, rollout *v1alpha1.Rollout, opts v1.UpdateOptions) (result *v1alpha1.Rollout, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rolloutsResource, c.ns, rollout), &v1alpha1.Rollout{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Rollout), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRollouts) UpdateStatus(ctx context.Context, rollout *v1alpha1.Rollout, opts v1.UpdateOptions) (*v1alpha1.Rollout, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(rolloutsResource, "status", c.ns, rollout), &v1alpha1.Rollout{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Rollout), err
}

// Delete takes name of the rollout and deletes it. Returns an error if one occurs.
func (c *FakeRollouts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(rolloutsResource, c.ns, name), &v1alpha1.Rollout{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRollouts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rolloutsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RolloutList{})
	return err
}
