SHELL = /bin/bash

.PHONY = module build-docker clean

TOPDIR := $(shell git rev-parse --show-toplevel)

SERVICENAME := api
TAG := v3

module: build-docker

build-docker:
	@GOOS="linux" go build -o $(SERVICENAME) $(TOPDIR)/cmd/api/main.go
	@docker build --rm -t $(SERVICENAME):$(TAG) .
	@rm $(SERVICENAME)

module-install: module
	@docker run --name $(SERVICENAME) -d $(SERVICENAME):$(TAG)

clean:
	@echo clean docker images
	@docker kill $(SERVICENAME)
	@docker rm $(SERVICENAME)
	@docker rmi $(SERVICENAME):$(TAG)
