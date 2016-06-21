# Go Commands
GOCMD=go
GODEP=godep $(GOCMD)
GOBUILD=$(GODEP) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GODEP) test -v
GOFMT=go fmt

DATE := $(shell date -u +%Y-%m-%d%Z%H:%M:%S)
APP := "winktest"

# Optional User Provided Parameters
ifndef CGO_ENABLED
	export CGO_ENABLED=0
endif

ifndef GOOS
	export GOOS=$(shell uname -s | tr '[:upper:]' '[:lower:]')
endif

ifndef BUILD_NUMBER
	BUILD_NUMBER := "0001"
endif

.PHONY: all build clean test fmt

all: test build

build: fmt ## Build the project
	$(GOBUILD) -ldflags "-X main.buildTime=$(DATE) -X main.buildNumber=$(BUILD_NUMBER)" -o $(APP) .

clean:  ## Clean out all generated files
	$(GOCLEAN)
	rm $(APP)-linux
	rm $(APP)-darwin
	rm $(APP)-windows

help:  ## Show This Help
	@for line in $$(cat Makefile | grep "##" | grep -v "grep" | sed  "s/:.*##/:/g" | sed "s/\ /!/g"); do verb=$$(echo $$line | cut -d ":" -f 1); desc=$$(echo $$line | cut -d ":" -f 2 | sed "s/!/\ /g"); printf "%-30s--%s\n" "$$verb" "$$desc"; done

release:  ## Build binaiers for the primary OS'
	GOOS=linux $(GOBUILD) -ldflags "-X main.buildTime=$(DATE) -X main.buildNumber=$(BUILD_NUMBER)" -o $(APP)-linux .
	GOOS=darwin $(GOBUILD) -ldflags "-X main.buildTime=$(DATE) -X main.buildNumber=$(BUILD_NUMBER)" -o $(APP)-darwin .
	GOOS=windows $(GOBUILD) -ldflags "-X main.buildTime=$(DATE) -X main.buildNumber=$(BUILD_NUMBER)" -o $(APP)-windows .

test: unit_test ## Run all available tests

unit_test:  ## Run unit tests
	$(GOTEST)

fmt:  ## Run go fmt
	$(GOFMT)
