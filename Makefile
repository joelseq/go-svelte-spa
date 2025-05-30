all: build test

build-server:
	@echo "Building Server..."
	@go build -o main examples/stdlib/main.go

build-ui:
	@echo "Building UI..."
	@cd ui && pnpm install --frozen-lockfile && pnpm run build

# Build the entire application
build: build-ui build-server

# Run the application
run:
	@go run examples/stdlib/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

.PHONY: all build build-ui build-all run test clean watch
