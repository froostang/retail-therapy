# Makefile for building the API service for Ubuntu 64-bit architecture

# Variables
BINARY_NAME=api-service
BUILD_DIR=build
OS=linux
ARCH=amd64

# Build command
build:
	GOOS=$(OS) GOARCH=$(ARCH) go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd

.PHONY: build
