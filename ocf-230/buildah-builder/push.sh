#!/bin/bash -x
exec buildah push --creds=openshift:$(cat /var/run/secrets/kubernetes.io/serviceaccount/token) $@
