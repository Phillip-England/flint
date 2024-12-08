package command

import "fmt"

type Help struct {
	Type string
}

func NewHelp() (*Help, error) {
	cmd := &Help{
		Type: KeyHelp,
	}
	return cmd, nil
}

func (cmd *Help) Print()          { fmt.Println(cmd.Type) }
func (cmd *Help) GetType() string { return cmd.Type }
