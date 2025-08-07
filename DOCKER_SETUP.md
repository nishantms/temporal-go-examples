# Docker Setup Guide

This document explains the Docker-based setup for Temporal Go Examples.

## Overview

Our Docker setup provides:
- 🐳 **Temporal Server** with SQLite backend (simple, no external dependencies)
- 🌐 **Temporal Web UI** for monitoring workflows
- 🐹 **Go Development Environment** with all examples ready to run
- 🔧 **Helper Scripts** for easy example execution
- 🔗 **Proper Networking** between all services
- 📊 **Health Checks** to ensure everything is ready

## Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                     Docker Network                         │
│  ┌─────────────────┐  ┌─────────────────┐  ┌──────────────┐│
│  │ temporal-sqlite │  │ temporal-web    │  │ go-examples  ││
│  │                 │  │                 │  │              ││
│  │ Port: 7233      │  │ Port: 8080      │  │ Development  ││
│  │ SQLite DB       │  │ Web UI          │  │ Container    ││
│  └─────────────────┘  └─────────────────┘  └──────────────┘│
└─────────────────────────────────────────────────────────────┘
              │                    │                    │
         ┌────▼────┐          ┌────▼────┐          ┌────▼────┐
         │localhost│          │localhost│          │ Your    │
         │:7233    │          │:8080    │          │ Code    │
         └─────────┘          └─────────┘          └─────────┘
```

## Files Explanation

### 1. `Dockerfile`
**Purpose**: Main development container with Temporal server integration

**What it does**:
- Builds Go application
- Includes helper scripts for running examples
- Can start Temporal server internally (all-in-one mode)
- Provides development environment

**Use case**: Single container with everything included

### 2. `Dockerfile.simple`
**Purpose**: Lightweight Go development container

**What it does**:
- Only builds Go application
- Expects external Temporal server
- Minimal dependencies
- Fast build times

**Use case**: When you have Temporal running elsewhere

### 3. `docker-compose.yml`
**Purpose**: Multi-service orchestration

**What it does**:
- Starts Temporal server (SQLite backend)
- Starts Temporal Web UI
- Starts Go development container
- Sets up networking between services
- Includes health checks
- Provides volume mounting for development

**Use case**: Complete development environment (recommended)

### 4. `scripts/docker-setup.sh`
**Purpose**: Automated setup script

**What it does**:
- Checks Docker availability
- Builds images
- Starts services with docker-compose
- Waits for health checks
- Provides usage instructions

**Use case**: One-command setup

## Setup Options

### Option 1: Automated Setup (Recommended)

```bash
# One command to rule them all
./scripts/docker-setup.sh
```

**Pros**:
- ✅ Fully automated
- ✅ Includes health checks
- ✅ Clear status messages
- ✅ Usage instructions

**Cons**:
- None really!

### Option 2: Manual Docker Compose

```bash
# Build and start
docker-compose up -d

# Check status
docker-compose ps

# Enter development container
docker-compose exec temporal-go-examples bash
```

**Pros**:
- ✅ Full control over the process
- ✅ Can see all logs
- ✅ Easy to customize

**Cons**:
- Manual steps required
- Need to wait for services manually

### Option 3: All-in-One Container

```bash
# Build and run single container
docker build -t temporal-go-all .
docker run -p 7233:7233 -p 8080:8080 temporal-go-all
```

**Pros**:
- ✅ Single container
- ✅ No external dependencies
- ✅ Portable

**Cons**:
- Heavier container
- Less flexibility
- Harder to debug individual services

### Option 4: Lightweight Container + External Temporal

```bash
# Build lightweight container
docker build -f Dockerfile.simple -t temporal-go-simple .

# Run with external Temporal
docker run -it --network host temporal-go-simple
```

**Pros**:
- ✅ Fast build
- ✅ Minimal overhead
- ✅ Can connect to any Temporal server

**Cons**:
- Requires external Temporal server
- Manual networking setup

## Services Details

### Temporal Server (temporal-sqlite)
- **Image**: `temporalio/temporal-server:latest`
- **Port**: 7233
- **Database**: SQLite (file-based, no external DB needed)
- **Health Check**: Workflow list command
- **Data**: Stored in Docker volume

### Temporal Web UI (temporal-web)
- **Image**: `temporalio/temporal-web:latest`
- **Port**: 8080
- **Purpose**: Visual interface for monitoring workflows
- **Dependencies**: Connects to temporal-sqlite service

### Go Development Container (temporal-go-examples)
- **Built from**: Our custom Dockerfile
- **Purpose**: Development environment with Go and examples
- **Features**:
  - All Go examples built and ready
  - Helper scripts installed
  - Volume mounting for live development
  - Network access to Temporal services

## Helper Scripts

### Inside Container: `run-example.sh`
```bash
# Usage
./run-example.sh <example-name> <worker|client>

