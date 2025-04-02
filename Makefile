CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})

APP_CMD_DIR=${CURRENT_DIR}/cmd

IMG_NAME=${APP}
REGISTRY=${REGISTRY}
TAG=${TAG}
TAG_LATEST=${TAG_LATEST}
GITLAB_URL=https://gitlab.hamkorbank.uz/
GITLAB_TOKEN=eefzzBNVwDb1-x8FR69H
PKG_LIST := $(shell go list ${CURRENT_DIR}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
YAML_FILES :=  $(shell find .  -iname *-ci.yml  | grep -v /vendor/)


.PHONY: all build build-image push-image swag-init dep build clean test coverage coverhtml lint-go lint-yaml

# # Including
# include .build_info

all: build

build: ## build
	CGO_ENABLED=0 GOOS=linux go build -mod=mod -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

swag-init: ## init swagger
	swag init -g cmd/api/main.go -o api/docs

run: ## run application
	go run cmd/api/main.go

test: ## Run unittests
	@go test -short ${PKG_LIST}

race: dep ## Run data race detector
	@go test -race -short ${PKG_LIST}

migrate-up:
	@export $(shell cat .env | xargs) && \
	migrate -path ./migrations -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable' up

migrate-down:
	@export $(shell cat .env | xargs) && \
	migrate -path ./migrations -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable' down

.PHONY: compose-up
compose-up: ## Bring up docker-compose services
	@echo "Bringing up docker-compose services..."
	@docker-compose down
	@docker-compose up -d --build

.PHONY: compose-down
compose-down: ## Bring down docker-compose services
	@echo "Bringing down docker-compose services..."
	@docker-compose down

