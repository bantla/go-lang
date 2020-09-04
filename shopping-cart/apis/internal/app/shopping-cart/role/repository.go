// Package role implements the detailed role of the API
// such as repository, service, route, middleware
package role

import (
	"github.com/bantla/internal/app/shopping-cart/domain/model"
	"gorm.io/gorm"
)

// roleRepository is a specific implementation of the role repository constract
type roleRepository struct {
	// Database gorm
	db *gorm.DB
}

// newRepository method creates an instance of the role repository constract
func newRepository(db *gorm.DB) roleRepository {
	return roleRepository{db}
}

// FindAll method returns all roles in store
func (r roleRepository) FindAll(roles *[]*model.Role) error {
	db := r.db.Find(roles)
	return db.Error
}
