include golang.mk
.DEFAULT_GOAL := test # override default goal set in library makefile

.PHONY: $(PKGS) all build clean codegen gen-client gen-client-python gen-server run test vendor
SHELL := /bin/bash
PKG := github.com/Clever/kayvee-logger-service
PKGS := $(shell go list ./... | grep -v /vendor | grep -v /restapi)
EXECUTABLE := $(shell basename $(PKG))

$(eval $(call golang-version-check,1.5))

all: codegen build test

build:
	go build -o bin/$(EXECUTABLE) $(PKG)/cmd/kayvee-logger-service-server

clean:
	rm bin/*

codegen: gen-server gen-client

gen-client: gen-client-python

gen-client-python:
	java -jar swagger-codegen-cli.jar generate -l python -i kayvee-logger-service.yaml -o client/python -c client/python/swagger_config.json -s

gen-server:
	swagger generate server -f kayvee-logger-service.yaml

py-deps:
	python client/python/setup.py develop

run:
	./bin/$(EXECUTABLE) --port=5020

test: $(PKGS)

$(PKGS): golang-test-all-deps
	$(call golang-test-all,$@)

vendor: golang-godep-vendor-deps
	$(call golang-godep-vendor,$(PKGS))

