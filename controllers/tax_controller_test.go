package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
)

func TestCalculateTax(t *testing.T) {
	tests := []struct {
		name        string
		totalIncome float64
		statusCode  int
	}{
		{"StatusBadRequest when Income -10", -10.0, http.StatusBadRequest},
		{"StatusOK when Income 0", 0.0, http.StatusOK},
		{"StatusOK when Income 70,000", 70000.0, http.StatusOK},
		{"StatusOK when Income 150,000", 150000.0, http.StatusOK},
		{"StatusOK when Income 300,000", 300000.0, http.StatusOK},
		{"StatusOK when Income 500,000", 500000.0, http.StatusOK},
		{"StatusOK when Income 750,000", 750000.0, http.StatusOK},
		{"StatusOK when Income 1,000,000", 1000000.0, http.StatusOK},
		{"StatusOK when Income 1,500,000", 1500000.0, http.StatusOK},
		{"StatusOK when Income 2,000,000", 2000000.0, http.StatusOK},
		{"StatusOK when Income 2,500,000", 2500000.0, http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/tax/calculations", strings.NewReader(fmt.Sprintf(`{"totalIncome": %v}`, tt.totalIncome)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			tCl := TaxController{}
			err := tCl.CalculateTax(c)

			if (err != nil && tt.statusCode != http.StatusBadRequest) || rec.Code != tt.statusCode {
				t.Errorf("%s failed. Income: %v, Expected Status Code: %d, Got: %d, Error: %v",
					tt.name, tt.totalIncome, tt.statusCode, rec.Code, err)
				return
			}

			fmt.Printf(color.GreenString("%s passed. Income: %v\n"), tt.name, tt.totalIncome)
		})
	}
}
