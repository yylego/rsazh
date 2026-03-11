.PHONY: test
COVERAGE_DIR ?= .coverage.out

test: test-with-flags

test-with-flags:
	mkdir -p $(COVERAGE_DIR)
	go test -v -race -covermode atomic -coverprofile $(COVERAGE_DIR)/combined.txt -bench=. -benchmem -timeout 20m ./...
