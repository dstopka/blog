BIN ?= $(CURDIR)/bin

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

.PHONY: build-query-post
build-query-post: $(BIN)
	@go build -o $(BIN)/post $(CURDIR)/cmd/post/post.go

.PHONY: build-query-posts-page
build-query-posts-page: $(BIN)
	@go build -o $(BIN)/posts $(CURDIR)/cmd/posts/posts.go

.PHONY: build
build: build-query-post build-query-posts-page

$(BIN):
	@mkdir -p $(BIN)
