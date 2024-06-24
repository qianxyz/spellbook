package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", indexHandler)
	e.GET("/spells", spellListHandler)
	e.GET("/spells/:id", spellDetailHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
