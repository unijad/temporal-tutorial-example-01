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

	w := worker.New(c, "weather", worker.Options{})
	w.RegisterWorkflow(workflow.WeatherWorkflow)
	w.RegisterActivity(activity.GetWeather)

	mux := http.NewServeMux()
	mux.HandleFunc("/weather", WeatherServiceHandler) // curl -X GET http://localhost:8080/weather?city=Cairo
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
