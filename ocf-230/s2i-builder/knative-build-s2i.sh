#!/bin/bash

echo "Creating Dockerfile"
/create-dockerfile.sh $1

echo "Building Image"
/build-image.sh $2
