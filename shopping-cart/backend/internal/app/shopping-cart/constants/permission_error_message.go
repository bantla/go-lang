package constants

const (
	// MessageErrorCheckIfUsingWithPermissionServiceMiddleware is used when using the echo request context to obtain the permission service
	MessageErrorCheckIfUsingWithPermissionServiceMiddleware = "Check if using WithPermissionService middleware"

	// MessageErrorCheckIfUsingBindPermissionMiddleware is used when using the echo request context to obtain the valid permission data
	MessageErrorCheckIfUsingBindPermissionMiddleware = "Check if using BindPermission middleware"

	// MessageErrorPermissionTitleAlreadyExists is used when checking existsing permission tile
	MessageErrorPermissionTitleAlreadyExists = "Title already exists"

	// MessageErrorPermissionSlugAlreadyExists is used when checking existsing permission slug
	MessageErrorPermissionSlugAlreadyExists = "Slug already exists"

	// MessageErrorInvalidPermission is used when permission data of request is invalid
	MessageErrorInvalidPermission = "Invalid permission"
)
