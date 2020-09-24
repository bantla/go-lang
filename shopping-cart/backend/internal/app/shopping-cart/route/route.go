// Package route combines feature routes
package route

import (
	"github.com/bantla/internal/app/shopping-cart/constants"
	permissionRouteV1 "github.com/bantla/internal/app/shopping-cart/permission/delivery/http/route/rv1"
	rolePermissionRouteV1 "github.com/bantla/internal/app/shopping-cart/role-permission/delivery/http/route/rv1"
	"github.com/bantla/internal/app/shopping-cart/role/delivery/http/rolev1"
	roleRouteV1 "github.com/bantla/internal/app/shopping-cart/role/delivery/http/route/rv1"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Register function register all routes
func Register(e *echo.Echo, db *gorm.DB) {
	routeV1 := e.Group(constants.PathPrefix + constants.PathV1)
	rolev1.RegisterRoute(routeV1, db)
	roleRouteV1.RegisterRoute(routeV1.Group(constants.PathRole), db)
	permissionRouteV1.RegisterRoute(routeV1)
	rolePermissionRouteV1.RegisterRoute(routeV1)
}
