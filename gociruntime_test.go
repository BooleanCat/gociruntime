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

	Describe("OCIRuntimeState", func() {
		It("generates a state command with the provided container ID", func() {
			cmd := gociruntime.OCI("railcar").State("super-container").Command()
			Expect(invocation(cmd)).To(Equal("railcar state super-container"))
		})

		When("a command is generated twice", func() {
			It("is unique", func() {
				state := gociruntime.OCI("railcar").State("super-container")
				cmd1 := state.Command()
				cmd2 := state.Command()

				Expect(invocation(cmd1)).To(Equal("railcar state super-container"))
				Expect(invocation(cmd2)).To(Equal("railcar state super-container"))
			})
		})
	})

	Describe("OCIRuntimeDelete", func() {
		It("generates a delete command with the provided container ID", func() {
			cmd := gociruntime.OCI("railcar").Delete("super-container").Command()
			Expect(invocation(cmd)).To(Equal("railcar delete super-container"))
		})

		When("a command is generated twice", func() {
			It("is unique", func() {
				state := gociruntime.OCI("railcar").Delete("super-container")
				cmd1 := state.Command()
				cmd2 := state.Command()

				Expect(invocation(cmd1)).To(Equal("railcar delete super-container"))
				Expect(invocation(cmd2)).To(Equal("railcar delete super-container"))
			})
		})
	})
})
