#!/usr/bin/bash

echo 'test' >> $DOCKER_TAG

if [[ 'test' == 'main' ]]; then DOCKER_TAG='latest'; fi
if [[ 'deploy' == 'deploy' ]]; then DOCKER_TAG='latest'; fi