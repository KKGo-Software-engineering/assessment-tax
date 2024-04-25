package constants

type TaxDeduction struct {
	Deduction float64 `json:"deduction"`
}

func TaxDeductionInit() *TaxDeduction {
	return &TaxDeduction{
		Deduction: 60000.0,
	}
}
