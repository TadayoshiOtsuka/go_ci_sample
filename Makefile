SHELL=/bin/bash

.PHONY: $(shell egrep -oh ^[a-zA-Z0-9][a-zA-Z0-9_-]+: $(MAKEFILE_LIST) | sed 's/://')

build:
	docker compose --env-file ./.env.dev build --no-cache

build-test:
	docker compose -f docker-compose.test.yaml --env-file ./.env.test build

up:
	docker compose --env-file ./.env.dev up

up-d:
	docker compose --env-file ./.env.dev up -d

seed: build
	docker compose --env-file ./.env.dev run --rm server go run app/cmd/seed/main.go

migrate: build
	docker compose --env-file ./.env.dev run --rm migrate

ent: build
	docker compose --env-file ./.env.dev run --rm server go generate ./app/pkg/rdb/orm/ent

test:
	docker compose -f docker-compose.test.yaml --env-file .env.test up -d
	docker compose -f docker-compose.test.yaml --env-file .env.test exec server_test go test -v -coverpkg=./... ./.../test

test-c:
	go test -v -coverprofile=cover_file.out -coverpkg=./... ./.../test
	go tool cover -html=cover_file.out -o cover_file.html
