#!/bin/bash

echo "🛑 Stopping Temporal Go Examples Docker services..."

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

# Check if docker-compose is available
if ! command -v docker-compose &> /dev/null; then
    print_error "docker-compose is not installed."
    exit 1
fi

# Stop all services
print_info "Stopping Docker Compose services..."
if docker-compose down; then
    print_status "Services stopped successfully!"
else
    print_warning "Some services may not have stopped cleanly"
fi

# Optional: Remove containers and volumes
if [ "$1" = "--clean" ] || [ "$1" = "-c" ]; then
    print_info "Cleaning up containers, networks, and volumes..."
    if docker-compose down -v --remove-orphans; then
        print_status "Cleanup completed!"
    else
        print_warning "Cleanup may not have completed fully"
    fi
    
    # Remove the built image if requested
    if [ "$2" = "--image" ] || [ "$1" = "--image" ]; then
        print_info "Removing Docker image..."
        if docker rmi temporal-go-examples 2>/dev/null; then
            print_status "Docker image removed!"
        else
            print_warning "Docker image may not exist or couldn't be removed"
        fi
    fi
fi

echo ""
print_info "Usage options:"
echo "   ./scripts/docker-stop.sh              # Stop services only"
echo "   ./scripts/docker-stop.sh --clean      # Stop and remove volumes"
echo "   ./scripts/docker-stop.sh --clean --image  # Stop, clean, and remove image"
echo ""
print_status "Docker services stopped!"
