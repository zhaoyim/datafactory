package testclient

import (
	kapi "k8s.io/kubernetes/pkg/api"
	ktestclient "k8s.io/kubernetes/pkg/client/unversioned/testclient"
	//"k8s.io/kubernetes/pkg/fields"
	//"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/watch"

	backingserviceinstanceapi "github.com/openshift/origin/pkg/backingserviceinstance/api"
)

// FakeBackingServiceInstances implements BackingServiceInstanceInterface. Meant to be embedded into a struct to get a default
// implementation. This makes faking out just the methods you want to test easier.
type FakeBackingServiceInstances struct {
	Fake *Fake
	Namespace string
}

func (c *FakeBackingServiceInstances) Get(name string) (*backingserviceinstanceapi.BackingServiceInstance, error) {
	obj, err := c.Fake.Invokes(ktestclient.NewGetAction("backingserviceinstances", c.Namespace, name), &backingserviceinstanceapi.BackingServiceInstance{})
	if obj == nil {
		return nil, err
	}

	return obj.(*backingserviceinstanceapi.BackingServiceInstance), err
}

func (c *FakeBackingServiceInstances) List(opts kapi.ListOptions) (*backingserviceinstanceapi.BackingServiceInstanceList, error) {
	obj, err := c.Fake.Invokes(ktestclient.NewListAction("backingserviceinstances", c.Namespace, opts), &backingserviceinstanceapi.BackingServiceInstanceList{})
	if obj == nil {
		return nil, err
	}

	return obj.(*backingserviceinstanceapi.BackingServiceInstanceList), err
}

func (c *FakeBackingServiceInstances) Create(inObj *backingserviceinstanceapi.BackingServiceInstance) (*backingserviceinstanceapi.BackingServiceInstance, error) {
	obj, err := c.Fake.Invokes(ktestclient.NewCreateAction("backingserviceinstances", c.Namespace, inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*backingserviceinstanceapi.BackingServiceInstance), err
}

func (c *FakeBackingServiceInstances) Update(inObj *backingserviceinstanceapi.BackingServiceInstance) (*backingserviceinstanceapi.BackingServiceInstance, error) {
	obj, err := c.Fake.Invokes(ktestclient.NewUpdateAction("backingserviceinstances", c.Namespace, inObj), inObj)
	if obj == nil {
		return nil, err
	}

	return obj.(*backingserviceinstanceapi.BackingServiceInstance), err
}

func (c *FakeBackingServiceInstances) Delete(name string) error {
	_, err := c.Fake.Invokes(ktestclient.NewDeleteAction("backingserviceinstances", c.Namespace, name), &backingserviceinstanceapi.BackingServiceInstance{})
	return err
}

func (c *FakeBackingServiceInstances) Watch(opts kapi.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(ktestclient.NewWatchAction("backingserviceinstances", c.Namespace, opts))
}

func (c *FakeBackingServiceInstances) CreateBinding(name string, bro *backingserviceinstanceapi.BindingRequestOptions) (err error) {
	_, err = c.Fake.Invokes(ktestclient.NewCreateAction("backingserviceinstances/binding", c.Namespace, bro), bro)
	return
}

func (c *FakeBackingServiceInstances) UpdateBinding(name string, bro *backingserviceinstanceapi.BindingRequestOptions) (err error) {
	_, err = c.Fake.Invokes(ktestclient.NewUpdateAction("backingserviceinstances/binding", c.Namespace, bro), bro)
	return
} 
