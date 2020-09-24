// Package middleware defines functions chained in the HTTP request-response cycle
package middleware

import (
	"fmt"
	"net/http"

	"github.com/bantla/internal/app/shopping-cart/constants"
	"github.com/bantla/internal/app/shopping-cart/domain/model"
	"github.com/bantla/internal/app/shopping-cart/role/service/sv1"
	"github.com/bantla/pkg/errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// BindRole function adds valid role to context
func BindRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		role := &model.Role{}

		if err := ctx.Bind(role); err != nil {
			return errors.NewHTTPError(
				http.StatusBadRequest,
				nil,
				constants.MessageErrorInvalidRole,
				err,
			)
		}

		ctx.Set(constants.RoleContextKeyValidRole, role)
		return next(ctx)
	}
}

// WithRoleService function adds an instance of role service to context
func WithRoleService(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var err error

		if db, ok := ctx.Get(constants.ContextKeyDB).(*gorm.DB); ok {
			roleService := sv1.InitializeRoleService(db)
			ctx.Set(constants.RoleContextKeyRoleService, roleService)
			return next(ctx)
		}

		err = errors.New(fmt.Sprintf(
			"`middleware.WithRoleService` %v",
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

// CheckDuplicateRoleTitle function checks if title already exists
func CheckDuplicateRoleTitle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var err error

		if role, ok := ctx.Get(constants.RoleContextKeyValidRole).(*model.Role); ok {
			if roleService, ok := ctx.Get(constants.RoleContextKeyRoleService).(sv1.RoleService); ok {
				if isDuplicate, err := roleService.IsDuplicateTitle(role.Title); err != nil {
					return errors.NewHTTPError(
						http.StatusInternalServerError,
						nil,
						constants.MessageStatusInternalServerError,
						err,
					)
				} else if isDuplicate {
					return errors.NewHTTPError(
						http.StatusBadRequest,
						nil,
						constants.MessageErrorRoleTitleAlreadyExists,
						nil,
					)
				}

				return next(ctx)
			}

			err = errors.New(fmt.Sprintf(
				"`middleware.CheckDuplicateRoleSlug` %v",
				constants.MessageErrorCheckIfUsingWithRoleServiceMiddleware,
			))
		} else {
			err = errors.New(fmt.Sprintf(
				"`middleware.CheckDuplicateRoleSlug` %v",
				constants.MessageErrorCheckIfUsingBindRoleMiddleware,
			))
		}

		return errors.NewHTTPError(
			http.StatusInternalServerError,
			nil,
			constants.MessageStatusInternalServerError,
			err,
		)
	}
}

// CheckDuplicateRoleSlug function checks if slug already exists
func CheckDuplicateRoleSlug(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var err error

		if role, ok := ctx.Get(constants.RoleContextKeyValidRole).(*model.Role); ok {
			if roleService, ok := ctx.Get(constants.RoleContextKeyRoleService).(sv1.RoleService); ok {
				if isDuplicate, err := roleService.IsDuplicateSlug(role.Slug); err != nil {
					return &errors.HTTPError{
						Status:  http.StatusInternalServerError,
						Message: constants.MessageStatusInternalServerError,
					}
				} else if isDuplicate {
					return errors.NewHTTPError(
						http.StatusBadRequest,
						nil,
						constants.MessageErrorRoleSlugAlreadyExists,
						nil,
					)
				}

				return next(ctx)
			}

			err = errors.New(fmt.Sprintf(
				"`middleware.CheckDuplicateRoleSlug` %v",
				constants.MessageErrorCheckIfUsingWithRoleServiceMiddleware,
			))
		} else {
			err = errors.New(fmt.Sprintf(
				"`middleware.CheckDuplicateRoleSlug` %v",
				constants.MessageErrorCheckIfUsingBindRoleMiddleware,
			))
		}

		return errors.NewHTTPError(
			http.StatusInternalServerError,
			nil,
			constants.MessageStatusInternalServerError,
			err,
		)
	}
}
