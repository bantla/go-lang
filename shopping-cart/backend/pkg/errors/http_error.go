package errors

import "github.com/labstack/echo"

// NewHTTPErrorHandler function customizes the default http error handler (echo)
func NewHTTPErrorHandler() echo.HTTPErrorHandler {
	// Do somethings before using DefaultHTTPErrorHandler
	// Examples: Generate error id, save to log file, print error message ...
	// Should seprate each feature into private function
	return echo.New().DefaultHTTPErrorHandler
}
