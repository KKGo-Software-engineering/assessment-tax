package errorx_test

import (
	"errors"
	"fmt"

	"github.com/wit-switch/assessment-tax/pkg/errorx"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Errorx", func() {
	var (
		errStd    = errors.New("new error")
		errPkg    = errorx.New("new error")
		errWrap   = errorx.Wrap(errStd, 0)
		errPrefix = errorx.WrapPrefix(errStd, "prefix", 0)
	)

	Describe("New", func() {
		It("should create an error", func() {
			Expect(errPkg).To(MatchError(errStd))
		})
	})

	Describe("Is", func() {
		It("should detect an error", func() {
			Expect(errorx.Is(errPkg, errStd)).To(BeFalse())
		})
	})

	Describe("As", func() {
		It("should match an error", func() {
			Expect(errorx.As(errPkg, &errStd)).To(BeTrue())
		})
	})

	Describe("Wrap", func() {
		It("should wrap an error", func() {
			Expect(errWrap.Error()).To(Equal(errStd.Error()))
		})
	})

	Describe("WrapPrefix", func() {
		It("should wrap an error with prefix", func() {
			Expect(errPrefix.Error()).To(Equal(fmt.Sprintf("prefix: %s", errStd.Error())))
		})
	})

	Describe("Unwrap", func() {
		It("should Unwrap an error", func() {
			Expect(errorx.Unwrap(errWrap)).To(MatchError(errors.Unwrap(errWrap)))
		})
	})

	Describe("Errorf", func() {
		It("should create an error with format", func() {
			Expect(errorx.Errorf("%s", "errorf")).To(MatchError(fmt.Errorf("%s", "errorf")))
		})
	})
})
