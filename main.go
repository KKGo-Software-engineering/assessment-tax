package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wit-switch/assessment-tax/config"
)

func main() {
	cfg := config.GetConfig()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})
	e.Logger.Fatal(e.Start(cfg.Server.HTTPAddress()))
}
