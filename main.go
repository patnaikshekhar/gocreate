// Package main is the entry point for the sample web application.
// It sets up an Echo web server with database connectivity, middleware,
// API routes, and static file serving.
package main

import (
	"os"
	"sample/internal/db"
	"sample/internal/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// main initializes and starts the web server with all necessary configurations.
// It performs the following steps:
// 1. Creates a new Echo instance
// 2. Configures logging and recovery middleware
// 3. Establishes database connection based on environment variables
// 4. Sets up database injection middleware
// 5. Registers API routes and static file serving
// 6. Starts the HTTP server
func main() {
	// Create a new Echo web framework instance
	e := echo.New()

	// Configure middleware for logging HTTP requests
	e.Use(middleware.Logger())
	// Configure middleware for recovering from panics
	e.Use(middleware.Recover())

	// Read database configuration from environment variables
	dbType := os.Getenv("DB_TYPE")
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")

	// Initialize database connection with the specified type and connection string
	conn, err := db.New(dbType, dbConnectionString)
	if err != nil {
		// Fatal error if database connection fails - application cannot proceed
		e.Logger.Fatal(err)
	}

	// Register custom middleware to inject database connection into request context
	// This makes the database connection available to all route handlers via context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", conn)
			return next(c)
		}
	})

	// Register API routes
	// Health check endpoint for monitoring application status
	e.GET("/api/v1/healthz", routes.Healthz)

	// Serve static files from the ui/dist directory at the root path
	// This typically serves the frontend application
	e.Static("/", "ui/dist")

	// Start the HTTP server on the configured port
	// Note: The port uses a template variable that should be replaced during deployment
	err = e.Start(":{{ .BackendPort }}")
	if err != nil {
		// Fatal error if server fails to start
		e.Logger.Fatal(err)
	}
}
