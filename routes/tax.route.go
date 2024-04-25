package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/khris-xp/assessment-tax/controllers"
)

func TaxRoutes(e *echo.Echo) {
	tax := e.Group("/tax")

	tCl := controllers.TaxController{}
	tax.POST("/calculations", tCl.CalculateTax)
}
