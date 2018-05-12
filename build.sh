#!/usr/bin/env bash
env GOOS=linux GOARCH=amd64 go build -o build/server
docker build . -t to-go-server
rm -rf build
