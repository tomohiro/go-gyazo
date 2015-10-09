help:
	@echo "Please type: make [target]"
	@echo "  test         Run tests"
	@echo "  deps         Install runtime dependencies"
	@echo "  updatedeps   Update runtime dependencies"
	@echo "  build        Build"
	@echo "  clean        Cleanup artifacts"
	@echo "  help         Show this help messages"

test: deps
	@echo "===> Running tests..."
	go test -v ./...

deps:
	@echo "===> Installing runtime dependencies..."
	go get -v ./...

updatedeps:
	@echo "===> Updating runtime dependencies..."
	go clean ./...
	go get -u -v ./...

build: deps
	@echo "===> Beginning compile..."
	go build

clean:
	go clean ./...

.PHONY: help test setup deps updatedeps build clean
