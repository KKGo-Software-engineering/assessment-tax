package errorx

type ErrCode int

func (e ErrCode) Int() int {
	return int(e)
}

var (
	CodeUnknown        ErrCode
	CodeValidationFail ErrCode = 1

	CodeTaxDeductNotFound ErrCode = 100
)

var (
	ErrValidationFail = NewInternalErr[any](CodeValidationFail)

	ErrTaxDeductNotFound = NewInternalErr[any](CodeTaxDeductNotFound)
)
