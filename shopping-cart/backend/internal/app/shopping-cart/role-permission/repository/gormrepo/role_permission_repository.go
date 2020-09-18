// Package gormrepo (gorm repository) implements repositories using gorm (go get -u gorm.io/gorm)
package gormrepo

import (
	"github.com/bantla/internal/app/shopping-cart/domain/model"
	"gorm.io/gorm"
)

// RolePermissionRepository is a specific implementation of the role permission repository constract
type RolePermissionRepository struct {
	// Database gorm
	db *gorm.DB
}

// NewRolePermissionRepository method creates an instance of the role permission repository
func NewRolePermissionRepository(db *gorm.DB) RolePermissionRepository {
	return RolePermissionRepository{db}
}

// Create method inserts a role permission into store
func (r RolePermissionRepository) Create(rolePermission *model.RolePermission) error {
	return r.db.Create(rolePermission).Error
}

// Delete method remove a role permission
func (r RolePermissionRepository) Delete(roleID uint, permissionID uint) error {
	return r.db.Where("role_id = ? and permission_id", roleID, permissionID).Delete(&model.RolePermission{}).Error
}

// FindByRoleID method retrieves list of role permission matching id of role
func (r RolePermissionRepository) FindByRoleID(rolePermission *[]*model.RolePermission, roleID uint) error {
	return r.db.Where("role_id = ?", roleID).Find(rolePermission).Error
}

// FindRolePermissionByIDs method retrieves a role permission using IDs
func (r RolePermissionRepository) FindRolePermissionByIDs(rolePermission *model.RolePermission, roleID uint, permissionID uint) error {
	return r.db.Where("role_id = ? and permission_id = ?", roleID, permissionID).Find(&model.RolePermission{}).Error
}
