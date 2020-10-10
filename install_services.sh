#!/bin/bash

kubectl create ns foo || true
kubectl create ns bar || true
kubectl label ns/foo linkerd.io/inject=enabled
kubectl label ns/bar linkerd.io/inject=enabled
helm install linkerd-service-tester .
