.PHONY: help

SHELL := /bin/bash
BUILD_ARGS=-ldflags "-X github.com/timo-reymann/deterministic-zip/pkg/buildinfo.GitSha=$(shell git rev-parse --short HEAD) -X github.com/timo-reymann/deterministic-zip/pkg/buildinfo.Version=$(shell git describe --tags `git rev-list --tags --max-count=1`) -X github.com/timo-reymann/deterministic-zip/pkg/buildinfo.BuildTime=$(NOW)"
NOW=$(shell date +'%y-%m-%d_%H:%M:%S')

help: ## Display this help page
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[33m%-30s\033[0m %s\n", $$1, $$2}'

coverage: ## Run tests and measure coverage
	@go test -covermode=count -coverprofile=/tmp/count.out -v ./...

test-coverage-report: coverage ## Run test and display coverage report in browser
	@go tool cover -html=/tmp/count.out

save-coverage-report: coverage ## Save coverage report to coverage.html
	@go tool cover -html=/tmp/count.out -o coverage.html

create-dist: ## Create dist folder if not already existent
	@mkdir -p dist/

build-linux: create-dist ## Build binaries for linux (amd64, arm64)
	@GOOS=linux GOARCH=amd64 go build -o dist/webp-utils-linux-amd64 $(BUILD_ARGS)
	@GOOS=linux GOARCH=arm64 go build -o dist/webp-utils-linux-arm64 $(BUILD_ARGS)

build-windows: create-dist ## Build binaries for windows (amd64)
	@GOOS=windows GOARCH=amd64 go build -o dist/webp-utils-windows-amd64.exe $(BUILD_ARGS)

build-darwin: create-dist  ## Build binaries for macOS (amd64, arm64)
	@GOOS=darwin GOARCH=amd64 go build -o dist/webp-utils-darwin-amd64 $(BUILD_ARGS)
	@GOOS=darwin GOARCH=amd64 go build -o dist/webp-utils-darwin-arm64 $(BUILD_ARGS)

build: build-linux build-darwin build-windows ## Build binaries for all platform
