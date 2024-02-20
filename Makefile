.PHONY: test
test:
	@echo "Running unit tests for secondary adapters"
	@go test -v ./pkg/adapters/secondary/...
	@echo "Running integration simulation tests"
	@go test -v ./tests/...