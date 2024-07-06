.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run: ## Run thumbnail using test-subject.jpg
	@go run main.go -s=160x224 test-subject.jpg

.PHONY: run

test: ## Run tests
	@go test ./...

.PHONY: test
