// Package validator implements custom validators
package validator

import (
	"fmt"

	"github.com/bantla/pkg/errors"
	"github.com/bantla/pkg/tag"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// Validator is custom validator struct for Echo
type Validator struct {
	validator *validator.Validate
}

// Validate method tells Echo#Validator how to validate a struct
func (cv *Validator) Validate(i interface{}) error {
	err := cv.validator.Struct(i)

	switch e := err.(type) {
	case *validator.InvalidValidationError:
		return errors.NewAPIValidatorError(err, nil)
	case validator.ValidationErrors:
		var validatorErrors []*errors.ValidatorError
		for _, v := range e {
			validatorTag := v.Tag()
			field := tag.GetFieldValueOfJSONTag(i, v.Field())
			validatorError := &errors.ValidatorError{
				Code: getErrorCodeByValidatorTag(validatorTag),
				// TODO: Should use translation to custom validator message
				// Then: `Message: e.Error()`
				Message: fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", field, validatorTag),
				Model: tag.GetStructTypeName(i),
				Field: field,
				RejectedValue: v.Value(),
			}

			validatorErrors = append(validatorErrors, validatorError)
		}

		return errors.NewAPIValidatorError(err, validatorErrors...)
	}

	return err
}

// SetValidator set custom validator for Echo
func SetValidator(e *echo.Echo) {
	// TODO: Add custom validator message e.Validator.validator.Register ...
	e.Validator = &Validator{
		validator.New(),
	}
}

func getErrorCodeByValidatorTag(tag string) (code string) {
	switch tag {
	case "required":
		code = errors.ErrorCodeValidatorRequired
	}

	return
}
