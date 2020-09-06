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
	db := r.db.Find(roles)
	return db.Error
}
