include golang.mk
.DEFAULT_GOAL := test # override default goal set in library makefile

.PHONY: $(PKGS) all build clean codegen deploy gen-client gen-client-python gen-server py-deps py-test run test vendor
SHELL := /bin/bash
PKG := github.com/Clever/kayvee-logger-service
PKGS := $(shell go list ./... | grep -v /vendor | grep -v /restapi)
EXECUTABLE := $(shell basename $(PKG))

$(eval $(call golang-version-check,1.10))

all: codegen build test

build:
	go build -o bin/$(EXECUTABLE) $(PKG)/cmd/kayvee-logger-service-server

clean:
	rm bin/*
	find -type f -name '*.pyc' -delete

codegen: gen-server gen-client

deploy:
	ark start kayvee-logger-service -e system

gen-client: gen-client-python

gen-client-python:
	java -jar swagger-codegen-cli.jar generate -l python -i kayvee-logger-service.yaml -o client/python -c client/python/swagger_config.json -s

gen-server:
	swagger generate server -f kayvee-logger-service.yaml

py-deps:
	sudo python setup.py develop

py-test: py-deps
	nosetests client/python/test

run:
	./bin/$(EXECUTABLE) --port=5020

test: $(PKGS) py-test

$(PKGS): golang-test-all-deps
	$(call golang-test-all,$@)




install_deps: golang-dep-vendor-deps
	$(call golang-dep-vendor)