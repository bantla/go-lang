// Package rolev1 implements the API version 1 of the role feature
package rolev1

import (
	"github.com/bantla/internal/app/shopping-cart/role/delivery/http/rolev1/handler"
	"github.com/bantla/internal/app/shopping-cart/role/delivery/http/rolev1/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// RegisterRoute function creates the role route
func RegisterRoute(e *echo.Group, db *gorm.DB) {
	roleService := service.InitializeRoleService(db)
	roleHandler := handler.NewRoleHandler(roleService)

	e.GET("/role-handler", roleHandler.Get)
	e.POST("/role-handler", roleHandler.Create)
}
