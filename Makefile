# Variables
APP_NAME := github.com/spleeroosh/go-translate
DOCKER_IMAGE := github.com/spleeroosh/go-translate
ENV_FILE := .env

.PHONY: all build run dockerize clean docker-up docker-down dev-local

all: build

generate:
	@templ generate

build:
	@echo "Building the Go application..."
	@go build -o $(APP_NAME) ./cmd

run:
	@echo "Running the application..."
	@./$(APP_NAME)

dockerize:
	@echo "Building Docker image..."
	@docker build --build-arg ENV_FILE=$(ENV_FILE) -t $(DOCKER_IMAGE) .

docker-up:
	@echo "Starting Docker containers..."
	@docker-compose up -d

docker-down:
	@echo "Stopping Docker containers..."
	@docker-compose down

dev-local:
	@echo "Building and running the application locally..."
	@go run cmd/main.go

clean:
	@echo "Cleaning up..."
	@rm -f $(APP_NAME)
