package main

import (
	"context"

	"github.com/bantla/internal/app/shopping-cart/constants"
	"github.com/bantla/internal/app/shopping-cart/route"
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
		errors.HandleError(errors.New(constants.MessageErrorDatabaseConnectionFailed))
	}

	ctx := context.Background()

	if err := migration.AutoMigrate(ctx, db); err != nil {
		errors.HandleError(errors.New(constants.MessageErrorAutomaticDatabaseMigrationFailed))
	}

	// Create echo instance
	e := echo.New()

	// Customize HTTPErrorHandler echo
	errors.HTTPErrorHandler(e)

	// Add global middlewares
	e.Use(middleware.WithDB(db))

	// Register routes
	route.Register(e)

	// Start server
	e.Logger.Fatal(e.Start("127.0.0.1:8080"))
}
