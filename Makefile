.PHONY: help

SHELL := /bin/bash
VERSION=$(shell git describe --tags `git rev-list --tags --max-count=1`)
NOW=$(shell date +'%y-%m-%d_%H:%M:%S')
COMMIT_REF=$(shell git rev-parse --short HEAD)
BUILD_ARGS=-ldflags "-s -w -X github.com/timo-reymann/deterministic-zip/pkg/buildinfo.GitSha=$(COMMIT_REF) -X github.com/timo-reymann/deterministic-zip/pkg/buildinfo.Version=$(VERSION) -X github.com/timo-reymann/deterministic-zip/pkg/buildinfo.BuildTime=$(NOW)"
BIN_PREFIX="dist/deterministic-zip_"

clean: ## Cleanup artifacts
	@rm -rf dist/

help: ## Display this help page
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[33m%-30s\033[0m %s\n", $$1, $$2}'

coverage: ## Run tests and measure coverage
	@CGO_ENABLED=0 go test -covermode=count -coverprofile=/tmp/count.out -v ./...

test-coverage-report: coverage ## Run test and display coverage report in browser
	@go tool cover -html=/tmp/count.out

save-coverage-report: coverage ## Save coverage report to coverage.html
	@go tool cover -html=/tmp/count.out -o coverage.html

create-dist: ## Create dist folder if not already existent
	@mkdir -p dist/

build-linux: create-dist ## Build binaries for linux
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BIN_PREFIX)linux-amd64 $(BUILD_ARGS)
	@CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o $(BIN_PREFIX)linux-i386 $(BUILD_ARGS)
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o $(BIN_PREFIX)linux-arm $(BUILD_ARGS)
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o $(BIN_PREFIX)linux-arm64 $(BUILD_ARGS)

build-windows: create-dist ## Build binaries for windows
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(BIN_PREFIX)windows-amd64.exe $(BUILD_ARGS)
	@CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o $(BIN_PREFIX)windows-i386.exe $(BUILD_ARGS)
	@CGO_ENABLED=0 GOOS=windows GOARCH=arm go build -o $(BIN_PREFIX)windows-arm.exe $(BUILD_ARGS)

build-darwin: create-dist  ## Build binaries for macOS
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(BIN_PREFIX)darwin-amd64 $(BUILD_ARGS)
	@CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o $(BIN_PREFIX)darwin-arm64 $(BUILD_ARGS)

build-freebsd: create-dist ## Build binaries for FreeBSD
	@CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -o $(BIN_PREFIX)freebsd-amd64 $(BUILD_ARGS)
    @CGO_ENABLED=0 GOOS=freebsd GOARCH=386 go build -o $(BIN_PREFIX)freebsd-i386 $(BUILD_ARGS)
    @CGO_ENABLED=0 GOOS=freebsd GOARCH=arm64 go build -o $(BIN_PREFIX)freebsd-arm64 $(BUILD_ARGS)
    @CGO_ENABLED=0 GOOS=freebsd GOARCH=arm go build -o $(BIN_PREFIX)freebsd-arm $(BUILD_ARGS)

build-openbsd: create-dist ## Build binaries for OpenBSD
	@CGO_ENABLED=0 GOOS=openbsd GOARCH=amd64 go build -o $(BIN_PREFIX)openbsd-amd64 $(BUILD_ARGS)
    @CGO_ENABLED=0 GOOS=openbsd GOARCH=386 go build -o $(BIN_PREFIX)openbsd-i386 $(BUILD_ARGS)

build-docker: ## Build docker image based on the built linux builds in the dist folder
	@docker buildx build --tag timoreymann/deterministic-zip:latest \
		--platform linux/amd64,linux/arm/v7,linux/arm64 \
		--build-arg BUILD_TIME="$(NOW)" \
		--build-arg BUILD_VERSION="$(VERSION)" \
		--build-arg BUILD_COMMIT_REF="$(COMMIT_REF)" \
		--push .
	@docker buildx build --tag timoreymann/deterministic-zip:$(VERSION) \
		--platform linux/amd64,linux/arm/v7,linux/arm64 \
		--build-arg BUILD_TIME="$(NOW)" \
		--build-arg BUILD_VERSION="$(VERSION)" \
		--build-arg BUILD_COMMIT_REF="$(COMMIT_REF)" \
		--push .

create-checksums: ## Create checksums for binaries
	@find ./dist -type f -exec sh -c 'sha256sum {} | cut -d " " -f 1 > {}.sha256' {} \;

build: build-linux build-darwin build-windows build-freebsd build-openbsd create-checksums ## Build binaries for all platform
