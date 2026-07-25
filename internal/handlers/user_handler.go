package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"userrestapigo/internal/errors"
	"userrestapigo/internal/logger"
	"userrestapigo/internal/models"
	"userrestapigo/internal/repository"
	"userrestapigo/internal/requests"
	"userrestapigo/internal/responses"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var validate = validator.New()

/*
* GetUsers handles the HTTP GET request to retrieve all users.
* It fetches the users from the repository and returns them as a JSON response.
@param w http.ResponseWriter - The HTTP response writer to send the response.
@param r *http.Request - The HTTP request object containing request details.
*/
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := repository.GetUsers()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(responses.APIResponse{
			Status:  false,
			Message: "Failed to fetch users.",
			Data:    nil,
			// Error:   err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// users empty then no records found
	if len(users) == 0 {
		json.NewEncoder(w).Encode(responses.APIResponse{
			Status:  true,
			Message: "User Records no found...! Please create new user records.",
			Data:    nil,
		})
		return
	} else {
		json.NewEncoder(w).Encode(responses.APIResponse{
			Status:  true,
			Message: "Users fetched successfully.",
			Data:    users,
		})
	}
}

/*
* GetUserByID handles the HTTP GET request to retrieve a user by their ID.
* It fetches the user from the repository and returns them as a JSON response.
@param w http.ResponseWriter - The HTTP response writer to send the response.
@param r *http.Request - The HTTP request object containing request details.
*/
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(responses.APIResponse{
			Status:  false,
			Message: "Invalid user ID",
			Data:    nil,
			// Error:   err.Error(),
		})
		return
	}

	user, err := repository.GetUserByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)

			json.NewEncoder(w).Encode(responses.APIResponse{
				Status:  false,
				Message: "User not found",
				Data:    nil,
			})
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(w).Encode(responses.APIResponse{
				Status:  false,
				Message: "Error fetching user",
				Data:    nil,
			})
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(responses.APIResponse{
		Status:  true,
		Message: "User fetched successfully.",
		Data:    user,
	})
}

/*
* CreateUser handles the HTTP POST request to create a new user.
* It decodes the request payload into a User struct and calls the repository to insert the user into the database.
@param w http.ResponseWriter - The HTTP response writer to send the response.
@param r *http.Request - The HTTP request object containing request details.
*/
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Log the incoming request details
	logger.Logger.Info(
		"HTTP CreateUser request",
		"method", r.Method,
		"path", r.URL.Path,
		"ip", r.RemoteAddr,
		"user_agent", r.UserAgent(),
	)
	// Decode JSON request body into the CreateUserRequest struct
	req, err := requests.DecodeCreateUserRequest(r)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(responses.APIResponse{
			Status:  false,
			Message: "Invalid request payload",
			Data:    nil,
			// Error:   err.Error(),
		})
		return
	}
	// create user request log
	logger.Logger.Info(
		"Create user request",
		"email", req.Email,
		"name", req.Name,
		"gender", req.Gender,
		"birthdate", req.Birthdate,
		"is_active", req.IsActive,
	)

	// Validate the request struct using the validator package
	if err := validate.Struct(req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(responses.APIResponse{
			Status:  false,
			Message: "Validation failed",
			Data:    nil,
			// Error:   err.Error(),
		})
		return
	}

	// Business validation
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

	// parse birthdate string to time.Time
	birthDate, err := time.Parse("2006-01-02", req.Birthdate)
	if err != nil {
		logger.Logger.Error("Invalid birthdate format", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.APIResponse{
			Status:  false,
			Message: "Birthdate must be in YYYY-MM-DD format",
			Data:    nil,
			// Error:   err.Error(),
		})
		return
	}

	// Convert request to model
	user := models.User{
		Name:      req.Name,
		Email:     req.Email,
		Gender:    req.Gender,
		Birthdate: birthDate,
		IsActive:  req.IsActive,
		Password:  req.Password, // Hash before saving
	}

	// call repository to create user
	id, err := repository.CreateUser(user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		if errors.IsDuplicateEmailError(err) {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(responses.APIResponse{
				Status:  false,
				Message: "A user with this email already exists",
				Data:    nil,
				// Error:   err.Error(),
			})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(responses.APIResponse{
			Status:  false,
			Message: "Error creating user",
			Data:    nil,
			// Error:   err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(responses.APIResponse{
		Status:  true,
		Message: "User created successfully.",
		Data: map[string]int{
			"id": id,
		},
	})
}

/*
* UpdateUser handles the HTTP PUT request to update an existing user.
* It decodes the request payload into a User struct and calls the repository to update the user in the database.
@param w http.ResponseWriter - The HTTP response writer to send the response.
@param r *http.Request - The HTTP request object containing request details.
*/
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(responses.APIResponse{
			Status:  false,
			Message: "Invalid request payload",
			Data:    nil,
			// Error:   err.Error(),
		})
		return
	}

	// Decode JSON request body into the UpdateUserRequest struct
	req, err := requests.DecodeUpdateUserRequest(r)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(responses.APIResponse{
			Status:  false,
			Message: "Invalid request payload",
			Data:    nil,
			// Error:   err.Error(),
		})
		return
	}

	req.ID = id
	birthDate, err := time.Parse("2006-01-02", req.Birthdate)
	if err != nil {
		logger.Logger.Error("Invalid birthdate format", "error", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(responses.APIResponse{
			Status:  false,
			Message: "Birthdate must be in YYYY-MM-DD format",
			Data:    nil,
		})
		return
	}

	user := models.User{
		ID:        req.ID,
		Name:      req.Name,
		Email:     req.Email,
		Gender:    req.Gender,
		Birthdate: birthDate,
		IsActive:  req.IsActive,
	}

	err = repository.UpdateUser(user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(responses.APIResponse{
			Status:  false,
			Message: "Error updating user",
			Data:    nil,
			// Error:   err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses.APIResponse{
		Status:  true,
		Message: "User updated successfully.",
		Data:    user,
	})
}

/*
* DeleteUser handles the HTTP DELETE request to remove a user by their ID.
* It calls the repository to delete the user from the database.
@param w http.ResponseWriter - The HTTP response writer to send the response.
@param r *http.Request - The HTTP request object containing request details.
*/
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(responses.APIResponse{
			Status:  false,
			Message: "Invalid user ID",
			Data:    nil,
			// Error:   err.Error(),
		})
		return
	}

	deleted, err := repository.DeleteUser(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(responses.APIResponse{
			Status:  false,
			Message: "Error deleting user",
			Data:    nil,
			// Error:   err.Error(),
		})
		return
	}
	if !deleted {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(responses.APIResponse{
			Status:  false,
			Message: "User id not found",
			Data:    nil,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses.APIResponse{
		Status:  true,
		Message: "User deleted successfully.",
		Data:    nil,
	})
}
