
GOPATH=$(shell go env GOPATH)
export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on

all: modules test check

.PHONY: modules
modules:
		go mod tidy

.PHONY: test
test:
		go test ./...
		go vet ./...

.PHONY: check
check: test
		golangci-lint run

.PHONY: coverage
coverage:
		mkdir -p coverage
		go test -cover ./...
		go test -coverprofile=coverage/coverage.out ./...
		go tool cover -func=coverage/coverage.out
		go tool cover -html=coverage/coverage.out
