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

# Wait for services to be ready
max_attempts=60
attempt=0

echo ""
print_info "Waiting for services to start (this may take 1-2 minutes)..."

while [ $attempt -lt $max_attempts ]; do
    if docker-compose ps | grep -q "Up"; then
        # Check if Temporal is responding
        if docker-compose exec -T temporal curl -f http://localhost:7233/ >/dev/null 2>&1; then
            break
        fi
    fi
    sleep 5
    attempt=$((attempt + 5))
    echo -n "."
done

echo ""

# Check if services are running
if docker-compose ps | grep -q "Up"; then
    print_status "Services are starting up!"
    echo ""
    echo -e "${BLUE}📊 Access points:${NC}"
    echo "   - Temporal Web UI: http://localhost:8080 (may take 1-2 minutes to be ready)"
    echo "   - Temporal Server: localhost:7233"
    echo ""
    echo -e "${BLUE}🎯 To run examples:${NC}"
    echo "   # Wait for Temporal to be fully ready, then:"
    echo "   docker-compose exec temporal-go-examples bash"
    echo "   ./check-temporal.sh  # Check if Temporal is ready"
    echo "   ./run-example.sh 01-hello-world worker"
    echo ""
    echo -e "${BLUE}📖 Monitor startup:${NC}"
    echo "   docker-compose logs -f temporal     # Watch Temporal startup"
    echo "   docker-compose logs postgres        # Check database logs"
    echo "   docker-compose ps                   # Check container status"
    echo ""
    echo -e "${GREEN}🎉 Setup initiated! Temporal may take 1-2 minutes to be fully ready.${NC}"
else
    print_error "Some services failed to start. Check logs:"
    echo "   docker-compose logs"
    exit 1
fi
