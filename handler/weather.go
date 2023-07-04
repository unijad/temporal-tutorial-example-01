package handler

import (
	"encoding/json"
	"goenv/messages"
	"goenv/workflow"
	"log"
	"net/http"

	"go.temporal.io/sdk/client"
)

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	// execute weather workflow with the city name from request query
	cityName := r.URL.Query().Get("city")
	if cityName == "" {
		http.Error(w, "city name is required", http.StatusBadRequest)
		return
	}

	// create a new temporal client
	// set up the worker
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	we, err := c.ExecuteWorkflow(r.Context(), client.StartWorkflowOptions{
		ID:        "weather_workflow",
		TaskQueue: "weather",
	}, workflow.WeatherWorkflow, cityName)
	if err != nil {
		http.Error(w, "unable to start workflow", http.StatusInternalServerError)
		return
	}

	// wait for workflow to complete
	var result []messages.WeatherData
	if err := we.Get(r.Context(), &result); err != nil {
		http.Error(w, "unable to get workflow result", http.StatusInternalServerError)
		return
	}

	// convert result to json in key-value pais
	response := make(map[string]interface{})
	for _, data := range result {
		response[cityName] = data
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "unable to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
