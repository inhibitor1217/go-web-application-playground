GIT_REF := $(shell git rev-parse --short HEAD)
PROJECT_PATH := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
STAGE ?= development
TARGET ?= ./target
TARGET_BIN ?= ${TARGET}/bin/$(shell basename ${PROJECT_PATH})

init:
	go mod download

build: init
	go build -o ${TARGET_BIN} ${PROJECT_PATH}/cmd

run:
	APP_BUILD=${GIT_REF} STAGE=${STAGE} ${TARGET_BIN}

.PHONY: docs
docs:
	swag init -g cmd/main.go -o docs

clean: clean/build clean/docs

clean/build:
	rm -rf ${TARGET}

clean/docs:
	rm -rf docs/
