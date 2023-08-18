GOBIN ?= $(shell go env GOPATH)/bin
GODIRS ?= $(shell go list -f "{{.Dir}}" ./... | grep -v "/pkg/apis")

all: proto test

$(GOBIN)/gofumpt:
	go install mvdan.cc/gofumpt@latest

$(GOBIN)/golangci-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

$(GOBIN)/buf:
	go install github.com/bufbuild/buf/cmd/buf@latest

$(GOBIN)/protoc-gen-go:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

$(GOBIN)/protoc-gen-connect-go:
	go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest

proto: $(GOBIN)/buf $(GOBIN)/protoc-gen-go $(GOBIN)/protoc-gen-connect-go
	buf format -w
	buf generate --template build/buf.gen.yaml

test: $(GOBIN)/gofumpt $(GOBIN)/golangci-lint
	gofumpt -l -w $(GODIRS)
	golangci-lint run
	go test ./...
	go mod tidy
