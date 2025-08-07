# Hello World Example

This is your first Temporal workflow! 🎉

## What You'll Learn

- What is a workflow
- How to define a simple workflow
- How to start a worker
- How to execute a workflow

## What This Example Does

1. **Workflow**: Takes a name and returns a greeting
2. **Worker**: Runs the workflow when requested
3. **Client**: Starts the workflow and waits for the result

## Files Explained

- `workflow.go` - Defines the workflow function
- `worker/main.go` - Starts a worker that can execute the workflow
- `client/main.go` - Starts the workflow and prints the result

## How to Run

### Step 1: Start the Worker
```bash
# In terminal 1
cd examples/01-hello-world
go run worker/main.go
```

You should see:
```
Starting worker on task queue: temporal-learning-queue
Worker started successfully
```

### Step 2: Execute the Workflow
```bash
# In terminal 2 (keep worker running)
cd examples/01-hello-world
go run client/main.go
```

You should see:
```
Starting workflow...
Workflow result: Hello, Temporal World!
Workflow completed successfully
```

## Try These Modifications

1. **Change the greeting**: Edit `workflow.go` and modify the return message
2. **Add a parameter**: Modify the workflow to take a custom name
3. **Add logging**: Use `workflow.GetLogger()` to add logging inside the workflow

## Key Concepts

### Workflow Function
```go
func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
    // This function defines what happens in the workflow
    // It must be deterministic (same input = same output)
    return fmt.Sprintf("Hello, %s!", name), nil
}
```

### Worker Registration
```go
w.RegisterWorkflow(GreetingWorkflow)  // Tell the worker about our workflow
```

### Workflow Execution
```go
we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, GreetingWorkflow, "Temporal World")
```

## Next Steps

Once this works, move to [Example 02 - Activities](../02-activities/) to learn about breaking work into activities.
