package executor

import (
	"flint/src/cli/command"

	"github.com/phillip-england/purse"
)

type Executor interface {
	Print()
	Run() error
}

func New(cmd command.Command) (Executor, error) {
	if cmd.GetType() == command.KeyHelp {
		executor, err := NewHelp()
		if err != nil {
			return nil, err
		}
		return executor, nil
	}
	if cmd.GetType() == command.KeySpark {
		executor, err := NewSpark()
		if err != nil {
			return nil, err
		}
		return executor, nil
	}
	return nil, purse.Err(`
		Command of type %s does not link to a valid Executor
	`)
}
