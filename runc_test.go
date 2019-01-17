package gociruntime_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/BooleanCat/gociruntime"
)

var _ = Describe("RuncRuntime", func() {
	It("generates a runc command", func() {
		cmd := gociruntime.Runc().Command()
		Expect(invocation(cmd)).To(Equal("runc"))
	})
})
