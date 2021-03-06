/*
Copyright 2018 caicloud authors. All rights reserved.
*/

package v1beta1

import (
	scheme "github.com/caicloud/clientset/kubernetes/scheme"
	v1beta1 "github.com/caicloud/clientset/pkg/apis/resource/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ConfigsGetter has a method to return a ConfigInterface.
// A group's client should implement this interface.
type ConfigsGetter interface {
	Configs() ConfigInterface
}

// ConfigInterface has methods to work with Config resources.
type ConfigInterface interface {
	Create(*v1beta1.Config) (*v1beta1.Config, error)
	Update(*v1beta1.Config) (*v1beta1.Config, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.Config, error)
	List(opts v1.ListOptions) (*v1beta1.ConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Config, err error)
	ConfigExpansion
}

// configs implements ConfigInterface
type configs struct {
	client rest.Interface
}

// newConfigs returns a Configs
func newConfigs(c *ResourceV1beta1Client) *configs {
	return &configs{
		client: c.RESTClient(),
	}
}

// Create takes the representation of a config and creates it.  Returns the server's representation of the config, and an error, if there is any.
func (c *configs) Create(config *v1beta1.Config) (result *v1beta1.Config, err error) {
	result = &v1beta1.Config{}
	err = c.client.Post().
		Resource("configs").
		Body(config).
		Do().
		Into(result)
	return
}

// Update takes the representation of a config and updates it. Returns the server's representation of the config, and an error, if there is any.
func (c *configs) Update(config *v1beta1.Config) (result *v1beta1.Config, err error) {
	result = &v1beta1.Config{}
	err = c.client.Put().
		Resource("configs").
		Name(config.Name).
		Body(config).
		Do().
		Into(result)
	return
}

// Delete takes name of the config and deletes it. Returns an error if one occurs.
func (c *configs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("configs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *configs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("configs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the config, and returns the corresponding config object, and an error if there is any.
func (c *configs) Get(name string, options v1.GetOptions) (result *v1beta1.Config, err error) {
	result = &v1beta1.Config{}
	err = c.client.Get().
		Resource("configs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Configs that match those selectors.
func (c *configs) List(opts v1.ListOptions) (result *v1beta1.ConfigList, err error) {
	result = &v1beta1.ConfigList{}
	err = c.client.Get().
		Resource("configs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested configs.
func (c *configs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("configs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched config.
func (c *configs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Config, err error) {
	result = &v1beta1.Config{}
	err = c.client.Patch(pt).
		Resource("configs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
