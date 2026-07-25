package router

import "github.com/gorilla/mux"

func registerV1Routes(route *mux.Router) {
	v1 := route.PathPrefix("/api/v1").Subrouter()

	registerUserV1Routes(v1)
	registerAuthV1Routes(v1)
}
