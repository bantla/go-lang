package constants

const (
	// MessageErrorCheckIfUsingWithRolePermissionServiceMiddleware is used when using the echo request context to obtain the role permission service
	MessageErrorCheckIfUsingWithRolePermissionServiceMiddleware = "Check if using WithRolePermissionService middleware"

	// MessageErrorRolePermissionAlreadyExists is used when checking existsing role tile
	MessageErrorRolePermissionAlreadyExists = "Role permission already exists"

	// MessageErrorInvalidRoleID is used when id role data of query param request is invalid
	MessageErrorInvalidRoleID = "Invalid role id"

	// MessageErrorInvalidRolePermission is used when id role permission data request is invalid
	MessageErrorInvalidRolePermission = "Invalid role permission id"

	// MessageErrorNoneRoleReference is used when role reference data request do not exists
	MessageErrorNoneRoleReference = "Do not exist Role Reference"

	// MessageErrorNonePermissionReference is used when permission reference data request do not exists
	MessageErrorNonePermissionReference = "Do not exist Permission Reference"

	// MessageErrorCheckIfUsingBindRolePermissionMiddleware is used when using the echo request context to obtain the valid role permission data
	MessageErrorCheckIfUsingBindRolePermissionMiddleware = "Check if using BindRolePermission middleware"
)
