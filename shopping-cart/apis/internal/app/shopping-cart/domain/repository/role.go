// Package repository is in charge of the data access and defined as abstract
package repository

import "github.com/bantla/internal/app/shopping-cart/domain/model"

// Role is repository as data access constract
type Role interface {
	// The FindAll method gets all roles in store
	FindAll(roles *[]*model.Role) error
}
