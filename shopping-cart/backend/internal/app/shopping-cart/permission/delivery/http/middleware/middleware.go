// Package middleware defines functions chained in the HTTP request-response cycle
package middleware

import (
	"fmt"
	"net/http"

	"github.com/bantla/internal/app/shopping-cart/constants"
	"github.com/bantla/internal/app/shopping-cart/domain/model"
	"github.com/bantla/internal/app/shopping-cart/permission/service/sv1"
	"github.com/bantla/pkg/errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// BindPermission function adds valid permission to context
func BindPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		permission := &model.Permission{}

		if err := ctx.Bind(permission); err != nil {
			return errors.NewHTTPError(
				http.StatusBadRequest,
				nil,
				constants.MessageErrorInvalidPermission,
				err,
			)
		}

		ctx.Set(constants.PermissionContextKeyValidPermission, permission)
		return next(ctx)
	}
}

// WithPermissionService function adds an instance of permission service to context
func WithPermissionService(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var err error

		if db, ok := ctx.Get(constants.ContextKeyDB).(*gorm.DB); ok {
			permissionService := sv1.InitializePermissionService(db)
			ctx.Set(constants.PermissionContextKeyPermissionService, permissionService)
			return next(ctx)
		}

		err = errors.New(fmt.Sprintf(
			"`middleware.WithPermissionService` %v",
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

// CheckDuplicatePermissionTitle function checks if title already exists
func CheckDuplicatePermissionTitle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var err error

		if permission, ok := ctx.Get(constants.PermissionContextKeyValidPermission).(*model.Permission); ok {
			if PermissionService, ok := ctx.Get(constants.PermissionContextKeyPermissionService).(sv1.PermissionService); ok {
				if isDuplicate, err := PermissionService.DuplicateTitle(permission.Title); err != nil {
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
						constants.MessageErrorPermissionTitleAlreadyExists,
						nil,
					)
				}

				return next(ctx)
			}

			err = errors.New(fmt.Sprintf(
				"`middleware.CheckDuplicatePermissionSlug` %v",
				constants.MessageErrorCheckIfUsingWithPermissionServiceMiddleware,
			))
		} else {
			err = errors.New(fmt.Sprintf(
				"`middleware.CheckDuplicatePermissionSlug` %v",
				constants.MessageErrorCheckIfUsingBindPermissionMiddleware,
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

// CheckDuplicatePermissionSlug function checks if slug already exists
func CheckDuplicatePermissionSlug(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var err error

		if permission, ok := ctx.Get(constants.PermissionContextKeyValidPermission).(*model.Permission); ok {
			if PermissionService, ok := ctx.Get(constants.PermissionContextKeyPermissionService).(sv1.PermissionService); ok {
				if isDuplicate, err := PermissionService.DuplicateSlug(permission.Slug); err != nil {
					return &errors.HTTPError{
						Status:  http.StatusInternalServerError,
						Message: constants.MessageStatusInternalServerError,
					}
				} else if isDuplicate {
					return errors.NewHTTPError(
						http.StatusBadRequest,
						nil,
						constants.MessageErrorPermissionSlugAlreadyExists,
						nil,
					)
				}

				return next(ctx)
			}

			err = errors.New(fmt.Sprintf(
				"`middleware.CheckDuplicatePermissionSlug` %v",
				constants.MessageErrorCheckIfUsingWithPermissionServiceMiddleware,
			))
		} else {
			err = errors.New(fmt.Sprintf(
				"`middleware.CheckDuplicatePermissionSlug` %v",
				constants.MessageErrorCheckIfUsingBindPermissionMiddleware,
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
