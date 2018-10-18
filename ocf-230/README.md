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
