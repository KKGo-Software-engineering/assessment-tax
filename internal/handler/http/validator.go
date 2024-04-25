package http

import (
	"github.com/wit-switch/assessment-tax/pkg/errorx"
	"github.com/wit-switch/assessment-tax/pkg/validator"
)

type Validator struct {
	validate validator.Validator
}

func NewValidator(validate validator.Validator) *Validator {
	return &Validator{
		validate: validate,
	}
}

func (v *Validator) Validate(req any) error {
	err := v.validate.Struct(req)
	if err := validator.IsFieldErr(err); err != nil {
		return errorx.ErrValidationFail.WithError(err).WithParams(err.Field())
	}

	if err != nil {
		return errorx.ErrValidationFail.WithError(err)
	}

	return nil
}
