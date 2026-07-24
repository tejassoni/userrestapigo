package main

import (
	"log"
	"net/http"
	"userrestapigo/config"
	"userrestapigo/database"
	"userrestapigo/routes"
	"userrestapigo/utils"

	"github.com/gorilla/mux"
)

func main() {

	// logs
	utils.InitLogger()
	utils.Logger.Info("Application started")

	// Initialize the validator
	config.InitValidator()

	config.LoadEnv()              // Load environment variables from .env file
	database.ConnectDB()          // Establish a connection to the database
	router := mux.NewRouter()     // Create a new router using Gorilla Mux
	routes.RegisterRoutes(router) // Register the routes for the application

	log.Println("Server started on http://localhost:8080") // Log the server start message

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