# Examples
./run-example.sh 01-hello-world worker
./run-example.sh 02-activities client
```

**Features**:
- ✅ Validates example exists
- ✅ Checks Temporal connectivity
- ✅ Provides clear error messages
- ✅ Shows execution details

### Inside Container: `check-temporal.sh`
```bash
# Check if Temporal server is reachable
./check-temporal.sh
```

**Features**:
- ✅ Tests network connectivity
- ✅ Shows connection details
- ✅ Provides troubleshooting tips

## Development Workflow

### 1. Initial Setup
```bash
# Clone and setup (once)
git clone <repo>
cd temporal-go-examples
./scripts/docker-setup.sh
```

### 2. Daily Development
```bash
# Start services (if not running)
docker-compose up -d

# Enter development container
docker-compose exec temporal-go-examples bash

# Work on examples
./run-example.sh 01-hello-world worker
```

### 3. Cleanup
```bash
# Stop services (preserves data)
docker-compose down

# Stop and remove data
docker-compose down -v
```

## Environment Variables

### Container Environment
- `TEMPORAL_HOSTPORT`: Address of Temporal server (default: localhost:7233)

### Docker Compose Environment
- Automatic service discovery via container names
- No manual configuration needed

## Volumes and Data

### Development Volume
```yaml
volumes:
  - .:/app  # Live code mounting
```
- Changes in your local files appear immediately in container
- No need to rebuild for code changes

### Go Modules Volume
```yaml
volumes:
  - go-modules:/go/pkg/mod  # Go module cache
```
- Speeds up builds by caching downloaded modules
- Persistent across container restarts

### Temporal Data Volume
```yaml
volumes:
  - temporal-data:/tmp  # SQLite database
```
- Persists workflow data across restarts
- Remove with `docker-compose down -v` to start fresh

## Networking

### Internal Network
- All services communicate via internal Docker network
- Service names resolve to container IPs
- No external network access required for inter-service communication

### External Access
- Temporal Server: `localhost:7233`
- Temporal Web UI: `localhost:8080`
- Your applications can connect from host system

## Health Checks

### Temporal Server Health Check
```yaml
healthcheck:
  test: ["CMD", "temporal", "workflow", "list", "--address", "temporal-sqlite:7233"]
  interval: 30s
  timeout: 10s
  retries: 5
  start_period: 40s
```

### Benefits
- ✅ Ensures services are ready before proceeding
- ✅ Prevents "connection refused" errors
- ✅ Automatic retries and recovery

## Troubleshooting

### Check Service Status
```bash
# View all services
docker-compose ps

# Check logs
docker-compose logs temporal-sqlite
docker-compose logs temporal-web
docker-compose logs temporal-go-examples
```

### Common Issues

#### Services Not Starting
```bash
# Check Docker
docker info

# Rebuild and restart
docker-compose down
docker-compose up -d --build
```

#### Port Conflicts
```bash
# Check what's using ports
netstat -an | grep :7233
netstat -an | grep :8080

# Stop conflicting services
docker-compose down
```

#### Network Issues
```bash
# Test connectivity from container
docker-compose exec temporal-go-examples ./check-temporal.sh

# Restart networking
docker-compose down
docker-compose up -d
```

## Customization

### Use PostgreSQL Instead of SQLite
1. Uncomment PostgreSQL sections in `docker-compose.yml`
2. Change `temporal-sqlite` references to `temporal-postgres`
3. Restart: `docker-compose down && docker-compose up -d`

### Add Custom Environment Variables
```yaml
services:
  temporal-go-examples:
    environment:
      - TEMPORAL_HOSTPORT=temporal-sqlite:7233
      - YOUR_CUSTOM_VAR=value
```

### Mount Additional Volumes
```yaml
services:
  temporal-go-examples:
    volumes:
      - .:/app
      - ./custom-config:/etc/config
```

## Best Practices

### 1. Use the Setup Script
- Always start with `./scripts/docker-setup.sh`
- It handles all the complexity for you

### 2. Check Health Status
- Wait for health checks before running examples
- Use `docker-compose ps` to verify

### 3. Use Helper Scripts
- Prefer `./run-example.sh` over direct `go run`
- It provides better error messages and connectivity checks

### 4. Clean Up Regularly
```bash
# Clean up unused resources
docker system prune -f

# Remove all data and start fresh
docker-compose down -v
./scripts/docker-setup.sh
```

### 5. Development Tips
- Keep containers running during development
- Use `docker-compose exec` for multiple terminals
- Check logs when things go wrong

## Advanced Usage

### Running Multiple Workers
```bash
# Terminal 1
docker-compose exec temporal-go-examples ./run-example.sh 01-hello-world worker

# Terminal 2 (different task queue)
docker-compose exec temporal-go-examples ./run-example.sh 02-activities worker

# Terminal 3 (client)
docker-compose exec temporal-go-examples ./run-example.sh 01-hello-world client
```

### Custom Temporal Configuration
1. Create custom config files in `./config/`
2. Mount them in docker-compose:
   ```yaml
   volumes:
     - ./config:/etc/temporal/config
   ```
3. Update Temporal server command to use custom config

### Scaling Services
```bash
# Run multiple instances of a service
docker-compose up -d --scale temporal-go-examples=3
```

## Summary

The Docker setup provides a complete, production-like Temporal development environment with minimal setup. Use `./scripts/docker-setup.sh` to get started quickly, then use the helper scripts inside the container for day-to-day development.

For questions or issues, check the troubleshooting section or refer to the main README.md.
