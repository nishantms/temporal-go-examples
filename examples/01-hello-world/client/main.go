package main

import (
	"context"
	"log"

	hello "temporal-go-examples/examples/01-hello-world"
	"temporal-go-examples/shared"
)

func main() {
	// Step 1: Create a Temporal client
	c, err := shared.CreateTemporalClient()
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// Step 2: Start the workflow
	shared.LogInfo("Starting GreetingWorkflow...")

	workflowRun, err := shared.ExecuteWorkflow(c, hello.GreetingWorkflow, "Temporal World")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	shared.LogInfo("Workflow started! WorkflowID: %s, RunID: %s",
		workflowRun.GetID(), workflowRun.GetRunID())

	// Step 3: Wait for the workflow to complete and get the result
	var result string
	err = workflowRun.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable to get workflow result", err)
	}

	// Step 4: Print the result
	shared.LogInfo("Workflow result: %s", result)
	shared.LogInfo("Workflow completed successfully! 🎉")
}
