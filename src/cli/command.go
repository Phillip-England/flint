package cli

import (
	"fmt"
	"os"

	"github.com/phillip-england/purse"
)

type Command interface {
	Print()
}

func NewCommand() (Command, error) {
	args := os.Args
	fmt.Println(args)
	return nil, purse.Err(`invalid command`)
}
