package repository

import "github.com/bantla/internal/app/shopping-cart/domain/model"

// Permission is repository as data access constract
type Permission interface {
	// The FindAll method gets all permissions in store
	FindAll(permissions *[]*model.Permission) error

	// The Create method inserts a permission into store
	Create(permission *model.Permission) error

	// The FindByID method retrieves a permission matching a id
	FindByID(permission *model.Permission, id uint) error

	// The FindByTitle method retrieves a permission matching a title
	FindByTitle(permission *model.Permission, title string) error

	// The FindBySlug method retrieves a permission matching a slug
	FindBySlug(permission *model.Permission, slug string) error

	// The FindByFields method retrieves a permission matching permission field values
	FindByFields(permission *model.Permission, permissionMatching model.Permission) error
}
