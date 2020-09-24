package rv1

import (
	"net/http"

	"github.com/bantla/internal/app/shopping-cart/constants"
	"github.com/bantla/internal/app/shopping-cart/domain/model"
	"github.com/bantla/internal/app/shopping-cart/role/service/sv1"
	"github.com/bantla/pkg/errors"
	"github.com/labstack/echo/v4"
)

// roleHandler define handlers of role route
type roleHandler struct {
	roleService sv1.RoleService
}

// newRoleHandler creates an instance of role handler
func newRoleHandler(roleService sv1.RoleService) roleHandler {
	return roleHandler{roleService}
}

// getRoles method gets a list of roles
func (h roleHandler) getRoles(ctx echo.Context) error {
	roles := []*model.Role{}

	if err := h.roleService.FindAll(&roles); err != nil {
		return errors.NewHTTPError(
			http.StatusInternalServerError,
			nil,
			constants.MessageStatusInternalServerError,
			err,
		)
	}

	return ctx.JSON(http.StatusOK, roles)
}
