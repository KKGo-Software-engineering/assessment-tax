package http_test

import (
	"net/http"
	"net/http/httptest"

	httphdl "github.com/wit-switch/assessment-tax/internal/handler/http"
	"github.com/wit-switch/assessment-tax/pkg/errorx"

	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Error", func() {
	type fieldMessage struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}

	Describe("as error response", func() {
		When("with error is nil", func() {
			It("should return response error with default value", func() {
				var arg error
				actual := httphdl.AsErrorResponse[any](arg)

				expected := &httphdl.ResponseError[any]{
					BaseResponse: httphdl.BaseResponse{
						Code:    httphdl.ResponseCode(http.StatusInternalServerError, errorx.CodeUnknown.Int()),
						Message: "unknown",
					},
				}

				Expect(actual).NotTo(BeNil())
				Expect(actual).To(Equal(expected))
			})
		})

		When("with error is internal", func() {
			Context("existing in map response list", func() {
				When("code validation fail", func() {
					params := []fieldMessage{
						{
							Field:   "totalIncome",
							Message: "required",
						},
						{
							Field:   "allowances",
							Message: "required",
						},
					}
					It("should return error existing code table", func() {
						actual := httphdl.AsErrorResponse[any](
							errorx.ErrValidationFail.WithParams(params),
						)
						expected := &httphdl.ResponseError[any]{
							BaseResponse: httphdl.BaseResponse{
								Code:    httphdl.ResponseCode(http.StatusBadRequest, errorx.CodeValidationFail.Int()),
								Message: "validation error",
							},
							Errors: []fieldMessage{
								{
									Field:   "totalIncome",
									Message: "required",
								},
								{
									Field:   "allowances",
									Message: "required",
								},
							},
						}

						Expect(actual).NotTo(BeNil())
						Expect(actual).To(Equal(expected))
					})

				})
			})

			Context("not existing in map response list", func() {
				When("with error message is empty", func() {
					It("should return 500 with unknown message", func() {
						arg := errorx.ErrTaxDeductNotFound.WithError(errorx.New(""))
						actual := httphdl.AsErrorResponse[any](arg)

						expected := &httphdl.ResponseError[any]{
							BaseResponse: httphdl.BaseResponse{
								Code:    httphdl.ResponseCode(http.StatusInternalServerError, arg.Code().Int()),
								Message: "unknown",
							},
						}

						Expect(actual).NotTo(BeNil())
						Expect(actual).To(Equal(expected))
					})
				})
				When("with wrap error", func() {
					It("should return 500 with unknown message", func() {
						arg := errorx.Wrap(errorx.ErrTaxDeductNotFound.WithError(errorx.New("error")), 0)
						actual := httphdl.AsErrorResponse[any](arg)

						expected := &httphdl.ResponseError[any]{
							BaseResponse: httphdl.BaseResponse{
								Code:    httphdl.ResponseCode(http.StatusInternalServerError, errorx.CodeTaxDeductNotFound.Int()),
								Message: "error",
							},
						}

						Expect(actual).NotTo(BeNil())
						Expect(actual).To(Equal(expected))
					})
				})

				It("should return 500 with error message", func() {
					arg := errorx.ErrTaxDeductNotFound.WithError(errorx.New("error"))
					actual := httphdl.AsErrorResponse[any](arg)

					expected := &httphdl.ResponseError[any]{
						BaseResponse: httphdl.BaseResponse{
							Code:    httphdl.ResponseCode(http.StatusInternalServerError, arg.Code().Int()),
							Message: "error",
						},
					}

					Expect(actual).NotTo(BeNil())
					Expect(actual).To(Equal(expected))
				})
			})
		})
	})

	Describe("http error handler", func() {
		var (
			app *echo.Echo
		)

		BeforeEach(func() {
			app = echo.New()
		})

		When("is echo error", func() {
			It("should return error response from builder", func() {
				req := httptest.NewRequest(http.MethodPost, "/", nil)
				rec := httptest.NewRecorder()
				c := app.NewContext(req, rec)
				httphdl.HTTPErrorHandler(echo.NewHTTPError(http.StatusInternalServerError, "echo error"), c)

				Expect(500).To(Equal(rec.Code))
				actual, err := compacJSON(rec.Body.String())
				Expect(err).NotTo(HaveOccurred())

				expectedResp := `{"code":"500000","message":"echo error"}`
				expected, err := compacJSON(expectedResp)
				Expect(err).NotTo(HaveOccurred())
				Expect(actual).To(Equal(expected))
			})
		})

		When("is not echo error", func() {
			It("should return error response", func() {
				req := httptest.NewRequest(http.MethodPost, "/", nil)
				rec := httptest.NewRecorder()
				c := app.NewContext(req, rec)
				httphdl.HTTPErrorHandler(errorx.ErrTaxDeductNotFound, c)

				Expect(500).To(Equal(rec.Code))
				actual, err := compacJSON(rec.Body.String())
				Expect(err).NotTo(HaveOccurred())

				expectedResp := `{"code":"500100","message":"error"}`
				expected, err := compacJSON(expectedResp)
				Expect(err).NotTo(HaveOccurred())
				Expect(actual).To(Equal(expected))
			})
		})
	})
})
