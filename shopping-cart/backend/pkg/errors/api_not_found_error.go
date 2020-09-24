package errors

import (
	"fmt"
	"net/http"
)

// NewAPINotFoundError returns an not found error
func NewAPINotFoundError(code string, err error, field string, rejectedValue interface{}) *APIError {
	message := fmt.Sprintf("Not found data for %v \"%v\"", field, rejectedValue)

	return &APIError{
		http.StatusNotFound,
		ErrorCodeNotFound,
		message,
		err,
		"",
		nil,
	}
}
