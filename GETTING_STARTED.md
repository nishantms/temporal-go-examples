# Getting Started Guide

Welcome to Temporal Go Examples! 🎉

This guide will walk you through your first steps with Temporal and Go.

## Quick Start (5 minutes)

### 1. Prerequisites Check

Make sure you have Go installed:
```bash
go version
```

Make sure you have Docker installed:
```bash
docker --version
```

### 2. Setup the Project

```bash
# Run the setup script
./scripts/setup.sh
```

### 3. Start Temporal Server

```bash
# Option 1: Use our script (recommended)
./scripts/run-temporal.sh

# Option 2: Manual setup
git clone https://github.com/temporalio/docker-compose.git temporal-docker
cd temporal-docker
docker-compose up
```

This starts:
- Temporal server (localhost:7233)
- Temporal Web UI (localhost:8080) 
- PostgreSQL database

### 4. Run Your First Workflow

Open 2 terminals:

**Terminal 1 - Start the Worker:**
```bash
cd examples/01-hello-world
go run worker/main.go
```

**Terminal 2 - Execute the Workflow:**
```bash
cd examples/01-hello-world  
go run client/main.go
```

You should see:
```
[12:34:56] INFO: Starting GreetingWorkflow...
[12:34:56] INFO: Workflow result: Hello, Temporal World! Welcome to Temporal! 🎉
[12:34:56] INFO: Workflow completed successfully! 🎉
```

### 5. View in Temporal Web UI

1. Open http://localhost:8080 in your browser
2. You'll see your workflow execution
3. Click on it to see the details

## What Just Happened?

1. **Worker** - A process that can execute workflows and activities
2. **Workflow** - A function that defines your business logic
3. **Client** - Started the workflow and waited for the result
4. **Temporal Server** - Orchestrated everything and stored the state

## Next Steps

Now try the other examples in order:

1. **[Hello World](examples/01-hello-world/)** ✅ (You just did this!)
2. **[Activities](examples/02-activities/)** - Learn about breaking work into activities
3. **[Signals](examples/03-signals/)** - Learn about communicating with workflows
4. **[Error Handling](examples/04-error-handling/)** - Learn about retry policies and compensation

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

### "Worker not receiving tasks"
- Check that the task queue name matches between client and worker
- Ensure the worker is running and registered properly

### "Connection refused"
- Make sure Temporal server is running (`docker-compose up`)
- Check that localhost:7233 is accessible

### "Workflow not found"
- Ensure the workflow is registered with the worker
- Check that imports and package names are correct

### "Activity not found"
- Ensure activities are registered with the worker
- Check function names and signatures

## Need Help?

- Check the example READMEs in each folder
- Visit the Temporal Web UI at http://localhost:8080
- Read the [Temporal documentation](https://docs.temporal.io/)
- Ask questions in the [Temporal community](https://community.temporal.io/)

Happy coding! 🚀
