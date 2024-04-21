package taxCalculator

import "errors"

type IncomeTaxCalculator struct {
	TotalIncome float64
	Wht         float64
	Allowances  []allowance
}

func (i *IncomeTaxCalculator) addAllowance(a allowance) {
	i.Allowances = append(i.Allowances, a)
}

func (i IncomeTaxCalculator) CalculateTax(personalAllowance float64) (float64, error) {

	netIncome := max(i.TotalIncome-personalAllowance, 0)

	if 0.0 <= netIncome && netIncome <= 150000 {
		return 0.0, nil
	} else if 150000 < netIncome && netIncome <= 500000 {
		netIncome -= 150000
		return netIncome * 0.1, nil
	} else if 500000 < netIncome && netIncome <= 1000000 {
		netIncome -= 500000
		cumelativeTax := (500000 - 150000) * 0.1
		return (netIncome * 0.15) + cumelativeTax, nil
	}

	// else if netIncome == 500001.0 {
	// 	return 3500.15, nil
	// } else if netIncome == 1000000 {
	// 	return 110000.0, nil
	// } else if netIncome == 500002 {
	// 	return 3500.30, nil
	// } else if netIncome == 940000 {
	// 	return 101000, nil
	// }

	return 0, errors.New("input out of range")

}

type allowance struct {
	AllowanceType string
	Amount        float64
}
