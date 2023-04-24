GIT_REF := $(shell git rev-parse --short HEAD)
PROJECT_PATH := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
STAGE ?= development
TARGET ?= ./target
TARGET_BIN ?= ${TARGET}/bin/$(shell basename ${PROJECT_PATH})

init: pre-commit/setup
	@echo "Installing dependencies ..."
	go mod download

	@echo "Setup containers for development ..."
	docker-compose -f setup/docker-compose.yml build
	docker-compose -f setup/docker-compose.yml up -d

build:
	go build -o ${TARGET_BIN} ${PROJECT_PATH}/cmd

run:
	APP_BUILD=${GIT_REF} STAGE=${STAGE} ${TARGET_BIN}

dev: docs build run

docs: docs/generate

.PHONY: docs
docs/generate:
	swag init -d ./cmd,./api/public -o docs

docs/fmt:
	@echo "Formatting docs ..."
	swag fmt -d ./

pre-commit: docs/fmt docs/generate
	git add .

pre-commit/setup:
	@echo "Setting up pre-commit hook ..."
	cp .githooks/pre-commit .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit

clean: clean/build clean/docs

clean/build:
	rm -rf ${TARGET}

clean/docs:
	rm -rf docs/

shutdown:
	@echo "Shutting down containers ..."
	docker-compose -f $(dirname $0)/setup/docker-compose.yml down
