package tax

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Err struct {
	Message string `json:"message"`
}

type Storer interface {
	TaxCalculation(TaxDetails) (Tax, error)
}

type Handler struct {
	store Storer
}

func New(store Storer) *Handler {
	return &Handler{store: store}
}

func (h *Handler) TaxHandler(c echo.Context) error {
	td := TaxDetails{}

	if err := c.Bind(&td); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	if td.TotalIncome <= 0 {
		return c.JSON(http.StatusBadRequest, Err{Message: "Total income must be greater than 0"})
	}

	t, err := h.store.TaxCalculation(td)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})

	}
	return c.JSON(http.StatusOK, t)
}
