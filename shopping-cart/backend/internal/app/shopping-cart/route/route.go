// Package route combines feature routes
package route

import (
	"github.com/bantla/internal/app/shopping-cart/constants"
	"github.com/bantla/internal/app/shopping-cart/role/delivery/http/route/rv1"
	"github.com/labstack/echo/v4"
)

// Register function register all routes
func Register(e *echo.Echo) {
	rv1.RegisterRoute(e.Group(constants.PathPrefix + constants.PathV1))
}
