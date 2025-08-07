#!/bin/bash

# Setup script for Temporal Go Examples
echo "🚀 Setting up Temporal Go Examples..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go first:"
    echo "   Ubuntu: sudo apt install golang-go"
    echo "   macOS: brew install go"
    echo "   Windows: Download from https://golang.org/dl/"
    exit 1
fi

echo "✅ Go is installed: $(go version)"

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo "❌ Docker is not installed. Please install Docker first:"
    echo "   Ubuntu: sudo apt install docker.io"
    echo "   macOS/Windows: Download Docker Desktop"
    exit 1
fi

echo "✅ Docker is installed: $(docker --version)"

# Install dependencies
echo "📦 Installing Go dependencies..."
go mod tidy

# Check if Temporal server is running
echo "🔍 Checking if Temporal server is running..."
if curl -s http://localhost:8080 > /dev/null; then
    echo "✅ Temporal server is already running"
else
    echo "⚠️  Temporal server is not running"
    echo "📋 To start Temporal server, run:"
    echo "   ./scripts/run-temporal.sh"
    echo "   OR"
    echo "   git clone https://github.com/temporalio/docker-compose.git temporal-docker"
    echo "   cd temporal-docker && docker-compose up"
fi

echo ""
echo "🎉 Setup complete! Next steps:"
echo ""
echo "1. Start Temporal server (if not running):"
echo "   ./scripts/run-temporal.sh"
echo ""
echo "2. Try the Hello World example:"
echo "   cd examples/01-hello-world"
echo "   go run worker/main.go    # In terminal 1"
echo "   go run client/main.go    # In terminal 2"
echo ""
echo "3. Visit http://localhost:8080 to see the Temporal Web UI"
echo ""
echo "Happy learning! 🚀"
