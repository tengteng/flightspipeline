.PHONY: default env install build doc test clean fmt vet vendor_clean

GOPATH := ${PWD}:${PWD}/_vendor
export GOPATH

default: install

env:
	@echo GOPATH is $$GOPATH

install: fmt vet
	go install ./...

build: fmt vet
	go build ./...

test:
	go test -short ./...

clean:
	go clean -i ./...
	rm -rf bin
	rm -rf pkg
	rm -rf _vendor/pkg

fmt:
	go fmt ./...

vet:
	go vet ./...

vendor_clean:
	rm -rf ./_vendor/src

