package handler

import (
	"context"
	"encoding/json"
	"errors"
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

	result, err := WeatherGet(context.Background(), c, cityName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	if err != nil {
		http.Error(w, "unable to write response", http.StatusInternalServerError)
		return
	}
}

func WeatherGet(ctx context.Context, temporalClient client.Client, cityName string) ([]*messages.WeatherData, error) {
	we, err := temporalClient.ExecuteWorkflow(ctx, client.StartWorkflowOptions{
		ID:        "weather_workflow",
		TaskQueue: "weather",
	}, workflow.WeatherWorkflow, cityName)
	if err != nil {
		return nil, err
	}

	// wait for workflow to complete
	var result []*messages.WeatherData
	if err := we.Get(ctx, &result); err != nil {
		err := &messages.Error{
			RunId:   we.GetRunID(),
			Message: err.Error(),
		}
		return nil, errors.New(string(err.ErrorJSON()))
	}

	return result, nil
}

func WeatherUpdate(ctx context.Context, temporalClient client.Client, cityName string) {
	data := messages.WeatherData{
		CityName: cityName,
	}
	err := temporalClient.SignalWorkflow(ctx, "your-workflow-id", "", "your-signal-name", data)
	if err != nil {
		log.Fatalln("Error sending the Signal", err)
		return
	}
}
