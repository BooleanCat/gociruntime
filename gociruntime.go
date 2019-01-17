package gociruntime

import "os/exec"

type OCIRuntime struct {
	path string
}

func OCI(path string) OCIRuntime {
	return OCIRuntime{path: path}
}

func (r OCIRuntime) Command() *exec.Cmd {
	return exec.Command(r.path)
}
