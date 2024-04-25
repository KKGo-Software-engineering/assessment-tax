package http_test

import (
	httphdl "github.com/wit-switch/assessment-tax/internal/handler/http"
	"github.com/wit-switch/assessment-tax/pkg/errorx"
	"github.com/wit-switch/assessment-tax/pkg/validator"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validator", func() {
	var (
		validate *httphdl.Validator
	)

	BeforeEach(func() {
		validate = httphdl.NewValidator(validator.New())
	})

	Describe("new validator", func() {
		It("should return validator", func() {
			Expect(validate).NotTo(BeNil())
		})
	})

	Describe("with validate", func() {
		When("with have validate tag", func() {
			Context("return validate fail error", func() {
				BeforeEach(func() {
					validate = httphdl.NewValidator(&mockValidator{
						mockStruct: func(any) error {
							return errorx.New("error")
						},
					})
				})
				It("should return validate fail error", func() {
					type arg struct {
						Text string `json:"text" validate:"required"`
					}
					actual := validate.Validate(arg{})
					Expect(actual).To(MatchError(errorx.ErrValidationFail))
				})
			})

			It("should return validate fail error with params", func() {
				type arg struct {
					Text string `json:"text" validate:"required"`
				}
				val := validate.Validate(arg{})
				var actual *errorx.InternalError[any]
				errorx.As(val, &actual)

				expected := []validator.Field{{
					Field:   "text",
					Message: "text is required",
				}}
				Expect(actual).To(MatchError(errorx.ErrValidationFail))
				Expect(actual.Params()).To(Equal(expected))
			})
		})

		It("should return nil error", func() {
			type arg struct {
				Text string
			}
			actual := validate.Validate(arg{})
			Expect(actual).To(BeNil())
		})
	})
})

type mockValidator struct {
	mockStruct func(any) error
}

func (m *mockValidator) Struct(v any) error {
	return m.mockStruct(v)
}
