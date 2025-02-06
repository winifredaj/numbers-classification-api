package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/winifredaj/number-classification-api/handlers"
	"github.com/winifredaj/number-classification-api/utils"
)

func main() {

	// Load environment variables from a .env file (if present).
	// This helps manage configuration settings like the server port.
	_ = godotenv.Load()

	// Retrieve the server port from the environment variables.
	port := os.Getenv("PORT")
	if port == "" {
		// If no port is specified in the environment variables, default to "8080".
		port = "8080"
	}

	// Register the endpoint "/api/classify-number" with its corresponding handler.
	// The handler is wrapped with the CORS middleware to allow cross-origin requests.
	http.HandleFunc("/api/classify-number", utils.CORS(handlers.ClassifyNumberHandler))

	// Log a message indicating that the server is starting and on which port.
	log.Printf("Starting server on :%s", port)
	

	// Start the HTTP server and listen on the specified port.
	// If the server encounters a fatal error, log it and exit the application.
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}

