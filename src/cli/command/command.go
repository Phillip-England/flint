package command

import (
	"os"

	"github.com/phillip-england/purse"
)

type Command interface {
	Print()
	GetType() string
}

func New() (Command, error) {
	args := os.Args
	if len(args) <= 1 {
		cmd, err := NewHelp()
		if err != nil {
			return nil, err
		}
		return cmd, nil
	}
	secondArg := args[1]
	if secondArg == "build" {
		cmd, err := NewBuild()
		if err != nil {
			return nil, err
		}
		return cmd, nil
	}
	return nil, purse.Err(`invalid command`)
}
