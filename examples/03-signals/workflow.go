package signals

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

// OrderStatus represents the current state of an order
type OrderStatus struct {
	Items   []string `json:"items"`
	Address string   `json:"address"`
	Status  string   `json:"status"`
}

// DeliveryOrderWorkflow demonstrates signals and queries
// This workflow can receive updates while running
func DeliveryOrderWorkflow(ctx workflow.Context, initialItem string, address string) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("DeliveryOrderWorkflow started", "initialItem", initialItem, "address", address)

	// Initialize order status
	orderStatus := OrderStatus{
		Items:   []string{initialItem},
		Address: address,
		Status:  "Preparing",
	}

	// Set up signal channels
	addItemSignal := workflow.GetSignalChannel(ctx, "add-item")
	updateAddressSignal := workflow.GetSignalChannel(ctx, "update-address")
	completeOrderSignal := workflow.GetSignalChannel(ctx, "complete-order")

	// Set up query handler - clients can query the current status
	err := workflow.SetQueryHandler(ctx, "get-status", func() (OrderStatus, error) {
		return orderStatus, nil
	})
	if err != nil {
		return "", err
	}

	logger.Info("Order initialized", "status", orderStatus)

	// Main workflow loop - wait for signals
	for {
		selector := workflow.NewSelector(ctx)

		// Handle add item signal
		selector.AddReceive(addItemSignal, func(c workflow.ReceiveChannel, more bool) {
			var newItem string
			c.Receive(ctx, &newItem)
			orderStatus.Items = append(orderStatus.Items, newItem)
			logger.Info("Item added to order", "item", newItem, "totalItems", len(orderStatus.Items))
		})

		// Handle address update signal
		selector.AddReceive(updateAddressSignal, func(c workflow.ReceiveChannel, more bool) {
			var newAddress string
			c.Receive(ctx, &newAddress)
			orderStatus.Address = newAddress
			logger.Info("Address updated", "newAddress", newAddress)
		})

		// Handle order completion signal
		selector.AddReceive(completeOrderSignal, func(c workflow.ReceiveChannel, more bool) {
			var message string
			c.Receive(ctx, &message)
			orderStatus.Status = "Completed"
			logger.Info("Order completion signal received", "message", message)
		})

		// Add a timeout to automatically progress the order
		selector.AddFuture(workflow.NewTimer(ctx, time.Second*10), func(f workflow.Future) {
			if orderStatus.Status == "Preparing" {
				orderStatus.Status = "Out for Delivery"
				logger.Info("Order status updated", "status", orderStatus.Status)
			} else if orderStatus.Status == "Out for Delivery" {
				orderStatus.Status = "Completed"
				logger.Info("Order delivered automatically")
			}
		})

		// Wait for one of the above to happen
		selector.Select(ctx)

		// Exit if order is completed
		if orderStatus.Status == "Completed" {
			break
		}
	}

	result := fmt.Sprintf("Order completed! Items: %v, Delivered to: %s",
		orderStatus.Items, orderStatus.Address)
	logger.Info("DeliveryOrderWorkflow completed", "result", result)
	return result, nil
}
