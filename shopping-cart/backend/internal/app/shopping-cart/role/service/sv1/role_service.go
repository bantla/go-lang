// Package sv1 (service version 1) implements use cases of role - API version 1
package sv1

import (
	"github.com/bantla/internal/app/shopping-cart/domain/model"
	"github.com/bantla/internal/app/shopping-cart/domain/repository"
)

// RoleService defines the application bussiness rules of the role model
type RoleService interface {
	// FindAll method gets all roles in store
	FindAll(roles *[]*model.Role) error
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
