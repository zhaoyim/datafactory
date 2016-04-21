DataFoundry
==============================

[![Build Status](https://travis-ci.org/asiainfoLDP/datafactory.svg?branch=feature%2Fenhance)](https://travis-ci.org/asiainfoLDP/datafactory)


**Features:**

* Easily build applications with integrated service discovery, DNS, load balancing, failover, health checking, persistent storage, and fast scaling
* Push source code to your Git repository and automatically roll out containerized applications (micro or macro)
* Easy to use client and web console for building and monitoring applications
* Centralized administration and management of an entire stack, team, or organization
  * Templatize the components of your system, reuse them, and iteratively deploy them over time
  * Roll out changes to software stacks to your entire organization in a controlled fashion
* Team and user isolation of containers, builds, and network communication in an easy multi-tenant system
  * Allow developers to run containers securely with fine-grained controls in production
  * Limit, track, and manage the developers and teams on the platform
* Integrated Docker registry, automatic edge load balancing, cluster logging, and integrated metrics


Getting Started
---------------

### Concepts

We highly recommend trying out the [Origin walkthrough](https://github.com/openshift/origin/blob/master/examples/sample-app/README.md) which covers the core concepts in Origin.  The walkthrough is accompanied by a blog series on [blog.openshift.com](https://blog.openshift.com/openshift-v3-deep-dive-docker-kubernetes/) that goes into more detail.  It's a great place to start.

### API

The API is located on each server at `https://<host>:8443/oapi/v1`. These APIs are described via [Swagger v1.2](https://www.swagger.io) at `https://<host>:8443/swaggerapi/oapi/v1`. For more, [see the API documentation](http://docs.openshift.org/latest/rest_api/openshift_v1.html).

### Kubernetes

If you're looking for more information about using Kubernetes or the lower level concepts that Origin depends on, see the following:

* [Kubernetes Getting Started](https://github.com/kubernetes/kubernetes/blob/master/README.md)
* [Kubernetes Documentation](https://github.com/kubernetes/kubernetes/blob/master/docs/README.md)
* [Kubernetes API](http://docs.openshift.org/latest/rest_api/kubernetes_v1.html)



FAQ
---

1. How does Origin relate to Kubernetes?

    Origin is a distribution of Kubernetes optimized for enterprise application development and deployment,
    used by OpenShift 3 and Atomic Enterprise.  Origin embeds Kubernetes and adds additional
    functionality to offer a simple, powerful, and easy-to-approach developer and operator experience
    for building applications in containers.  Our goal is to do most of that work upstream, with
    integration and final packaging occurring in Origin.

    You can run the core Kubernetes server components with `openshift start kube`, use `kubectl` via
    `openshift kube`, and the Origin release zips include versions of `kubectl`, `kubelet`,
    `kube-apiserver`, and other core components. You can see the version of Kubernetes included with
    Origin via `openshift version`.


2. Why doesn't my Docker image run on OpenShift?

    Security! Origin runs with the following security policy by default:

    * Containers run as a non-root unique user that is separate from other system users
      * They cannot access host resources, run privileged, or become root
      * They are given CPU and memory limits defined by the system administrator
      * Any persistent storage they access will be under a unique SELinux label, which prevents others from seeing their content
      * These settings are per project, so containers in different projects cannot see each other by default
    * Regular users can run Docker, source, and custom builds
      * By default, Docker builds can (and often do) run as root. You can control who can create Docker builds through the `builds/docker` and `builds/custom` policy resource.
    * Regular users and project admins cannot change their security quotas.

    Many Docker containers expect to run as root (and therefore edit all the contents of the filesystem). The [Image Author's guide](https://docs.openshift.org/latest/creating_images/guidelines.html#openshift-specific-guidelines) gives recommendations on making your image more secure by default:

    * Don't run as root
    * Make directories you want to write to group-writable and owned by group id 0
    * Set the net-bind capability on your executables if they need to bind to ports &lt;1024

    Otherwise, you can see the [security documentation](https://docs.openshift.org/latest/admin_guide/manage_scc.html) for descriptions on how to relax these restrictions.

3. How do I get networking working?

    The Origin and Kubernetes network model assigns each pod (group of containers) an IP that is expected to be reachable from all nodes in the cluster. The default setup is through a simple SDN plugin with OVS - this plugin expects the port 4679 to be open between nodes in the cluster. Also, the Origin master processes need to be able to reach pods via the network, so they may require the SDN plugin.

    Other networking options are available such as Calico, Flannel, Nuage, and Weave. For a non-overlay networking solution, existing networks can be used by assigning a different subnet to each host, and ensuring routing rules deliver packets bound for that subnet to the host it belongs to. This is called [host subnet routing](https://docs.openshift.org/latest/admin_guide/native_container_routing.html).

    Versions of Docker distributed by the Docker team don't allow containers to mount volumes on the host and write to them (mount propagation is private). Kubernetes manages volumes and uses them to expose secrets into containers, which Origin uses to give containers the tokens they need to access the API and run deployments and builds. Until mount propagation is configurable in Docker you must use Docker on Fedora, CentOS, or RHEL (which have a patch to allow mount propagation) or run Origin outside of a container. Tracked in [this issue](https://github.com/openshift/origin/issues/3072).


Testing
------------

If you want to run the test suite, make sure you have your environment set up, and from the `datafactory` directory run:

```
# run the unit tests
$ make check

# run a command-line integration test suite
$ hack/test-cmd.sh

# run the integration server test suite
$ hack/test-integration.sh

# run the end-to-end test suite
$ hack/test-end-to-end.sh

# run all of the tests above
$ make test
```

You'll need [etcd](https://github.com/coreos/etcd) installed and on your path for the integration and end-to-end tests to run, and Docker must be installed to run the end-to-end tests.  To install etcd you should be able to run:

```
$ hack/install-etcd.sh
```

Some of the components of Origin run as Docker images, including the builders and deployment tools in `images/builder/docker/*` and 'images/deploy/*`.  To build them locally run

```
$ hack/build-images.sh
```

To hack on the web console, check out the [assets/README.md](assets/README.md) file for instructions on testing the console and building your changes.


License
-------

DataFoundry is licensed under the [Apache License, Version 2.0](http://www.apache.org/licenses/).
