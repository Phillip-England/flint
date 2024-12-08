package command

import (
	"os"
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
	if secondArg == "spark" {
		cmd, err := NewSpark()
		if err != nil {
			return nil, err
		}
		return cmd, nil
	}
	if secondArg == "help" {
		cmd, err := NewHelp()
		if err != nil {
			return nil, err
		}
		return cmd, nil
	}
	cmd, err := NewHelp()
	if err != nil {
		return nil, err
	}
	return cmd, nil
}
