package activities

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"go.temporal.io/sdk/activity"
)

// ValidateOrder checks if an order is valid
// Activities can perform non-deterministic operations like database calls
func ValidateOrder(ctx context.Context, order Order) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Validating order", "orderID", order.ID)

	// Simulate some validation logic
	if order.ID == "" {
		return fmt.Errorf("order ID cannot be empty")
	}
	if order.Amount <= 0 {
		return fmt.Errorf("order amount must be positive")
	}
	if order.Email == "" {
		return fmt.Errorf("email cannot be empty")
	}

	// Simulate some processing time
	time.Sleep(time.Millisecond * 100)

	// Simulate occasional failures (10% chance)
	if rand.Float32() < 0.1 {
		return fmt.Errorf("validation service temporarily unavailable")
	}

	logger.Info("Order validation successful", "orderID", order.ID)
	return nil
}

// ProcessPayment processes the payment for an order
// Returns a payment ID if successful
func ProcessPayment(ctx context.Context, order Order) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Processing payment", "orderID", order.ID, "amount", order.Amount)

	// Simulate payment processing time
	time.Sleep(time.Millisecond * 200)

	// Simulate payment failures (5% chance)
	if rand.Float32() < 0.05 {
		return "", fmt.Errorf("payment gateway error: insufficient funds")
	}

	// Generate a mock payment ID
	paymentID := fmt.Sprintf("pay_%d", time.Now().Unix())

	logger.Info("Payment processed successfully", "orderID", order.ID, "paymentID", paymentID)
	return paymentID, nil
}

// SendConfirmationEmail sends a confirmation email to the customer
func SendConfirmationEmail(ctx context.Context, order Order, paymentID string) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Sending confirmation email", "orderID", order.ID, "email", order.Email)

	// Simulate email sending time
	time.Sleep(time.Millisecond * 150)

	// Simulate email service failures (3% chance)
	if rand.Float32() < 0.03 {
		return fmt.Errorf("email service temporarily unavailable")
	}

	logger.Info("Confirmation email sent successfully",
		"orderID", order.ID,
		"email", order.Email,
		"paymentID", paymentID,
	)
	return nil
}

// Pro tip: Activities should be idempotent when possible
// This means running them multiple times with the same input
// should produce the same result without side effects
