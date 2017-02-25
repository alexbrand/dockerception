#!/bin/bash

docker build . -t nested-docker

CONTAINER_ID=$(docker run --privileged -v /sys/fs/cgroup:/sys/fs/cgroup:ro -d nested-docker)

# Wait for docker daemon to start up
sleep 3 

docker exec $CONTAINER_ID docker run busybox echo nested-docker!
