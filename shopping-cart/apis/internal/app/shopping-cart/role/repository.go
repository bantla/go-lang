// Package role implements the detailed role of the API
// such as repository, service, route, middleware
package role

import (
	"github.com/bantla/internal/app/shopping-cart/domain/model"
	"gorm.io/gorm"
)

// Repository is a specific implementation of the role repository constract
type Repository struct {
	// Database gorm
	db *gorm.DB
}

// NewRepository method creates an instance of the role repository constract
func NewRepository(db *gorm.DB) Repository {
	return Repository{db}
}

// FindAll method returns all roles in store
func (r Repository) FindAll(roles *[]*model.Role) error {
	db := r.db.Find(roles)
	return db.Error
}
