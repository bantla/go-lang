//+build wireinject

package role

import (
	"github.com/bantla/internal/app/shopping-cart/domain/repository"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// InitializeService initializes the role service
func InitializeService(db *gorm.DB) Service {
	panic(wire.Build(
		newRepository,
		wire.Bind(new(repository.Role), new(roleRepository)),
		newService,
		wire.Bind(new(Service), new(service)),
	))
}
