package activities

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// Order represents an order to be processed
type Order struct {
	ID      string  `json:"id"`
	UserID  string  `json:"user_id"`
	Email   string  `json:"email"`
	Amount  float64 `json:"amount"`
	Product string  `json:"product"`
}

// OrderProcessingWorkflow orchestrates the order processing steps
// This workflow calls multiple activities in sequence
func OrderProcessingWorkflow(ctx workflow.Context, order Order) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("OrderProcessingWorkflow started", "orderID", order.ID)

	// Activity options - configure timeouts and retry policies
	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 2, // Max time for activity to complete
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Second * 30,
			MaximumAttempts:    3,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, activityOptions)

	// Step 1: Validate the order
	logger.Info("Validating order", "orderID", order.ID)
	err := workflow.ExecuteActivity(ctx, ValidateOrder, order).Get(ctx, nil)
	if err != nil {
		logger.Error("Order validation failed", "error", err)
		return "", fmt.Errorf("order validation failed: %w", err)
	}

	// Step 2: Process payment
	logger.Info("Processing payment", "amount", order.Amount)
	var paymentID string
	err = workflow.ExecuteActivity(ctx, ProcessPayment, order).Get(ctx, &paymentID)
	if err != nil {
		logger.Error("Payment processing failed", "error", err)
		return "", fmt.Errorf("payment processing failed: %w", err)
	}

	// Step 3: Send confirmation email
	logger.Info("Sending confirmation email", "email", order.Email)
	err = workflow.ExecuteActivity(ctx, SendConfirmationEmail, order, paymentID).Get(ctx, nil)
	if err != nil {
		logger.Error("Failed to send confirmation email", "error", err)
		// Note: We don't fail the workflow if email fails
		// This is a business decision - order is still processed
		logger.Warn("Order processed but confirmation email failed")
	}

	result := fmt.Sprintf("Order %s processed successfully! Payment ID: %s", order.ID, paymentID)
	logger.Info("OrderProcessingWorkflow completed", "result", result)
	return result, nil
}
