.PHONY: dev grpc rest buf mocks lint test


help: ## Show Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

dev: ## run docker in dev mode
	docker-compose up -d

down: ## stop docker
	docker-compose down

grpc: ## run gRPC server
	./debug-grpc.sh

rest: ## run REST server
	./debug-rest.sh

buf: ## generate buf code
	./pb/compile.sh

buf-lent: ## lint buf code
	buf lint

buf-breaking: ## check breaking changes
	buf breaking --against '.git#branch=master'

mocks: ## generate mocks
	mockery --all --recursive --keeptree

lint: ## lint go code
	golangci-lint run

test: ## run tests
	go test ./...