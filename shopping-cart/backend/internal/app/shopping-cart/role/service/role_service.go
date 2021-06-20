// Package service implements use cases of role feature
package service

import (
	"errors"

	"github.com/bantla/internal/app/shopping-cart/domain/model"
	"github.com/bantla/internal/app/shopping-cart/domain/repository"
	"gorm.io/gorm"
)

// RoleService defines the application bussiness rules of the role model
type RoleService interface {
	// FindAll method gets all roles in store
	FindAll(roles *[]*model.Role) error

	// Create method creates a new role
	Create(roles *model.Role) error

	// FindByTitle method retrieves a role matching a title
	FindByTitle(role *model.Role, title string) error

	// FindBySlug method retrieves a role matching a slug
	FindBySlug(role *model.Role, slug string) error

	// FindByFields method retrieves a role matching role field values
	FindByFields(role *model.Role, roleMatching model.Role) error

	// IsDuplicateTitle method checks whether or not the role title exists
	IsDuplicateTitle(title string) (bool, error)

	// DuplicateSlug method checks whether or not the role slug exists
	IsDuplicateSlug(slug string) (bool, error)

	// Delete method removes a role matching id
	Delete(id uint) (uint, error)
}

type roleService struct {
	// The roleRepository is role repository
	roleRepository repository.Role
}

// newRoleService function returns an instance of RoleService
func newRoleService(roleRepository repository.Role) roleService {
	return roleService{roleRepository}
}

// FindAll method gets all roles in store
func (rs roleService) FindAll(roles *[]*model.Role) error {
	return rs.roleRepository.FindAll(roles)
}

// Create method creates a new role
func (rs roleService) Create(roles *model.Role) error {
	return rs.roleRepository.Create(roles)
}

// FindByTitle method retrieves a role matching a title
func (rs roleService) FindByTitle(role *model.Role, title string) error {
	return rs.roleRepository.FindByTitle(role, title)
}

// FindByTitle method retrieves a role matching a slug
func (rs roleService) FindBySlug(role *model.Role, slug string) error {
	return rs.roleRepository.FindBySlug(role, slug)
}

// FindByFields method retrieves a role matching role field values
func (rs roleService) FindByFields(role *model.Role, roleMatching model.Role) error {
	return rs.roleRepository.FindByFields(role, roleMatching)
}

// IsDuplicateTitle method checks whether or not the role title exists
func (rs roleService) IsDuplicateTitle(title string) (bool, error) {
	err := rs.roleRepository.FindByFields(&model.Role{}, model.Role{
		Title: title,
	})

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return true, err
}

// IsDuplicateSlug method checks whether or not the role slug exists
func (rs roleService) IsDuplicateSlug(slug string) (bool, error) {
	err := rs.roleRepository.FindByFields(&model.Role{}, model.Role{
		Slug: slug,
	})

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return true, err
}

// Delete method removes a role matching a id
func (rs roleService) Delete(id uint) (uint, error) {
	if err := rs.roleRepository.Delete(id); err != nil {
		return 0, err
	}

	return id, nil
}
