FUNCTIONS_BIN := $(CURDIR)/functions

.PHONY: build-lambda
build-lambda:
	make -C $(CURDIR)/lambda build BIN=$(FUNCTIONS_BIN)

.PHONY: build-frontend
build-frontend:
	@cd frontend && npm install && npm run build

.PHONY: build
build: build-lambda build-frontend