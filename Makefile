BINARY_NAME=gitai

all: build

build:
	go build -o bin/$(BINARY_NAME) main.go

test: 
	go test ./...

lint:
	golangci-lint run

clean:
	rm -f bin/$(BINARY_NAME)

deps:
	go mod tidy

fmt:
	go fmt ./...

run:
	go run main.go

.PHONY: all build test lint clean deps fmt run