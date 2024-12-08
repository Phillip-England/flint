package command

import "fmt"

type Build struct {
	Type string
}

func NewBuild() (*Build, error) {
	cmd := &Build{
		Type: KeyBuild,
	}
	return cmd, nil
}

func (cmd *Build) Print()          { fmt.Println(cmd.Type) }
func (cmd *Build) GetType() string { return cmd.Type }
