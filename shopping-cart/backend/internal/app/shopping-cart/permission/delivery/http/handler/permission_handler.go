// Package handler package implements handlers of the permission route
package handler

import (
	"net/http"
	"time"

	"github.com/bantla/internal/app/shopping-cart/domain/model"
	"github.com/bantla/internal/app/shopping-cart/role/service"
	"github.com/bantla/pkg/errors"
	"github.com/bantla/pkg/tag"
	"github.com/labstack/echo/v4"
)

// RoleHandler define handlers of role route
type RoleHandler interface {
	// Get method retrives a list of roles
	Get(ctx echo.Context) error

	// Create method create a new role
	Create(ctx echo.Context) error
}

// roleHandler define handlers of role route
type roleHandler struct {
	roleService service.RoleService
}

// NewRoleHandler creates an instance of role handler
func NewRoleHandler(roleService service.RoleService) RoleHandler {
	return roleHandler{roleService}
}

// Get method retrives a list of roles
func (rh roleHandler) Get(ctx echo.Context) error {
	roles := []*model.Role{}

	if err := rh.roleService.FindAll(&roles); err != nil {
		return errors.NewAPIUnexpectedError(err)
	}

	return ctx.JSON(http.StatusOK, roles)
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
// @Success 200 {object} model.Role
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /roles [post]
func (rh roleHandler) Create(ctx echo.Context) error {
	role := &model.Role{}
	roleModel := tag.GetStructTypeName(role)

	// Retrive role data from the request
	if err := ctx.Bind(role); err != nil {
		return errors.NewAPIValidatorError(err, errors.NewInvalidBinderDataValidatorError(err, roleModel))
	}

	// Set createAt time
	role.CreatedAT = time.Now()

	// Validate role data
	if err := ctx.Validate(role); err != nil {
		return err
	}

	// Check for db Error
	var validatorErrors []*errors.ValidatorError

	// Check for duplicate title
	if isDuplicate, err := rh.roleService.IsDuplicateTitle(role.Title); err != nil {
		return errors.NewAPIUnexpectedError(err)
	} else if isDuplicate {
		validatorErrors = append(validatorErrors, errors.NewUniqueValidatorError(
			roleModel,
			tag.GetFieldValueOfJSONTag(role, "Title"),
			role.Title,
		))
	}

	// Check for duplicate slug
	if isDuplicate, err := rh.roleService.IsDuplicateSlug(role.Slug); err != nil {
		return errors.NewAPIUnexpectedError(err)
	} else if isDuplicate {
		validatorErrors = append(validatorErrors, errors.NewUniqueValidatorError(
			roleModel,
			tag.GetFieldValueOfJSONTag(role, "Slug"),
			role.Slug,
		))
	}

	if validatorErrors != nil {
		return errors.NewAPIValidatorError(nil, validatorErrors...)
	}

	// Create a new role
	if err := rh.roleService.Create(role); err != nil {
		return errors.NewAPIUnexpectedError(err)
	}

	return ctx.JSON(http.StatusOK, role)
}
