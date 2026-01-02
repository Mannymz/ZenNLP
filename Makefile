.PHONY: all proto go-proto python-proto clean deps-go deps-python run-server test-client

# Default target
all: proto deps-go deps-python

# Generate protobuf files for both Go and Python
proto: go-proto python-proto

# Generate Go protobuf files
go-proto:
	@echo "Generating Go protobuf files..."
	protoc --go_out=go-sdk/api --go_opt=paths=source_relative \
		--go-grpc_out=go-sdk/api --go-grpc_opt=paths=source_relative \
		go-sdk/api/nlp.proto
	@echo "Go protobuf files generated successfully"

# Generate Python protobuf files
python-proto:
	@echo "Generating Python protobuf files..."
	python -m grpc_tools.protoc \
		--proto_path=api \
		--python_out=nlp-engine \
		--grpc_python_out=nlp-engine \
		api/nlp.proto
	@echo "Python protobuf files generated successfully"

# Install Go dependencies
deps-go:
	@echo "Installing Go dependencies..."
	cd go-sdk && go mod tidy
	cd go-sdk && go mod download
	@echo "Go dependencies installed"

# Install Python dependencies
deps-python:
	@echo "Installing Python dependencies..."
	cd nlp-engine && pip install -r requirements.txt
	@echo "Python dependencies installed"

# Run the Python gRPC server
run-server:
	@echo "Starting NLP gRPC server..."
	cd nlp-engine && python server.py

# Test the Go client
test-client:
	@echo "Testing Go client..."
	cd go-sdk && go run client.go

# Clean generated files
clean:
	@echo "Cleaning generated files..."
	rm -f api/*.pb.go
	rm -f nlp-engine/*_pb2.py
	rm -f nlp-engine/*_pb2_grpc.py
	rm -f go-sdk/*.pb.go
	@echo "Clean completed"

# Help target
help:
	@echo "Available targets:"
	@echo "  all          - Generate protobuf files and install dependencies"
	@echo "  proto        - Generate protobuf files for both Go and Python"
	@echo "  go-proto     - Generate Go protobuf files only"
	@echo "  python-proto - Generate Python protobuf files only"
	@echo "  deps-go      - Install Go dependencies"
	@echo "  deps-python  - Install Python dependencies"
	@echo "  run-server   - Start the Python gRPC server"
	@echo "  test-client  - Test the Go client"
	@echo "  clean        - Clean generated files"
	@echo "  help         - Show this help message"
