FUNCTIONS_BIN := $(CURDIR)/functions

.PHONY: build-lambda
build-lambda:
	make -C $(CURDIR)/lambda build BIN=$(FUNCTIONS_BIN)