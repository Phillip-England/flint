package main

import (
	"flint/src/cli"
	"fmt"
	"os"
)

func main() {

	cmd, err := cli.NewCommand()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}

	cmd.Print()

}
