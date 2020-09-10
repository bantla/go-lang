package main

import (
	"context"

	"github.com/bantla/internal/app/shopping-cart/route"
	"github.com/bantla/migration"
	"github.com/bantla/pkg/configuration"
	"github.com/bantla/pkg/database"
	"github.com/bantla/pkg/errors"
	"github.com/bantla/pkg/middleware"
	"github.com/labstack/echo"
)

func main() {
	// TODO: Should check env: dev or prod
	config, err := configuration.New("shopping_cart_dev", "config")

	if err != nil {
		// Debug: path should be ../../config
		if config, err = configuration.New("shopping_cart_dev", "../../config"); err != nil {
			errors.HandleError(err)
			return
		}
	}

	db, err := database.ConnectMySQL(config.Database.GetConnectionURI())

	if err != nil {
		errors.HandleError(err)
		return
	}

	ctx := context.Background()

	if err := migration.AutoMigrate(ctx, db); err != nil {
		errors.HandleError(err)
		return
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
	e.Logger.Fatal(e.Start(config.Server.GetAddress()))
}
