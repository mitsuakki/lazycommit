GO ?= go
EXEC := $(shell basename $(CURDIR))
GOFILES := $(shell find . -name "*.go" -type f)
MAIN_PKG := ./cmd

all: lint test run

lint: $(GOFILES)
	$(GO) fmt ./...
	$(GO) vet ./...
#   $(GO) golangci-lint run

test: $(GOFILES)
	$(GO) test -v ./...

build: $(GOFILES)
	$(GO) build -o $(EXEC) $(MAIN_PKG)

run: build
	chmod +x ./$(EXEC); ./$(EXEC)

clean:
	rm -f $(EXEC)