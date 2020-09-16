package errors

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

// HTTPError defines struct of Http Error
type HTTPError struct {
	// Status is the status of error response
	Status int

	// Code is the unique id of error
	Code *int

	// Message descripes error details
	Message string

	// ErrorStackTrace is error stack trace
	ErrorStackTrace error
}

// Error method defines error string
func (he *HTTPError) Error() string {
	return fmt.Sprintf(
		"status: %v, code: %v, message: %v, ErrorStackTrace: %v",
		he.Status,
		he.Code,
		he.Message,
		he.ErrorStackTrace,
	)
}

// NewHTTPError function creates an instance of HTTPError
func NewHTTPError(status int, code *int, message string, err error) *HTTPError {
	return &HTTPError{status, code, message, err}
}

// HTTPErrorResponse defines struct of http error response
type HTTPErrorResponse struct {
	// Code is the unique id of error
	Code *int `json:"code,omitempty"`

	// Message descripes error details
	Message string `json:"message"`

	// TrackingID is tracking id of error
	TrackingID *int `json:"tracking_id,omitempty"`
}

// HTTPErrorHandler customizes the DefaultHTTPErrorHandler (echo)
func HTTPErrorHandler(e *echo.Echo) {
	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		var (
			status = http.StatusInternalServerError
			msg = http.StatusText(status)
			code *int
			trackingID *int
		)

		// TODO: Generate unique tracking id, then store tracking id,
		// error stack, API endpoint, request method in database or log files: log.Output
		HandleError(err)
		log.Println(fmt.Sprintf("Tracking ID: %v", trackingID))

		if he, ok := err.(*HTTPError); ok {
			// Create status code
			if he.Status != 0 {
				status = he.Status
			}

			// Create error code
			code = he.Code

			// Create error message
			if he.Message != "" {
				msg = he.Message
			} else if e.Debug && err != nil {
				msg = err.Error()
			}
		} else if heEcho, ok := err.(*echo.HTTPError); ok {
			status = heEcho.Code
			if heMsg, ok := heEcho.Message.(string); ok {
				msg = heMsg
			} else {
				msg = http.StatusText(status)
			}
		}

		// Send response
		if !ctx.Response().Committed {
			if ctx.Request().Method == http.MethodHead {
				err = ctx.NoContent(status)
			} else {
				err = ctx.JSON(status, HTTPErrorResponse{
					Code: code,
					Message: msg,
					TrackingID: trackingID,
				})
			}

			// Log err if any
			if err != nil {
				e.Logger.Error(err)
			}
		}
	}
}

// https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_dup_entry
func handleMySQLError(err *mysql.MySQLError) {
	switch err.Number {
	case 1062:
		err.Message = "Duplicate field"
	}
}
