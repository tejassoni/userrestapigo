package router

import (
	"encoding/json"
	"net/http"
	"userrestapigo/internal/config"

	"github.com/gorilla/mux"
)

func registerHealthRoutes(router *mux.Router) {
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
}
