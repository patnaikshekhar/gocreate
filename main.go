package main

import (
	"os"
	"sample/internal/db"
	"sample/internal/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// main wires up the HTTP server, shared dependencies, and starts listening.
func main() {
	e := echo.New()

	// Log every request and recover from panics to keep the process alive.
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Pull database configuration from the environment.
	dbType := os.Getenv("DB_TYPE")
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")

	conn, err := db.New(dbType, dbConnectionString)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Inject the database connection into each request's context.
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", conn)
			return next(c)
		}
	})

	e.GET("/api/v1/healthz", routes.Healthz)

	// Serve the built frontend assets.
	e.Static("/", "ui/dist")

	// Start the HTTP server on the templated backend port.
	err = e.Start(":{{ .BackendPort }}")
	if err != nil {
		e.Logger.Fatal(err)
	}
}
