#!/usr/bin/env bash

./swag init -d cmd/webserver/ -o cmd/webserver/docs/
#gofmt -s -w cmd pkg
rm -f webserver
go build -o webserver ./cmd/webserver
./webserver
