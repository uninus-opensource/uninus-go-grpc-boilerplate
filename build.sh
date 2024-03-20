#!/usr/bin/env bash

set -e
set -o pipefail

if [ $# -eq 0 ]
  then
    echo "Usage: build.sh [version] [config file]"
    exit 1
fi

## export go module
# export GO111MODULE=off

## export gosumb
# export GOSUMDB=off

# go clean && CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo
cp $2 service.conf

docker build --no-cache -t asia.gcr.io/$NAMESPACE/$SERVICE:$1 --build-arg SSH_KEY="$(cat ~/.ssh/id_ecdsa)" .
docker push asia.gcr.io/$NAMESPACE/$SERVICE:$1
docker rmi asia.gcr.io/$NAMESPACE/$SERVICE:$1
