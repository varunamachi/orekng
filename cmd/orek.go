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

//ArgGetter - this struct and its method are helpers to combine getting args
//from commandline arguments or from reading from console. Also handles errors
//when required arguments are not provided
type ArgGetter struct {
	Ctx *cli.Context
	Err error
}

//GetString - gives a string argument either from commandline or from blocking
//user input, this method doesnt complain even if the arg-value is empty
func (retriever *ArgGetter) GetString(key string) (val string) {
	if retriever.Err != nil {
		return val
	}
	val = retriever.Ctx.String(key)
	if len(val) == 0 {
		fmt.Print(key + ": ")
		_, err := fmt.Scanln(&val)
		if err != nil {
			val = ""
		}
	}
	return val
}

//GetRequiredString - gives a string argument either from commandline or from
//blocking user input, this method sets the error if required arg-val is empty
func (retriever *ArgGetter) GetRequiredString(key string) (val string) {
	if retriever.Err != nil {
		return val
	}
	val = retriever.Ctx.String(key)
	if len(val) == 0 {
		fmt.Print(key + ": ")
		_, err := fmt.Scanln(&val)
		if err != nil || len(val) == 0 {
			val = ""
			retriever.Err = fmt.Errorf("Required argument %s not provided", key)
		}
	}
	return val
}
