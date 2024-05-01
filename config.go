package main

import (
	"log"
	"os"

	openfga "github.com/openfga/go-sdk"
)

// AppConfig holds all the configuration that the application uses.
type AppConfig struct {
	FGAAPIURL            string
	FGAAPIKey            string
	AuthorizationModelID string
}

// LoadConfig reads environment variables and forms the application configuration.
func LoadConfig() *AppConfig {
	fgaAPIURL := os.Getenv("OPENFGA_API_URL")
	if fgaAPIURL == "" {
		log.Fatal("OPENFGA_API_URL must be set")
	}

	fgaAPIKey := os.Getenv("OPENFGA_API_KEY")
	if fgaAPIKey == "" {
		log.Fatal("OPENFGA_API_KEY must be set")
	}

	authorizationModelID := os.Getenv("OPENFGA_AUTH_MODEL_ID")
	if authorizationModelID == "" {
		log.Fatal("OPENFGA_AUTH_MODEL_ID must be set")
	}

	return &AppConfig{
		FGAAPIURL:            fgaAPIURL,
		FGAAPIKey:            fgaAPIKey,
		AuthorizationModelID: authorizationModelID,
	}
}

// NewFGAClient initializes and returns an OpenFGA client using the AppConfig.
func NewFGAClient(config *AppConfig) *openfga.Client {
	client, err := openfga.NewClient(openfga.Config{
		APIURL: config.FGAAPIURL,
		APIKey: config.FGAAPIKey,
	})
	if err != nil {
		log.Fatal("Failed to create FGA client: ", err)
	}
	return client
}
