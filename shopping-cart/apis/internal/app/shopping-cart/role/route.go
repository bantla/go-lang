package role

import (
	"fmt"
	"net/http"

	"github.com/bantla/internal/app/shopping-cart/domain/model"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

const (
	path = "/roles"
)

// RegisterRoute function creates the role route
func RegisterRoute(e *echo.Echo) {
	e.GET(path, welcome)
}

// Welcome
func welcome(ctx echo.Context) error {
	if db, ok := ctx.Get("DB").(*gorm.DB); ok {
		// TODO: FAKE DB
		// role := model.Role{
		// 	Title: "test role",
		// 	Slug: "role slug",
		// 	CreatedAT: time.Now(),
		// }
		// result := db.Create(&role)
		// fmt.Println(result)
		service := InitializeService(db)
		roles := []*model.Role{}
		err := service.FindAll(&roles)

		if err != nil {
			fmt.Println(err)
			ctx.String(http.StatusNotFound, "Has error - DB Query")
		}

		return ctx.JSON(http.StatusOK, roles)
	}

	return ctx.String(http.StatusInternalServerError, "Has error - DB context")
}
