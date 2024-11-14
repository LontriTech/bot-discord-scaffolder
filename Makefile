# https://www.gnu.org/software/make/

APP_NAME := app
CMD_DIR := ./cmd/$(APP_NAME)
OUTPUT_DIR := dist
OUTPUT_NAME := $(APP_NAME)
OUTPUT_PATH := $(OUTPUT_DIR)/$(OUTPUT_NAME)
GO_FILES := $(shell find . -type f -name '*.go')

.PHONY: all build run clean test fmt lint deps install

all: build

build: $(OUTPUT_PATH)

$(OUTPUT_PATH): $(GO_FILES)
	@echo "Building the application..."
	@mkdir -p $(OUTPUT_DIR)
	@go build -o $(OUTPUT_PATH) $(CMD_DIR)
	@echo "Build completed: $(OUTPUT_PATH)"

run: build
	@echo "Running the application..."
	@$(OUTPUT_PATH)

clean:
	@echo "Cleaning up..."
	@rm -rf $(OUTPUT_DIR)
	@echo "Clean completed."

test:
	@echo "Running tests..."
	@go test ./...

fmt:
	@echo "Formatting code..."
	@go fmt ./...

lint:
	@echo "Linting code..."
	@golangci-lint run

deps:
	@echo "Tidying dependencies..."
	@go mod tidy

install:
	@echo "Installing application..."
	@go install $(CMD_DIR)
