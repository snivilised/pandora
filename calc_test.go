package pandora_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/snivilised/pandora"
)

var _ = Describe("Calc", func() {
	Context("dummy", func() {
		It("should: foo-bar", func() {
			Expect(1).To(Equal(1))
		})
	})

	Context("Add", func() {
		It("should: add", func() {
			calc := pandora.Calc{}
			Expect(calc.Add(42).Result()).To(Equal(42))
			Expect(calc.Clear().Add(42).Add(18).Result()).To(Equal(60))
		})
	})

	Context("Subtract", func() {
		It("should: subtract", func() {
			calc := pandora.Calc{
				Accumulator: 27,
			}
			Expect(calc.Subtract(42).Result()).To(Equal(-15))
		})
	})

	Context("Multiply", func() {
		It("should: multiply", func() {
			calc := pandora.Calc{
				Accumulator: 9,
			}
			Expect(calc.Multiply(42).Result()).To(Equal(378))
		})
	})
})
