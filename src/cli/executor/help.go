package executor

import "fmt"

type Help struct {
	Type string
}

func NewHelp() (*Help, error) {
	executor := &Help{
		Type: KeyHelp,
	}
	return executor, nil
}

func (exe *Help) Print() { fmt.Println(exe.Type) }
func (exe *Help) Run() error {
	return nil
}
