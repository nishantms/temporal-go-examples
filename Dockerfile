# Dockerfile for Temporal Go Examples Development
FROM golang:1.24-alpine

# Install necessary packages
RUN apk add --no-cache \
    git \
    bash \
    curl \
    netcat-openbsd \
    ca-certificates

# Set working directory
WORKDIR /app

# Copy go modules first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build all examples
RUN go build -o bin/ ./...

# Copy helper scripts
COPY scripts/run-example.sh /app/run-example.sh
COPY scripts/check-temporal.sh /app/check-temporal.sh

# Make scripts executable
RUN chmod +x /app/run-example.sh /app/check-temporal.sh

# Default environment variables
ENV TEMPORAL_HOSTPORT=localhost:7233

# Expose any ports your application might use
EXPOSE 8081

# Default command shows usage
CMD ["bash", "-c", "echo '🎯 Temporal Go Examples Development Container'; echo ''; echo 'Available commands:'; echo '  ./run-example.sh <example> <worker|client>'; echo '  ./check-temporal.sh'; echo '  bash  # Interactive shell'; echo ''; echo 'Make sure Temporal server is running at localhost:7233'; echo 'Use docker-compose for complete setup.'; echo ''; exec bash"]
