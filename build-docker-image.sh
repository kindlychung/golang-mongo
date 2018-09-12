#!/usr/bin/env bash

# docker rmi $(docker images -q)
dockerKillall.jl
docker rmi keenzien/golang-mongo-example
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .
docker build -t keenzien/golang-mongo-example .
docker push keenzien/golang-mongo-example