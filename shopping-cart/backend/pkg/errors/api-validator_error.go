package errors

import (
	"fmt"
	"net/http"
)

// ValidatorError returns validator error
type ValidatorError struct {
	Code string `json:"code"`

	Message string `json:"message"`

	Model string `json:"model,omitempty"`

	Field string `json:"field,omitempty"`

	RejectedValue interface{} `json:"rejectedValue,omitempty"`
}

// ValidatorErrors returns a list of validator error
type ValidatorErrors []*ValidatorError

// NewAPIValidatorError returns an API validator error
func NewAPIValidatorError(
	err error,
	validatorErrors ...*ValidatorError,
) *APIError {
	// Error message for the general validator error
	validatorMassage := "The request data is invalid"

	return &APIError{
		http.StatusBadRequest,
		ErrorCodeFailedValidator,
		validatorMassage,
		err,
		"",
		validatorErrors,
	}
}

// NewRequiredValidatorError returns a required error
func NewRequiredValidatorError(
	model string,
	field string,
) *ValidatorError {
	message := fmt.Sprintf("%v must be not null", field)

	return &ValidatorError{
		ErrorCodeValidatorRequired,
		message,
		model,
		field,
		nil,
	}
}

// NewUniqueValidatorError returns a violated unique constrant
func NewUniqueValidatorError(
	model string,
	field string,
	rejectedValue interface{},
) *ValidatorError {
	message := fmt.Sprintf("%v '%v' already exists", field, rejectedValue)

	return &ValidatorError{
		ErrorCodeValidatorRequired,
		message,
		model,
		field,
		rejectedValue,
	}
}

// NewNotFoundReferenceValidatorError returns a violated foreign key constrant
func NewNotFoundReferenceValidatorError(
	model string,
	field string,
	rejectedValue interface{},
) *ValidatorError {
	message := fmt.Sprintf("%v '%v' reference is not found", field, rejectedValue)

	return &ValidatorError{
		ErrorCodeValidatorRequired,
		message,
		model,
		field,
		rejectedValue,
	}
}

// NewInvalidBinderDataValidatorError is used when Echo#Binder can not bind the request data
func NewInvalidBinderDataValidatorError(err error, model string) *ValidatorError {
	message := fmt.Sprintf("Can't bind data for '%v'", model)

	return &ValidatorError{
		ErrorCodeValidatorRequired,
		message,
		model,
		"",
		nil,
	}
}

// NewInvalidFieldValidatorError is used when a ID can't convert to uint
func NewInvalidFieldValidatorError(
	err error,
	model string,
	field string,
	rejectedValue interface{},
) *ValidatorError {
	message := fmt.Sprintf("ID is invalid '%v'", rejectedValue)

	return &ValidatorError{
		ErrorCodeValidatorInvalidField,
		message,
		model,
		field,
		rejectedValue,
	}
}
