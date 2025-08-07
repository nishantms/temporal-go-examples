# Error Handling Example

Learn how to build resilient workflows! 🛡️

## What You'll Learn

- Activity retry policies
- Different types of errors
- How to handle failures gracefully
- Compensation patterns

## What This Example Does

This example simulates a bank transfer that can fail at various points:

1. **Validate Accounts** - Check if accounts exist (can fail)
2. **Debit Source** - Withdraw money from source account (can fail)
3. **Credit Destination** - Add money to destination account (can fail)
4. **Compensation** - If any step fails, reverse previous steps

## Key Concepts

### Activity Retry Policies
- **InitialInterval**: Time before first retry
- **BackoffCoefficient**: How much to increase wait time each retry
- **MaximumInterval**: Maximum time between retries
- **MaximumAttempts**: Maximum number of retry attempts

### Error Types
- **ApplicationError**: Business logic errors (don't retry by default)
- **TimeoutError**: Activity took too long
- **CanceledError**: Workflow was canceled
- **TemporalError**: Infrastructure errors (retry automatically)

## Files Explained

- `workflow.go` - Transfer workflow with error handling
- `activities.go` - Activities that can fail and be retried
- `worker/main.go` - Worker setup
- `client/main.go` - Starts transfers (some will fail)

## How to Run

### Step 1: Start the Worker

```bash
cd examples/04-error-handling
go run worker/main.go
```

### Step 2: Execute the Workflow

```bash
# In another terminal
cd examples/04-error-handling
go run client/main.go
```

## Expected Output

```
[12:34:56] INFO: Starting MoneyTransferWorkflow...
[12:34:56] INFO: Validating accounts...
[12:34:57] INFO: Debiting $100.00 from account-123...
[12:34:57] ERROR: Credit failed, compensating...
[12:34:58] INFO: Compensation successful
[12:34:58] INFO: Transfer failed but system is consistent
```

## Next Steps

Congratulations! You've completed all the basic examples. Next, explore:
- Advanced Temporal features
- Building real applications
- Production deployment
