package handlers

import (
	"encoding/json"
	"net/http"
	"userrestapigo/internal/logger"
	"userrestapigo/internal/models"
	"userrestapigo/internal/repository"
	"userrestapigo/internal/requests"
	"userrestapigo/internal/responses"
)

func Register(w http.ResponseWriter, r *http.Request) {
	// Log the incoming request details
	logger.Logger.Info(
		"HTTP Register request",
		"method", r.Method,
		"path", r.URL.Path,
		"ip", r.RemoteAddr,
		"user_agent", r.UserAgent(),
	)

	var req requests.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(responses.APIResponse{
			Status:  false,
			Message: "Invalid request payload",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	if err := validate.Struct(req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(responses.APIResponse{
			Status:  false,
			Message: "Validation failed",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	if req.Password != req.ConfirmPassword {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(responses.APIResponse{
			Status:  false,
			Message: "Password and Confirm Password do not match",
			Data:    nil,
		})
		return
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password, // Hash before saving
	}

	id, err := repository.CreateUser(user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(responses.APIResponse{
			Status:  false,
			Message: "Error creating user",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(responses.APIResponse{
		Status:  true,
		Message: "User registered successfully.",
		Data: map[string]int{
			"id": id,
		},
	})
}
