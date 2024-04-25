package http_test

import (
	"net/http"

	httphdl "github.com/wit-switch/assessment-tax/internal/handler/http"
	"github.com/wit-switch/assessment-tax/pkg/errorx"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Builder", func() {
	type fieldMessage struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}

	var (
		builder *httphdl.ResponseErrorBuilder[fieldMessage]
	)

	BeforeEach(func() {
		builder = httphdl.NewResponseErrorBuilder[fieldMessage]()
	})

	Describe("new response error builder", func() {
		It("should return response error builder with default value", func() {
			Expect(builder).NotTo(BeNil())
		})
	})

	Describe("with code", func() {
		It("should set code into builder", func() {
			builder.WithCode(999)
			Expect(builder.Build().Code).To(Equal("500999"))
		})
	})

	Describe("with message", func() {
		It("should set message into builder", func() {
			builder.WithMesssage("ok")
			Expect(builder.Build().Message).To(Equal("ok"))
		})
	})

	Describe("with http status", func() {
		It("should set http status into builder", func() {
			builder.WithHTTPStatus(http.StatusServiceUnavailable)
			Expect(builder.Build().StatusCode()).To(Equal(503))
		})
	})

	Describe("with error", func() {
		It("should set error into builder", func() {
			err := errorx.New("error")
			builder.WithMesssage("")
			builder.WithError(err)
			Expect(builder.Build().Error()).To(Equal(err.Error()))
		})
	})

	Describe("with params", func() {
		It("should set params into builder", func() {
			params := fieldMessage{
				Field:   "json tag",
				Message: "error message",
			}
			builder.WithParams(params)
			Expect(builder.Build().Errors).To(Equal(params))
		})
	})

	Describe("build response error", func() {
		When("with message", func() {
			It("should return response with message", func() {
				actual := builder.
					WithMesssage("ok").
					Build()

				expected := &httphdl.ResponseError[fieldMessage]{
					BaseResponse: httphdl.BaseResponse{
						Code:    httphdl.ResponseCode(http.StatusInternalServerError, errorx.CodeUnknown.Int()),
						Message: "ok",
					},
					Errors: fieldMessage{},
				}
				Expect(actual).To(Equal(expected))
			})
		})

		When("with empty message", func() {
			Context("code is validation fail", func() {
				It("should return response with validation fail message", func() {
					actual := builder.
						WithMesssage("").
						WithCode(errorx.CodeValidationFail.Int()).
						Build()

					expected := &httphdl.ResponseError[fieldMessage]{
						BaseResponse: httphdl.BaseResponse{
							Code:    httphdl.ResponseCode(http.StatusInternalServerError, errorx.CodeValidationFail.Int()),
							Message: "validation error",
						},
						Errors: fieldMessage{},
					}
					Expect(actual).To(Equal(expected))
				})
			})

			Context("code is not validation fail and have error", func() {
				It("should return response with error message", func() {
					actual := builder.
						WithMesssage("").
						WithError(errorx.New("error")).
						Build()

					expected := &httphdl.ResponseError[fieldMessage]{
						BaseResponse: httphdl.BaseResponse{
							Code:    httphdl.ResponseCode(http.StatusInternalServerError, errorx.CodeUnknown.Int()),
							Message: "error",
						},
						Errors: fieldMessage{},
					}
					Expect(actual).To(Equal(expected))
				})
			})

			It("should return response with unknown message", func() {
				actual := builder.
					WithMesssage("").
					Build()

				expected := &httphdl.ResponseError[fieldMessage]{
					BaseResponse: httphdl.BaseResponse{
						Code:    httphdl.ResponseCode(http.StatusInternalServerError, errorx.CodeUnknown.Int()),
						Message: "unknown",
					},
					Errors: fieldMessage{},
				}
				Expect(actual).To(Equal(expected))
			})
		})
	})
})
