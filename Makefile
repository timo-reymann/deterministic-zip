.PHONY: help

SHELL := /bin/bash

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
