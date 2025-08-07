package main

import (
	"context"
	"log"
	"time"

	errors "temporal-go-examples/examples/04-error-handling"
	"temporal-go-examples/shared"
)

func main() {
	// Create Temporal client
	c, err := shared.CreateTemporalClient()
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// Test different scenarios
	testScenarios := []struct {
		name    string
		request errors.TransferRequest
	}{
		{
			name: "Normal Transfer",
			request: errors.TransferRequest{
				FromAccount: "account-123",
				ToAccount:   "account-456",
				Amount:      100.50,
				Reference:   "Payment for services",
			},
		},
		{
			name: "Invalid Account (will fail immediately)",
			request: errors.TransferRequest{
				FromAccount: "invalid-account",
				ToAccount:   "account-456",
				Amount:      50.00,
				Reference:   "Test invalid account",
			},
		},
		{
			name: "Insufficient Funds (will fail)",
			request: errors.TransferRequest{
				FromAccount: "broke-account",
				ToAccount:   "account-456",
				Amount:      1000.00,
				Reference:   "Test insufficient funds",
			},
		},
	}

	// Run the MoneyTransferWorkflow tests
	shared.LogInfo("=== Testing MoneyTransferWorkflow ===")
	for _, scenario := range testScenarios {
		shared.LogInfo("🧪 Testing scenario: %s", scenario.name)

		workflowRun, err := shared.ExecuteWorkflow(c, errors.MoneyTransferWorkflow, scenario.request)
		if err != nil {
			shared.LogError("Failed to start workflow: %v", err)
			continue
		}

		// Wait for result
		var result string
		err = workflowRun.Get(context.Background(), &result)
		if err != nil {
			shared.LogError("Workflow failed: %v", err)
		} else {
			shared.LogInfo("✅ Workflow succeeded: %s", result)
		}

		time.Sleep(time.Second) // Give some time between tests
	}

	// Run the RetryableTransferWorkflow tests
	shared.LogInfo("\n=== Testing RetryableTransferWorkflow ===")

	riskyRequest := errors.TransferRequest{
		FromAccount: "risky-account",
		ToAccount:   "target-account",
		Amount:      75.25,
		Reference:   "Risky transfer test",
	}

	// Run multiple attempts to see different retry behaviors
	for i := 1; i <= 3; i++ {
		shared.LogInfo("🧪 Running risky transfer attempt %d", i)

		workflowRun, err := shared.ExecuteWorkflow(c, errors.RetryableTransferWorkflow, riskyRequest)
		if err != nil {
			shared.LogError("Failed to start workflow: %v", err)
			continue
		}

		var result string
		err = workflowRun.Get(context.Background(), &result)
		if err != nil {
			shared.LogError("Risky transfer failed: %v", err)
		} else {
			shared.LogInfo("✅ Risky transfer succeeded: %s", result)
		}

		time.Sleep(time.Second * 2)
	}

	shared.LogInfo("\n🎉 Error handling examples completed!")
	shared.LogInfo("💡 Check the Temporal Web UI at http://localhost:8080 to see:")
	shared.LogInfo("   - Retry attempts and their timings")
	shared.LogInfo("   - Activity failures and compensations")
	shared.LogInfo("   - Different error types and how they're handled")
}
