all: run

ensure-oapi-codegen:
ifeq (, $(shell type -p oapi-codegen))
	$(warning oapi-codegen not found, installing...)
	@go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
endif
.PHONY: ensure-oapi-codegen

gen: ensure-oapi-codegen
	@oapi-codegen -old-config-style -generate types,spec,server -package main -o openapi.gen.go openapi.yaml
.PHONY: gen

test:
	@go test ./...
.PHONY: test

run:
	@go run .
.PHONY: run
