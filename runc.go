package gociruntime

import "os/exec"

type RuncRuntime struct {
	OCIRuntime

	debug         bool
	log           string
	logFormat     string
	root          string
	criu          string
	systemdCgroup bool
	help          bool
	version       bool
}

func Runc() RuncRuntime {
	return RuncRuntime{OCIRuntime: OCIRuntime{path: "runc"}}
}

func (r RuncRuntime) Command() *exec.Cmd {
	cmd := r.OCIRuntime.Command()

	cmd.Args = appendBoolArg(cmd.Args, "--debug", r.debug)
	cmd.Args = appendStringArg(cmd.Args, "--log", r.log)
	cmd.Args = appendStringArg(cmd.Args, "--log-format", r.logFormat)
	cmd.Args = appendStringArg(cmd.Args, "--root", r.root)
	cmd.Args = appendStringArg(cmd.Args, "--criu", r.criu)
	cmd.Args = appendBoolArg(cmd.Args, "--systemd-cgroup", r.systemdCgroup)
	cmd.Args = appendBoolArg(cmd.Args, "--help", r.help)
	cmd.Args = appendBoolArg(cmd.Args, "--version", r.version)

	return cmd
}

func (r RuncRuntime) RawArgs(args ...string) RuncRuntime {
	r.OCIRuntime = r.OCIRuntime.RawArgs(args...)
	return r
}

func (r RuncRuntime) Debug() RuncRuntime {
	r.debug = true
	return r
}

func (r RuncRuntime) Log(path string) RuncRuntime {
	r.log = path
	return r
}

func (r RuncRuntime) LogFormat(path string) RuncRuntime {
	r.logFormat = path
	return r
}

func (r RuncRuntime) Root(path string) RuncRuntime {
	r.root = path
	return r
}

func (r RuncRuntime) Criu(path string) RuncRuntime {
	r.criu = path
	return r
}

func (r RuncRuntime) SystemCgroup() RuncRuntime {
	r.systemdCgroup = true
	return r
}

func (r RuncRuntime) Help() RuncRuntime {
	r.help = true
	return r
}

func (r RuncRuntime) Version() RuncRuntime {
	r.version = true
	return r
}

func (r RuncRuntime) Delete(containerID string) RuncRuntimeDelete {
	return RuncRuntimeDelete{OCIRuntimeDelete: r.OCIRuntime.Delete(containerID)}
}

func (r RuncRuntime) Kill(containerID, signal string) RuncRuntimeKill {
	return RuncRuntimeKill{OCIRuntimeKill: r.OCIRuntime.Kill(containerID, signal)}
}

type RuncRuntimeDelete struct {
	OCIRuntimeDelete

	force bool
}

func (r RuncRuntimeDelete) Command() *exec.Cmd {
	cmd := r.OCIRuntimeDelete.Command()

	cmd.Args = appendBoolArg(cmd.Args, "--force", r.force)

	return cmd
}

func (r RuncRuntimeDelete) Force() RuncRuntimeDelete {
	r.force = true
	return r
}

type RuncRuntimeKill struct {
	OCIRuntimeKill

	all bool
}

func (r RuncRuntimeKill) Command() *exec.Cmd {
	cmd := r.OCIRuntimeKill.Command()

	cmd.Args = appendBoolArg(cmd.Args, "--all", r.all)

	return cmd
}

func (r RuncRuntimeKill) All() RuncRuntimeKill {
	r.all = true
	return r
}

func appendStringArg(args []string, flag, value string) []string {
	if value == "" {
		return args
	}

	return append(args, flag, value)
}

func appendBoolArg(args []string, flag string, value bool) []string {
	if !value {
		return args
	}

	return append(args, flag)
}
