GOBIN ?= $(shell go env GOPATH)/bin

all: test

$(GOBIN)/gofumpt:
	go install mvdan.cc/gofumpt@latest

$(GOBIN)/golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin latest

test: $(GOBIN)/gofumpt $(GOBIN)/golangci-lint
	gofumpt -l -w .
	golangci-lint run
	go test ./...
	go mod tidy
