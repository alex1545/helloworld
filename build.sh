#!/bin/bash

set -ex

# Compile the binary.
CGO_ENABLED=0 GOOS=linux go build -a helloworld.go

# Create a Docker image with this binary in it.
docker build -t "${PUSH_REGISTRY:-staging-k8s.gcr.io}/helloworld:$(git describe --always --dirty)" .
