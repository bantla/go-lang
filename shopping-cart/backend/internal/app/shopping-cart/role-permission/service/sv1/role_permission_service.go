// Package sv1 (service version 1) implements use cases of role permission - API version 1
package sv1

import (
	"errors"

	"github.com/bantla/internal/app/shopping-cart/domain/model"
	"github.com/bantla/internal/app/shopping-cart/domain/repository"
	"gorm.io/gorm"
)

// RolePermissionService defines the application bussiness permissions of the role permission model
type RolePermissionService interface {
	// The Create method inserts a role permission into store
	Create(rolePermission *model.RolePermission) error

	// The Delete method remove a role permission
	Delete(roleID uint, permissionID uint) error

	// The GetByRoleID method retrieves list of role permission matching id of role
	GetByRoleID(rolePermission *[]*model.RolePermission, roleID uint) error

	// DuplicateRolePermission method checks whether or not the role permission exists
	DuplicateRolePermission(rolePermission model.RolePermission) (bool, error)

	// ExistRoleReference method checks whether or not the role reference exists
	ExistRoleReference(id uint) (bool, error)

	// ExistPermissionReference method checks whether or not the permission reference exists
	ExistPermissionReference(id uint) (bool, error)
}

type rolePermissionService struct {
	// The rolePermissionRepository is role permission repository
	rolePermissionRepository repository.RolePermission

	// The roleRepository is role repository
	roleRepository repository.Role

	// The permissionRepository is permission repository
	permissionRepository repository.Permission
}

// newPermissionService function returns an instance of permissionService
func newPermissionService(rolePermissionRepo repository.RolePermission, roleRepo repository.Role, permissionRepo repository.Permission) rolePermissionService {
	return rolePermissionService{rolePermissionRepo, roleRepo, permissionRepo}
}

// Create method creates a new role permission
func (s rolePermissionService) Create(permissions *model.RolePermission) error {
	return s.rolePermissionRepository.Create(permissions)
}

// Delete method removes a role permission
func (s rolePermissionService) Delete(roleID uint, permissionID uint) error {
	return s.rolePermissionRepository.Delete(roleID, permissionID)
}

// GetByRoleID method retrieves list of role permission matching id of role
func (s rolePermissionService) GetByRoleID(rolePermission *[]*model.RolePermission, roleID uint) error {
	return s.rolePermissionRepository.FindByRoleID(rolePermission, roleID)
}

// DuplicateRolePermission method checks whether or not the role permission exists
func (s rolePermissionService) DuplicateRolePermission(rolePermission model.RolePermission) (bool, error) {
	err := s.rolePermissionRepository.FindRolePermissionByIDs(&model.RolePermission{}, rolePermission.RoleID, rolePermission.PermissionID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return true, err
}

// ExistRoleReference method checks whether or not the role reference exists
func (s rolePermissionService) ExistRoleReference(id uint) (bool, error) {
	role := &model.Role{}
	err := s.roleRepository.FindByID(role, id)

	if err == nil {
		return true, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return false, err
}

// ExistPermissionReference method checks whether or not the permission reference exists
func (s rolePermissionService) ExistPermissionReference(id uint) (bool, error) {
	err := s.permissionRepository.FindByID(&model.Permission{}, id)

	if err == nil {
		return true, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return false, err
}
