package cmd

import (
	"gopkg.in/urfave/cli.v1"
)

//ClientCommandProvider - Providers commands for running Orek app as client
type ClientCommandProvider struct{}

//GetCommands - gives commands for running Orek app as client to Orek Service
func (ccp *ClientCommandProvider) GetCommands() []cli.Command {
	return []cli.Command{
		cli.Command{
			Name: "client",
		},
	}
}

func (ccp *ClientCommandProvider) clientSubCommands() []cli.Command {
	return []cli.Command{
		cli.Command{
			Name: "client",
		},
	}
}
