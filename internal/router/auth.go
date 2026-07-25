package router

import (
	"userrestapigo/internal/handlers"

	"github.com/gorilla/mux"
)

func registerAuthV1Routes(r *mux.Router) {
	auth := r.PathPrefix("/auth").Subrouter()

	auth.HandleFunc("/register", handlers.Register).Methods("POST")
	// Add /login and /refresh here after their handlers are implemented.
}
