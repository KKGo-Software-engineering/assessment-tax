package dto

type TaxRequest struct {
	TotalIncome float64          `json:"totalIncome"`
	Wht         float64          `json:"wht"`
	Allowances  []AllowancesType `json:"allowances"`
}

type AllowancesType struct {
	AllowancesType string  `json:"allowancesType"`
	Amount         float64 `json:"amount"`
}
