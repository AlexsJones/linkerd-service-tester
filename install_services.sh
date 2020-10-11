#!/bin/bash

kubectl create ns foo || true
kubectl annotate ns foo linkerd.io/inject=enabled
kubectl create ns bar || true
kubectl annotate ns bar linkerd.io/inject=enabled
helm install linkerd-service-tester .
