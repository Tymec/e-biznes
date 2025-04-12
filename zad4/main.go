package main

import (
	"zad4/lib"
	"zad4/route"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	lib.InitDB()
	route.InitApiRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
