#!/usr/bin/env bash

./swag init -d cmd/webserver/ -o cmd/webserver/docs/
#gofmt -s -w cmd pkg
rm -f webserver-linux
env GOOS=linux GOARCH=amd64 go build -o webserver-linux ./cmd/webserver
