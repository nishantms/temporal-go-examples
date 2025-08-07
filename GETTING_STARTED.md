# Getting Started Guide

Welcome to Temporal Go Examples! 🎉

This guide will walk you through your first steps with Temporal and Go using our Docker-based setup.

## Quick Start (3 minutes)

### 1. Prerequisites Check

Make sure you have Docker installed:
```bash
docker --version
docker-compose --version
```

Optionally, check if Go is installed for local development:
```bash
go version
```

### 2. Setup the Project

```bash
# Clone the repository
git clone <your-repo-url>
cd temporal-go-examples

# Run the automated setup script
./scripts/docker-setup.sh
```

This will:
- Build your Go application Docker image
- Start Temporal server with SQLite backend
- Start Temporal Web UI
- Set up all networking and health checks

### 3. Verify Everything is Running

Check that all services are healthy:
```bash
docker-compose ps
```

You should see services with "Up" status and health checks passing.

### 4. Run Your First Workflow

**Terminal 1 - Start the Worker:**
```bash
# Enter the development container
docker-compose exec temporal-go-examples bash

# Run the worker using our helper script
./run-example.sh 01-hello-world worker
```

**Terminal 2 - Execute the Workflow:**
```bash
# In a new terminal, enter the container again
docker-compose exec temporal-go-examples bash

# Run the client
./run-example.sh 01-hello-world client
```

You should see:
```
🚀 Running 01-hello-world worker...
📂 Working directory: examples/01-hello-world
🔗 Temporal server: localhost:7233

▶️  Executing: go run worker/main.go
----------------------------------------
[12:34:56] INFO: Starting GreetingWorkflow...
[12:34:56] INFO: Workflow result: Hello, Temporal World! Welcome to Temporal! 🎉
[12:34:56] INFO: Workflow completed successfully! 🎉
```

### 5. View in Temporal Web UI

1. Open http://localhost:8080 in your browser
2. You'll see your workflow execution
3. Click on it to see the details, timeline, and logs

## What Just Happened?

1. **Docker Environment** - Everything runs in containers with proper networking
2. **Worker** - A process that can execute workflows and activities
3. **Workflow** - A function that defines your business logic
4. **Client** - Started the workflow and waited for the result
5. **Temporal Server** - Orchestrated everything and stored the state

## Next Steps

Now try the other examples in order:

1. **[Hello World](examples/01-hello-world/)** ✅ (You just did this!)
2. **[Activities](examples/02-activities/)** - Learn about breaking work into activities
3. **[Signals](examples/03-signals/)** - Learn about communicating with workflows
4. **[Error Handling](examples/04-error-handling/)** - Learn about retry policies and compensation

## Docker Commands Reference

### Essential Commands
```bash
# Start everything
./scripts/docker-setup.sh

# Enter development container
docker-compose exec temporal-go-examples bash

# Run examples with helper script
./run-example.sh <example-name> <worker|client>

# Check Temporal connection
./check-temporal.sh

# View logs
docker-compose logs temporal-go-examples
docker-compose logs temporal-sqlite

# Stop services
docker-compose down

# Stop and remove all data
docker-compose down -v
```

### Development Workflow
```bash
# 1. Start services (once)
./scripts/docker-setup.sh

# 2. Development loop
docker-compose exec temporal-go-examples bash
./run-example.sh 01-hello-world worker    # Terminal 1
./run-example.sh 01-hello-world client    # Terminal 2

# 3. Check results
# - View terminal output
# - Visit http://localhost:8080 for Web UI

# 4. Clean up (when done)
docker-compose down
```

## Key Concepts to Understand

### Workflows
- Define the steps of your business process
- Are durable (survive crashes and restarts)
- Must be deterministic
- Can wait for signals, timers, and activity results

### Activities  
- Do the actual work (database calls, API calls, etc.)
- Can fail and be retried automatically
- Are not deterministic (can have side effects)
- Have timeouts and retry policies

### Workers
- Execute workflows and activities
- Run your code
- Poll Temporal for work to do

### Task Queues
- Named queues where work is scheduled
- Workers listen to specific task queues
- Provide routing and load balancing

## Common Patterns

### Sequential Processing
```go
// Execute activities one after another
err := workflow.ExecuteActivity(ctx, Activity1, input1).Get(ctx, &result1)
err = workflow.ExecuteActivity(ctx, Activity2, result1).Get(ctx, &result2)
```

### Parallel Processing
```go
// Execute activities in parallel
future1 := workflow.ExecuteActivity(ctx, Activity1, input1)
future2 := workflow.ExecuteActivity(ctx, Activity2, input2)

var result1, result2 string
future1.Get(ctx, &result1)
future2.Get(ctx, &result2)
```

### Waiting for Signals
```go
signalChan := workflow.GetSignalChannel(ctx, "my-signal")
var signalData string
signalChan.Receive(ctx, &signalData)
```

### Handling Errors
```go
err := workflow.ExecuteActivity(ctx, MyActivity, input).Get(ctx, &result)
if err != nil {
    // Handle or compensate for the error
    return err
}
```

## Troubleshooting

### "Cannot connect to Temporal server"
```bash
# Check if services are running
docker-compose ps

# Check Temporal server logs
docker-compose logs temporal-sqlite

# Test connection from container
docker-compose exec temporal-go-examples ./check-temporal.sh
```

### "Worker not receiving tasks"
- Check that the task queue name matches between client and worker
- Ensure the worker is running and registered properly
- Verify network connectivity between containers

### "Docker permission issues"
- On Linux, add user to docker group: `sudo usermod -aG docker $USER`
- Restart terminal or log out/in

### "Port already in use"
```bash
# Check what's using the ports
docker ps

# Stop existing containers
docker-compose down

# Force cleanup if needed
docker-compose down -v
```

### "Module not found" or "Build errors"
```bash
# Enter container and rebuild
docker-compose exec temporal-go-examples bash
go mod tidy
go build ./...
```

### Starting Fresh
```bash
# Complete cleanup and restart
docker-compose down -v
docker system prune -f
./scripts/docker-setup.sh
```

## Need Help?

- Check the example READMEs in each folder
- Visit the Temporal Web UI at http://localhost:8080
- Read the [Temporal documentation](https://docs.temporal.io/)
- Ask questions in the [Temporal community](https://community.temporal.io/)

Happy coding! 🚀
