package main

import (
	"context"
	"fmt"

	"github.com/bantla/internal/app/shopping-cart/role"
	"github.com/bantla/migration"
	"github.com/bantla/pkg/dbconn"
	"github.com/labstack/echo"
)

func main() {
	// TODO: Data source name should be obtained from the configuration
	dsn := "root:root-admin@(127.0.0.1:3306)/shopping_cart?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := dbconn.ConnectMySQL(dsn)

	if err != nil {
		fmt.Println(err, "So failed")
		return
	}

	ctx := context.Background()

	if err := migration.AutoMigrate(ctx, db); err != nil {
		fmt.Println(err)
		return
	}

	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Set("DB", db)

			return next(ctx)
		}
	})
	role.RegisterRoute(e)
	e.Logger.Fatal(e.Start(":4200"))
}
