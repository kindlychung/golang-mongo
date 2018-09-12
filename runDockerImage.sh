#!/usr/bin/env bash

docker run -d --name golang-mongo-example -p 8000:8000 keenzien/golang-mongo-example
docker run 
docker logs golang-mongo-example