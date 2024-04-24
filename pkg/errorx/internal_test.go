package errorx_test

import (
	"errors"

	"github.com/wit-switch/assessment-tax/pkg/errorx"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Internal", func() {
	DescribeTable("IsInternalErr", func(err error, expected *errorx.InternalError[any]) {
		actual := errorx.IsInternalErr[any](err)
		if expected != nil {
			Expect(actual).To(HaveOccurred())
			return
		}

		Expect(actual).NotTo(HaveOccurred())
	},
		Entry("with nil error", nil, nil),
		Entry("with internal error", &errorx.InternalError[any]{}, &errorx.InternalError[any]{}),
		Entry("with wrap internal error", errorx.Wrap(&errorx.InternalError[any]{}, 0), &errorx.InternalError[any]{}),
		Entry("with standard error", errors.New("standard error"), nil),
		Entry("with wrap standard error", errorx.Wrap(errors.New("standard error"), 0), nil),
	)

	DescribeTable("InternalCode", func(err error, expected errorx.ErrCode) {
		actual := errorx.InternalCode[any](err)
		Expect(actual).To(Equal(expected))
	},
		Entry("with internal error", errorx.NewInternalErr[any](errorx.CodeValidationFail), errorx.CodeValidationFail),
		Entry("with standard error", errors.New("standard error"), errorx.CodeUnknown),
	)

	Describe("InternalError", func() {
		errBase := errorx.NewInternalErr[any](errorx.CodeUnknown)

		When("get code", func() {
			It("should get code from error", func() {
				Expect(errBase.Code()).To(Equal(errorx.CodeUnknown))
				Expect(errBase.Error()).To(Equal("0"))
			})
		})

		When("add param", func() {
			It("should get param from error", func() {
				err := errBase.WithParams("params")
				Expect(err.Params()).To(Equal("params"))
			})
		})

		When("add error", func() {
			It("should get error from error", func() {
				errStd := errors.New("standard error")
				err := errBase.WithError(errStd)
				Expect(err.Error()).To(Equal(errStd.Error()))
			})
		})

		When("get stack error", func() {
			Context("with err is nil", func() {
				It("should return nil", func() {
					actual := errBase.StackError()
					Expect(actual).To(BeNil())
				})
			})

			Context("with standard error", func() {
				It("should return nil", func() {
					actual := errBase.WithError(errors.New("standard error")).StackError()
					Expect(actual).To(BeNil())
				})
			})

			It("should return error with stacktrace", func() {
				errPkg := errorx.New("package error")
				actual := errBase.WithError(errPkg).StackError()
				Expect(actual).To(MatchError(errPkg))
			})
		})
	})
})
