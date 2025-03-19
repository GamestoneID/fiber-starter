package validator

import "github.com/go-playground/validator/v10"

// Provide a Validator instance
func ProvideValidator() *validator.Validate {
	return validator.New()
}
