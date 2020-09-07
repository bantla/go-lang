// Package errors handles http errors, writes error log
package errors

import "errors"

// New function creates an instance of error
func New(msg string) error {
	return errors.New(msg)
}
