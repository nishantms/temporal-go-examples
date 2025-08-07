#!/bin/bash

# Script to start Temporal server using Docker
echo "🚀 Starting Temporal server..."

# Check if Docker is running
if ! docker info >/dev/null 2>&1; then
    echo "❌ Docker is not running. Please start Docker first."
    exit 1
fi

# Create a temporary directory for Temporal
TEMP_DIR="/tmp/temporal-docker"

if [ ! -d "$TEMP_DIR" ]; then
    echo "📥 Downloading Temporal Docker Compose configuration..."
    git clone https://github.com/temporalio/docker-compose.git "$TEMP_DIR"
fi

cd "$TEMP_DIR"

echo "🐳 Starting Temporal services..."
echo "📍 This will start:"
echo "   - Temporal server (localhost:7233)"
echo "   - Temporal Web UI (localhost:8080)"
echo "   - PostgreSQL database"
echo "   - Elasticsearch (optional)"
echo ""

# Start the services
docker-compose up

echo ""
echo "🎉 Temporal server is running!"
echo "📱 Web UI: http://localhost:8080"
echo "🔌 gRPC: localhost:7233"
echo ""
echo "Press Ctrl+C to stop the server"
