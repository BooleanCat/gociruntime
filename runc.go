package gociruntime

type RuncRuntime struct {
	OCIRuntime
}

func Runc() RuncRuntime {
	return RuncRuntime{OCIRuntime: OCIRuntime{path: "runc"}}
}
