package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator interface {
	Struct(s any) error
}

type wrapValidator struct {
	*validator.Validate
}

func New() Validator {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})

	return &wrapValidator{v}
}

func (w *wrapValidator) Struct(s any) error {
	if err := w.Validate.Struct(s); err != nil {
		return &FieldError{err: err}
	}

	return nil
}
