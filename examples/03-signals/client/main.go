package main

import (
	"context"
	"log"
	"time"

	signals "temporal-go-examples/examples/03-signals"
	"temporal-go-examples/shared"
)

func main() {
	// Create Temporal client
	c, err := shared.CreateTemporalClient()
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// Start the workflow
	shared.LogInfo("Starting DeliveryOrderWorkflow...")

	workflowRun, err := shared.ExecuteWorkflow(c, signals.DeliveryOrderWorkflow, "Pizza", "123 Main St")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	workflowID := workflowRun.GetID()
	shared.LogInfo("Workflow started! WorkflowID: %s", workflowID)

	// Give the workflow a moment to start
	time.Sleep(time.Second)

	// Send some signals to interact with the running workflow
	shared.LogInfo("Sending signals to update the order...")

	// Add an item
	err = c.SignalWorkflow(context.Background(), workflowID, "", "add-item", "Coke")
	if err != nil {
		log.Fatalln("Unable to signal workflow", err)
	}
	shared.LogInfo("✅ Added item: Coke")

	time.Sleep(time.Second)

	// Add another item
	err = c.SignalWorkflow(context.Background(), workflowID, "", "add-item", "Fries")
	if err != nil {
		log.Fatalln("Unable to signal workflow", err)
	}
	shared.LogInfo("✅ Added item: Fries")

	time.Sleep(time.Second)

	// Update address
	err = c.SignalWorkflow(context.Background(), workflowID, "", "update-address", "456 Oak Avenue")
	if err != nil {
		log.Fatalln("Unable to signal workflow", err)
	}
	shared.LogInfo("✅ Updated address: 456 Oak Avenue")

	// Query the current status
	time.Sleep(time.Second)
	resp, err := c.QueryWorkflow(context.Background(), workflowID, "", "get-status")
	if err != nil {
		log.Fatalln("Unable to query workflow", err)
	}

	var status signals.OrderStatus
	err = resp.Get(&status)
	if err != nil {
		log.Fatalln("Unable to decode query result", err)
	}
	shared.LogInfo("📊 Current status: %s, Items: %v, Address: %s",
		status.Status, status.Items, status.Address)

	// Wait a bit more and then complete the order
	time.Sleep(time.Second * 3)

	err = c.SignalWorkflow(context.Background(), workflowID, "", "complete-order", "Customer confirmed delivery")
	if err != nil {
		log.Fatalln("Unable to signal workflow", err)
	}
	shared.LogInfo("✅ Sent completion signal")

	// Wait for the workflow to complete
	var result string
	err = workflowRun.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Workflow failed", err)
	}

	// Print the result
	shared.LogInfo("🎉 Final result: %s", result)
	shared.LogInfo("Signals example completed successfully!")
}
