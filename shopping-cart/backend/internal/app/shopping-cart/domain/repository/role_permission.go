package repository

import "github.com/bantla/internal/app/shopping-cart/domain/model"

// RolePermission is repository as data access constract
type RolePermission interface {
	// The Create method inserts a role permission into store
	Create(rolePermission *model.RolePermission) error

	// The Delete method remove a role permission
	Delete(roleID uint, permissionID uint) error

	// The FindByRoleID method retrieves list of role permission matching id of role
	FindByRoleID(roleID uint) error
}
