// Package gormrepo (gorm repository) implements repositories using gorm (go get -u gorm.io/gorm)
package gormrepo

import (
	"github.com/bantla/internal/app/shopping-cart/domain/model"
	"gorm.io/gorm"
)

// RoleRepository is a specific implementation of the role repository constract
type RoleRepository struct {
	// Database gorm
	db *gorm.DB
}

// NewRoleRepository method creates an instance of the role repository
func NewRoleRepository(db *gorm.DB) RoleRepository {
	return RoleRepository{db}
}

// FindAll method returns all roles in store
func (r RoleRepository) FindAll(roles *[]*model.Role) error {
	return r.db.Preload("Permissions").Find(roles).Error
}

// Create method inserts a role into store
func (r RoleRepository) Create(role *model.Role) error {
	return r.db.Create(role).Error
}

// FindByID method retrieves a role matching a id
func (r RoleRepository) FindByID(role *model.Role, id uint) error {
	return r.db.Where("id = ?", id).First(role).Error
}

// FindByTitle method retrieves a role matching a title
func (r RoleRepository) FindByTitle(role *model.Role, title string) error {
	return r.db.Where("title = ?", title).First(role).Error
}

// FindBySlug method retrieves a role matching a slug
func (r RoleRepository) FindBySlug(role *model.Role, slug string) error {
	return r.db.Where("slug = ?", slug).First(role).Error
}

// FindByFields method retrieves a role matching role field values
func (r RoleRepository) FindByFields(role *model.Role, roleMatching model.Role) error {
	return r.db.Where(roleMatching).First(&role).Error
}
