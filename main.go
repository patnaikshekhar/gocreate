// Package main is the entry point for the sample application.
// It sets up an Echo web server with database connectivity and routes.
package main

import (
	"os"
	"sample/internal/db"
	"sample/internal/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// main initializes and starts the web server with all required middleware and routes.
func main() {
	// Create a new Echo instance
	e := echo.New()

	// Add middleware for logging HTTP requests
	e.Use(middleware.Logger())
	// Add middleware to recover from panics and return a 500 error
	e.Use(middleware.Recover())

	// Get database configuration from environment variables
	dbType := os.Getenv("DB_TYPE")
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")

	// Initialize database connection
	conn, err := db.New(dbType, dbConnectionString)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Add custom middleware to inject database connection into request context
	// This makes the database connection available to all route handlers
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", conn)
			return next(c)
		}
	})

	// Register API routes
	e.GET("/api/v1/healthz", routes.Healthz)

	// Serve static files from the ui/dist directory
	e.Static("/", "ui/dist")

	// Start the server on the specified port
	err = e.Start(":{{ .BackendPort }}")
	if err != nil {
		e.Logger.Fatal(err)
	}
}
