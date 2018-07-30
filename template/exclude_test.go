package template_test

import (
	t "github.com/BooleanCat/go-functional/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ExcludeIter", func() {
	It("excludes items from the Iterator that pass the exclusion check", func() {
		lessThanFive := func(i interface{}) bool {
			value, ok := i.(int)
			Expect(ok).To(BeTrue())
			return value < 5
		}

		iter := t.Exclude(NewCounter(), lessThanFive)
		next := resultValue(iter.Next())
		Expect(next).To(Equal(5))
	})
})
