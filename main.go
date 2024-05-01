package main

import (
	"log"
	"net/http"

	openfga "github.com/openfga/go-sdk"
)

func main() {
	config := LoadConfig()
	fgaClient := NewFGAClient(config)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleRequest(w, r, fgaClient, config.AuthorizationModelID)
	})

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request, fgaClient *openfga.Client) {
	// Example authorization check
	userID := "exampleUser"
	object := "document:exampleDocument"
	action := "read"

	authorizationModelID := "your-model-id" // replace with your actual model ID
	response, err := fgaClient.Check(openfga.CheckRequest{
		TupleKey: openfga.TupleKey{
			Object:   object,
			Relation: action,
			User:     userID,
		},
		AuthorizationModelID: authorizationModelID,
	})
	if err != nil {
		http.Error(w, "Failed to check permissions", http.StatusInternalServerError)
		return
	}

	if !response.Allowed {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Access granted: Welcome to the OpenFGA powered application!"))
}
