#!/bin/env bash

BUILD_TIME=$(date +%Y%m%d%H%M%S)
BUILD_ID=$(git rev-parse HEAD)
VERSION="1.0.1"

go get -u && \
go mod tidy && \
export CGO_ENABLED=1 && \
export GO111MODULE=on && \
go install -v -tags="sqlite_userauth" -ldflags "-w -s -X main.VERSION=$VERSION -X main.BUILD_TIME=$BUILD_TIME -X main.BUILD_ID=$BUILD_ID"