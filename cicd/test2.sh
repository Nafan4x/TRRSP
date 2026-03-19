#!/bin/bash

echo "Running test2.sh - checking dependencies"

# Проверяем наличие Go
if command -v go &> /dev/null; then
    echo "✅ Go is installed: $(go version)"
else
    echo "❌ Go is not installed"
    exit 1
fi

# Проверяем наличие make
if command -v make &> /dev/null; then
    echo "✅ Make is installed"
else
    echo "❌ Make is not installed"
    exit 1
fi

echo "All dependencies satisfied!"
exit 0

