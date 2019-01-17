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

	Describe(".Debug", func() {
		It("does not include the debug flag", func() {
			cmd := gociruntime.Runc().Command()
			Expect(cmd.Args).NotTo(ContainElement("--debug"))
		})

		When("debug is configured", func() {
			It("includes the debug flag", func() {
				cmd := gociruntime.Runc().Debug().Command()
				Expect(invocation(cmd)).To(Equal("runc --debug"))
			})
		})
	})

	Describe(".Log", func() {
		It("does not include the log flag", func() {
			cmd := gociruntime.Runc().Command()
			Expect(cmd.Args).NotTo(ContainElement("--log"))
		})

		When("log is configured", func() {
			It("includes the log flag", func() {
				cmd := gociruntime.Runc().Log("/my/log").Command()
				Expect(invocation(cmd)).To(Equal("runc --log /my/log"))
			})
		})
	})

	Describe(".LogFormat", func() {
		It("does not include the log-format flag", func() {
			cmd := gociruntime.Runc().Command()
			Expect(cmd.Args).NotTo(ContainElement("--log-format"))
		})

		When("log-format is configured", func() {
			It("includes the log-format flag", func() {
				cmd := gociruntime.Runc().LogFormat("myFormat").Command()
				Expect(invocation(cmd)).To(Equal("runc --log-format myFormat"))
			})
		})
	})

	Describe(".Root", func() {
		It("does not include the root flag", func() {
			cmd := gociruntime.Runc().Command()
			Expect(cmd.Args).NotTo(ContainElement("--root"))
		})

		When("root is configured", func() {
			It("includes the root flag", func() {
				cmd := gociruntime.Runc().Root("/my/root").Command()
				Expect(invocation(cmd)).To(Equal("runc --root /my/root"))
			})
		})
	})

	Describe(".Criu", func() {
		It("does not include the criu flag", func() {
			cmd := gociruntime.Runc().Command()
			Expect(cmd.Args).NotTo(ContainElement("--criu"))
		})

		When("criu is configured", func() {
			It("includes the criu flag", func() {
				cmd := gociruntime.Runc().Criu("/my/criu").Command()
				Expect(invocation(cmd)).To(Equal("runc --criu /my/criu"))
			})
		})
	})

	Describe(".SystemCgroup", func() {
		It("does not include the systemd-cgroup flag", func() {
			cmd := gociruntime.Runc().Command()
			Expect(cmd.Args).NotTo(ContainElement("--systemd-cgroup"))
		})

		When("systemd-cgroup is configured", func() {
			It("includes the systemd-cgroup flag", func() {
				cmd := gociruntime.Runc().SystemCgroup().Command()
				Expect(invocation(cmd)).To(Equal("runc --systemd-cgroup"))
			})
		})
	})

	Describe(".Help", func() {
		It("does not include the help flag", func() {
			cmd := gociruntime.Runc().Command()
			Expect(cmd.Args).NotTo(ContainElement("--help"))
		})

		When("help is configured", func() {
			It("includes the help flag", func() {
				cmd := gociruntime.Runc().Help().Command()
				Expect(invocation(cmd)).To(Equal("runc --help"))
			})
		})
	})

	Describe(".Version", func() {
		It("does not include the version flag", func() {
			cmd := gociruntime.Runc().Command()
			Expect(cmd.Args).NotTo(ContainElement("--version"))
		})

		When("version is configured", func() {
			It("includes the version flag", func() {
				cmd := gociruntime.Runc().Version().Command()
				Expect(invocation(cmd)).To(Equal("runc --version"))
			})
		})
	})

	Describe(".RawArgs", func() {
		It("prepends raw args to the command args", func() {
			cmd := gociruntime.Runc().Debug().RawArgs("--foo", "bar").Command()
			Expect(invocation(cmd)).To(Equal("runc --foo bar --debug"))
		})
	})

	Describe("Chaining", func() {
		It("generates a runc command with multiple args", func() {
			cmd := gociruntime.Runc().Debug().Root("/my/root").Criu("/my/criu").Command()
			Expect(invocation(cmd)).To(Equal("runc --debug --root /my/root --criu /my/criu"))
		})
	})

	Describe("OCIRuntimeState", func() {
		It("generates a state command with the provided container ID", func() {
			cmd := gociruntime.Runc().State("super-container").Command()
			Expect(invocation(cmd)).To(Equal("runc state super-container"))
		})

		When("a command is generated twice", func() {
			It("is unique", func() {
				state := gociruntime.Runc().State("super-container")
				cmd1 := state.Command()
				cmd2 := state.Command()

				Expect(invocation(cmd1)).To(Equal("runc state super-container"))
				Expect(invocation(cmd2)).To(Equal("runc state super-container"))
			})
		})
	})

	Describe("OCIRuntimeDelete", func() {
		It("generates a delete command with the provided container ID", func() {
			cmd := gociruntime.Runc().Delete("super-container").Command()
			Expect(invocation(cmd)).To(Equal("runc delete super-container"))
		})

		When("a command is generated twice", func() {
			It("is unique", func() {
				state := gociruntime.Runc().Delete("super-container")
				cmd1 := state.Command()
				cmd2 := state.Command()

				Expect(invocation(cmd1)).To(Equal("runc delete super-container"))
				Expect(invocation(cmd2)).To(Equal("runc delete super-container"))
			})
		})

		Describe(".Force", func() {
			It("does not include the force flag", func() {
				cmd := gociruntime.Runc().Command()
				Expect(cmd.Args).NotTo(ContainElement("--force"))
			})

			When("force is configured", func() {
				It("includes the force flag", func() {
					cmd := gociruntime.Runc().Delete("super-container").Force().Command()
					Expect(invocation(cmd)).To(Equal("runc delete super-container --force"))
				})
			})
		})
	})

	Describe("OCIRuntimeStart", func() {
		It("generates a start command with the provided container ID", func() {
			cmd := gociruntime.Runc().Start("super-container").Command()
			Expect(invocation(cmd)).To(Equal("runc start super-container"))
		})

		When("a command is generated twice", func() {
			It("is unique", func() {
				state := gociruntime.Runc().Start("super-container")
				cmd1 := state.Command()
				cmd2 := state.Command()

				Expect(invocation(cmd1)).To(Equal("runc start super-container"))
				Expect(invocation(cmd2)).To(Equal("runc start super-container"))
			})
		})
	})
})
