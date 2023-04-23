.PHONY: test
test:
	@go test ./...

.PHONY: fmt
fmt:
	@goimports -w .

.PHONY: generate
generate:
	@go generate ./...

.PHONY: lint
lint:
	@golangci-lint run