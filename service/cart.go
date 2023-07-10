package service

import (
	"context"
	"encoding/json"
	"goenv/activity"
	"goenv/messages"
	"log"
	"net/http"
	"strings"

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
	stringArr := strings.Split(productString, "")
	err = setCart(r.Context(), c, &stringArr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = setCart(r.Context(), c, &[]string{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getCart(ctx context.Context, temporalClient client.Client) (*[]messages.Product, error) {
	data, err := activity.GetCart(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func setCart(ctx context.Context, temporalClient client.Client, in *[]string) error {
	cart := &messages.Cart{
		Products: []string{"1", "2"},
	}
	err := activity.SetCart(ctx, cart)
	if err != nil {
		return err
	}
	return err
}
