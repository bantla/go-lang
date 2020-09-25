// Package route implements route version 1 of the role feature
package route

import (
	"github.com/bantla/internal/app/shopping-cart/role/delivery/http/handler"
	"github.com/bantla/internal/app/shopping-cart/role/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Register function creates the role route
func Register(e *echo.Group, db *gorm.DB) {
	roleService := service.InitializeRoleService(db)
	roleHandler := handler.NewRoleHandler(roleService)

	e.GET("", roleHandler.Get)
	e.POST("", roleHandler.Create)
}
