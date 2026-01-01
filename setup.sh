#!/bin/bash

# Setup script for ZenNLP
# This script sets up the development environment

set -e

echo "Setting up ZenNLP development environment..."

# Build demo.go
echo "Building demo.go..."
go build -o demo demo.go

echo "Setup complete!"
