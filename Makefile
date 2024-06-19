# Variables
APP_NAME := redplanet-bridge
CMD_DIR := cmd
MAIN_FILE := $(CMD_DIR)/main.go
BUILD_DIR := bin
BUILD_FILE := $(BUILD_DIR)/$(APP_NAME)

# Comandos
.PHONY: all build clean run air prod test 

all: build

# Construir el proyecto
build:
	@echo "Building the project..."
	@go build -o $(BUILD_FILE) $(MAIN_FILE)

# Limpiar archivos binarios
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@rm -rf pb_data

# Ejecutar el proyecto
run: build
	@echo "Running the project..."
	@$(BUILD_FILE) serve --dir $(BUILD_DIR)/pb_data

prod:
	@echo "Deploying on fly.io"
	@fly deploy
# Ejecutar pruebas
test:
	@echo "Running tests..."
	@go test ./...

# Ejecutar formateo de código
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Ejecutar análisis estático del código
lint:
	@echo "Linting code..."
	@golangci-lint run

# Actualizar dependencias
deps:
	@echo "Downloading dependencies..."
	@go mod tidy

# Generar binario
release:
	@echo "Building for release..."
	@GOOS=linux GOARCH=amd64 go build -o $(BUILD_FILE)-linux $(MAIN_FILE)
	@GOOS=windows GOARCH=amd64 go build -o $(BUILD_FILE)-windows.exe $(MAIN_FILE)
	@GOOS=darwin GOARCH=amd64 go build -o $(BUILD_FILE)-darwin $(MAIN_FILE)

# Ayuda
help:
	@echo "Usage:"
	@echo "  make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  all       - Build the project (default)"
	@echo "  build     - Build the project"
	@echo "  clean     - Clean up generated files"
	@echo "  run       - Run the project"
	@echo "  test      - Run tests"
	@echo "  fmt       - Format code"
	@echo "  lint      - Lint code"
	@echo "  deps      - Download dependencies"
	@echo "  release   - Build for release (cross-compile)"
	@echo "  help      - Display this help message"

