package activity

import (
	"goenv/repository"

	"go.temporal.io/sdk/workflow"
)

func StartOrder(ctx workflow.Context) (orderId string, err error) {
	c := &repository.Order{}
	// create order record
	print("StartOrder", c)
	return orderId, err
}

func CreateTransaction(ctx workflow.Context) (err error) {
	c := &repository.Order{}
	// create transaction record
	print("CreateTransaction", c)
	return err
}

func CreateShipping(ctx workflow.Context) (err error) {
	c := &repository.Order{}
	// create shipping record
	print("CreateShipping", c)
	return err
}

func ConfirmShipping(ctx workflow.Context) (err error) {
	c := &repository.Order{}
	// keep retying until shipping is confirmed
	print("ConfirmShipping", c)
	return err
}
