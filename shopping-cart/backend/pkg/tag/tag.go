// Package tag define common functions that handle tag
package tag

import (
	"reflect"
	"strings"
)

// GetValueOfTag function returns tag value string pointer, returns string zero value if it's not found
func GetValueOfTag(tag string, concreteValue interface{}, field string) string {
	// Get a new Value initialized to the concrete value
	v := reflect.ValueOf(concreteValue)

	// Get value of pointer if the concreteValue is a pointer
	if v.Kind() == reflect.Ptr {
		v = reflect.Indirect(v)
	}

	if v.Kind() == reflect.Struct {
		// Get struct field
		if f, ok := v.Type().FieldByName(field); ok {
			// Return value of tag name
			return f.Tag.Get(tag)
		}
	}

	return ""
}

// GetFieldValueOfJSONTag function returns json tag value string pointer, returns string zero value if
// the JSON field must be the first value in the value tag
func GetFieldValueOfJSONTag(concreteValue interface{}, field string) string {
	// Ignore tag "options" like omitempty, etc.
	jsonField := strings.Split(GetValueOfTag("json", concreteValue, field), ",")[0]
	return jsonField
}

// GetStructTypeName function returns struct type name of a concrete value (None pointer)
func GetStructTypeName(concreteValue interface{}) string {
	// Get struct type of a concrete value
	v := reflect.ValueOf(concreteValue)

	// Get value of pointer if the concreteValue is a pointer
	if v.Kind() == reflect.Ptr {
		v = reflect.Indirect(v)
	}

	return v.Type().Name()
}
