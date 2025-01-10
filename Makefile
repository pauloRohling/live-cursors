PACKAGES = ./internal/application/... ./internal/domain/...

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## install: install all dependencies
.PHONY: install
install:
	go mod tidy -e

## container: build docker image
.PHONY: container
container:
	-docker rmi live-cursors
	docker build -t live-cursors .
	docker run --rm -p 8080:8080 live-cursors

## test: run all tests
.PHONY: test
test:
	go test -race -failfast -buildvcs $(PACKAGES)

## test/c: run all tests and display coverage
.PHONY: test/c
test/c:
	go test -v -race -buildvcs -coverprofile=./tmp/coverage.out $(PACKAGES)
	go tool cover -html=./tmp/coverage.out

## test/v: run all tests in verbose mode
.PHONY: test/v
test/v:
	go test -v -race -failfast -buildvcs $(PACKAGES)

## run: run the application
.PHONY: run
run:
	air -c .air.toml