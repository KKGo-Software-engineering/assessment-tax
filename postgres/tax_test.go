package postgres

import (
	"fmt"
	"testing"
)

func TestTaxCalculation(t *testing.T) {

	tests := []struct {
		income float64
		tax    float64
	}{
		{0.0, 0.0},
		{150000.0, 0.0},
		{150001.0, 0.0},
		{500000.0, 35000.0},
		{500001.0, 35000.0},
		{1000000.0, 110000.0},
		{1000001.0, 110000.0},
		{2000000.0, 310000.0},
		{3000000.0, 660000.0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("final income %v should return %v", tt.income, tt.tax), func(t *testing.T) {
			want := tt.tax

			got := calculateTax(tt.income)

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}

}

func TestAllowancesCalculation(t *testing.T) {
	t.Run("should return 60000", func(t *testing.T) {
		want := 60000.0

		got := calculateAllowances()

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestNetIncomeCalculation(t *testing.T) {
	t.Run("should return 60000", func(t *testing.T) {
		want := 60000.0

		got := calculateNetIncome(120000.0)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
