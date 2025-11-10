// Package main is the entry point for the sample web application.
// It sets up an Echo web server with database connectivity and routing.
package main

import (
	"os"
	"sample/internal/db"
	"sample/internal/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// main initializes and starts the web server with all necessary middleware and routes.
func main() {
	// Create a new Echo instance
	e := echo.New()

	// Configure middleware
	// Logger middleware logs HTTP requests
	e.Use(middleware.Logger())
	// Recover middleware recovers from panics anywhere in the chain
	e.Use(middleware.Recover())

	// Load database configuration from environment variables
	dbType := os.Getenv("DB_TYPE")
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")

	// Initialize database connection
	conn, err := db.New(dbType, dbConnectionString)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Add database connection to context for all requests
	// This middleware makes the database connection available to all handlers
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", conn)
			return next(c)
		}
	})

	// Register API routes
	// Health check endpoint for monitoring service availability
	e.GET("/api/v1/healthz", routes.Healthz)

	// Serve static files from the UI distribution folder
	// This serves the frontend application
	e.Static("/", "ui/dist")

	// Start the server on the configured port
	err = e.Start(":{{ .BackendPort }}")
	if err != nil {
		e.Logger.Fatal(err)
	}
}
