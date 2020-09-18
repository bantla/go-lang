// Package gormrepo (gorm repository) implements repositories using gorm (go get -u gorm.io/gorm)
package gormrepo

import (
	"github.com/bantla/internal/app/shopping-cart/domain/model"
	"gorm.io/gorm"
)

// PermissionRepository is a specific implementation of the permission repository constract
type PermissionRepository struct {
	// Database gorm
	db *gorm.DB
}

// NewPermissionRepository method creates an instance of the permission repository
func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return PermissionRepository{db}
}

// FindAll method returns all permissions in store
func (r PermissionRepository) FindAll(permissions *[]*model.Permission) error {
	return r.db.Preload("Roles").Find(permissions).Error
}

// Create method inserts a permission into store
func (r PermissionRepository) Create(permission *model.Permission) error {
	return r.db.Create(permission).Error
}

// FindByID method retrieves a permission matching a id
func (r PermissionRepository) FindByID(permission *model.Permission, id uint) error {
	return r.db.Where("id = ?", id).First(permission).Error
}

// FindByTitle method retrieves a permission matching a title
func (r PermissionRepository) FindByTitle(permission *model.Permission, title string) error {
	return r.db.Where("title = ?", title).First(permission).Error
}

// FindBySlug method retrieves a permission matching a slug
func (r PermissionRepository) FindBySlug(permission *model.Permission, slug string) error {
	return r.db.Where("slug = ?", slug).First(permission).Error
}

// FindByFields method retrieves a permission matching permission field values
func (r PermissionRepository) FindByFields(permission *model.Permission, permissionMatching model.Permission) error {
	return r.db.Where(permissionMatching).First(&permission).Error
}
