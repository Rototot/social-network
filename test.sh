#!/usr/bin/env bash

set -ex

kubectl create secret docker-registry regcred \
    --docker-server=registry.digitalocean.com/rototot \
    --docker-username=d883240a68f12827aca3a02e2576dbee002cb81a1512708f622581b1979f1f0d \
    --docker-password=d883240a68f12827aca3a02e2576dbee002cb81a1512708f622581b1979f1f0d \
    --docker-email=




{"auths":{"registry.digitalocean.com":{"auth":"ZjdkZGQxYjllNDIxNGM0NTBmOWJjZTA2MWQzZDdkYTNmZTQzZjI0NTg1OTE1YzI3Y2M3ZmY1NWQ0ZmQxZTIyZTpmN2RkZDFiOWU0MjE0YzQ1MGY5YmNlMDYxZDNkN2RhM2ZlNDNmMjQ1ODU5MTVjMjdjYzdmZjU1ZDRmZDFlMjJl"}}}