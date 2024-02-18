# Variables
APP_NAME := github.com/spleeroosh/go-translate
DOCKER_IMAGE := github.com/spleeroosh/go-translate
DOCKER_APP_IMAGE := go-translate-app
ENV_FILE := .env
TAG := latest

.PHONY: all build run dockerize clean docker-up docker-down docker-update dev-local

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
	@docker build --build-arg ENV_FILE=$(ENV_FILE) -t $(DOCKER_IMAGE):$(TAG) .

docker-up:
	@echo "Starting Docker containers..."
	@docker-compose up -d

docker-down:
	@echo "Stopping Docker containers..."
	@docker-compose down

docker-update: generate
	@docker-compose down
	@docker rmi ${DOCKER_IMAGE} ${DOCKER_APP_IMAGE}
	@docker image prune -f
	@echo "Building Docker image..."
	@docker build --build-arg ENV_FILE=$(ENV_FILE) --build-arg TIMESTAMP=$(shell date +%s) -t $(DOCKER_IMAGE):$(TAG) .
	@echo "Restarting Docker containers..."
	@docker-compose up -d
	@echo "Docker containers updated with the latest image."

dev-local:
	@echo "Building and running the application locally..."
	@go run cmd/main.go

clean:
	@echo "Cleaning up..."
	@rm -f $(APP_NAME)
