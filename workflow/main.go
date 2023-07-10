package workflow

import (
	"goenv/activity"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func WorkflowDifnition() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, "weather", worker.Options{})

	// register workflows
	w.RegisterWorkflow(GetCartWorkflow)
	w.RegisterWorkflow(SetCartWorkflow)
	w.RegisterWorkflow(StartOrderWorkflow)

	// register activities
	w.RegisterActivity(activity.GetCart)
	w.RegisterActivity(activity.SetCart)
	w.RegisterActivity(activity.StartOrder)
	w.RegisterActivity(activity.CreateTransaction)
	w.RegisterActivity(activity.CreateShipping)
	w.RegisterActivity(activity.ConfirmShipping)

	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
