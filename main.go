package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "ui/dist")

	err := e.Start(":1323")
	if err != nil {
		e.Logger.Fatal(err)
	}
}
