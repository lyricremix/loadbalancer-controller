/*
Copyright 2018 caicloud authors. All rights reserved.
*/

package v1alpha1

import (
	scheme "github.com/caicloud/clientset/kubernetes/scheme"
	v1alpha1 "github.com/caicloud/clientset/pkg/apis/release/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CanaryReleasesGetter has a method to return a CanaryReleaseInterface.
// A group's client should implement this interface.
type CanaryReleasesGetter interface {
	CanaryReleases(namespace string) CanaryReleaseInterface
}

// CanaryReleaseInterface has methods to work with CanaryRelease resources.
type CanaryReleaseInterface interface {
	Create(*v1alpha1.CanaryRelease) (*v1alpha1.CanaryRelease, error)
	Update(*v1alpha1.CanaryRelease) (*v1alpha1.CanaryRelease, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CanaryRelease, error)
	List(opts v1.ListOptions) (*v1alpha1.CanaryReleaseList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CanaryRelease, err error)
	CanaryReleaseExpansion
}

// canaryReleases implements CanaryReleaseInterface
type canaryReleases struct {
	client rest.Interface
	ns     string
}

// newCanaryReleases returns a CanaryReleases
func newCanaryReleases(c *ReleaseV1alpha1Client, namespace string) *canaryReleases {
	return &canaryReleases{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a canaryRelease and creates it.  Returns the server's representation of the canaryRelease, and an error, if there is any.
func (c *canaryReleases) Create(canaryRelease *v1alpha1.CanaryRelease) (result *v1alpha1.CanaryRelease, err error) {
	result = &v1alpha1.CanaryRelease{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("canaryreleases").
		Body(canaryRelease).
		Do().
		Into(result)
	return
}

// Update takes the representation of a canaryRelease and updates it. Returns the server's representation of the canaryRelease, and an error, if there is any.
func (c *canaryReleases) Update(canaryRelease *v1alpha1.CanaryRelease) (result *v1alpha1.CanaryRelease, err error) {
	result = &v1alpha1.CanaryRelease{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("canaryreleases").
		Name(canaryRelease.Name).
		Body(canaryRelease).
		Do().
		Into(result)
	return
}

// Delete takes name of the canaryRelease and deletes it. Returns an error if one occurs.
func (c *canaryReleases) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("canaryreleases").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *canaryReleases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("canaryreleases").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the canaryRelease, and returns the corresponding canaryRelease object, and an error if there is any.
func (c *canaryReleases) Get(name string, options v1.GetOptions) (result *v1alpha1.CanaryRelease, err error) {
	result = &v1alpha1.CanaryRelease{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("canaryreleases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CanaryReleases that match those selectors.
func (c *canaryReleases) List(opts v1.ListOptions) (result *v1alpha1.CanaryReleaseList, err error) {
	result = &v1alpha1.CanaryReleaseList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("canaryreleases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested canaryReleases.
func (c *canaryReleases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("canaryreleases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched canaryRelease.
func (c *canaryReleases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CanaryRelease, err error) {
	result = &v1alpha1.CanaryRelease{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("canaryreleases").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
