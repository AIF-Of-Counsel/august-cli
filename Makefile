.PHONY: build test lint install clean

build:
	go build -o bin/august-pp-cli ./cmd/august-pp-cli

test:
	go test ./...

lint:
	golangci-lint run

install:
	go install ./cmd/august-pp-cli

clean:
	rm -rf bin/

build-mcp:
	go build -o bin/august-pp-mcp ./cmd/august-pp-mcp

install-mcp:
	go install ./cmd/august-pp-mcp

build-all: build build-mcp
