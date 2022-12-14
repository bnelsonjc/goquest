#!/bin/bash

cd "$(dirname "$0")"

export GOFLAGS=-mod=mod
go build -o goquest cmd/goquest/main.go
