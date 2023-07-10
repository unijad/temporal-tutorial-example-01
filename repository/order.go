package repository

import "context"

type Order struct{}

func (o *Order) GetByWorkflowId(ctx context.Context) (workflowId string, err error) {
	return workflowId, err
}

func (o *Order) Create(ctx context.Context, workflowId string) (err error) {
	return err
}

func (o *Order) Update(ctx context.Context, workflowId string) (err error) {
	return err
}
