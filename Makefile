.PHONY: build
build:
	DOCKER_DEFAULT_PLATFORM=linux/amd64 docker compose build

.PHONY: start
start:
	DOCKER_DEFAULT_PLATFORM=linux/amd64 docker compose up -d

.PHONY: down
down:
	docker compose down -v
