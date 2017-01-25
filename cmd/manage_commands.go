package cmd

import cli "gopkg.in/urfave/cli.v1"

//ManageCommandProvider - provides command for managing Orek instance
type ManageCommandProvider struct {
}

//GetCommands - gives commands for managing orek instance
func (ccp *ManageCommandProvider) GetCommands(orek *OrekApp) cli.Command {
	subcmds := []cli.Command{}
	return cli.Command{
		Name:        "manage",
		Subcommands: subcmds,
		Flags:       []cli.Flag{},
	}
}

func initCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "init",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			if err = argetr.Err; err == nil {
			}
			return err
		},
	}
	return cmd
}

func reInitCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "re-init",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			if err = argetr.Err; err == nil {
			}
			return err
		},
	}
	return cmd
}

func deleteDatabaseCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "destroy",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			if err = argetr.Err; err == nil {
			}
			return err
		},
	}
	return cmd
}

func clearDataCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "clear-data",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			if err = argetr.Err; err == nil {
			}
			return err
		},
	}
	return cmd
}
