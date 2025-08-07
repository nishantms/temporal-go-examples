package shared

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomID generates a random ID for workflow instances
func RandomID() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%d", rand.Intn(1000000))
}

// LogInfo prints an informational message with timestamp
func LogInfo(message string, args ...interface{}) {
	timestamp := time.Now().Format("15:04:05")
	fmt.Printf("[%s] INFO: %s\n", timestamp, fmt.Sprintf(message, args...))
}

// LogError prints an error message with timestamp
func LogError(message string, args ...interface{}) {
	timestamp := time.Now().Format("15:04:05")
	fmt.Printf("[%s] ERROR: %s\n", timestamp, fmt.Sprintf(message, args...))
}
