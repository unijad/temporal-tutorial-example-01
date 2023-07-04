package workflow

import (
	"goenv/activity"
	"goenv/messages"
	"time"

	"go.temporal.io/sdk/workflow"
)

// define the workflow function
func WeatherWorkflow(ctx workflow.Context, cityName string) ([]messages.WeatherData, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	// start the activities
	currentWeatherFuture := workflow.ExecuteActivity(ctx, activity.GetWeather, cityName)

	// wait for activities to complete
	var current messages.WeatherData
	if err := currentWeatherFuture.Get(ctx, &current); err != nil {
		return nil, err
	}

	var response []messages.WeatherData
	// combine results
	response = append(response, current)

	return response, nil
}
