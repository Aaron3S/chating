
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

BASEPATH := $(shell pwd)
BUILDDIR=$(BASEPATH)/dist
SERVER_MAIN= $(BASEPATH)/cmd/server/server.go
CLIENT_MAIN= $(BASEPATH)/cmd/client/client.go
SERVER_NAME=chat-server
CLIENT_NAME=chat

build_server:
	GOOS=linux GOARCH=amd64 $(GOBUILD)  -o $(BUILDDIR)/server/$(SERVER_NAME)_linux_amd64 $(SERVER_MAIN)

build_client:
	GOOS=darwin GOARCH=arm64 $(GOBUILD)  -o $(BUILDDIR)/client/$(CLIENT_NAME)_darwin_arm64 $(CLIENT_MAIN)
	GOOS=darwin GOARCH=amd64 $(GOBUILD)  -o $(BUILDDIR)/client/$(CLIENT_NAME)_darwin_amd64 $(CLIENT_MAIN)

build_all: build_server build_client