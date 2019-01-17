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

func (r OCIRuntime) Delete(containerID string) OCIRuntimeDelete {
	return OCIRuntimeDelete{args: r.Command().Args, containerID: containerID}
}

type OCIRuntimeState struct {
	args        []string
	containerID string
}

func (r OCIRuntimeState) Command() *exec.Cmd {
	cmd := exec.Command(r.args[0], r.args[1:]...)
	cmd.Args = append(cmd.Args, "state", r.containerID)
	return cmd
}

type OCIRuntimeDelete struct {
	args        []string
	containerID string
}

func (r OCIRuntimeDelete) Command() *exec.Cmd {
	cmd := exec.Command(r.args[0], r.args[1:]...)
	cmd.Args = append(cmd.Args, "delete", r.containerID)
	return cmd
}
