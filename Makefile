include golang.mk
.DEFAULT_GOAL := test # override default goal set in library makefile

.PHONY: clean all gen-server build test $(PKGS) vendor run
SHELL := /bin/bash
PKG := github.com/Clever/kayvee-logger-service
PKGS := $(shell go list ./... | grep -v /vendor | grep -v /restapi)
EXECUTABLE := $(shell basename $(PKG))

$(eval $(call golang-version-check,1.5))

clean:
	rm bin/*

all: gen-server build test

gen-server:
	swagger generate server -f kayvee-logger-service.yaml

build:
	go build -o bin/$(EXECUTABLE) $(PKG)/cmd/kayvee-logger-service-server

test: $(PKGS)

$(PKGS): golang-test-all-deps
	$(call golang-test-all,$@)

vendor: golang-godep-vendor-deps
	$(call golang-godep-vendor,$(PKGS))

run:
	./bin/$(EXECUTABLE) --port=5020
