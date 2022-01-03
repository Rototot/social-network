#!/usr/bin/env sh

set -ex

#trap "docker-compose rm -f -s -v"

go test -v ./pkg/...