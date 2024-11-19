GO ?= /usr/local/go/bin/go
EXEC := $(shell basename $(CURDIR))
GOFILES := $(shell find . -name "*.go" -type f)

all: lint test build

lint: $(GOFILES)
	$(GO) fmt ./...
	$(GO) vet ./...

test: $(GOFILES)
	$(GO) test -v ./...

build: $(GOFILES)
	$(GO) mod tidy
	$(GO) mod download

	mkdir -p bin
	$(GO) build -o bin/$(EXEC) .

docs: $(GOFILES)
	doxygen doxygen/doxygen.conf

clean:
	rm -f bin/