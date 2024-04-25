package validator

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type Field struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type FieldError struct {
	err error
}

func (f *FieldError) Error() string {
	return f.err.Error()
}

func (f *FieldError) Field() []Field {
	var vErr validator.ValidationErrors
	if ok := errors.As(f.err, &vErr); ok {
		out := make([]Field, len(vErr))
		for i, v := range vErr {
			out[i] = Field{
				Field:   v.Field(),
				Message: getValidatorMessage(v),
			}
		}

		return out
	}

	return nil
}

func IsFieldErr(err error) *FieldError {
	if err == nil {
		return nil
	}

	var vErr *FieldError
	if ok := errors.As(err, &vErr); ok {
		return vErr
	}

	return nil
}
