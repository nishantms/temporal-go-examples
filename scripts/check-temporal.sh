#!/bin/bash

TEMPORAL_HOST=${TEMPORAL_HOSTPORT:-localhost:7233}

echo "🔍 Checking Temporal server status..."
echo "📡 Temporal endpoint: $TEMPORAL_HOST"

if nc -z ${TEMPORAL_HOST/:/ } 2>/dev/null; then
    echo "✅ Temporal server is reachable!"
else
    echo "❌ Cannot reach Temporal server at $TEMPORAL_HOST"
    echo ""
    echo "💡 Make sure Temporal server is running:"
    echo "   - Using Docker Compose: docker-compose up -d"
    echo "   - Using setup script: ./scripts/docker-setup.sh"
    echo "   - Using Temporal CLI: temporal server start-dev"
    exit 1
fi
