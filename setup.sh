#!/bin/bash

echo "Setting up ZenNLP Project..."

# Install Python dependencies
echo "Installing Python dependencies..."
pip install grpcio grpcio-tools transformers torch numpy

# Generate Python protobuf files
echo "Generating Python protobuf files..."
python -m grpc_tools.protoc --proto_path=api --python_out=nlp-engine --grpc_python_out=nlp-engine api/nlp.proto

# Setup Go modules
echo "Setting up Go modules..."
cd api && go mod tidy
cd ../go-sdk && go mod tidy
cd ../examples && go mod tidy

echo "Setup complete!"
echo ""
echo "To run the project:"
echo "1. Start the Python server: cd nlp-engine && python server.py"
echo "2. Run the Go example: cd examples && go run test.go"
