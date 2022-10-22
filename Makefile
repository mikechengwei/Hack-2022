GOPATH := $(shell go env GOPATH)
GOVER_MAJOR := $(shell go version | sed -E -e "s/.*go([0-9]+)[.]([0-9]+).*/\1/")
GOVER_MINOR := $(shell go version | sed -E -e "s/.*go([0-9]+)[.]([0-9]+).*/\2/")
GO113 := $(shell [ $(GOVER_MAJOR) -gt 1 ] || [ $(GOVER_MAJOR) -eq 1 ] && [ $(GOVER_MINOR) -ge 13 ]; echo $$?)
ifeq ($(GO113), 1)
$(error Please upgrade your Go compiler to 1.13 or higher version)
endif

# Enable GO111MODULE=on explicitly, disable it with GO111MODULE=off when necessary.
export GO111MODULE := on
GOOS := $(if $(GOOS),$(GOOS),linux)
GOARCH := $(if $(GOARCH),$(GOARCH),amd64)
GOENV  := GO15VENDOREXPERIMENT="1" CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH)
GO     := $(GOENV) go
GO_BUILD := $(GO) build -trimpath
GO_SUBMODULE_DIRS = pkg/apis pkg/client

DOCKER_REGISTRY ?= localhost:5000
DOCKER_REPO ?= ${DOCKER_REGISTRY}/mcloud
IMAGE_TAG ?= latest

# NOTE: coverage report generated for E2E tests (with `-c`) may not stable, see
# https://github.com/golang/go/issues/23883#issuecomment-381766556
GO_TEST := $(GO) test -cover -covermode=atomic

default: build



build: server_client server

server_client:
	$(GO_BUILD) -o images/bin/client client/main.go
server:
	$(GO_BUILD) -o images/bin/server main.go



