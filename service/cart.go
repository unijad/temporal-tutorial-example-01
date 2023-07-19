package service

import (
	"context"
	"encoding/json"
	"goenv/messages"
	"goenv/workflow"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
)

func CartGetHandler(w http.ResponseWriter, r *http.Request) {
	// create a new temporal client
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	// get the cart
	result, err := getCart(r.Context(), c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "unable to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "unable to write response", http.StatusInternalServerError)
		return
	}
}

func CartSetHandler(w http.ResponseWriter, r *http.Request) {
	productString := r.URL.Query().Get("products")
	// create a new temporal client
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	// split productString
	stringArr := strings.Split(productString, ",")
	err = setCart(r.Context(), c, &messages.Cart{
		Products: stringArr,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	jsonResponse, err := json.Marshal("{status: 'ok'}")
	if err != nil {
		http.Error(w, "unable to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "unable to write response", http.StatusInternalServerError)
		return
	}
}

func getCart(ctx context.Context, temporalClient client.Client) (*[]messages.Product, error) {
	we, err := temporalClient.ExecuteWorkflow(ctx, client.StartWorkflowOptions{
		ID:        "GetCartWorkflow_" + uuid.New().String(),
		TaskQueue: "cart",
	}, workflow.GetCartWorkflow)
	if err != nil {
		return nil, err
	}

	result := &[]messages.Product{}
	if err := we.Get(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func setCart(ctx context.Context, temporalClient client.Client, in *messages.Cart) error {
	we, err := temporalClient.ExecuteWorkflow(ctx, client.StartWorkflowOptions{
		ID:        "SetCartWorkflow_" + uuid.New().String(),
		TaskQueue: "cart",
	}, workflow.SetCartWorkflow, in)
	if err != nil {
		return err
	}

	if err := we.Get(ctx, nil); err != nil {
		return err
	}

	return nil
}
