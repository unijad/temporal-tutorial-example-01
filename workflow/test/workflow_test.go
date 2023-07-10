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

func TestCartWorkflow(t *testing.T) {
	// Set up the test suite and testing execution environment
	testSuite := &testsuite.WorkflowTestSuite{}

	t.Run("SetCart", func(t *testing.T) {
		env := testSuite.NewTestWorkflowEnvironment()
		// Mock activity implementation
		env.OnActivity(activity.SetCart, mock.Anything, mock.Anything).Return(nil)
		env.ExecuteWorkflow(workflow.SetCartWorkflow)

		// Verify that the SetCart activity was executed
		if err := env.GetWorkflowError(); err != nil {
			t.Fatalf("Workflow failed: %v", err)
		}
	})

	t.Run("GetCart", func(t *testing.T) {
		env := testSuite.NewTestWorkflowEnvironment()
		// Mock activity implementation
		products := &[]messages.Product{
			{
				Name:  "Product 1",
				Price: 1.1,
			},
			{
				Name:  "Product 2",
				Price: 1.1,
			},
			{
				Name:  "Product 3",
				Price: 1.1,
			},
		}
		env.OnActivity(activity.GetCart, mock.Anything).Return(products, nil)
		env.ExecuteWorkflow(workflow.GetCartWorkflow)

		require.True(t, env.IsWorkflowCompleted())
		require.NoError(t, env.GetWorkflowError())

		var data *[]messages.Product
		require.NoError(t, env.GetWorkflowResult(&data))
		require.Equal(t, products, data)
	})
}
