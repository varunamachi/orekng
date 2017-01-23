package cmd

import (
	"gopkg.in/urfave/cli.v1"
)

//CliCommandProvider - gives commands supported by the application
type CliCommandProvider interface {
	GetCommands() []cli.Command
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
		app.Commands = append(app.Commands, cmdp.GetCommands()...)
	}
	err = app.Run(args)
	return err
}
