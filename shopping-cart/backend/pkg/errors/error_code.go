package errors

const (
	// ErrorCodeUnexpectedError is the internal server error - status code 500
	ErrorCodeUnexpectedError = "UnexpectedError"

	// ErrorCodeNotFound is the not found - status code 404
	ErrorCodeNotFound = "NotFound"

	//ErrorCodeFailedValidator is used when bad request data
	ErrorCodeFailedValidator = "FailedValidator"

	// ErrorCodeValidatorRequired is used when the value of field (field of model, query parameter, path parameter) is required
	// sub error of ErrorCodeFailedValidator error
	ErrorCodeValidatorRequired = "FailedValidator/Required"

	// ErrorCodeValidatorUnique is used when the unique constraint is violated
	// sub error of ErrorCodeFailedValidator error
	ErrorCodeValidatorUnique = "FailedValidator/Unique"

	// ErrorCodeValidatorNotFoundReference is used when the foreign key constraint is violated. There is no foreign key value
	// sub error of ErrorCodeFailedValidator error
	ErrorCodeValidatorNotFoundReference = "FailedValidator/NotFoundReference"

	// ErrorCodeValidatorInvalidField is used when the fields of data model violdates "validator" tag,
	// should be more detailed about the validator type: min, max, equal, required, isBool, isInt ...
	// sub error of ErrorCodeFailedValidator error
	ErrorCodeValidatorInvalidField = "FailedValidator/InvalidField"

	// ErrorCodeValidatorInvalidBinderData is used when Echo#Binder can not bind the request data
	// sub error of ErrorCodeFailedValidator error
	ErrorCodeValidatorInvalidBinderData = "FailedValidator/InvalidBinderData"
)
