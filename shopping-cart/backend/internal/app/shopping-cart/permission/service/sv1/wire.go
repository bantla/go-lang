//+build wireinject

package sv1

import (
	"github.com/bantla/internal/app/shopping-cart/domain/repository"
	"github.com/bantla/internal/app/shopping-cart/permission/repository/gormrepo"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// InitializePermissionService initializes the permission service
func InitializePermissionService(db *gorm.DB) PermissionService {
	panic(wire.Build(
		gormrepo.NewPermissionRepository,
		wire.Bind(new(repository.Permission), new(gormrepo.PermissionRepository)),
		newPermissionService,
		wire.Bind(new(PermissionService), new(permissionService)),
	))
}
