package main

import (
	"net/http"

	"github.com/khris-xp/assessment-tax/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})
	routes.TaxRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
