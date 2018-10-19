#!/bin/bash

STI_SCRIPTS_PATH=/usr/libexec/s2i
baseImage=$1

echo "Generating Dockerfile"
echo "Base Image: $baseImage"

echo "FROM $1" >> ./Dockerfile
echo "" >> ./Dockerfile
echo "USER root"
echo "COPY . /tmp/src" >> ./Dockerfile
echo "RUN $STI_SCRIPTS_PATH/assemble" >> ./Dockerfile
echo "USER 1001"
echo "CMD $STI_SCRIPTS_PATH/run" >> ./Dockerfile

echo "Dockerfile created"
echo "------------------------------------------------------------------------"
cat ./Dockerfile
echo "------------------------------------------------------------------------"

