# Variables
BINARY_NAME= gomyloader
DOCKER_IMAGE_NAME = gomyloader

.PHONY: all build test clean

all: build test

build:
	@echo "Building the project..."
	go build -o bin/$(BINARY_NAME) cmd/main.go

test-unit:
	@echo "Running unit tests..."
	go test -v ./test/unit

test-integration:
	@echo "Running integration tests..."
	go test -v ./test/integration

clean:
	@echo "Cleaning up..."
	rm -rf bin

docker-build: build
	docker build -t $(DOCKER_IMAGE_NAME) .

docker-run:
	docker compose up -d

run: build
	./bin/$(BINARY_NAME)