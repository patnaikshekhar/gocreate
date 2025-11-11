package main

import (
	"os"
	"sample/internal/db"
	"sample/internal/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// main bootstraps the Echo HTTP server and wires up dependencies.
func main() {
	e := echo.New() // Create a fresh Echo instance.

	// Attach middleware for structured logging and panic recovery.
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Read database configuration from the environment so the binary stays 12-factor friendly.
	dbType := os.Getenv("DB_TYPE")
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")

	// Establish a database connection using the resolved configuration.
	conn, err := db.New(dbType, dbConnectionString)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Share the database connection with downstream handlers through the Echo context.
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", conn)
			return next(c)
		}
	})

	// Health check endpoint used by monitors and load balancers.
	e.GET("/api/v1/healthz", routes.Healthz)

	// Serve the compiled frontend assets from the UI build output.
	e.Static("/", "ui/dist")

	// Start the HTTP server on the templated backend port.
	err = e.Start(":{{ .BackendPort }}")
	if err != nil {
		e.Logger.Fatal(err)
	}
}
