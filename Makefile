## Server commands
.PHONY: lint
lint:
	@echo "Linting code using golangci-lint"
	golangci-lint run ./...

.PHONY: run-server
run-server:
	@echo "Starting Mirai server"
	go run cmd/main.go

## Development Server docker-compose commands
DOCKER_COMPOSE_DEV_FILE=docker-compose-dev.yml
DOCKER_COMPOSE_DEV_SERVER_SERVICE=mirai_api

.PHONY: compose-dev-build
compose-dev-build:
	@echo "Running Mirai API docker compose dev build"
	docker compose -f $(DOCKER_COMPOSE_DEV_FILE) build

.PHONY: compose-dev-up
compose-dev-up:
	@echo "Running Mirai API docker compose dev up in detach mode"
	docker compose -f $(DOCKER_COMPOSE_DEV_FILE) up -d

.PHONY: compose-dev-logs
compose-dev-logs:
	@echo "Running Mirai API docker compose dev logs"
	docker compose -f $(DOCKER_COMPOSE_DEV_FILE) logs $(DOCKER_COMPOSE_DEV_SERVER_SERVICE) -f

.PHONY: compose-dev-down
compose-dev-down:
	@echo "Running Mirai API docker compose dev down"
	docker compose -f $(DOCKER_COMPOSE_DEV_FILE) down
