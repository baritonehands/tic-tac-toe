# Variables
BINARY_NAME := gh-pages/tic-tac-toe/web/app.wasm

# Default target
all: build

# Build the application
build:
	GOARCH=wasm GOOS=js go build -o $(BINARY_NAME) main_wasm.go

	go run main.go

wgo:
	GOARCH=wasm GOOS=js wgo go build -o $(BINARY_NAME) main_wasm.go \
		:: GOARCH=arm64 GOOS=darwin wgo run main.go

# Clean the binary
clean:
	go clean
	rm -f $(BINARY_NAME)
	rm gh-pages/tic-tac-toe/*

# Format code
fmt:
	go fmt ./...

serve:
	http-server static
