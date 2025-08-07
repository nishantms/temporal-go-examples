# Signals Example

Learn how to communicate with running workflows! 📡

## What You'll Learn

- What are signals and why use them
- How to send signals to workflows
- How to receive signals in workflows
- Building interactive workflows

## What This Example Does

This example simulates a food delivery order that can be updated while in progress:

1. **Start Order** - Customer places an order
2. **Send Signals** - Customer can update the order (add items, change address)
3. **Complete Order** - Order is delivered

## Key Concepts

### Signals
- **Purpose**: Send data to a running workflow
- **Async**: Don't block the sender
- **Persistent**: Stored in workflow history
- **Examples**: Update order, cancel request, user input

### Queries  
- **Purpose**: Get data from a running workflow
- **Sync**: Return data immediately
- **Read-only**: Don't change workflow state
- **Examples**: Get status, check progress

## Files Explained

- `workflow.go` - Workflow that receives signals
- `worker/main.go` - Starts the worker
- `client/main.go` - Starts workflow and sends signals

## How to Run

### Step 1: Start the Worker

```bash
cd examples/03-signals
go run worker/main.go
```

### Step 2: Execute the Workflow

```bash
# In another terminal
cd examples/03-signals
go run client/main.go
```

## Expected Output

```
[12:34:56] INFO: Starting DeliveryOrderWorkflow...
[12:34:56] INFO: Order received: Pizza
[12:34:56] INFO: Sending signal to add item: Coke
[12:34:57] INFO: Sending signal to update address: 456 Oak St
[12:34:58] INFO: Order status: In Progress (Pizza, Coke) -> 456 Oak St
[12:35:00] INFO: Order completed successfully!
```

## Next Steps

Move to [Example 04 - Error Handling](../04-error-handling/) to learn about robust error handling.
