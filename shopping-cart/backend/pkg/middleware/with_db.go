// Package middleware defines functions chained in the HTTP request-response cycle
package middleware

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// WithDB function attaches a database connection (*gorm.DB) to each request context
func WithDB(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Set("DB", db)

			return next(ctx)
		}
	}
}
