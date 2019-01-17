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

func (r OCIRuntime) State(containerID string) OCIRuntimeState {
	return OCIRuntimeState{args: r.Command().Args, containerID: containerID}
}

type OCIRuntimeState struct {
	args        []string
	containerID string
}

func (s OCIRuntimeState) Command() *exec.Cmd {
	cmd := exec.Command(s.args[0], s.args[1:]...)
	cmd.Args = append(cmd.Args, "state", s.containerID)
	return cmd
}
