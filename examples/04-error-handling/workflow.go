package errors

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// TransferRequest represents a money transfer request
type TransferRequest struct {
	FromAccount string  `json:"from_account"`
	ToAccount   string  `json:"to_account"`
	Amount      float64 `json:"amount"`
	Reference   string  `json:"reference"`
}

// MoneyTransferWorkflow demonstrates error handling and compensation
func MoneyTransferWorkflow(ctx workflow.Context, request TransferRequest) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("MoneyTransferWorkflow started", "request", request)

	// Configure retry policy for activities
	retryPolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Second,      // Wait 1 second before first retry
		BackoffCoefficient: 2.0,              // Double the wait time each retry
		MaximumInterval:    time.Second * 30, // Don't wait more than 30 seconds
		MaximumAttempts:    3,                // Try maximum 3 times
	}

	// Configure activity options
	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 2,
		RetryPolicy:         retryPolicy,
	}
	ctx = workflow.WithActivityOptions(ctx, activityOptions)

	// Step 1: Validate accounts
	logger.Info("Validating accounts")
	err := workflow.ExecuteActivity(ctx, ValidateAccounts, request.FromAccount, request.ToAccount).Get(ctx, nil)
	if err != nil {
		logger.Error("Account validation failed", "error", err)
		return "", fmt.Errorf("account validation failed: %w", err)
	}

	// Step 2: Debit source account
	logger.Info("Debiting source account", "account", request.FromAccount, "amount", request.Amount)
	var debitTxnID string
	err = workflow.ExecuteActivity(ctx, DebitAccount, request.FromAccount, request.Amount, request.Reference).Get(ctx, &debitTxnID)
	if err != nil {
		logger.Error("Debit failed", "error", err)
		return "", fmt.Errorf("debit failed: %w", err)
	}

	// Step 3: Credit destination account
	logger.Info("Crediting destination account", "account", request.ToAccount, "amount", request.Amount)
	var creditTxnID string
	err = workflow.ExecuteActivity(ctx, CreditAccount, request.ToAccount, request.Amount, request.Reference).Get(ctx, &creditTxnID)
	if err != nil {
		logger.Error("Credit failed, starting compensation", "error", err)

		// Compensation: Reverse the debit
		logger.Info("Compensating: reversing debit", "debitTxnID", debitTxnID)
		compensateErr := workflow.ExecuteActivity(ctx, CompensateDebit, request.FromAccount, request.Amount, debitTxnID).Get(ctx, nil)
		if compensateErr != nil {
			logger.Error("CRITICAL: Compensation failed", "error", compensateErr)
			return "", fmt.Errorf("transfer failed and compensation failed: credit_error=%v, compensation_error=%v", err, compensateErr)
		}

		logger.Info("Compensation successful")
		return "", fmt.Errorf("transfer failed but system is consistent: %w", err)
	}

	// Success!
	result := fmt.Sprintf("Transfer successful: $%.2f from %s to %s (Debit: %s, Credit: %s)",
		request.Amount, request.FromAccount, request.ToAccount, debitTxnID, creditTxnID)
	logger.Info("MoneyTransferWorkflow completed successfully", "result", result)
	return result, nil
}

// RetryableTransferWorkflow demonstrates handling retryable vs non-retryable errors
func RetryableTransferWorkflow(ctx workflow.Context, request TransferRequest) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("RetryableTransferWorkflow started", "request", request)

	// Special retry policy for different types of errors
	retryPolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    time.Second * 30,
		MaximumAttempts:    5,
		// Don't retry application errors (business logic errors)
		NonRetryableErrorTypes: []string{
			"InvalidAccount",
			"InsufficientFunds",
		},
	}

	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 2,
		RetryPolicy:         retryPolicy,
	}
	ctx = workflow.WithActivityOptions(ctx, activityOptions)

	// Try the risky transfer operation
	var result string
	err := workflow.ExecuteActivity(ctx, RiskyTransferActivity, request).Get(ctx, &result)
	if err != nil {
		// Check if it's a non-retryable error
		var appErr *temporal.ApplicationError
		if temporal.IsApplicationError(err) {
			appErr = err.(*temporal.ApplicationError)
			logger.Error("Non-retryable error occurred", "type", appErr.Type(), "message", appErr.Error())
		} else {
			logger.Error("Retryable error occurred", "error", err)
		}
		return "", err
	}

	logger.Info("RetryableTransferWorkflow completed successfully", "result", result)
	return result, nil
}
