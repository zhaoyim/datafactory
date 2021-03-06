package client

import (
	applicationapi "github.com/openshift/origin/pkg/application/api"
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/watch"
)

// ApplicationsInterface has methods to work with Application resources in a namespace
type ApplicationsInterface interface {
	Applications(namespace string) ApplicationInterface
}

// ApplicationInterface exposes methods on project resources.
type ApplicationInterface interface {
	Create(p *applicationapi.Application) (*applicationapi.Application, error)
	Delete(name string) error
	Update(p *applicationapi.Application) (*applicationapi.Application, error)
	Get(name string) (*applicationapi.Application, error)
	List(opts kapi.ListOptions) (*applicationapi.ApplicationList, error)
	Watch(opts kapi.ListOptions) (watch.Interface, error)
}

type applications struct {
	r  *Client
	ns string
}

// newUsers returns a project
func newApplications(c *Client, namespace string) *applications {
	return &applications{
		r:  c,
		ns: namespace,
	}
}

// Get returns information about a particular project or an error
func (c *applications) Get(name string) (result *applicationapi.Application, err error) {
	result = &applicationapi.Application{}
	err = c.r.Get().Namespace(c.ns).Resource("applications").Name(name).Do().Into(result)
	return
}

// List returns all applications matching the label selector
func (c *applications) List(opts kapi.ListOptions) (result *applicationapi.ApplicationList, err error) {
	result = &applicationapi.ApplicationList{}
	err = c.r.Get().
		Namespace(c.ns).
		Resource("applications").
		VersionedParams(&opts, kapi.ParameterCodec).
		Do().
		Into(result)
	return
}

// Create creates a new Application
func (c *applications) Create(p *applicationapi.Application) (result *applicationapi.Application, err error) {
	result = &applicationapi.Application{}
	err = c.r.Post().Namespace(c.ns).Resource("applications").Body(p).Do().Into(result)
	return
}

// Update updates the project on server
func (c *applications) Update(p *applicationapi.Application) (result *applicationapi.Application, err error) {
	result = &applicationapi.Application{}
	err = c.r.Put().Namespace(c.ns).Resource("applications").Name(p.Name).Body(p).Do().Into(result)
	return
}

// Delete removes the project on server
func (c *applications) Delete(name string) (err error) {
	err = c.r.Delete().Namespace(c.ns).Resource("applications").Name(name).Do().Error()
	return
}

// Watch returns a watch.Interface that watches the requested applications
func (c *applications) Watch(opts kapi.ListOptions) (watch.Interface, error) {
	return c.r.Get().
		Prefix("watch").
		Namespace(c.ns).
		Resource("applications").
		VersionedParams(&opts, kapi.ParameterCodec).
		Watch()
}
