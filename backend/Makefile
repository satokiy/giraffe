GO				    ?= go
GOIMPORTS		  ?= goimports
GOLANGCI_LINT	?= golangci-lint
AIR           ?= ${GO} run github.com/cosmtrek/air

## $make build
.PHONY: build
build: giraffe

giraffe: $(shell find . -name '*.go')
		${GO} build -o ./cmd/main main.go

## $make run
.PHONY: run
run:
	${AIR} -c .air.toml

## $ make fmt
# ソースコードの自動整形をする
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
