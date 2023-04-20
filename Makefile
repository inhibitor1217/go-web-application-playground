PROJECT_PATH := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
STAGE ?= development
TARGET ?= ./target
TARGET_BIN ?= ${TARGET}/bin/$(shell basename ${PROJECT_PATH})

init:
	go mod download

build: init
	go build -o ${TARGET_BIN} ${PROJECT_PATH}/cmd

run:
	STAGE=${STAGE} ${TARGET_BIN}

clean:
	rm -rf ${TARGET}