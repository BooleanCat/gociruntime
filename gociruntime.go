package gociruntime

import "os/exec"

type OCIRuntime struct {
	path    string
	rawArgs []string
}

func OCI(path string) OCIRuntime {
	return OCIRuntime{path: path}
}

func (r OCIRuntime) RawArgs(args ...string) OCIRuntime {
	r.rawArgs = args
	return r
}

func (r OCIRuntime) Command() *exec.Cmd {
	return exec.Command(r.path, r.rawArgs...)
}
