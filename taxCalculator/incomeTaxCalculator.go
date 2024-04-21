package taxCalculator

type IncomeTaxCalculator struct {
	TotalIncome float64
	Wht         float64
	Allowances  []allowance
}

func (i *IncomeTaxCalculator) addAllowance(a allowance) {
	i.Allowances = append(i.Allowances, a)
}

func (i IncomeTaxCalculator) CalculateTax(personalAllowance float64) float64 {

	netIncome := max(i.TotalIncome-personalAllowance, 0)

	if 0.0 <= netIncome && netIncome <= 150000 {
		return 0.0
	} else if 150000 < netIncome && netIncome <= 500000 {
		netIncome -= 150000
		return netIncome * 0.1
	}
	return 0.0
}

type allowance struct {
	AllowanceType string
	Amount        float64
}
