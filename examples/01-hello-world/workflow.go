package hello

import (
	"fmt"

	"go.temporal.io/sdk/workflow"
)

// GreetingWorkflow is a simple workflow that takes a name and returns a greeting
//
// Key Points:
// - Must take workflow.Context as first parameter
// - Must be deterministic (same input always produces same output)
// - Can call activities, sleep, wait for signals, etc.
// - Automatically retried on failure
func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	// Get a logger that's safe to use in workflows
	logger := workflow.GetLogger(ctx)
	logger.Info("GreetingWorkflow started", "name", name)

	// Create the greeting message
	greeting := fmt.Sprintf("Hello, %s! Welcome to Temporal! 🎉", name)

	logger.Info("GreetingWorkflow completed", "greeting", greeting)
	return greeting, nil
}

// Pro tip: You can add more workflows in the same file
// Just make sure to register them all in your worker!
