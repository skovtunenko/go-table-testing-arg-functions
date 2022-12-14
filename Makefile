export GO111MODULE=on

PROJECT_PATH=$(shell dirname $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST)))))

.PHONY: test
test: ## Run tests for all packages (test cache is disabled)
	go test -count=1 -race -trimpath -cover ./...

.PHONY: code-quality
code-quality: ## Print code smells using Golangci-lint
	golangci-lint --timeout 1m --out-format tab run ./...

.PHONY: generate
generate: ## Generate mock files
	go generate ./...

.PHONY: install-tools
install-tools: ## Install dependencies locally
	@echo "Installing tools locally..."
	@go install github.com/golang/mock/mockgen@v1.6.0
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.47.3
	@echo "Done."

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'