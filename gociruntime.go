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

func (r OCIRuntime) Start(containerID string) OCIRuntimeStart {
	return OCIRuntimeStart{args: r.Command().Args, containerID: containerID}
}

func (r OCIRuntime) Kill(containerID, signal string) OCIRuntimeKill {
	return OCIRuntimeKill{args: r.Command().Args, containerID: containerID, signal: signal}
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

type OCIRuntimeStart struct {
	args        []string
	containerID string
}

func (r OCIRuntimeStart) Command() *exec.Cmd {
	cmd := exec.Command(r.args[0], r.args[1:]...)
	cmd.Args = append(cmd.Args, "start", r.containerID)
	return cmd
}

type OCIRuntimeKill struct {
	args        []string
	containerID string
	signal      string
}

func (r OCIRuntimeKill) Command() *exec.Cmd {
	cmd := exec.Command(r.args[0], r.args[1:]...)
	cmd.Args = append(cmd.Args, "kill", r.containerID, r.signal)
	return cmd
}
