package main

import (
	"awesomeProject/api/controller"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	controller.Makeroutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
