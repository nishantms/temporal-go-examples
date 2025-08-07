package main

import (
	"log"

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

	// Create worker
	w := shared.CreateTemporalWorker(c)

	// Register workflows and activities
	w.RegisterWorkflow(errors.MoneyTransferWorkflow)
	w.RegisterWorkflow(errors.RetryableTransferWorkflow)
	w.RegisterActivity(errors.ValidateAccounts)
	w.RegisterActivity(errors.DebitAccount)
	w.RegisterActivity(errors.CreditAccount)
	w.RegisterActivity(errors.CompensateDebit)
	w.RegisterActivity(errors.RiskyTransferActivity)

	// Start the worker
	shared.LogInfo("Worker is starting...")
	shared.LogInfo("Registered workflows: MoneyTransferWorkflow, RetryableTransferWorkflow")
	shared.LogInfo("Registered activities: ValidateAccounts, DebitAccount, CreditAccount, CompensateDebit, RiskyTransferActivity")
	shared.LogInfo("This example demonstrates error handling and compensation patterns")
	shared.LogInfo("Press Ctrl+C to stop the worker")
	shared.StartWorker(w)
}
