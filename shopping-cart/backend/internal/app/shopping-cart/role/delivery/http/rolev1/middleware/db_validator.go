// Package middleware implements functions chained in the HTTP request-response cycle
package middleware

import (
	"github.com/bantla/internal/app/shopping-cart/role/delivery/http/rolev1/service"
)

// DBValidator define middleware to validate role data in the database
type DBValidator interface {
	// UniqueRoleTitle method returns an error response if title of role is already exists
	// UniqueRoleTitle(next echo.HandlerFunc) echo.HandlerFunc

	// UniqueRoleSlug method returns an error response if slug of role is already exists
	// UniqueRoleSlug(next echo.HandlerFunc) echo.HandlerFunc
}

type dbValidator struct {
	// roleService is a role service
	roleService service.RoleService
}

// NewDBValidator function returns an instance of db validator
func NewDBValidator(roleService service.RoleService) DBValidator {
	return dbValidator{roleService}
}
