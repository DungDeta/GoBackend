GOOSE_DBSTRING ?= root:root1234@tcp(localhost:33306)/shopdevgo
GOOSE_MIGRATION_DIR ?= ./sql/schema
GOOSE_DRIVER ?= mysql

APP_NAME := server
PI_USER = pi
PI_HOST = pi.local
PI_DEPLOY_DIR = /home/pi/GoLandProject/API_Backend
PI_BINARY_NAME = myapp

build-pi:
	@echo "Building for Raspberry Pi (ARM64)..."
	set GOOS=linux&& set GOARCH=arm64&& set CGO_ENABLED=0&& go build -ldflags="-s -w" -o ./bin/$(PI_BINARY_NAME) ./cmd/$(APP_NAME)
	@echo "Build completed: .\bin\$(PI_BINARY_NAME)"

deploy-pi: build-pi
	@echo "Deploy lên Raspberry Pi..."
	scp ./bin/$(PI_BINARY_NAME) $(PI_USER)@$(PI_HOST):$(PI_DEPLOY_DIR)/ || exit 1
	@echo "Deployed thành công $(PI_DEPLOY_DIR)/$(PI_BINARY_NAME)"

run-pi:
	@echo "Chạy app trên Raspberry Pi..."
	ssh $(PI_USER)@$(PI_HOST) "cd $(PI_DEPLOY_DIR) && chmod +x $(PI_BINARY_NAME) && ./$(PI_BINARY_NAME)"

deploy-run-pi: deploy-pi run-pi
docker_build:
	docker-compose up -d --build
	docker-compose ps

docker_stop:
	docker-compose down

dev:
	go run ./cmd/$(APP_NAME)

build:
	go build -o ./bin/$(APP_NAME) ./cmd/$(APP_NAME)

docker_up:
	docker-compose -f /enviroment/docker-compose-dev.yml compose up

up_by_one:
	@goose -dir=$(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) "$(GOOSE_DBSTRING)" up-by-one

create_migration:
	@goose -dir=$(GOOSE_MIGRATION_DIR) create $(name) sql

upse:
	@goose -dir=$(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) "$(GOOSE_DBSTRING)" up

downse:
	@goose -dir=$(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) "$(GOOSE_DBSTRING)" down

resetse:
	@goose -dir=$(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) "$(GOOSE_DBSTRING)" reset

sqlgen:
	sqlc generate

swag:
	swag init -g ./cmd/server/main.go -o ./cmd/swag/docs

.PHONY: dev downse upse resetse docker_build docker_stop docker_up swag

.PHONY: air