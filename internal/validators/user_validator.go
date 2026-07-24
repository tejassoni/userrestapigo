package validators

import (
	"userrestapigo/repository"

	"github.com/go-playground/validator/v10"
)

/*
* UniqueEmail is a custom validation function that checks if the provided email is unique in the database.
* It queries the database to determine if the email already exists.
@param fl validator.FieldLevel - The field level information provided by the validator.
@return bool - Returns true if the email is unique (does not exist in the database), otherwise returns false.
*/
func UniqueEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	exists, err := repository.EmailExists(email)
	if err != nil {
		return false
	}

	return !exists
}

func BirthDateValidation(fl validator.FieldLevel) bool {
	birthDate := fl.Field().String()

	exists, err := repository.BirthDateExists(birthDate)
	if err != nil {
		return false
	}

	return !exists
}
