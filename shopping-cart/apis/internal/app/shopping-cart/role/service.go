package role

import (
	"github.com/bantla/internal/app/shopping-cart/domain/model"
	"github.com/bantla/internal/app/shopping-cart/domain/repository"
)

// Service defines the application bussiness rules of the role model
type Service interface {
	// FindAll method gets all roles in store
	FindAll(roles *[]*model.Role) error
}

type service struct {
	// The roleRepository is role repository
	roleRepository repository.Role
}

// newService function returns an instance of RoleService
func newService(r repository.Role) service {
	return service{roleRepository: r}
}

// FindAll method gets all roles in store
func (s service) FindAll(roles *[]*model.Role) error {
	return s.roleRepository.FindAll(roles)
}
