package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/hello", func(c echo.Context) error{
		return c.String(http.StatusOK, "Hello, World.\n")
	})

	e.Logger.Fatal(e.Start(":4000"))
}