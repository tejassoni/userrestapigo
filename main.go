package main

import (
	"log"
	"net/http"
	"userrestapigo/config"
	"userrestapigo/routes"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv()              // Load environment variables from .env file
	config.ConnectDB()            // Establish a connection to the database
	router := mux.NewRouter()     // Create a new router using Gorilla Mux
	routes.RegisterRoutes(router) // Register the routes for the application

	log.Fatal(http.ListenAndServe(":8080", router)) // Start the HTTP server on port 8080 and log any fatal errors
}
