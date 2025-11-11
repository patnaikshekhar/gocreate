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

// main initializes and starts the web application server.
// It configures middleware, establishes database connection, sets up routes,
// and serves both API endpoints and static frontend files.
func main() {
	// Create a new Echo instance
	e := echo.New()

	// Register middleware for logging HTTP requests
	e.Use(middleware.Logger())
	// Register middleware for recovering from panics
	e.Use(middleware.Recover())

	// Read database configuration from environment variables
	dbType := os.Getenv("DB_TYPE")
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")

	// Initialize database connection with the configured type and connection string
	conn, err := db.New(dbType, dbConnectionString)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Register custom middleware to inject database connection into request context
	// This makes the database connection available to all route handlers
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", conn)
			return next(c)
		}
	})

	// Register API routes
	e.GET("/api/v1/healthz", routes.Healthz)

	// Serve static files from the ui/dist directory at the root path
	e.Static("/", "ui/dist")

	// Start the server on the configured port
	// Note: Port is templated and should be replaced during build/deployment
	err = e.Start(":{{ .BackendPort }}")
	if err != nil {
		e.Logger.Fatal(err)
	}
}
