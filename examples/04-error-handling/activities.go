package errors

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/temporal"
)

// ValidateAccounts checks if both accounts exist and are valid
func ValidateAccounts(ctx context.Context, fromAccount, toAccount string) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Validating accounts", "from", fromAccount, "to", toAccount)

	// Simulate validation time
	time.Sleep(time.Millisecond * 100)

	// Simulate validation failures
	if fromAccount == "invalid-account" {
		return temporal.NewNonRetryableApplicationError("account not found", "InvalidAccount", nil)
	}
	if toAccount == "invalid-account" {
		return temporal.NewNonRetryableApplicationError("account not found", "InvalidAccount", nil)
	}

	// Simulate temporary service issues (will be retried)
	if rand.Float32() < 0.2 {
		return fmt.Errorf("account service temporarily unavailable")
	}

	logger.Info("Account validation successful")
	return nil
}

// DebitAccount withdraws money from an account
func DebitAccount(ctx context.Context, account string, amount float64, reference string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Debiting account", "account", account, "amount", amount)

	// Simulate processing time
	time.Sleep(time.Millisecond * 200)

	// Simulate insufficient funds (non-retryable)
	if account == "broke-account" {
		return "", temporal.NewNonRetryableApplicationError("insufficient funds", "InsufficientFunds", nil)
	}

	// Simulate network issues (retryable)
	if rand.Float32() < 0.15 {
		return "", fmt.Errorf("database connection failed")
	}

	// Generate transaction ID
	txnID := fmt.Sprintf("debit_%d", time.Now().Unix())
	logger.Info("Debit successful", "txnID", txnID)
	return txnID, nil
}

// CreditAccount adds money to an account
func CreditAccount(ctx context.Context, account string, amount float64, reference string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Crediting account", "account", account, "amount", amount)

	// Simulate processing time
	time.Sleep(time.Millisecond * 200)

	// Simulate account service issues (retryable)
	if rand.Float32() < 0.3 {
		return "", fmt.Errorf("credit service temporarily unavailable")
	}

	// Generate transaction ID
	txnID := fmt.Sprintf("credit_%d", time.Now().Unix())
	logger.Info("Credit successful", "txnID", txnID)
	return txnID, nil
}

// CompensateDebit reverses a debit transaction
func CompensateDebit(ctx context.Context, account string, amount float64, originalTxnID string) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Compensating debit", "account", account, "amount", amount, "originalTxn", originalTxnID)

	// Simulate compensation time
	time.Sleep(time.Millisecond * 150)

	// Compensation should rarely fail, but simulate occasional issues
	if rand.Float32() < 0.05 {
		return fmt.Errorf("compensation service failed - manual intervention required")
	}

	logger.Info("Compensation successful", "account", account, "reversedTxn", originalTxnID)
	return nil
}

// RiskyTransferActivity demonstrates different types of errors
func RiskyTransferActivity(ctx context.Context, request TransferRequest) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Attempting risky transfer", "request", request)

	// Simulate processing time
	time.Sleep(time.Millisecond * 300)

	// Simulate various types of failures
	random := rand.Float32()

	switch {
	case random < 0.2:
		// Non-retryable: Invalid account
		return "", temporal.NewNonRetryableApplicationError(
			fmt.Sprintf("account %s does not exist", request.FromAccount),
			"InvalidAccount",
			nil,
		)
	case random < 0.3:
		// Non-retryable: Insufficient funds
		return "", temporal.NewNonRetryableApplicationError(
			"insufficient funds in account",
			"InsufficientFunds",
			nil,
		)
	case random < 0.6:
		// Retryable: Network error
		return "", fmt.Errorf("network timeout during transfer")
	case random < 0.7:
		// Retryable: Database error
		return "", fmt.Errorf("database connection lost")
	default:
		// Success!
		result := fmt.Sprintf("Transfer successful: $%.2f from %s to %s",
			request.Amount, request.FromAccount, request.ToAccount)
		logger.Info("Transfer completed successfully")
		return result, nil
	}
}
