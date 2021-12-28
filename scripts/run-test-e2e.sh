#!/usr/bin/env sh

set -ex

PWD=$(pwd)
#trap "docker-compose down -v --rmi local"

cd ./backend

docker-compose -f docker-compose.tests.yaml build --pull
docker-compose -f docker-compose.tests.yaml run tests-e2e