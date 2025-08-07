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

# Check if Temporal server is reachable
TEMPORAL_HOST=${TEMPORAL_HOSTPORT:-localhost:7233}
if ! nc -z ${TEMPORAL_HOST/:/ } 2>/dev/null; then
    echo "⚠️  Warning: Cannot reach Temporal server at $TEMPORAL_HOST"
    echo "   Make sure Temporal server is running first!"
    echo ""
fi

echo "🚀 Running $EXAMPLE $TYPE..."
echo "📂 Working directory: examples/$EXAMPLE"
echo "🔗 Temporal server: $TEMPORAL_HOST"
echo ""

cd examples/$EXAMPLE

if [ ! -f "$TYPE/main.go" ]; then
    echo "❌ File $TYPE/main.go not found in examples/$EXAMPLE/"
    exit 1
fi

echo "▶️  Executing: go run $TYPE/main.go"
echo "----------------------------------------"
go run $TYPE/main.go
