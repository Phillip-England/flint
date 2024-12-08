package executor

import (
	"fmt"
	"strings"

	"github.com/phillip-england/purse"
)

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
	title := purse.Fmt(`
		#################################################
		##          ##  ###  ####   ######  ##        ###
		##  ##########  #########    #####  #####  ######
		##       #####  ###  ####  #  ####  #####  ######
		##  ##########  ###  ####  ##  ###  #####  ######
		##  ##########  ###  ####  ###  ##  #####  ######
		##  ##########  #########  ####  #  #####  ######
		##  ##########         ##  #####    #####  ######
	`)
	title = strings.ReplaceAll(title, " ", "$")
	title = strings.ReplaceAll(title, "#", " ")
	title = strings.ReplaceAll(title, "$", "#")

	intro := purse.Fmt(`
		%s
		-----------------------------------------------
		Language-Agnostic Static Sites

		- serve your app locally
		- setup your flint.json
		- run 'flint spark'
		- deploy your static assets

		Read the docs: https://github.com/phillip-england/flint
	`, title)

	fmt.Println(intro)

	return nil
}
