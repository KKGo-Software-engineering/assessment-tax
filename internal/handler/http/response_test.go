package http_test

import (
	"net/http"

	httphdl "github.com/wit-switch/assessment-tax/internal/handler/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Response", func() {
	DescribeTable("response code", func(httpStatus, code int, expected string) {
		actual := httphdl.ResponseCode(httpStatus, code)
		Expect(actual).To(Equal(expected))
	},
		Entry("with http status = 500, code=1", http.StatusInternalServerError, 1, "500001"),
		Entry("with http status = 400, code=10", http.StatusBadRequest, 10, "400010"),
		Entry("with http status = 404, code=100", http.StatusNotFound, 100, "404100"),
		Entry("with http status = 401, code=1000", http.StatusUnauthorized, 1000, "4011000"),
	)

	Describe("base response", func() {
		DescribeTable("status code", func(resp httphdl.BaseResponse, expected int) {
			actual := resp.StatusCode()
			Expect(actual).To(Equal(expected))
		},
			Entry("with code=4011000", httphdl.BaseResponse{
				Code: "4011000",
			}, 401),
			Entry("with code=404100", httphdl.BaseResponse{
				Code: "404100",
			}, 404),
			Entry("with code=40001", httphdl.BaseResponse{
				Code: "40001",
			}, 400),
			Entry("with code=2000", httphdl.BaseResponse{
				Code: "2000",
			}, 200),
			Entry("with code=501", httphdl.BaseResponse{
				Code: "501",
			}, 501),
			Entry("with code=50", httphdl.BaseResponse{
				Code: "50",
			}, 50),
			Entry("with code=5", httphdl.BaseResponse{
				Code: "5",
			}, 5),
			Entry("with code is empty", httphdl.BaseResponse{
				Code: "",
			}, 500),
			Entry("with code is not number", httphdl.BaseResponse{
				Code: "xxx",
			}, 500),
		)
	})

	Describe("response error", func() {
		var responseErr httphdl.ResponseError[any]
		BeforeEach(func() {
			responseErr = httphdl.ResponseError[any]{}
		})
		When("with new response error", func() {
			It("should return response error with default value", func() {
				expected := httphdl.ResponseError[any]{}
				Expect(responseErr).To(Equal(expected))
			})
		})

		When("with error", func() {
			It("should return message from base response", func() {
				actual := responseErr.Error()
				expected := ""
				Expect(actual).To(Equal(expected))
			})
		})
	})
})
