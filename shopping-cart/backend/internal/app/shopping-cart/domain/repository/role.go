// Package repository is in charge of the data access and defined as abstract
package repository

import "github.com/bantla/internal/app/shopping-cart/domain/model"

// Role is repository as data access constract
type Role interface {
	// The FindAll method gets all roles in store
	FindAll(roles *[]*model.Role) error

	// The Create method inserts a role into store
	Create(role *model.Role) error

	// The FindByID method retrieves a role matching a id
	FindByID(role *model.Role, id uint) error

	// The FindByTitle method retrieves a role matching a title
	FindByTitle(role *model.Role, title string) error

	// The FindBySlug method retrieves a role matching a slug
	FindBySlug(role *model.Role, slug string) error

	// The FindByFields method retrieves a role matching role field values
	FindByFields(role *model.Role, roleMatching model.Role) error

	// The delete method removes a role matching a id
	Delete(id uint) error
}
