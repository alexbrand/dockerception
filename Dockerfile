FROM golang:1.8.0

RUN curl https://get.docker.com/builds/Linux/x86_64/docker-1.12.2.tgz -O && tar -xf docker-1.12.2.tgz

RUN mv docker/docker /usr/bin/docker

ADD . /go/src/github.com/alexbrand/dockerception
