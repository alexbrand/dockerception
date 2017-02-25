#!/bin/bash

docker build -t dockerception .

docker run                                                         \
	--rm                                                       \
	-v /var/run/docker.sock:/var/run/docker.sock               \
	--privileged                                               \
	dockerception                                              \
        go run /go/src/github.com/alexbrand/dockerception/main.go
