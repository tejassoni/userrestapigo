package router

import (
	"userrestapigo/internal/handlers"

	"github.com/gorilla/mux"
)

/*
Routes for the API are registered version 1
*/
func registerUserV1Routes(route *mux.Router) {
	// Register user-related routes
	route.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	route.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET")
	route.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	route.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	route.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
}

/*
Routes for the API are registered version 2
*/
func registerUserV2Routes(route *mux.Router) {
	// Register user-related routes
	route.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	route.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET")
	route.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	route.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	route.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
}
