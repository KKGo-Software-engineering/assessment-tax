package tax

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

type stubTax struct {
	TaxDetails TaxDetails
	Tax        Tax
	err        error
}

func (s *stubTax) TaxCalculation(td TaxDetails) (Tax, error) {
	return s.Tax, s.err
}

func TestTaxHandler(t *testing.T) {
	t.Run("total income must be greather than 0", func(t *testing.T) {
		mockTaxDetails := TaxDetails{
			TotalIncome: -1.0,
			Wht:         0.0,
			Allowances: []Allowance{
				{
					AllowanceType: "donation",
					Amount:        0.0,
				},
			},
		}

		mockTaxDetailsJSON, _ := json.Marshal(mockTaxDetails)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/tax/calculations", bytes.NewBuffer(mockTaxDetailsJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		st := stubTax{
			TaxDetails: mockTaxDetails,
			err:        errors.New("Total income must be greater than 0"),
		}
		p := New(&st)
		err := p.TaxHandler(c)

		if err != nil {
			t.Errorf("got some error %v", err)
		}

		var gotErr Err
		json.Unmarshal(rec.Body.Bytes(), &gotErr)

		if gotErr.Message != st.err.Error() {
			t.Errorf("expected error message %v but got %v", st.err, err)
		}

		if rec.Code != http.StatusBadRequest {
			t.Errorf("expected status code %v but got %v", http.StatusBadRequest, rec.Code)
		}

	})

	t.Run("should return 500 and an error message if the tax calculation fails", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/tax/calculations", bytes.NewBufferString(`{"totalIncome": 5000}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		st := stubTax{err: errors.New("tax calculation fails")}
		p := New(&st)
		err := p.TaxHandler(c)

		if err != nil {
			t.Errorf("got some error %v", err)
		}

		var gotErr Err
		json.Unmarshal(rec.Body.Bytes(), &gotErr)

		if gotErr.Message != st.err.Error() {
			t.Errorf("expected error message %v but got %v", st.err, err)
		}

		if rec.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %v but got %v", http.StatusInternalServerError, rec.Code)
		}
	})

	t.Run("should return a tax of 29000.0 if the income is 500000.0", func(t *testing.T) {
		mockTaxDetails := TaxDetails{
			TotalIncome: 500000.0,
			Wht:         0.0,
			Allowances: []Allowance{
				{
					AllowanceType: "donation",
					Amount:        0.0,
				},
			},
		}

		e := echo.New()
		mockTaxDetailsJSON, _ := json.Marshal(mockTaxDetails)
		req := httptest.NewRequest(http.MethodPost, "/tax/calculations", bytes.NewBuffer(mockTaxDetailsJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		want := Tax{Tax: "29000.0"}

		st := stubTax{
			TaxDetails: mockTaxDetails,
			Tax:        want,
		}

		p := New(&st)
		err := p.TaxHandler(c)

		if err != nil {
			t.Errorf("got some error %v", err)
		}

		var got Tax
		json.Unmarshal(rec.Body.Bytes(), &got)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
