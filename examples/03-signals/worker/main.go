package main

import (
	"log"

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

	// Create worker
	w := shared.CreateTemporalWorker(c)

	// Register workflow
	w.RegisterWorkflow(signals.DeliveryOrderWorkflow)

	// Start the worker
	shared.LogInfo("Worker is starting...")
	shared.LogInfo("Registered workflow: DeliveryOrderWorkflow")
	shared.LogInfo("This workflow supports signals: add-item, update-address, complete-order")
	shared.LogInfo("This workflow supports queries: get-status")
	shared.LogInfo("Press Ctrl+C to stop the worker")
	shared.StartWorker(w)
}
