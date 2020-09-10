// Package constants defines constants as route path, message ...
package constants

const (
	// MessageErrorDatabaseConnectionFailed error message
	MessageErrorDatabaseConnectionFailed = "Database connection failed"

	// MessageErrorAutomaticDatabaseMigrationFailed error message
	MessageErrorAutomaticDatabaseMigrationFailed = "Automatic database migration failed"

	// MessageStatusInternalServerError error message
	MessageStatusInternalServerError = "Internal server error"

	// MessageErrorCheckIfUsingWithDBMiddleware is used when using the echo request context to obtain the db connection
	MessageErrorCheckIfUsingWithDBMiddleware = "Check if using WithDB middleware"
)
