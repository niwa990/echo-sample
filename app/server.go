package app

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start() {
	e := echo.New()
	e.Validator = &Validator{validator: validator.New()}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	Route(e)

	e.Logger.Fatal(e.Start(":8080"))
	return
}
