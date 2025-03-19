package utils

import (
	"fmt"
	"net/http"

	"gamestone/errors"

	"github.com/go-playground/validator/v10"
)

// ErrorResponse adalah format respons standar
type ErrorResponse struct {
	Errors  []string `json:"errors"`
	Message string   `json:"message"`
}

// HandleErrors menangani berbagai jenis error
func HandleErrors(errs ...error) (ErrorResponse, int) {
	var errorsList []string
	var code int
	var message string

	for _, err := range errs {
		if err == nil {
			continue
		}

		switch e := err.(type) {
		case validator.ValidationErrors:
			for _, ve := range e {
				errorsList = append(errorsList, fmt.Sprintf("Field '%s' failed on the '%s' tag", ve.Field(), ve.Tag()))
			}
			code = http.StatusBadRequest
			message = "Invalid data"
		case *errors.HttpError:
			code = e.Code
			message = e.Message
		default:
			errorsList = append(errorsList, err.Error())
			code = http.StatusInternalServerError
			message = "Internal server error"
		}
	}

	errorResponse := ErrorResponse{
		Errors:  errorsList,
		Message: message,
	}

	return errorResponse, code
}
