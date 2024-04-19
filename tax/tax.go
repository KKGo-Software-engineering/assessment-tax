package tax

type Allowance struct {
	AllowanceType string
	Amount        float64
}

type TaxDetails struct {
	TotalIncome float64
	Wht         float64
	Allowances  []Allowance
}

type Tax struct {
	Tax string `json:"tax"`
}
