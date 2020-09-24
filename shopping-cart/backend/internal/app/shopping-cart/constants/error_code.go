package constants

const (
	// ErrorCodeInternalServerError is internal server error code
	ErrorCodeInternalServerError = "InternalServerError"

	// ErrorCodeValidatorUniqueField is used when the unique constraint is violated
	ErrorCodeValidatorUniqueField = "Validator/UniqueField"

	// ErrorCodeValidatorNoneReference is used when the foreign key constraint is violated. There is no foreign key value
	ErrorCodeValidatorNoneReference = "Validator/NoneReference"

	// ErrorCodeValidatorInvalidField is used when the fields of data model violdates "validator" tag
	ErrorCodeValidatorInvalidField = "Validator/InvalidField"

	// ErrorCodeValidatorInvalidQueryParameter is used when the query parameter of the request is invalid
	ErrorCodeValidatorInvalidQueryParameter = "Validator/InvalidQueryParameters"

	// ErrorCodeValidatorInvalidPathParameter is used when the path parameter of the request is invalid
	ErrorCodeValidatorInvalidPathParameter = "Validator/InvalidQueryParameters"

	// ErrorCodeValidatorInvalidBinderData is used when Echo#Binder can not bind the request data
	ErrorCodeValidatorInvalidBinderData = "Validator/InvalidBinderData"
)
