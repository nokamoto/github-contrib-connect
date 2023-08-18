GOBIN ?= $(shell go env GOPATH)/bin

all: test

$(GOBIN)/gofumpt:
	go install mvdan.cc/gofumpt@latest

$(GOBIN)/golangci-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

test: $(GOBIN)/gofumpt $(GOBIN)/golangci-lint
	gofumpt -l -w .
	golangci-lint run
	go test ./...
	go mod tidy
