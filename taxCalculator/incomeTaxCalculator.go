package taxCalculator

import "strings"

type IncomeTaxCalculator struct {
	TotalIncome float64
	Wht         float64
	Allowances  []allowance
}

func (i *IncomeTaxCalculator) addAllowance(a allowance) {
	i.Allowances = append(i.Allowances, a)
}

func (i IncomeTaxCalculator) CalculateTax(personalAllowance float64, adminKrcp float64) float64 {

	netIncome := max(i.TotalIncome-personalAllowance, 0)
	allowanceMap := make(map[string]float64)
	allowanceMap["donation"] = 100000.0
	allowanceMap["k-receipt"] = adminKrcp

	for _, a := range i.Allowances {
		netIncome -= min(a.Amount, allowanceMap[strings.ToLower(a.AllowanceType)])
	}

	out := sum(taxStep1(netIncome), taxStep2(netIncome), taxStep3(netIncome),
		taxStep4(netIncome))

	return out - i.Wht

}

func sum(tax ...float64) float64 {
	sum := 0.0
	for _, v := range tax {
		sum += v
	}
	return sum
}

func taxStep1(netIncome float64) float64 {
	if 150000 < netIncome && netIncome <= 500000 {
		return (netIncome - 150000) * 0.1
	} else if netIncome > 500000 {
		return (500000 - 150000) * 0.1
	}
	return 0
}

func taxStep2(netIncome float64) float64 {
	if 500000 < netIncome && netIncome <= 1000000 {
		return (netIncome - 500000) * 0.15
	} else if netIncome > 1000000 {
		return (1000000 - 500000) * 0.15
	}
	return 0
}

func taxStep3(netIncome float64) float64 {
	if 1000000 < netIncome && netIncome <= 2000000 {
		return (netIncome - 1000000) * 0.2
	} else if netIncome > 1000000 {
		return (2000000 - 1000000) * 0.2
	}
	return 0
}

func taxStep4(netIncome float64) float64 {
	if netIncome > 2000000 {
		return (netIncome - 2000000) * 0.35
	}
	return 0

}

type allowance struct {
	AllowanceType string
	Amount        float64
}
