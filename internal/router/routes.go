package router

import "github.com/gorilla/mux"

func New() *mux.Router {
	router := mux.NewRouter()

	registerHealthRoutes(router)
	registerV1Routes(router)
	registerV2Routes(router)

	return router
}
