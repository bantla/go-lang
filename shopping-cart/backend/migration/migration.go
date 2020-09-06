// Package migration migrate schemas, to keep them update to date
package migration

import (
	"context"

	"github.com/bantla/internal/app/shopping-cart/domain/model"
	"gorm.io/gorm"
)

// AutoMigrate function will will create tables corresponding to the models
func AutoMigrate(c context.Context, db *gorm.DB) error {
	err := db.WithContext(c).AutoMigrate(
		&model.Role{},
		&model.Permission{},
		&model.RolePermission{},
	)

	return err
}
