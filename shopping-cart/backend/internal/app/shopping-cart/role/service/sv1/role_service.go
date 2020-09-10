// Package sv1 (service version 1) implements use cases of role - API version 1
package sv1

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

	// DuplicateTitle method checks whether or not the role title exists
	DuplicateTitle(title string) (bool, error)

	// DuplicateSlug method checks whether or not the role slug exists
	DuplicateSlug(slug string) (bool, error)
}

type roleService struct {
	// The roleRepository is role repository
	roleRepository repository.Role
}

// newRoleService function returns an instance of RoleService
func newRoleService(r repository.Role) roleService {
	return roleService{roleRepository: r}
}

// FindAll method gets all roles in store
func (s roleService) FindAll(roles *[]*model.Role) error {
	return s.roleRepository.FindAll(roles)
}

// Create method creates a new role
func (s roleService) Create(roles *model.Role) error {
	return s.roleRepository.Create(roles)
}

// FindByTitle method retrieves a role matching a title
func (s roleService) FindByTitle(role *model.Role, title string) error {
	return s.roleRepository.FindByTitle(role, title)
}

// FindByTitle method retrieves a role matching a slug
func (s roleService) FindBySlug(role *model.Role, slug string) error {
	return s.roleRepository.FindBySlug(role, slug)
}

// FindByFields method retrieves a role matching role field values
func (s roleService) FindByFields(role *model.Role, roleMatching model.Role) error {
	return s.roleRepository.FindByFields(role, roleMatching)
}

// DuplicateTitle method checks whether or not the role title exists
func (s roleService) DuplicateTitle(title string) (bool, error) {
	err := s.roleRepository.FindByFields(&model.Role{}, model.Role{
		Title: title,
	})

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return true, err
}

// DuplicateSlug method checks whether or not the role slug exists
func (s roleService) DuplicateSlug(slug string) (bool, error) {
	err := s.roleRepository.FindByFields(&model.Role{}, model.Role{
		Slug: slug,
	})

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return true, err
}
