package libs

import (
	"github.com/khris-xp/assessment-tax/common/dto"
)

func CalculateTax(totalIncome float64, wht float64, allowances []dto.AllowancesType) float64 {
	var totalAllowances float64
	for _, allowance := range allowances {
		totalAllowances += allowance.Amount
	}

	return totalIncome - totalAllowances - wht
}

func CalculateTaxRate(income float64) float64 {
	if income >= 0 && income <= 150000 {
		return income
	}
	if income > 150000 && income <= 500000 {
		return (income - 150000) * 0.1
	}
	if income > 500000 && income <= 1000000 {
		return (income - 500000) * 0.15
	}
	if income > 1000000 && income <= 2000000 {
		return (income - 1000000) * 0.2
	}
	if income > 2000000 {
		return (income - 2000000) * 0.35
	}
	return 0
}
