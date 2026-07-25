package validator

import (
	"userrestapigo/internal/repository"

	"github.com/go-playground/validator/v10"
)

/*
* New initializes the validator and registers custom validation functions.
* It sets up the validator to be used for validating request payloads and other data structures.
 */
func New() {
	validate := validator.New()
	validate.RegisterValidation("unique_email", UniqueEmail)
	validate.RegisterValidation("user_id_exists", UserIDExists)
}

/*
* UniqueEmail is a custom validation function that checks if the provided email is unique in the database.
* It queries the database to determine if the email already exists.
@param fl validator.FieldLevel - The field level information provided by the validator.
@return bool - Returns true if the email is unique (does not exist in the database), otherwise returns false.
*/
func UniqueEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	exists, err := repository.UserEmailExists(email)
	if err != nil {
		return false
	}

	return !exists
}

func UserIDExists(fl validator.FieldLevel) bool {
	id := fl.Field().Int()

	exists, err := repository.UserIDExists(int(id))
	if err != nil {
		return false
	}

	return exists
}
