package errors

import "net/http"

// NewAPIUnexpectedError returns an unexpected error
func NewAPIUnexpectedError(err error) *APIError {
	return &APIError{
		http.StatusInternalServerError,
		ErrorCodeUnexpectedError,
		"Unexpected error",
		err,
		"",
		nil,
	}
}
