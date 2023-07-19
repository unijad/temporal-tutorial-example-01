package service

import (
	"goenv/activity"
	"goenv/workflow"
	"log"
	"net/http"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func Server() {
	// set up the worker
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, "cart", worker.Options{})
	w.RegisterWorkflow(workflow.SetCartWorkflow)
	w.RegisterWorkflow(workflow.GetCartWorkflow)
	w.RegisterActivity(activity.SetCart)
	w.RegisterActivity(activity.GetCart)

	mux := http.NewServeMux()
	mux.HandleFunc("/cart/set", CartSetHandler) // curl -X POST http://localhost:5000/cart/set\?products\=1,2,3
	mux.HandleFunc("/cart", CartGetHandler)     // curl -X GET http://localhost:5000/cart

	server := &http.Server{Addr: ":5000", Handler: mux}

	// start the worker and the web server
	go func() {
		err = w.Run(worker.InterruptCh())
		if err != nil {
			log.Fatalln("unable to start Worker", err)
		}
	}()

	log.Fatal(server.ListenAndServe())
}
