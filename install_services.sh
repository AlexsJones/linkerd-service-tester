#!/bin/bash

kubectl create ns foo || true
kubectl create ns bar || true
helm install linkerd-service-tester .
