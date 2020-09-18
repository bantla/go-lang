//+build wireinject

package sv1

import (
	"github.com/bantla/internal/app/shopping-cart/domain/repository"
	permissionGormRepo "github.com/bantla/internal/app/shopping-cart/permission/repository/gormrepo"
	rolePermissionGormRepo "github.com/bantla/internal/app/shopping-cart/role-permission/repository/gormrepo"
	roleGormRepo "github.com/bantla/internal/app/shopping-cart/role/repository/gormrepo"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// InitializeRolePermissionService initializes the permission service
func InitializeRolePermissionService(db *gorm.DB) RolePermissionService {
	panic(wire.Build(
		rolePermissionGormRepo.NewRolePermissionRepository,
		wire.Bind(new(repository.RolePermission), new(rolePermissionGormRepo.RolePermissionRepository)),
		roleGormRepo.NewRoleRepository,
		wire.Bind(new(repository.Role), new(roleGormRepo.RoleRepository)),
		permissionGormRepo.NewPermissionRepository,
		wire.Bind(new(repository.Permission), new(permissionGormRepo.PermissionRepository)),
		newPermissionService,
		wire.Bind(new(RolePermissionService), new(rolePermissionService)),
	))
}
