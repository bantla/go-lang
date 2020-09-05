//+build wireinject

package v1

import (
	"github.com/bantla/internal/app/shopping-cart/domain/repository"
	"github.com/bantla/internal/app/shopping-cart/role"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// InitializeRoleService initializes the role service
func InitializeRoleService(db *gorm.DB) RoleService {
	panic(wire.Build(
		role.NewRepository,
		wire.Bind(new(repository.Role), new(role.Repository)),
		newRoleService,
		wire.Bind(new(RoleService), new(roleService)),
	))
}
