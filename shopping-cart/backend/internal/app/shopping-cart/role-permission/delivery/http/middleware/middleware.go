// Package middleware defines functions chained in the HTTP request-response cycle
package middleware

import (
	"fmt"
	"net/http"

	"github.com/bantla/internal/app/shopping-cart/constants"
	"github.com/bantla/internal/app/shopping-cart/domain/model"
	"github.com/bantla/internal/app/shopping-cart/role-permission/service/sv1"
	"github.com/bantla/pkg/errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// BindRolePermission function adds valid role permission to context
func BindRolePermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		rolePermission := &model.RolePermission{}

		if err := ctx.Bind(rolePermission); err != nil {
			return errors.NewHTTPError(
				http.StatusBadRequest,
				nil,
				constants.MessageErrorInvalidRole,
				err,
			)
		}

		ctx.Set(constants.RolePermissionContextKeyValidRolePermission, rolePermission)
		return next(ctx)
	}
}

// WithRolePermissionService function adds an instance of permission service to context
func WithRolePermissionService(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var err error

		if db, ok := ctx.Get(constants.ContextKeyDB).(*gorm.DB); ok {
			rolePermissionService := sv1.InitializeRolePermissionService(db)
			ctx.Set(constants.RolePermissionContextKeyRolePermissionService, rolePermissionService)
			return next(ctx)
		}

		err = errors.New(fmt.Sprintf(
			"`middleware.WithRolePermissionService` %v",
			constants.MessageErrorCheckIfUsingWithDBMiddleware,
		))
		return errors.NewHTTPError(
			http.StatusInternalServerError,
			nil,
			constants.MessageStatusInternalServerError,
			err,
		)
	}
}

// CheckDuplicatedRolePermission function checks if the role permission already exists
func CheckDuplicatedRolePermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// Use context to get role permission that's set by BindRolePermission middleware
		rolePermission, ok := ctx.Get(constants.RolePermissionContextKeyValidRolePermission).(*model.RolePermission)

		if !ok {
			return errors.NewHTTPError(
				http.StatusInternalServerError,
				nil,
				constants.MessageStatusInternalServerError,
				errors.New(fmt.Sprintf(
					"`middleware.CheckDuplicateRolePermission` %v",
					constants.MessageErrorCheckIfUsingBindRolePermissionMiddleware,
				)),
			)
		}

		// Use context to get role permission service that's set by WithRolePermissionService middleware
		rolePermissionService, ok := ctx.Get(constants.RolePermissionContextKeyRolePermissionService).(sv1.RolePermissionService)

		if !ok {
			return errors.NewHTTPError(
				http.StatusInternalServerError,
				nil,
				constants.MessageStatusInternalServerError,
				errors.New(fmt.Sprintf(
					"`middleware.CheckDuplicateRolePermission` %v",
					constants.MessageErrorCheckIfUsingWithRolePermissionServiceMiddleware,
				)),
			)
		}

		// Check if the role permission already exists
		isDuplicated, err := rolePermissionService.DuplicateRolePermission(*rolePermission)

		if err != nil {
			return errors.NewHTTPError(
				http.StatusInternalServerError,
				nil,
				constants.MessageStatusInternalServerError,
				err,
			)
		}

		if isDuplicated {
			return errors.NewHTTPError(
				http.StatusBadRequest,
				nil,
				constants.MessageErrorRolePermissionAlreadyExists,
				errors.New(fmt.Sprintf(
					"`middleware.CheckDuplicateRolePermission` %v",
					constants.MessageErrorRolePermissionAlreadyExists,
				)),
			)
		}

		return next(ctx)
	}
}
