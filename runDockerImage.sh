#!/usr/bin/env bash

dockerKillall.jl
docker run -d --name mongodb1 mongo:3-stretch
docker run -d --name golang-mongo-example -p 8000:8000 --link mongodb1:mongo keenzien/golang-mongo-example
docker logs golang-mongo-example
curl -v -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"name":"Amila","gender":"Female", "Age":34}' http://localhost:8000/users