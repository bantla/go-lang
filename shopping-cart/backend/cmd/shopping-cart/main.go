package main

import (
	"context"

	"github.com/bantla/internal/app/shopping-cart/constant"
	roleRV1 "github.com/bantla/internal/app/shopping-cart/role/delivery/http/route/rv1"
	"github.com/bantla/migration"
	"github.com/bantla/pkg/database"
	"github.com/bantla/pkg/errors"
	"github.com/bantla/pkg/middleware"
	"github.com/labstack/echo"
)

func main() {
	// TODO: Data source name should be obtained from the configuration
	dsn := "root:root-admin@(127.0.0.1:3306)/shopping_cart?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := database.ConnectMySQL(dsn)

	if err != nil {
		errors.HandleError(errors.New(constant.DatabaseConnectionFailed))
	}

	ctx := context.Background()

	if err := migration.AutoMigrate(ctx, db); err != nil {
		errors.HandleError(errors.New(constant.AutomaticDatabaseMigrationFailed))
	}

	// Create echo instance
	e := echo.New()

	// Customize echo
	e.HTTPErrorHandler = errors.NewHTTPErrorHandler()

	// Add middlewares
	e.Use(middleware.WithDB(db))

	// Register routes of API version 1
	roleRV1.RegisterRoute(e.Group(constant.PathV1))

	// Start server
	e.Logger.Fatal(e.Start("127.0.0.1:8080"))
}
