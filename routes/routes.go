package routes

import (
	"encoding/json"
	"net/http"
	"userrestapigo/config"
	"userrestapigo/controllers"

	"github.com/gorilla/mux"
)

// RegisterRoutes registers the routes for the application with the provided router
func RegisterRoutes(router *mux.Router) {

	// Define a simple health check endpoint
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"status":      true,
			"message":     "User REST API is running",
			"data":        nil,
			"app_name":    config.APPNAME,
			"app_version": config.APPVERSION,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}).Methods("GET")

	// API Version 1
	api := router.PathPrefix("/api").Subrouter()
	v1 := api.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	v1.HandleFunc("/users/{id}", controllers.GetUserByID).Methods("GET")
	v1.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	v1.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	v1.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
}
