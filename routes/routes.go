package routes

import (
	"userrestapigo/controllers"

	"github.com/gorilla/mux"
)

// RegisterRoutes registers the routes for the application with the provided router
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetUserByID).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
}
