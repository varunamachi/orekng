package cmd

import (
	"github.com/varunamachi/orekng/data"
	"gopkg.in/urfave/cli.v1"
)

//ClientCommandProvider - Providers commands for running Orek app as client
type ClientCommandProvider struct{}

//GetCommands - gives commands for running Orek app as client to Orek Service
func (ccp *ClientCommandProvider) GetCommands() cli.Command {
	subcmds := []cli.Command{
		listUsersCommand(),
		showUserCommand(),
		showUserWithEmailCommand(),
		createUserCommand(),
		updateUserCommand(),
		deleteUserCommand(),

		listEndpointsCommand(),
		showEndpointCommand(),
		createEndpointCommand(),
		updateEndpointCommand(),
		deleteEndpointCommand(),

		listVariablesCommand(),
		listVariablesForEndpointCommand(),
		showVariableCommand(),
		createVariableCommand(),
		updateVariableCommand(),
		deleteVariableCommand(),

		listParametersCommand(),
		listParametersForEndpointCommand(),
		showParameterCommand(),
		createParameterCommand(),
		updateParameterCommand(),
		deleteParameterCommand(),

		listUserGroupsCommand(),
		showUserGroupCommand(),
		createUserGroupCommand(),
		updateUserGroupCommand(),
		deleteUserGroupCommand(),

		addUserToGroupCommand(),
		removeUserFromGroupCommand(),
		getUsersInGroupCommand(),
		getGroupsForUserCommand(),

		clearValuesForVariableCommand(),
		getValuesForVariableCommand(),

		setPasswordCommand(),
		updatePasswordCommand(),
	}
	return cli.Command{
		Name:        "client",
		Subcommands: subcmds,
		Flags:       []cli.Flag{},
	}
}

func listUsersCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "list-users",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func showUserCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name: "show-user",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "user-name",
				Value: "",
				Usage: "The unique user_name for the user",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func showUserWithEmailCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func createUserCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name: "create-user",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "user-name",
				Value: "",
				Usage: "The unique user_name for the user",
			},
			cli.StringFlag{
				Name:  "first-name",
				Value: "",
				Usage: "The first name of the user",
			},
			cli.StringFlag{
				Name:  "second-name",
				Value: "",
				Usage: "The second name of the user",
			},
			cli.StringFlag{
				Name:  "email",
				Value: "",
				Usage: "Email of the user",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			userName := argetr.GetRequiredString("user-name")
			email := argetr.GetRequiredString("email")
			firstName := argetr.GetString("first-name")
			secondName := argetr.GetString("second-name")
			if argetr.Err == nil {
				//Below should only run if it is local mode otherwise should use
				//the not yet implemented REST client mode
				err = data.GetDataStore().CreateUser(&data.User{
					Name:       userName,
					FirstName:  firstName,
					SecondName: secondName,
					Email:      email,
				})

			}
			return err
		},
	}
	return cmd
}

func updateUserCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func deleteUserCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func listEndpointsCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func showEndpointCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func createEndpointCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func updateEndpointCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func deleteEndpointCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func listVariablesCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func listVariablesForEndpointCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func showVariableCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func createVariableCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func updateVariableCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func deleteVariableCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func listParametersCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func listParametersForEndpointCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func showParameterCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func createParameterCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func updateParameterCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func deleteParameterCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func listUserGroupsCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func showUserGroupCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func createUserGroupCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func updateUserGroupCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func deleteUserGroupCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func addUserToGroupCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func removeUserFromGroupCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func getUsersInGroupCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func getGroupsForUserCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func clearValuesForVariableCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func getValuesForVariableCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func setPasswordCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func updatePasswordCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}
