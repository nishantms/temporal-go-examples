# Activities Example

Learn how to break your workflow into activities! 🔧

## What You'll Learn

- What are activities and why use them
- How to define activities
- How to call activities from workflows
- Activity retry policies

## What This Example Does

This example simulates an order processing workflow:

1. **Validate Order** (activity) - Check if the order is valid
2. **Process Payment** (activity) - Process the payment
3. **Send Confirmation** (activity) - Send confirmation email
4. **Order Processing Workflow** - Orchestrates all the steps

## Key Differences from Hello World

- **Activities**: Can make network calls, access databases, etc.
- **Workflows**: Only orchestrate, can't do "side effects"
- **Retries**: Activities can be retried independently
- **Timeouts**: Activities have execution timeouts

## Files Explained

- `activities.go` - Defines all activity functions
- `workflow.go` - Defines the workflow that uses activities
- `worker/main.go` - Registers both workflows and activities
- `client/main.go` - Starts the workflow

## How to Run

### Option 1: Using Docker (Recommended)
```bash
# Make sure everything is set up
./scripts/docker-setup.sh

# In terminal 1 - start worker
docker-compose exec temporal-go-examples ./run-example.sh 02-activities worker

# In terminal 2 - execute workflow
docker-compose exec temporal-go-examples ./run-example.sh 02-activities client
```

### Option 2: Local Go Installation
```bash
# Make sure Temporal server is running first!

# In terminal 1
cd examples/02-activities
go run worker/main.go

# In terminal 2 (keep worker running)
cd examples/02-activities
go run client/main.go
```

## Expected Output

```
[12:34:56] INFO: Starting OrderProcessingWorkflow...
[12:34:56] INFO: Validating order...
[12:34:56] INFO: Processing payment for $99.99...
[12:34:56] INFO: Sending confirmation email to user@example.com...
[12:34:56] INFO: Order processing completed successfully!
```

## Try These Modifications

1. **Add a new activity**: Create a "UpdateInventory" activity
2. **Simulate failures**: Make an activity return an error sometimes
3. **Add timeouts**: Configure activity timeouts
4. **Add input validation**: Validate the order data

## Key Concepts

### Activity vs Workflow

| Workflow | Activity |
|----------|----------|
| Orchestration only | Can do "real work" |
| Must be deterministic | Can be non-deterministic |
| Can't make network calls | Can make network calls |
| Retried from beginning | Retried individually |

### Activity Definition
```go
func ValidateOrder(ctx context.Context, order Order) error {
    // Can access databases, make API calls, etc.
    // Will be retried if it fails
}
```

### Calling Activities from Workflows
```go
err := workflow.ExecuteActivity(ctx, ValidateOrder, order).Get(ctx, nil)
```

## Next Steps

Move to [Example 03 - Signals](../03-signals/) to learn about communicating with running workflows.
