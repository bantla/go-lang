package errors

import "log"

// HandleError function prints errors and exits program
// TODO: Break down errors into more types,
// creating error handlers for each (Using type assertion to check error type)
func HandleError(err error) {
	log.Fatal(err)
}
