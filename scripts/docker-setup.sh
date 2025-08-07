#!/bin/bash

echo "🐳 Setting up Temporal Go Examples with Docker..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
    print_error "Docker is not running. Please start Docker first."
    exit 1
fi

print_info "Docker is running!"

# Check if docker-compose is available
if ! command -v docker-compose &> /dev/null; then
    print_error "docker-compose is not installed. Please install it first."
    exit 1
fi

print_info "Building Docker images..."

# Build the application image
if docker build -t temporal-go-examples .; then
    print_status "Docker image built successfully!"
else
    print_error "Failed to build Docker image"
    exit 1
fi

echo ""
print_info "Starting services with Docker Compose..."

# Start all services
if docker-compose up -d; then
    print_status "Services started!"
else
    print_error "Failed to start services"
    exit 1
fi

echo ""
print_info "Waiting for services to be ready..."

# Wait for services to be healthy
max_attempts=30
attempt=0

while [ $attempt -lt $max_attempts ]; do
    if docker-compose ps | grep -q "healthy"; then
        break
    fi
    sleep 2
    attempt=$((attempt + 1))
    echo -n "."
done

echo ""

# Check if services are running
if docker-compose ps | grep -q "Up"; then
    print_status "Services are running!"
    echo ""
    echo -e "${BLUE}📊 Access points:${NC}"
    echo "   - Temporal Web UI: http://localhost:8080"
    echo "   - Temporal Server: localhost:7233"
    echo ""
    echo -e "${BLUE}🎯 To run examples:${NC}"
    echo "   # Method 1: Use the helper script"
    echo "   docker-compose exec temporal-go-examples ./run-example.sh 01-hello-world worker"
    echo ""
    echo "   # Method 2: Interactive development"
    echo "   docker-compose exec temporal-go-examples bash"
    echo "   cd examples/01-hello-world"
    echo "   go run worker/main.go"
    echo ""
    echo -e "${BLUE}📖 Available commands:${NC}"
    echo "   docker-compose exec temporal-go-examples bash              # Enter container"
    echo "   docker-compose logs temporal-go-examples                   # View app logs"
    echo "   docker-compose logs temporal-sqlite                        # View Temporal logs"
    echo "   docker-compose down                                        # Stop all services"
    echo "   docker-compose down -v                                     # Stop and remove volumes"
    echo ""
    echo -e "${GREEN}🎉 Setup complete! You can now run your Temporal workflows.${NC}"
else
    print_error "Some services failed to start. Check logs:"
    echo "   docker-compose logs"
    exit 1
fi
