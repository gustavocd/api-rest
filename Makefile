SHELL := /bin/bash

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## run: run the api
.PHONY: run
run:
	@go run api-rest.go

## test: execute the tests
.PHONY: test
test:
	@go test -v -count=1 ./...
