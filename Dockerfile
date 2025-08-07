# Multi-stage Dockerfile for Temporal Go Examples
FROM golang:1.21-alpine AS builder

# Install necessary packages
RUN apk add --no-cache \
    git \
    curl \
    bash \
    build-base

# Set working directory
WORKDIR /app

# Copy go modules first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build all examples
RUN go build -o bin/ ./...

# Production stage
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache \
    bash \
    curl \
    docker \
    docker-compose \
    ca-certificates

# Create app directory
WORKDIR /app

# Copy built binaries and source code
COPY --from=builder /app /app

# Create Temporal docker-compose setup
RUN mkdir -p /app/temporal-setup

# Create a docker-compose file for Temporal
RUN cat > /app/temporal-setup/docker-compose.yml << 'EOF'
version: '3.8'

services:
  temporal:
    image: temporalio/temporal-server:latest
    ports:
      - "7233:7233"
      - "8080:8080"
    environment:
      - TEMPORAL_CLI_SHOW_STACKS=1
    command: temporal-server start --env development --db sqlite --filename /tmp/temporal.db --ip 0.0.0.0
    healthcheck:
      test: ["CMD", "temporal", "workflow", "list", "--address", "temporal:7233"]
      interval: 30s
      timeout: 10s
      retries: 5
      start_period: 40s

  temporal-web:
    image: temporalio/temporal-web:latest
    ports:
      - "8088:8088"
    environment:
      - TEMPORAL_GRPC_ENDPOINT=temporal:7233
      - TEMPORAL_PERMIT_WRITE_API=true
    depends_on:
      temporal:
        condition: service_healthy
EOF

# Create startup script
RUN cat > /app/start-all.sh << 'EOF'
#!/bin/bash
set -e

echo "🚀 Starting Temporal Go Examples Environment..."

# Function to check if a port is available
wait_for_service() {
    local host=$1
    local port=$2
    local service=$3
    echo "⏳ Waiting for $service at $host:$port..."
    
    timeout=60
    while ! nc -z $host $port 2>/dev/null; do
        sleep 2
        timeout=$((timeout - 2))
        if [ $timeout -le 0 ]; then
            echo "❌ Timeout waiting for $service"
            exit 1
        fi
    done
    echo "✅ $service is ready!"
}

# Start Temporal with docker-compose
echo "🐳 Starting Temporal server..."
cd /app/temporal-setup
docker-compose up -d

# Wait for Temporal to be ready
wait_for_service localhost 7233 "Temporal server"

echo ""
echo "🎉 Environment is ready!"
echo ""
echo "📊 Services running:"
echo "   - Temporal Server: localhost:7233"
echo "   - Temporal Web UI: http://localhost:8080"
echo "   - Temporal Web (alternative): http://localhost:8088"
echo ""
echo "📚 Available examples:"
ls -1 /app/examples/ | grep -E '^[0-9]' | while read example; do
    echo "   - $example"
done
echo ""
echo "🎯 To run an example:"
echo "   1. Start worker: cd examples/01-hello-world && go run worker/main.go"
echo "   2. In another terminal: go run client/main.go"
echo ""
echo "🔧 Development commands:"
echo "   - ./run-example.sh <example-name> worker"
echo "   - ./run-example.sh <example-name> client"
echo ""

# Keep container running and show logs
echo "📋 Following Temporal logs (Ctrl+C to exit)..."
docker-compose logs -f
EOF

# Create helper script to run examples
RUN cat > /app/run-example.sh << 'EOF'
#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: ./run-example.sh <example-name> [worker|client]"
    echo ""
    echo "Available examples:"
    ls -1 examples/ | grep -E '^[0-9]' | sed 's/^/  - /'
    echo ""
    echo "Example usage:"
    echo "  ./run-example.sh 01-hello-world worker"
    echo "  ./run-example.sh 01-hello-world client"
    exit 1
fi

EXAMPLE=$1
TYPE=${2:-worker}

if [ ! -d "examples/$EXAMPLE" ]; then
    echo "❌ Example '$EXAMPLE' not found!"
    echo ""
    echo "Available examples:"
    ls -1 examples/ | grep -E '^[0-9]' | sed 's/^/  - /'
    exit 1
fi

if [[ "$TYPE" != "worker" && "$TYPE" != "client" ]]; then
    echo "❌ Type must be 'worker' or 'client'"
    exit 1
fi

echo "🚀 Running $EXAMPLE $TYPE..."
echo "📂 Working directory: examples/$EXAMPLE"
echo ""

cd examples/$EXAMPLE

if [ ! -f "$TYPE/main.go" ]; then
    echo "❌ File $TYPE/main.go not found in examples/$EXAMPLE/"
    exit 1
fi

echo "▶️  Executing: go run $TYPE/main.go"
echo "----------------------------------------"
go run $TYPE/main.go
EOF

# Make scripts executable
RUN chmod +x /app/start-all.sh /app/run-example.sh

# Install netcat for port checking
RUN apk add --no-cache netcat-openbsd

# Create a simple script to just run Go examples (without Docker setup)
RUN cat > /app/run-examples-only.sh << 'EOF'
#!/bin/bash
echo "🎯 Temporal Go Examples - Development Mode"
echo ""
echo "📋 Make sure Temporal server is running:"
echo "   - Temporal Server should be at localhost:7233"
echo "   - Start it with: ./start-all.sh (in another terminal)"
echo ""
echo "🚀 Available examples:"
ls -1 examples/ | grep -E '^[0-9]' | sed 's/^/  - /'
echo ""
echo "💡 Usage:"
echo "  ./run-example.sh <example-name> worker   # Start worker"
echo "  ./run-example.sh <example-name> client   # Start client"
echo ""
echo "📖 For detailed instructions, see README.md"

# Keep container running for interactive use
exec bash
EOF

RUN chmod +x /app/run-examples-only.sh

# Expose Temporal ports
EXPOSE 7233 8080 8088

# Set the default command
CMD ["/app/start-all.sh"]
