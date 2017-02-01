package cmd

import cli "gopkg.in/urfave/cli.v1"
import "github.com/varunamachi/orekng/olog"

//ServerCommandProvider - Providers commands for running Orek app as client
type ServerCommandProvider struct{}

//GetCommand - gives commands for running Orek app as client to Orek Service
func (ccp *ServerCommandProvider) GetCommand() cli.Command {
	subcmds := []cli.Command{
		serveCommand(),
	}
	return cli.Command{
		Name:        "service",
		Usage:       "Run Orek as a service",
		Subcommands: subcmds,
		Flags:       []cli.Flag{},
	}
}

func serveCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "start",
		Usage: "Starts the Orek server with given parameters",
		Flags: []cli.Flag{

			cli.IntFlag{
				Name:  "port",
				Value: 8000,
				Usage: "Port for server to bind to",
			},
			cli.BoolFlag{
				Name:  "silent",
				Usage: "Tells if console logging is required",
			},
		},
		Action: func(ctx *cli.Context) (err error) {

			argetr := ArgGetter{Ctx: ctx}
			silent := argetr.GetRequiredBool("silent")
			if err = argetr.Err; err == nil {
				if !silent {
					olog.GetLogger().RegisterWriter(olog.NewConsoleWriter())
				}
			}
			return err
		},
	}
	return cmd
}
