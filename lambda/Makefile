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
	@go build -o $(BIN)/query_post $(CURDIR)/cmd/query_post/query_post.go

.PHONY: build-query-posts-page
build-query-posts-page: $(BIN)
	@go build -o $(BIN)/query_posts_page $(CURDIR)/cmd/query_posts_page/query_posts_page.go

.PHONY: build
build: build-query-post build-query-posts-page

$(BIN):
	@mkdir -p $(BIN)
