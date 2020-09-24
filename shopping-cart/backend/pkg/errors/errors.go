// Package errors handles http errors, writes error log
package errors

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
)

// New function creates an instance of error
func New(msg string) error {
	return fmt.Errorf(msg)
}

// HandleError function prints errors and exits program
// TODO: Break down errors into more types,
// creating error handlers for each (Using type assertion to check error type)
func HandleError(err error) {
	if apiErr, ok := err.(*APIError); ok {
		log.Error(apiErr.Unwrap())
	}

	log.Error(err)
}

// GetMySQLErrorMessage ...
func GetMySQLErrorMessage(err error) string {
	var msg string

	if v, ok := err.((*mysql.MySQLError)); ok {
		switch v.Number {
		case 1062:
			msg = "Duplicate field"
		}
	}

	return msg
}
