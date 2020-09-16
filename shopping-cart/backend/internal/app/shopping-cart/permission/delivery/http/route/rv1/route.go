// Package rv1 (route version 1) defines version 1 of the permission API
package rv1

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bantla/internal/app/shopping-cart/constants"
	"github.com/bantla/internal/app/shopping-cart/domain/model"
	"github.com/bantla/internal/app/shopping-cart/permission/delivery/http/middleware"
	"github.com/bantla/internal/app/shopping-cart/permission/service/sv1"
	"github.com/bantla/pkg/errors"
	"github.com/labstack/echo/v4"
)

// RegisterRoute function creates the permission route
func RegisterRoute(e *echo.Group) {
	e.GET("", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "api/v1")
	})

	e.GET(
		constants.PathPermission,
		getPermissions,
		middleware.WithPermissionService,
	)

	e.POST(
		constants.PathPermission,
		createPermission,
		middleware.BindPermission,
		middleware.WithPermissionService,
		middleware.CheckDuplicatePermissionTitle,
		middleware.CheckDuplicatePermissionSlug,
	)
}

// getPermissions handler gets list of Permissions
func getPermissions(ctx echo.Context) error {
	var err error

	if permissionService, ok := ctx.Get(constants.PermissionContextKeyPermissionService).(sv1.PermissionService); ok {
		permissions := []*model.Permission{}

		if err = permissionService.FindAll(&permissions); err != nil {
			return errors.NewHTTPError(
				http.StatusInternalServerError,
				nil,
				constants.MessageStatusInternalServerError,
				err,
			)
		}

		return ctx.JSON(http.StatusOK, permissions)
	}

	err = errors.New(fmt.Sprintf(
		"`rv1.createPermission handler` %v",
		constants.MessageErrorCheckIfUsingWithPermissionServiceMiddleware,
	))
	return errors.NewHTTPError(
		http.StatusInternalServerError,
		nil,
		constants.MessageStatusInternalServerError,
		err,
	)
}

// createPermission handler
// @Summary create a new Permission
// @Description create a new Permission
// @Tags Permissions
// @Accept */*
// @Produce json
// @Param Permission body domain.Permission
// @Context ValidPermission WithValidatePermission middleware domain.Permission
// @Context PermissionService WithPermissionService middleware sv1.PermissionService
// @Success 200 {object} domain.Summaries
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /Permissions [post]
func createPermission(ctx echo.Context) error {
	var err error

	if permission, ok := ctx.Get(constants.PermissionContextKeyValidPermission).(*model.Permission); ok {
		permission.CreatedAT = time.Now()

		if PermissionService, ok := ctx.Get(constants.PermissionContextKeyPermissionService).(sv1.PermissionService); ok {
			if err = PermissionService.Create(permission); err != nil {
				return errors.NewHTTPError(
					http.StatusInternalServerError,
					nil,
					constants.MessageStatusInternalServerError,
					err,
				)
			}

			return ctx.JSON(http.StatusOK, permission)
		}

		err = errors.New(fmt.Sprintf(
			"`rv1.createPermission handler` %v",
			constants.MessageErrorCheckIfUsingWithPermissionServiceMiddleware,
		))
		} else {
			err = errors.New(fmt.Sprintf(
				"`rv1.createPermission` handler %v",
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
