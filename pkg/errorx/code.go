package errorx

type ErrCode int

func (e ErrCode) Int() int {
	return int(e)
}

var (
	CodeUnknown        ErrCode
	CodeValidationFail ErrCode = 1
)

var (
	ErrValidationFail = NewInternalErr[any](CodeValidationFail)
)
