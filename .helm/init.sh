#!/usr/bin/env bash


kubectl create secret generic do-registry \
  --from-file=.dockerconfigjson=docker-config.json \
  --type=kubernetes.io/dockerconfigjson
secret/do-registry created