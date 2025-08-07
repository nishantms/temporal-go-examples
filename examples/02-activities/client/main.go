package main

import (
	"context"
	"log"

	activities "temporal-go-examples/examples/02-activities"
	"temporal-go-examples/shared"
)

func main() {
	// Create Temporal client
	c, err := shared.CreateTemporalClient()
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// Create a sample order
	order := activities.Order{
		ID:      "order-12345",
		UserID:  "user-67890",
		Email:   "customer@example.com",
		Amount:  99.99,
		Product: "Premium Subscription",
	}

	// Start the workflow
	shared.LogInfo("Starting OrderProcessingWorkflow...")
	shared.LogInfo("Order details: ID=%s, Amount=$%.2f, Email=%s",
		order.ID, order.Amount, order.Email)

	workflowRun, err := shared.ExecuteWorkflow(c, activities.OrderProcessingWorkflow, order)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	shared.LogInfo("Workflow started! WorkflowID: %s, RunID: %s",
		workflowRun.GetID(), workflowRun.GetRunID())

	// Wait for the workflow to complete
	var result string
	err = workflowRun.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Workflow failed", err)
	}

	// Print the result
	shared.LogInfo("Workflow result: %s", result)
	shared.LogInfo("Order processing completed! 🎉")

	// Pro tip: You can also get workflow execution info
	shared.LogInfo("Check the Temporal Web UI at http://localhost:8080 to see the workflow execution!")
}
