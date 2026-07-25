package router

import "github.com/gorilla/mux"

func registerV2Routes(route *mux.Router) {
	v2 := route.PathPrefix("/api/v2").Subrouter()

	registerUserV2Routes(v2)
	// registerAuthV2Routes(v2)
}
