ENV				?= dev
GOIMPORTS		  ?= goimports
GOLANGCI_LINT	?= golangci-lint
AIR           ?= go run github.com/cosmtrek/air

## $make build
.PHONY: build
build: giraffe

giraffe: $(shell find . -name '*.go')
		go build -o ./cmd/main main.go

## $make run
.PHONY: run
run:
	GO_ENV=${ENV} ${AIR} -c .air.toml

## $ make fmt
.PHONY: fmt
fmt:
	${GOIMPORTS} -format-only -w .

## $ make lint
.PHONY: lint
lint: golangci-lint

## $ make golangci-lint
.PHONY: golangci-lint
golangci-lint:
	${GOLANGCI_LINT} run --allow-parallel-runners
