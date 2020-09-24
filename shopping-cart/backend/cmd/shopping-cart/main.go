package main

import (
	"context"
	"os"
	"strconv"

	"github.com/bantla/internal/app/shopping-cart/route"
	"github.com/bantla/migration"
	"github.com/bantla/pkg/configuration"
	"github.com/bantla/pkg/database"
	"github.com/bantla/pkg/errors"
	"github.com/bantla/pkg/middleware"
	"github.com/bantla/pkg/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	var configPath string
	var configName string

	switch os.Getenv("ENV") {
	case "DEV":
	case "PROD":
	default:
		configPath = "./config"
		configName = "shopping_cart_dev"
	}

	// Viper has issues configPath with vscode debugging. So we need to override configPath
	if isDebugging, _ := strconv.ParseBool(os.Getenv("DEBUG")); isDebugging {
		configPath = "../../config"
	}

	config, err := configuration.New(configName, configPath)

	if err != nil {
		errors.HandleError(err)
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
	e.HideBanner = true

	// Customize HTTPErrorHandler Echo
	errors.HTTPErrorHandler(e)

	// Customize validator Echo
	validator.SetValidator(e)

	// Add global middlewares
	e.Use(middleware.WithDB(db))

	// Register routes
	route.Register(e, db)

	// Start server
	e.Logger.Fatal(e.Start(config.Server.GetAddress()))
}
