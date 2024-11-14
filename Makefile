GO ?= go
EXEC := $(shell basename $(CURDIR))
GOFILES := $(shell find . -name "*.go" -type f)

all: lint test run

lint: $(GOFILES)
	$(GO) fmt ./...
	$(GO) vet ./...
#   $(GO) golangci-lint run

test: $(GOFILES)
	$(GO) test -v ./...

build: $(GOFILES)
	$(GO) build -o $(EXEC) .

run: build
	chmod +x ./$(EXEC); ./$(EXEC)

clean:
	rm -f $(EXEC)