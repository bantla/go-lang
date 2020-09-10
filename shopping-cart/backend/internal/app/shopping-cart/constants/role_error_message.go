package constants

const (
	// MessageErrorCheckIfUsingWithRoleServiceMiddleware is used when using the echo request context to obtain the service role
	MessageErrorCheckIfUsingWithRoleServiceMiddleware = "Check if using WithRoleService middleware"

	// MessageErrorCheckIfUsingBindRoleMiddleware is used when using the echo request context to obtain the valid role data
	MessageErrorCheckIfUsingBindRoleMiddleware = "Check if using BindRole middleware"

	// MessageErrorRoleTitleAlreadyExists is used when checking existsing role tile
	MessageErrorRoleTitleAlreadyExists = "Title already exists"

	// MessageErrorRoleSlugAlreadyExists is used when checking existsing role slug
	MessageErrorRoleSlugAlreadyExists = "Slug already exists"

	// MessageErrorInvalidRole is used when role data of request is invalid
	MessageErrorInvalidRole = "Invalid role"
)
