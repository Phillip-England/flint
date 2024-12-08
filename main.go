package main

import (
	"flint/src/cli/command"
	"flint/src/cli/executor"
	"fmt"
	"os"
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
