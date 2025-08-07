package main

import (
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

	// Create worker
	w := shared.CreateTemporalWorker(c)

	// Register workflow and activities
	// Both need to be registered for the worker to execute them
	w.RegisterWorkflow(activities.OrderProcessingWorkflow)
	w.RegisterActivity(activities.ValidateOrder)
	w.RegisterActivity(activities.ProcessPayment)
	w.RegisterActivity(activities.SendConfirmationEmail)

	// Start the worker
	shared.LogInfo("Worker is starting...")
	shared.LogInfo("Registered workflow: OrderProcessingWorkflow")
	shared.LogInfo("Registered activities: ValidateOrder, ProcessPayment, SendConfirmationEmail")
	shared.LogInfo("Press Ctrl+C to stop the worker")
	shared.StartWorker(w)
}
