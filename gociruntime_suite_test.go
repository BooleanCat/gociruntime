package gociruntime_test

import (
	"os/exec"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGociruntime(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gociruntime Suite")
}

func invocation(cmd *exec.Cmd) string {
	return strings.Join(cmd.Args, " ")
}
