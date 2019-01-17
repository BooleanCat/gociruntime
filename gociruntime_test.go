package gociruntime_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/BooleanCat/gociruntime"
)

var _ = Describe("OCIRuntime", func() {
	It("generates a command using the provided runtime", func() {
		cmd := gociruntime.OCI("railcar").Command()
		Expect(invocation(cmd)).To(Equal("railcar"))
	})

	Describe(".RawArgs", func() {
		It("appends raw args to the command", func() {
			cmd := gociruntime.OCI("railcar").RawArgs("--foo", "bar").Command()
			Expect(invocation(cmd)).To(Equal("railcar --foo bar"))
		})
	})
})
