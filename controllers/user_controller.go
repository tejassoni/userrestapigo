package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"userrestapigo/models"
	"userrestapigo/repository"

	"github.com/gorilla/mux"
)

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

		json.NewEncoder(w).Encode(models.APIResponse{
			Status:  false,
			Message: "Failed to fetch users.",
			Data:    nil,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// users empty then no records found
	if len(users) == 0 {
		json.NewEncoder(w).Encode(models.APIResponse{
			Status:  true,
			Message: "User Records no found...! Please create new user records.",
			Data:    nil,
		})
		return
	} else {
		json.NewEncoder(w).Encode(models.APIResponse{
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

		json.NewEncoder(w).Encode(models.APIResponse{
			Status:  false,
			Message: "Invalid user ID",
			Data:    nil,
		})
		return
	}

	user, err := repository.GetUserByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)

			json.NewEncoder(w).Encode(models.APIResponse{
				Status:  false,
				Message: "User not found",
				Data:    nil,
			})
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(w).Encode(models.APIResponse{
				Status:  false,
				Message: "Error fetching user",
				Data:    nil,
			})
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.APIResponse{
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
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(models.APIResponse{
			Status:  false,
			Message: "Invalid request payload",
			Data:    nil,
		})
		return
	}

	id, err := repository.CreateUser(user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(models.APIResponse{
			Status:  false,
			Message: "Error creating user",
			Data:    nil,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(models.APIResponse{
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
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	user.ID = id
	err = repository.UpdateUser(user)
	if err != nil {
		http.Error(w, "Error updating user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
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
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = repository.DeleteUser(id)
	if err != nil {
		http.Error(w, "Error deleting user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
