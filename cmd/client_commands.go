package cmd

import (
	"fmt"

	"github.com/varunamachi/orekng/data"
	"gopkg.in/urfave/cli.v1"
)

//ClientCommandProvider - Providers commands for running Orek app as client
type ClientCommandProvider struct{}

//GetCommands - gives commands for running Orek app as client to Orek Service
func (ccp *ClientCommandProvider) GetCommands() cli.Command {
	subcmds := []cli.Command{
		createUserSubCommand(),
	}
	return cli.Command{
		Name:        "client",
		Subcommands: subcmds,
		Flags:       []cli.Flag{},
	}
}

func createUserSubCommand() cli.Command {
	return cli.Command{
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
			userName := ctx.String("user-name")
			email := ctx.String("email")
			if userName != "" && email != "" {
				firstName := ctx.String("first-name")
				secondName := ctx.String("second-name")

				//Below should only run if it is local mode otherwise should use
				//the not yet implemented REST client mode
				err = data.GetDataStore().CreateUser(&data.User{
					Name:       userName,
					FirstName:  firstName,
					SecondName: secondName,
					Email:      email,
				})

			} else {
				pbm := "user name"
				if userName != "" {
					pbm = "email"
				}
				err = fmt.Errorf("Create User: %s not provided", pbm)
			}
			return err
		},
	}
}
