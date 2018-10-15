# OCF-196/197 Knative build and openshift build

This is a PoC to connect knative/build and Openshift build/buildconfig.
The idea is to be able to use Openshift BuildConfig (and Builds) for knative (build/serving…).

knative/build can use a buid that is not knative/build *but* it has to respect a contract, mainly in the returned status.

```
GET /apis/build.knative.dev/v1alpha1/namespaces/default/builds/build-1acub3
...
status:
  # Link to log stream; could be ELK or Stackdriver, for example
  buildLogsLink: "http://logging.infra.mycompany.com/...?filter=..."
  conditions:
  - type: Failed
    status: True
    reason: BuildStepFailed  # could also be SourceMissing, etc
    message: "Step XYZ failed with error message: $LASTLOGLINE"
```

This is described [here](https://github.com/knative/serving/blob/master/docs/spec/errors.md#build-failed), and resources should follow [the k8s API conventions for `condition`](https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#typical-status-properties).

Openshift `Build` and `BuildConfig` objects do not support that as of 3.11, thus `knative/serving` service do not pick up an Openshift build.

This folder holds a PoC to *adapt*/*bridge* knative/builds and Openshift build. What it provides is :

- a small tool (`build`) that can create/update and start openshift build from some simple arguments. *it needs love and polish*.
- a knative `BuildTemplate` that uses this tools to run a build from `knative/build`

In a nutshell, the templates will create builds with the following steps:

1. create or update the buildconfig
2. start a build from this buildconfig and wait for it (and fail if the status is wrong).

- `knative/0-build-template.yaml` holds the knative build template
- `knative/1-build.yaml` holds an example that creates a build using the template above
- `knative/2-serving.yaml` holds a service example that creates a service based on a knative build (base on the templates). It will schedule a build config and wait for it to complete to run.

The `Makefile` allows to build the image if needed, but so far it's tied to my user `vdemeester`.
The Go code in `cmd` is what is running inside `vdemeester/oc-builder` used in the build/template.

## Example

```bash
# Get a working minishift with knative (serving and build at least) on it
# […]
# We will need our user to have the correct right
# FIXME(vdemeester) this rights are too much
$ oc policy add-role-to-user admin system:serviceaccount:myproject:default
# We will need to be able to push image to the dockerhub
# assuming you're logged on to docker locally
$ oc secrets new dockerhub ~/.docker/config.json
# Create the build template
$ oc apply -f knative/0-build-template.yaml
# Create a build (and validate it work)
$ oc apply -f knative/1-build.yaml
# […]
$ oc get pods
$ oc get build
$ oc get buildconfig
# Clean the build (for now)
$ oc delete -f knative/1-build.yaml
# Create a serving that will depend on a build (creating this one too)
$ oc apply -f knative/2-serving.yaml
# […]
$ oc get pods
# […] builds and pod should be created and running, …
# Once the build is done and service running
$ curl -H "Host: helloworld-python.myproject.example.com" http://$(minishift ip):32380
Hello Openshift!
```