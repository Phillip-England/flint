package main

import (
	"fmt"
	"os"

	"github.com/phillip-england/flint/src/cli/command"
	"github.com/phillip-england/flint/src/cli/executor"
)

func main() {

	cmd, err := command.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}

	executor, err := executor.New(cmd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}

	err = executor.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}

}
