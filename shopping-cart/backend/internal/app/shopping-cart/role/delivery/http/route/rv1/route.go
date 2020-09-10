// Package rv1 (route version 1) defines version 1 of the role API
package rv1

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bantla/internal/app/shopping-cart/constants"
	"github.com/bantla/internal/app/shopping-cart/domain/model"
	"github.com/bantla/internal/app/shopping-cart/role/delivery/http/middleware"
	"github.com/bantla/internal/app/shopping-cart/role/service/sv1"
	"github.com/bantla/pkg/errors"
	"github.com/labstack/echo"
)

// RegisterRoute function creates the role route
func RegisterRoute(e *echo.Group) {
	e.GET("", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "api/v1")
	})

	e.GET(
		constants.PathRole,
		getRoles,
		middleware.WithRoleService,
	)

	e.POST(
		constants.PathRole,
		createRole,
		middleware.BindRole,
		middleware.WithRoleService,
		middleware.CheckDuplicateRoleTitle,
		middleware.CheckDuplicateRoleSlug,
	)
}

// getRoles handler gets list of roles
func getRoles(ctx echo.Context) error {
	var err error

	if roleService, ok := ctx.Get(constants.RoleContextKeyRoleService).(sv1.RoleService); ok {
		roles := []*model.Role{}

		if err = roleService.FindAll(&roles); err != nil {
			return errors.NewHTTPError(
				http.StatusInternalServerError,
				nil,
				constants.MessageStatusInternalServerError,
				err,
			)
		}

		return ctx.JSON(http.StatusOK, roles)
	}

	err = errors.New(fmt.Sprintf(
		"`rv1.createRole handler` %v",
		constants.MessageErrorCheckIfUsingWithRoleServiceMiddleware,
	))
	return errors.NewHTTPError(
		http.StatusInternalServerError,
		nil,
		constants.MessageStatusInternalServerError,
		err,
	)
}

// createRole handler
// @Summary create a new role
// @Description create a new role
// @Tags roles
// @Accept */*
// @Produce json
// @Param Role body domain.Role
// @Context ValidRole WithValidateRole middleware domain.Role
// @Context RoleService WithRoleService middleware sv1.RoleService
// @Success 200 {object} domain.Summaries
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /roles [post]
func createRole(ctx echo.Context) error {
	var err error

	if role, ok := ctx.Get(constants.RoleContextKeyValidRole).(*model.Role); ok {
		role.CreatedAT = time.Now()

		if roleService, ok := ctx.Get(constants.RoleContextKeyRoleService).(sv1.RoleService); ok {
			if err = roleService.Create(role); err != nil {
				return errors.NewHTTPError(
					http.StatusInternalServerError,
					nil,
					constants.MessageStatusInternalServerError,
					err,
				)
			}

			return ctx.JSON(http.StatusOK, role)
		}

		err = errors.New(fmt.Sprintf(
			"`rv1.createRole handler` %v",
			constants.MessageErrorCheckIfUsingWithRoleServiceMiddleware,
		))
		} else {
			err = errors.New(fmt.Sprintf(
				"`rv1.createRole` handler %v",
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
