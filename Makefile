# Variables
BINARY_NAME := gh-pages/tic-tac-toe/web/app.wasm

# Default target
all: build

# Build the application
build:
	GOARCH=wasm GOOS=js go build -o $(BINARY_NAME) main.go

	go build
	./tic-tac-toe

# Clean the binary
clean:
	go clean
	rm -f $(BINARY_NAME)

# Format code
fmt:
	go fmt ./...

serve:
	http-server static
