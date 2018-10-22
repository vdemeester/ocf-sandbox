# OCF-230 – knative build with s2i (binary) and buildah (on any k8s)

Taking the opposite direction of OCF-196 (using openshift
build/buildconfig), we want to be able to run a s2i build (using s2i
binary and `buildah`) to build an image that can be later on used for
a knative service.

This means:
- Create a `builder` image that can generate a `Dockerfile` from `s2i`
- Create a `builder` image that use `buildah` to build that
  `Dockerfile` (could be the same builder image)
- Provide a build-template (most likely upstream in
  https://github.com/knative/build-templates`) for that

# Builders

This folder contains 2 builders:

- [`buildah-builder`](./buildah-builder/), that contains everything
  required for building dockerfiles without Docker, using
  `buildah`. It is auto-built on the hub :
  [`vdemeester/buildah-builder`](https://hub.docker.com/r/vdemeester/buildah-builder/)
- [`s2i-builder`](./s2i-builder/), that contains uses [`s2i`
  a.k.a. source-to-image](https://github.com/openshift/source-to-image)
  to generate a `Dockerfile` from a source file and a base image. It
  can be combined with `buildah-builder`. It is auto-built on the hub :
  [`vdemeester/s2i-builder`](https://hub.docker.com/r/vdemeester/s2i-builder/)

# Knative templates

There is 2 templates too (that are or will be upstream in
[`knative/build-templates`](https://github.com/knative/build-templates))
:

- [`buildah.template.yml`](./buildah.template.yml), that uses the
  `buildah-builder` to build an image in `knative`.
  
  Example:
  ```yaml
  apiVersion: build.knative.dev/v1alpha1
  kind: Build
  metadata:
    name: buildah-build-my-repo
  spec:
    timeout: 50m
    serviceAccountName: build-bot
    source:
      git:
        url: https://github.com/vdemeester/os-sample-python.git
        revision: dockerfile
    template:
      name: buildah
      arguments:
      - name: IMAGE
        value: vdemeester/my-app
  ```
  
- [`s2i.template.yml`](./s2i.template.yml), that uses `s2i-builder`
  and `buildah-builder` to build an image, using `s2i`, in `knative`.
  
  Example:
  ```yaml
  apiVersion: build.knative.dev/v1alpha1
  kind: Build
  metadata:
    name: s2i-build-my-repo
  spec:
    timeout: 50m
    serviceAccountName: build-bot
    source:
      git:
        url: https://github.com/vdemeester/os-sample-python
        revision: master
    template:
      name: s2i
      arguments:
      - name: BASE_IMAGE
        value: centos/python-36-centos7
      - name: IMAGE
        value: docker.io/vdemeester/helloworld-python
  ```
  
