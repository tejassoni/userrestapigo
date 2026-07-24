package config

import (
	"userrestapigo/validators"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()
	Validate.RegisterValidation("unique_email", validators.UniqueEmail)
}
