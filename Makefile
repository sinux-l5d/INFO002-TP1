BINARY_NAME ?= rbt
ENTRYPOINT ?= ./cmd/$(BINARY_NAME)

.PHONY: build
build:
	@echo "Building $(BINARY_NAME)..."
	go build -o ./bin/$(BINARY_NAME) -v $(ENTRYPOINT)

.PHONY: download
download:
	@echo "Downloading dependencies..."
	go mod download