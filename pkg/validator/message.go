package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func getValidatorMessage(v validator.FieldError) string {
	switch v.Tag() {
	case "gt", "gtfield":
		return fmt.Sprintf("%s should greater than %s", v.Field(), v.Param())
	case "gte", "gtefield":
		return fmt.Sprintf("%s should greater than or equal %s", v.Field(), v.Param())
	default:
		return fmt.Sprintf("%s is %s", v.Field(), v.Tag())
	}
}
