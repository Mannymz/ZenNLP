#!/bin/bash
# Project Validation Script for ZenNLP
# Generated: 2026-01-01 17:21:45 UTC

set -e

echo "================================"
echo "ZenNLP Project Validation"
echo "================================"
echo ""

# Check if Python is installed
if ! command -v python3 &> /dev/null; then
    echo "❌ Python 3 is not installed"
    exit 1
fi
echo "✓ Python 3 found: $(python3 --version)"

# Check if pip is installed
if ! command -v pip3 &> /dev/null; then
    echo "❌ pip3 is not installed"
    exit 1
fi
echo "✓ pip3 found: $(pip3 --version)"

# Check project structure
echo ""
echo "Checking project structure..."

if [ ! -f "setup.py" ] && [ ! -f "pyproject.toml" ]; then
    echo "⚠ Warning: No setup.py or pyproject.toml found"
fi

if [ -d "tests" ]; then
    echo "✓ tests directory found"
else
    echo "⚠ Warning: tests directory not found"
fi

if [ -f "README.md" ]; then
    echo "✓ README.md found"
else
    echo "⚠ Warning: README.md not found"
fi

# Check for requirements files
if [ -f "requirements.txt" ]; then
    echo "✓ requirements.txt found"
else
    echo "⚠ Warning: requirements.txt not found"
fi

echo ""
echo "================================"
echo "Validation Complete"
echo "================================"
