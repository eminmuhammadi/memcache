#!/bin/env bash

BUILD_TIME=$(date +%Y%m%d%H%M%S)
BUILD_ID=$(git rev-parse HEAD)
VERSION="1.0.0"

go install -tags="github.com/eminmuhammadi/memcache" -ldflags "-w -s -X main.VERSION=$VERSION -X main.BUILD_TIME=$BUILD_TIME -X main.BUILD_ID=$BUILD_ID"