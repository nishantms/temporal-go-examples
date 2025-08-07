package shared

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

const (
	// TaskQueue is the default task queue name used across examples
	TaskQueue = "temporal-learning-queue"

	// Namespace is the Temporal namespace (use "default" for local development)
	Namespace = "default"

	// HostPort is the Temporal server address
	HostPort = "localhost:7233"
)

// CreateTemporalClient creates and returns a Temporal client
// This is used by both workers and clients to connect to Temporal
func CreateTemporalClient() (client.Client, error) {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{
		HostPort:  HostPort,
		Namespace: Namespace,
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
		return nil, err
	}
	return c, nil
}

// CreateTemporalWorker creates and returns a Temporal worker
// Workers are responsible for executing workflows and activities
func CreateTemporalWorker(c client.Client) worker.Worker {
	w := worker.New(c, TaskQueue, worker.Options{})
	return w
}

// StartWorker starts a worker and blocks until it's stopped
// This is typically called in your worker main function
func StartWorker(w worker.Worker) {
	log.Println("Starting worker on task queue:", TaskQueue)
	err := w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}

// ExecuteWorkflow is a helper function to start a workflow execution
func ExecuteWorkflow(c client.Client, workflowFunc interface{}, args ...interface{}) (client.WorkflowRun, error) {
	workflowOptions := client.StartWorkflowOptions{
		ID:        "learning-workflow-" + RandomID(),
		TaskQueue: TaskQueue,
	}

	return c.ExecuteWorkflow(context.Background(), workflowOptions, workflowFunc, args...)
}
