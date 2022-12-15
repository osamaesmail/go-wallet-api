.PHONY: grpc rest buf mocks lint test

grpc:
	./debug-grpc.sh

rest:
	./debug-rest.sh

buf:
	./pb/compile.sh

mocks:
	mockery --all --recursive --keeptree

lint:
	golangci-lint run

test:
	go test -v -cover ./...