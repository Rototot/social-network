#!/usr/bin/env sh

set -ex

PWD=$(pwd)
#trap "docker-compose rm -f -s -v"

cd ./backend

go test -v ./pkg/...