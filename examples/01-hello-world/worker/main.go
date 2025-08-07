package main

import (
	"log"

	hello "temporal-go-examples/examples/01-hello-world"
	"temporal-go-examples/shared"
)

func main() {
	// Step 1: Create a Temporal client
	// This connects to the Temporal server
	c, err := shared.CreateTemporalClient()
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// Step 2: Create a worker
	// Workers execute workflows and activities
	w := shared.CreateTemporalWorker(c)

	// Step 3: Register our workflow
	// We import the workflow from the parent package
	w.RegisterWorkflow(hello.GreetingWorkflow)

	// Step 4: Start the worker
	// This will block and keep the worker running
	shared.LogInfo("Worker is starting...")
	shared.LogInfo("Press Ctrl+C to stop the worker")
	shared.StartWorker(w)
}
