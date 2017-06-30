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
	dep ensure

updatedeps:
	@echo "===> Updating runtime dependencies..."
	dep ensure -update

.PHONY: help test deps updatedeps
