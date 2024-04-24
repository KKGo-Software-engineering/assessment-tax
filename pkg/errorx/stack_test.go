package errorx_test

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/wit-switch/assessment-tax/pkg/errorx"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stack", func() {
	Describe("ReplaceAttr", func() {
		When("attr is noy error", func() {
			It("should not replace value then return slog.Attr", func() {
				actual := errorx.ReplaceAttr(nil, slog.String("string", "text"))

				Expect(actual.String()).To(Equal("string=text"))
			})
		})

		When("attr is error", func() {
			Context("with standard error", func() {
				When("wrap error", func() {
					It("should not replace value then return slog.Attr", func() {
						actual := errorx.ReplaceAttr(nil, slog.Any("err", fmt.Errorf("wrap %w", errors.New("standard error"))))

						Expect(actual.String()).To(Equal("err=[msg=wrap standard error]"))
					})
				})

				It("should not replace value then return slog.Attr", func() {
					actual := errorx.ReplaceAttr(nil, slog.Any("err", errors.New("standard error")))

					Expect(actual.String()).To(Equal("err=[msg=standard error]"))
				})
			})

			Context("with internal error", func() {
				When("only code", func() {
					It("should not replace value then return slog.Attr", func() {
						actual := errorx.ReplaceAttr(nil, slog.Any("err", errorx.NewInternalErr[any](errorx.CodeUnknown)))

						Expect(actual.String()).To(Equal("err=[msg=0]"))
					})
				})

				When("with error", func() {
					It("should replace value then return slog.Attr", func() {
						actual := errorx.ReplaceAttr(
							nil,
							slog.Any("err", errorx.NewInternalErr[any](errorx.CodeUnknown).WithError(errorx.New("package error"))),
						)

						Expect(actual.String()).To(MatchRegexp(`err=\[msg=package error trace=`))
					})
				})
			})

			Context("with wrap error", func() {
				It("should replace value then return slog.Attr", func() {
					actual := errorx.ReplaceAttr(nil, slog.Any("err", errorx.Wrap(errors.New("standard error"), 0)))

					Expect(actual.String()).To(MatchRegexp(`err=\[msg=standard error trace=`))
				})
			})

			It("should replace value then return slog.Attr", func() {
				actual := errorx.ReplaceAttr(nil, slog.Any("err", errorx.New("package error")))

				Expect(actual.String()).To(MatchRegexp(`err=\[msg=package error trace=`))
			})
		})
	})
})
