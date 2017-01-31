package cmd

import (
	"github.com/varunamachi/orekng/data"
	"github.com/varunamachi/orekng/olog"
	cli "gopkg.in/urfave/cli.v1"
)

//ManageCommandProvider - provides command for managing Orek instance
type ManageCommandProvider struct {
}

//GetCommand - gives commands for managing orek instance
func (ccp *ManageCommandProvider) GetCommand() cli.Command {
	subcmds := []cli.Command{
		initCommand(),
		deleteDatabaseCommand(),
		clearDataCommand(),
	}
	return cli.Command{
		Name:        "manage",
		Subcommands: subcmds,
		Flags:       []cli.Flag{},
	}
}

func initCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "init",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			if err = argetr.Err; err == nil {
				err = data.GetStore().Init()
				if err != nil {
					olog.Print("Manage", "Data source initialized")
				}
			}
			return err
		},
		Usage: "Initializes the database if it is not initialized already",
	}
	return cmd
}

func deleteDatabaseCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "destroy",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			if err = argetr.Err; err == nil {
				err = data.GetStore().DeleteSchema()
				if err != nil {
					olog.Print("Manage", "Data source destroyed")
				}
			}
			return err
		},
		Usage: "Deletes the Orek's database schema",
	}
	return cmd
}

func clearDataCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "clear-data",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			if err = argetr.Err; err == nil {
				err = data.GetStore().ClearData()
				if err != nil {
					olog.Print("Manage", "Data source cleared")
				}
			}
			return err
		},
		Usage: "Clears all the data from all of the orek's table",
	}
	return cmd
}
