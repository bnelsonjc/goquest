#!/bin/bash

cd "$(dirname "$0")"

go build -o goquest cmd/goquest/main.go
