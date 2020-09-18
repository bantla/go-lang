// Package route combines feature routes
package route

import (
	"github.com/bantla/internal/app/shopping-cart/constants"
	permissionRouteV1 "github.com/bantla/internal/app/shopping-cart/permission/delivery/http/route/rv1"
	rolePermissionRouteV1 "github.com/bantla/internal/app/shopping-cart/role-permission/delivery/http/route/rv1"
	roleRouteV1 "github.com/bantla/internal/app/shopping-cart/role/delivery/http/route/rv1"
	"github.com/labstack/echo/v4"
)

// Register function register all routes
func Register(e *echo.Echo) {
	routeV1 := e.Group(constants.PathPrefix + constants.PathV1)
	roleRouteV1.RegisterRoute(routeV1)
	permissionRouteV1.RegisterRoute(routeV1)
	rolePermissionRouteV1.RegisterRoute(routeV1)
}
