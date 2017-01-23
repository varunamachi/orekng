package cmd

import cli "gopkg.in/urfave/cli.v1"
import "fmt"

//CliCommandProvider - gives commands supported by the application
type CliCommandProvider interface {
	GetCommands() cli.Command
}

//OrekApp - contains command providers and runs the app
type OrekApp struct {
	CommandProviders []CliCommandProvider
}

//RegisterCommandProvider - registers a command provider
func (orek *OrekApp) RegisterCommandProvider(cmdProvider CliCommandProvider) {
	if cmdProvider != nil {
		orek.CommandProviders = append(orek.CommandProviders, cmdProvider)
	}
}

//Run - runs the application
func (orek *OrekApp) Run(args []string) (err error) {
	app := cli.NewApp()
	app.Name = "Orek"
	app.Version = "0.0.1"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Varun Amachi",
			Email: "varunamachi@github.com",
		},
	}
	app.Commands = make([]cli.Command, 0, 30)
	for _, cmdp := range orek.CommandProviders {
		app.Commands = append(app.Commands, cmdp.GetCommands())
	}
	err = app.Run(args)
	return err
}

//AskString - gets string flag for key. If the context doesnt have it, it asks
//the user to enter it into console
func AskString(ctx *cli.Context, key string) (val string) {
	val = ctx.String(key)
	if len(val) == 0 {
		fmt.Print(key + ": ")
		_, err := fmt.Scanln(&val)
		if err != nil {
			val = ""
		}
	}
	return val
}
