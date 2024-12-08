package executor

import "fmt"

type Build struct {
	Type string
}

func NewBuild() (*Build, error) {
	executor := &Build{
		Type: KeyHelp,
	}
	return executor, nil
}

func (exe *Build) Print() { fmt.Println(exe.Type) }
func (exe *Build) Run() error {
	return nil
}
