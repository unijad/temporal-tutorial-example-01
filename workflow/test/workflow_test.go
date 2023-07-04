package test

import (
	"goenv/activity"
	"goenv/messages"
	"goenv/workflow"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func TestWeatherWorkflow(t *testing.T) {
	// Set up the test suite and testing execution environment
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	// Mock activity implementation
	env.OnActivity(activity.GetWeather, mock.Anything, mock.Anything).Return(messages.WeatherData{
		Temperature: 41,
		Humidity:    80,
		WindSpeed:   4,
	}, nil)

	env.ExecuteWorkflow(workflow.WeatherWorkflow, "Cairo")

	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	var data []messages.WeatherData
	require.NoError(t, env.GetWorkflowResult(&data))
	require.Equal(t, []messages.WeatherData{
		{
			Temperature: 41,
			Humidity:    80,
			WindSpeed:   4,
		},
	}, data)
}
