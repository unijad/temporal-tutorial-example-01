package workflow

import (
	"goenv/activity"
	"goenv/messages"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// define the workflow function
func SetCartWorkflow(ctx workflow.Context, cart *messages.Cart) error {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	// start the activities
	result := &messages.Cart{}
	err := workflow.ExecuteActivity(ctx, activity.SetCart, cart).Get(ctx, result)
	if err != nil {
		return temporal.NewApplicationError(err.Error(), "error")
	}

	return nil
}

// define the workflow function
func GetCartWorkflow(ctx workflow.Context) (*[]messages.Product, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	// start the activities
	result := &[]messages.Product{}
	err := workflow.ExecuteActivity(ctx, activity.GetCart).Get(ctx, result)
	if err != nil {
		return nil, temporal.NewApplicationError(err.Error(), "error")
	}

	return result, nil
}

// start order workflow
func StartOrderWorkflow(ctx workflow.Context) error {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:        time.Second,
			BackoffCoefficient:     2.0,
			MaximumInterval:        time.Minute * 60,
			MaximumAttempts:        0,
			NonRetryableErrorTypes: []string{},
		},
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	// register workflow signal to cancel workflow using orderId

	// start the activities
	orderId := workflow.ExecuteActivity(ctx, activity.StartOrder)
	workflow.ExecuteActivity(ctx, activity.CreateTransaction, orderId)
	workflow.ExecuteActivity(ctx, activity.CreateShipping, orderId)
	workflow.ExecuteActivity(ctx, activity.ConfirmShipping, orderId)

	// wait for activities to complete

	return nil
}
