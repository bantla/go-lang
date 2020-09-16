// Package sv1 (service version 1) implements use cases of permission - API version 1
package sv1

import (
	"errors"

	"github.com/bantla/internal/app/shopping-cart/domain/model"
	"github.com/bantla/internal/app/shopping-cart/domain/repository"
	"gorm.io/gorm"
)

// PermissionService defines the application bussiness permissions of the permission model
type PermissionService interface {
	// FindAll method gets all permissions in store
	FindAll(permissions *[]*model.Permission) error

	// Create method creates a new permission
	Create(permissions *model.Permission) error

	// FindByTitle method retrieves a permission matching a title
	FindByTitle(permission *model.Permission, title string) error

	// FindBySlug method retrieves a permission matching a slug
	FindBySlug(permission *model.Permission, slug string) error

	// FindByFields method retrieves a permission matching permission field values
	FindByFields(permission *model.Permission, permissionMatching model.Permission) error

	// DuplicateTitle method checks whether or not the permission title exists
	DuplicateTitle(title string) (bool, error)

	// DuplicateSlug method checks whether or not the permission slug exists
	DuplicateSlug(slug string) (bool, error)
}

type permissionService struct {
	// The permissionRepository is permission repository
	permissionRepository repository.Permission
}

// newPermissionService function returns an instance of permissionService
func newPermissionService(r repository.Permission) permissionService {
	return permissionService{permissionRepository: r}
}

// FindAll method gets all permissions in store
func (s permissionService) FindAll(permissions *[]*model.Permission) error {
	return s.permissionRepository.FindAll(permissions)
}

// Create method creates a new permission
func (s permissionService) Create(permissions *model.Permission) error {
	return s.permissionRepository.Create(permissions)
}

// FindByTitle method retrieves a permission matching a title
func (s permissionService) FindByTitle(permission *model.Permission, title string) error {
	return s.permissionRepository.FindByTitle(permission, title)
}

// FindByTitle method retrieves a permission matching a slug
func (s permissionService) FindBySlug(permission *model.Permission, slug string) error {
	return s.permissionRepository.FindBySlug(permission, slug)
}

// FindByFields method retrieves a permission matching permission field values
func (s permissionService) FindByFields(permission *model.Permission, permissionMatching model.Permission) error {
	return s.permissionRepository.FindByFields(permission, permissionMatching)
}

// DuplicateTitle method checks whether or not the permission title exists
func (s permissionService) DuplicateTitle(title string) (bool, error) {
	err := s.permissionRepository.FindByFields(&model.Permission{}, model.Permission{
		Title: title,
	})

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return true, err
}

// DuplicateSlug method checks whether or not the permission slug exists
func (s permissionService) DuplicateSlug(slug string) (bool, error) {
	err := s.permissionRepository.FindByFields(&model.Permission{}, model.Permission{
		Slug: slug,
	})

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return true, err
}
