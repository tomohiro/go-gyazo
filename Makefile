# Tasks
.PHONY: help
help:
	@echo "Please type: make [target]"
	@echo "  deps         Install runtime dependencies"
	@echo "  updatedeps   Update runtime dependencies"
	@echo "  lint 		  Lint codes"
	@echo "  test         Run tests"
	@echo "  help         Show this help messages"

.PHONY: deps
deps:
	@echo "===> Installing runtime dependencies..."
	go mod download

.PHONY: updatedeps
updatedeps:
	@echo "===> Updating runtime dependencies..."
	go get -u ./...

.PHONY: lint
lint: deps
	@echo "===> Running lint..."
	go vet ./...
	golint -set_exit_status ./...

.PHONY: test
test: deps
	@echo "===> Running tests..."
	go test -v ./... -cover
