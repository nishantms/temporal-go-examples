# Docker Quick Reference

## 🚀 Quick Start
```bash
./scripts/docker-setup.sh
```

## 📋 Essential Commands

### Setup & Start
```bash
# Initial setup (run once)
./scripts/docker-setup.sh

# Start services (if stopped)
docker-compose up -d

# Enter development container
docker-compose exec temporal-go-examples bash
```

### Running Examples
```bash
# Inside container - general format
./run-example.sh <example-name> <worker|client>

# Specific examples
./run-example.sh 01-hello-world worker     # Terminal 1
./run-example.sh 01-hello-world client     # Terminal 2
./run-example.sh 02-activities worker      # Terminal 1
./run-example.sh 02-activities client      # Terminal 2
```

### Monitoring & Debugging
```bash
# Check service status
docker-compose ps

# View logs
docker-compose logs temporal-go-examples
docker-compose logs temporal-sqlite
docker-compose logs temporal-web

# Test Temporal connection
docker-compose exec temporal-go-examples ./check-temporal.sh

# Follow logs in real-time
docker-compose logs -f temporal-sqlite
```

### Cleanup
```bash
# Stop services (preserves data)
docker-compose down

# Stop and remove data
docker-compose down -v

# Clean up Docker resources
docker system prune -f
```

## 🌐 Access Points
- **Temporal Web UI**: http://localhost:8080
- **Temporal Server**: localhost:7233

## 🔧 Development Workflow

### Daily Development
```bash
# 1. Start services (if not running)
docker-compose up -d

# 2. Enter container for development
docker-compose exec temporal-go-examples bash

# 3. Run examples
./run-example.sh 01-hello-world worker    # Terminal 1
./run-example.sh 01-hello-world client    # Terminal 2

# 4. View results in Web UI
# Open http://localhost:8080
```

### Multiple Terminals
```bash
# Terminal 1: Worker
docker-compose exec temporal-go-examples bash
./run-example.sh 01-hello-world worker

# Terminal 2: Client  
docker-compose exec temporal-go-examples bash
./run-example.sh 01-hello-world client

# Terminal 3: Another worker
docker-compose exec temporal-go-examples bash
./run-example.sh 02-activities worker
```

## 🐛 Troubleshooting

### Services Not Starting
```bash
# Check Docker
docker info

# Restart everything
docker-compose down
docker-compose up -d

# Check logs for errors
docker-compose logs
```

### Can't Connect to Temporal
```bash
# Test from inside container
docker-compose exec temporal-go-examples ./check-temporal.sh

# Check if Temporal is healthy
docker-compose ps | grep temporal

# Restart Temporal service
docker-compose restart temporal-sqlite
```

### Port Conflicts
```bash
# Check what's using ports
sudo netstat -tulpn | grep :7233
sudo netstat -tulpn | grep :8080

# Stop conflicting services
docker-compose down
```

### Starting Fresh
```bash
# Nuclear option - clean everything
docker-compose down -v
docker system prune -f
./scripts/docker-setup.sh
```

## 📁 File Structure
```
temporal-go-examples/
├── Dockerfile              # Main development container
├── Dockerfile.simple       # Lightweight container
├── docker-compose.yml      # Multi-service setup
├── scripts/
│   └── docker-setup.sh     # Automated setup script
└── examples/
    ├── 01-hello-world/
    ├── 02-activities/
    ├── 03-signals/
    └── 04-error-handling/
```

## 💡 Tips

### Use Helper Scripts
- Always use `./run-example.sh` instead of direct `go run`
- It provides better error messages and connectivity checks

### Keep Containers Running
- Leave services running during development
- Use `docker-compose exec` for multiple terminals
- Only stop when you're done for the day

### Check Health Status
- Always wait for health checks: `docker-compose ps`
- Green "healthy" status means ready to use

### Monitor Resources
```bash
# Check resource usage
docker stats

# Check disk usage
docker system df
```

### Development Tips
```bash
# Inside container - useful commands
go mod tidy                    # Update dependencies
go build ./...                 # Build all examples
go test ./...                  # Run tests
go fmt ./...                   # Format code

# Check available examples
ls -la examples/
```

## 🔄 Common Patterns

### New Example Development
```bash
# 1. Enter container
docker-compose exec temporal-go-examples bash

# 2. Create new example directory
mkdir examples/05-my-example
cd examples/05-my-example

# 3. Develop (worker.go, client.go, etc.)

# 4. Test
cd /app
./run-example.sh 05-my-example worker
./run-example.sh 05-my-example client
```

### Testing Different Examples
```bash
# Quick testing of all examples
for example in 01-hello-world 02-activities 03-signals 04-error-handling; do
    echo "Testing $example..."
    ./run-example.sh $example worker &
    sleep 2
    ./run-example.sh $example client
    pkill -f "worker/main.go"
done
```

### Viewing Workflow History
1. Run any example
2. Open http://localhost:8080
3. Click on workflow execution
4. Explore timeline, events, and details

This quick reference should cover 90% of your daily Docker usage needs!
