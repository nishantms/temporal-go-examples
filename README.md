# Temporal Go Learning Examples

A comprehensive repository to learn Temporal.io with Go from zero to hero! 🚀

## What is Temporal?

Temporal is a durable execution system that allows you to write code as if failure doesn't exist. It:
- Automatically retries failed operations
- Handles timeouts and network issues
- Provides state management for long-running processes
- Enables complex workflow orchestration

Think of it as a way to write reliable, distributed applications without worrying about the complexity of failure handling.

## What is Go?

Go (Golang) is a programming language developed by Google. It's:
- Simple and easy to learn
- Fast and efficient
- Great for concurrent programming
- Perfect for building reliable services

## Prerequisites

### Installing Go (if not already installed)

1. **Ubuntu/Debian:**
   ```bash
   sudo apt update
   sudo apt install golang-go
   ```

2. **macOS:**
   ```bash
   brew install go
   ```

3. **Windows:**
   Download from [golang.org](https://golang.org/dl/)

4. **Verify installation:**
   ```bash
   go version
   ```

### Installing Temporal Server

You'll need a Temporal server running locally. We provide multiple Docker-based options:

1. **Install Docker** (if not installed):
   - Ubuntu: `sudo apt install docker.io docker-compose`
   - macOS: Download Docker Desktop
   - Windows: Download Docker Desktop

2. **Run Temporal Server (Choose one option):**

   **Option A: All-in-One Setup (Recommended):**
   ```bash
   # Use our automated setup script
   ./scripts/docker-setup.sh
   ```

   **Option B: Docker Compose:**
   ```bash
   # Build and start all services
   docker-compose up -d
   ```

3. **Access Temporal Web UI:**
   Open http://localhost:8080 in your browser

## Project Structure

```
temporal-go-examples/
├── README.md                 # This file  
├── GETTING_STARTED.md        # Quick start guide
├── go.mod                   # Go module definition
├── go.sum                   # Go module checksums
├── scripts/                 # Helper scripts
│   ├── docker-setup.sh     # Docker setup script  
│   ├── docker-stop.sh      # Docker stop script
│   ├── run-example.sh      # Run examples script
│   └── check-temporal.sh   # Check Temporal connectivity
├── shared/                  # Shared utilities
│   ├── temporal.go         # Common Temporal setup
│   └── utils.go            # Utility functions
├── examples/
│   ├── 01-hello-world/     # Basic workflow example
│   │   ├── README.md       # Example documentation
│   │   ├── workflow.go     # Workflow definition
│   │   ├── worker/main.go  # Worker implementation
│   │   └── client/main.go  # Client to start workflow
│   ├── 02-activities/      # Activities and workflows  
│   │   ├── README.md       # Example documentation
│   │   ├── workflow.go     # Workflow with activities
│   │   ├── activities.go   # Activity implementations
│   │   ├── worker/main.go  # Worker implementation
│   │   └── client/main.go  # Client to start workflow
│   ├── 03-signals/         # Signals and queries
│   │   ├── README.md       # Example documentation
│   │   ├── workflow.go     # Workflow with signals
│   │   ├── worker/main.go  # Worker implementation
│   │   └── client/main.go  # Client with signal sending
│   └── 04-error-handling/  # Error handling patterns
│       ├── README.md       # Example documentation
│       ├── workflow.go     # Workflow with error handling
│       ├── activities.go   # Activities that can fail
│       ├── worker/main.go  # Worker implementation
│       └── client/main.go  # Client testing error scenarios
```

## Getting Started

### Quick Start (3 minutes)

1. **Clone and setup:**
   ```bash
   git clone <your-repo>
   cd temporal-go-examples
   ```

2. **Start everything with Docker:**
   ```bash
   ./scripts/docker-setup.sh
   ```

3. **Run your first example:**
   ```bash
   
   # In container terminal 1 - start worker
   docker-compose exec temporal-go-examples ./scripts/run-example.sh 01-hello-world worker
   
   # In container terminal 2 - start client
   docker-compose exec temporal-go-examples ./scripts/run-example.sh 01-hello-world client
   ```

4. **See results:**
   - Check your terminal for output
   - Visit http://localhost:8080 for the Temporal Web UI

5. **Stop services when done:**
   ```bash
   ./scripts/docker-stop.sh              # Stop services
   ./scripts/docker-stop.sh --clean      # Stop and remove volumes
   ./scripts/docker-stop.sh --clean --image  # Stop, clean, and remove image
   ```

📖 **For detailed instructions, see [GETTING_STARTED.md](GETTING_STARTED.md)**
🐳 **For Docker setup details, see [DOCKER_SETUP.md](DOCKER_SETUP.md)**
⚡ **For quick Docker commands, see [DOCKER_QUICK_REFERENCE.md](DOCKER_QUICK_REFERENCE.md)**

### Alternative Setup Methods

#### Method 1: Docker Compose (Manual)
```bash
git clone <your-repo-url>
cd temporal-go-examples
docker-compose up -d
```

#### Method 2: Local Development
```bash
git clone <your-repo-url>
cd temporal-go-examples
go mod tidy

# Start Temporal server manually
docker-compose up
```

## Learning Path

### 🟢 Beginner Level

1. **[Hello World](examples/01-hello-world/)** - Your first Temporal workflow
2. **[Activities](examples/02-activities/)** - Breaking work into activities

### 🟡 Intermediate Level

3. **[Signals](examples/03-signals/)** - Communicating with running workflows
4. **[Error Handling](examples/04-error-handling/)** - Retry policies and error management

### 🔴 Advanced Level

Coming soon...

## Key Concepts Explained

### Workflow
A workflow is a sequence of steps that can be paused, resumed, and retried automatically. Think of it as a recipe that Temporal follows, even if some steps fail temporarily.

### Activity
An activity is a single unit of work within a workflow. Activities can fail and be retried independently. Examples: sending an email, calling an API, processing a file.

### Worker
A worker is a service that executes your workflows and activities. It polls Temporal for work to do.

### Task Queue
A named queue where workflows and activities are scheduled. Workers listen to specific task queues.

## Common Go Concepts You'll Need

### Structs
```go
type Person struct {
    Name string
    Age  int
}
```

### Interfaces
```go
type Worker interface {
    Start() error
    Stop()
}
```

### Goroutines (concurrent execution)
```go
go myFunction() // Runs in the background
```

### Channels (communication between goroutines)
```go
ch := make(chan string)
ch <- "message"    // Send
msg := <-ch        // Receive
```

## Useful Commands

### Docker Commands
```bash
# Start everything
./scripts/docker-setup.sh

# Enter development container
docker-compose exec temporal-go-examples bash

# Run examples with helper script
docker-compose exec temporal-go-examples ./run-example.sh 01-hello-world worker
docker-compose exec temporal-go-examples ./run-example.sh 01-hello-world client

# View logs
docker-compose logs temporal-go-examples
docker-compose logs temporal-sqlite

# Stop services
docker-compose down

# Stop and remove all data
docker-compose down -v
```

### Go Development Commands
```bash
# Build all examples
go build ./...

# Run tests
go test ./...

# Format code
go fmt ./...

# Get dependencies
go mod tidy

# Update dependencies
go get -u ./...

# Run specific example (inside container or with local Go)
cd examples/01-hello-world
go run worker/main.go    # Terminal 1
go run client/main.go    # Terminal 2
```

## Troubleshooting

### Common Issues

1. **"Temporal server not running"**
   - Make sure Docker is running
   - Start Temporal server: `./scripts/docker-setup.sh` or `docker-compose up`
   - Stop services: `./scripts/docker-stop.sh` or `docker-compose down`
   - Check if services are healthy: `docker-compose ps`

2. **"Module not found"**
   - Run `go mod tidy` to download dependencies
   - Or enter the container: `docker-compose exec temporal-go-examples bash`

3. **"Port already in use"**
   - Check if Temporal is already running: `docker ps`
   - Stop existing containers: `./scripts/docker-stop.sh` or `docker-compose down`

4. **"Cannot connect to Temporal server"**
   - Check if Temporal server is ready: `docker-compose logs temporal-sqlite`
   - Wait for health checks to pass: `docker-compose ps`
   - Verify connection: `docker-compose exec temporal-go-examples ./check-temporal.sh`

5. **Docker permission issues**
   - On Linux, add user to docker group: `sudo usermod -aG docker $USER`
   - Restart terminal or log out/in

### Getting Help

- [Temporal Documentation](https://docs.temporal.io/)
- [Go Documentation](https://golang.org/doc/)
- [Temporal Community Forum](https://community.temporal.io/)
- [Go Community](https://golang.org/community/)

## Next Steps

1. Start with the Hello World example
2. Understand the basic concepts
3. Try modifying the examples
4. Build your own workflow
5. Explore advanced features

Happy learning! 🎉
