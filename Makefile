BINARY_NAME = txplorer

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## install: install all dependencies
.PHONY: install
install:
	go mod tidy -e

## test: run all tests
.PHONY: test
test:
	go test -v -race -failfast -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=./tmp/coverage.out ./...
	go tool cover -html=./tmp/coverage.out

## run: run the application
.PHONY: run
run:
	go run ./cmd/main.go