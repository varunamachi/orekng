package cmd

import cli "gopkg.in/urfave/cli.v1"

//ServerCommandProvider - Providers commands for running Orek app as client
type ServerCommandProvider struct{}

//GetCommands - gives commands for running Orek app as client to Orek Service
func (ccp *ServerCommandProvider) GetCommands(orek *OrekApp) cli.Command {
	subcmds := []cli.Command{
		serveCommand(orek),
	}
	return cli.Command{
		Name:        "server",
		Subcommands: subcmds,
		Flags:       []cli.Flag{},
	}
}

func serveCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "serve",
		Flags: []cli.Flag{

			cli.IntFlag{
				Name:  "port",
				Value: 8000,
				Usage: "Port for server to bind to",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			if err = argetr.Err; err == nil {
			}
			return err
		},
	}
	return cmd
}
