package main

import (
	"os"

	// _ "github.com/mattn/go-sqlite3"
	"github.com/varunamachi/orekng/cmd"
)

func main() {
	orek := cmd.OrekApp{
		CommandProviders: []cmd.CliCommandProvider{
			&cmd.ManageCommandProvider{},
			&cmd.ServerCommandProvider{},
			&cmd.ClientCommandProvider{},
		},
	}
	orek.Run(os.Args)
}
