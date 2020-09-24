//+build wireinject

package service

import (
	"github.com/bantla/internal/app/shopping-cart/domain/repository"
	"github.com/bantla/internal/app/shopping-cart/role/repository/gormrepo"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// InitializeRoleService initializes the role service
func InitializeRoleService(db *gorm.DB) RoleService {
	panic(wire.Build(
		gormrepo.NewRoleRepository,
		wire.Bind(new(repository.Role), new(gormrepo.RoleRepository)),
		newRoleService,
		wire.Bind(new(RoleService), new(roleService)),
	))
}
