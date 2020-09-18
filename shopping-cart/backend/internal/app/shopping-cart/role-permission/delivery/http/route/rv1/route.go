// Package rv1 (route version 1) defines version 1 of the permission API
package rv1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bantla/internal/app/shopping-cart/constants"
	"github.com/bantla/internal/app/shopping-cart/domain/model"
	"github.com/bantla/internal/app/shopping-cart/role-permission/delivery/http/middleware"
	"github.com/bantla/internal/app/shopping-cart/role-permission/service/sv1"
	"github.com/bantla/pkg/errors"
	"github.com/labstack/echo/v4"
)

// RegisterRoute function creates the role permission route
func RegisterRoute(e *echo.Group) {
	e.GET(
		constants.PathRolePermission,
		getRolePermissionByRoleID,
		middleware.WithRolePermissionService,
	)

	e.POST(
		constants.PathRolePermission,
		createRolePermission,
		middleware.BindRolePermission,
		middleware.WithRolePermissionService,
	)
}

// getRolePermissionByRoleID handler retrives a list role permission matching a role id
func getRolePermissionByRoleID(ctx echo.Context) error {
	roleID, err := strconv.ParseUint(ctx.QueryParam("role_id"), 0, 64)

	if err != nil {
		return errors.NewHTTPError(
			http.StatusBadRequest,
			nil,
			constants.MessageErrorInvalidRoleID,
			err,
		)
	}

	rolePermissionService, ok := ctx.Get(constants.RolePermissionContextKeyRolePermissionService).(sv1.RolePermissionService)

	if !ok {
		return errors.NewHTTPError(
			http.StatusInternalServerError,
			nil,
			constants.MessageStatusInternalServerError,
			errors.New(fmt.Sprintf(
				"`rv1.getRolePermissionByRoleID handler` %v",
				constants.MessageErrorCheckIfUsingWithRolePermissionServiceMiddleware,
			)),
		)
	}

	rolePermissions := &[]*model.RolePermission{}

	if err = rolePermissionService.GetByRoleID(rolePermissions, uint(roleID)); err != nil {
		return errors.NewHTTPError(
			http.StatusInternalServerError,
			nil,
			constants.MessageStatusInternalServerError,
			err,
		)
	}

	return ctx.JSON(http.StatusOK, rolePermissions)
}

// createRolePermission handler creates new a role permission
func createRolePermission(ctx echo.Context) error {
	// Use context to get role permission that is set by BindRolePermission middleware
	rolePermission, ok := ctx.Get(constants.RolePermissionContextKeyValidRolePermission).(*model.RolePermission)

	if !ok {
		err := errors.New(fmt.Sprintf(
			"`rv1.createRolePermission handler` %v",
			constants.MessageErrorCheckIfUsingBindRolePermissionMiddleware,
		))

		return errors.NewHTTPError(
			http.StatusInternalServerError,
			nil,
			constants.MessageStatusInternalServerError,
			err,
		)
	}

	// Use context to get role permission service that is set by WithRolePermissionService middleware
	rolePermissionService, ok := ctx.Get(constants.RolePermissionContextKeyRolePermissionService).(sv1.RolePermissionService)

	if !ok {
		err := errors.New(fmt.Sprintf(
			"`rv1.createRolePermission handler` %v",
			constants.MessageErrorCheckIfUsingWithRolePermissionServiceMiddleware,
		))

		return errors.NewHTTPError(
			http.StatusInternalServerError,
			nil,
			constants.MessageStatusInternalServerError,
			err,
		)
	}

	// Check if role reference exists
	if hasRoleReference, err := rolePermissionService.ExistRoleReference(rolePermission.RoleID); err != nil {
		return errors.NewHTTPError(
			http.StatusInternalServerError,
			nil,
			constants.MessageStatusInternalServerError,
			err,
		)
	} else if !hasRoleReference {
		return errors.NewHTTPError(
			http.StatusBadRequest,
			nil,
			constants.MessageErrorNoneRoleReference,
			err,
		)
	}

	// Check if permission reference exists
	if hasPermissionReference, err := rolePermissionService.ExistPermissionReference(rolePermission.PermissionID); err != nil {
		return errors.NewHTTPError(
			http.StatusInternalServerError,
			nil,
			constants.MessageStatusInternalServerError,
			err,
		)
	} else if !hasPermissionReference {
		return errors.NewHTTPError(
			http.StatusBadRequest,
			nil,
			constants.MessageErrorNonePermissionReference,
			err,
		)
	}

	// Create new role reference
	if err := rolePermissionService.Create(rolePermission); err != nil {
		return errors.NewHTTPError(
			http.StatusInternalServerError,
			nil,
			constants.MessageStatusInternalServerError,
			err,
		)
	}

	return ctx.JSON(http.StatusOK, rolePermission)
}
