# Project information
PACKAGE=gyazo

# Tasks
help:
	@echo "Please type: make [target]"
	@echo "  test         Run tests"
	@echo "  deps         Install runtime dependencies"
	@echo "  updatedeps   Update runtime dependencies"
	@echo "  help         Show this help messages"

test: deps
	@echo "===> Running tests..."
	go test -v ./${PACKAGE} -cover

deps:
	@echo "===> Installing runtime dependencies..."
	go get -v ./...

updatedeps:
	@echo "===> Updating runtime dependencies..."
	go clean ./...
	go get -u -v ./...

.PHONY: help test deps updatedeps
