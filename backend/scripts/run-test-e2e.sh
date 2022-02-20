#!/usr/bin/env sh

set -ex

trap "docker-compose down -v --rmi local --remove-orphans" SIGINT SIGTERM EXIT ERR

docker-compose -f docker-compose.tests.yaml build --pull
docker-compose -f docker-compose.tests.yaml run tests-e2e