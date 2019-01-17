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
})
