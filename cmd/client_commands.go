package cmd

import (
	"fmt"

	"github.com/varunamachi/orekng/data"
	"gopkg.in/urfave/cli.v1"
)

//ClientCommandProvider - Providers commands for running Orek app as client
type ClientCommandProvider struct{}

func isLocalMode(ctx *cli.Context) bool {
	return true
}

//GetCommands - gives commands for running Orek app as client to Orek Service
func (ccp *ClientCommandProvider) GetCommands(orek *OrekApp) cli.Command {
	subcmds := []cli.Command{
		listUsersCommand(orek),
		showUserCommand(orek),
		showUserWithEmailCommand(orek),
		createUserCommand(orek),
		updateUserCommand(orek),
		deleteUserCommand(orek),
		listEndpointsCommand(orek),
		showEndpointCommand(orek),
		createEndpointCommand(orek),
		updateEndpointCommand(orek),
		deleteEndpointCommand(orek),
		listVariablesCommand(orek),
		listVariablesForEndpointCommand(orek),
		showVariableCommand(orek),
		createVariableCommand(orek),
		updateVariableCommand(orek),
		deleteVariableCommand(orek),
		listParametersCommand(orek),
		listParametersForEndpointCommand(orek),
		showParameterCommand(orek),
		createParameterCommand(orek),
		updateParameterCommand(orek),
		deleteParameterCommand(orek),
		listUserGroupsCommand(orek),
		showUserGroupCommand(orek),
		createUserGroupCommand(orek),
		updateUserGroupCommand(orek),
		deleteUserGroupCommand(orek),
		addUserToGroupCommand(orek),
		removeUserFromGroupCommand(orek),
		getUsersInGroupCommand(orek),
		getGroupsForUserCommand(orek),
		clearValuesForVariableCommand(orek),
		getValuesForVariableCommand(orek),
		setPasswordCommand(orek),
		updatePasswordCommand(orek),
	}
	return cli.Command{
		Name:        "client",
		Subcommands: subcmds,
		Flags:       []cli.Flag{},
	}
}

func listUsersCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "list-users",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			var users []*data.User
			if isLocalMode(ctx) {
				users, err = data.GetDataStore().GetAllUsers()
			} else {
				users, err = orek.RestClient.GetAllUsers()
			}
			//Make it better
			if err == nil {
				for user := range users {
					fmt.Println(user)
				}
			}
			return err
		},
	}
	return cmd
}

func showUserCommand(orek *OrekApp) (cmd cli.Command) {
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
			argetr := ArgGetter{Ctx: ctx}
			userName := argetr.GetRequiredString("user-name")
			err = argetr.Err
			var user *data.User
			if err == nil {
				if isLocalMode(ctx) {
					user, err = data.GetDataStore().GetUser(userName)
				} else {
					user, err = orek.RestClient.GetUser(userName)
				}
			}
			if err == nil {
				fmt.Println(user)
			}
			return err
		},
	}
	return cmd
}

func showUserWithEmailCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			email := argetr.GetRequiredString("email")
			err = argetr.Err
			var user *data.User
			if err == nil {
				if isLocalMode(ctx) {
					user, err = data.GetDataStore().GetUserWithEmail(email)
				} else {
					user, err = orek.RestClient.GetUserWithEmail(email)
				}
			}
			if err == nil {
				fmt.Println(user)
			}
			return err
		},
	}
	return cmd
}

func createUserCommand(orek *OrekApp) (cmd cli.Command) {
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
				user := &data.User{
					Name:       userName,
					FirstName:  firstName,
					SecondName: secondName,
					Email:      email,
				}
				if isLocalMode(ctx) {
					err = data.GetDataStore().CreateUser(user)
				} else {
					err = orek.RestClient.CreateUser(user)
				}

			}
			return err
		},
	}
	return cmd
}

func updateUserCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "update-user",
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
				user := &data.User{
					Name:       userName,
					FirstName:  firstName,
					SecondName: secondName,
					Email:      email,
				}
				if isLocalMode(ctx) {
					err = data.GetDataStore().UpdateUser(user)
				} else {
					err = orek.RestClient.UpdateUser(user)
				}
			} else {
				err = argetr.Err
			}
			return err
		},
	}
	return cmd
}

func deleteUserCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "delete-user",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "user-name",
				Value: "",
				Usage: "The unique user_name for the user",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			userName := argetr.GetRequiredString("user-name")
			err = argetr.Err
			if err == nil {
				if isLocalMode(ctx) {
					err = data.GetDataStore().DeleteUser(userName)
				} else {
					err = orek.RestClient.DeleteUser(userName)
				}
			}
			return err
		},
	}
	return cmd
}

func listEndpointsCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "list-endpoint",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			var endpoints []*data.Endpoint
			if isLocalMode(ctx) {
				endpoints, err = data.GetDataStore().GetAllEndpoints()
			} else {
				endpoints, err = orek.RestClient.GetAllEndpoints()
			}
			if err == nil {
				for ep := range endpoints {
					fmt.Println(ep)
				}
			}
			return err
		},
	}
	return cmd
}

func showEndpointCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "show-endpoint",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "endpoint-id",
				Value: "",
				Usage: "The unique id of the endpoint",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			endpointID := argetr.GetRequiredString("endpoint-id")
			err = argetr.Err
			var ep *data.Endpoint
			if err == nil {
				if isLocalMode(ctx) {
					ep, err = data.GetDataStore().GetEndpoint(endpointID)
				} else {
					ep, err = orek.RestClient.GetEndpoint(endpointID)
				}
			}
			if err == nil {
				fmt.Println(ep)
			}
			return err
		},
	}
	return cmd
}

func createEndpointCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "create-endpoint",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "endpoint-id",
				Value: "",
				Usage: "The unique id of the endpoint",
			},
			cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: "Name of the endpoint",
			},
			cli.StringFlag{
				Name:  "owner",
				Value: "",
				Usage: "User who owns the endpoint",
			},
			cli.StringFlag{
				Name:  "owner-group",
				Value: "",
				Usage: "Group which owns the endpoint",
			},
			cli.StringFlag{
				Name:  "description",
				Value: "",
				Usage: "Description of the endpoint",
			},
			cli.StringFlag{
				Name:  "location",
				Value: "",
				Usage: "Loacation of the endpoint",
			},
			cli.StringFlag{
				Name:  "visibility",
				Value: "public",
				Usage: "Visibility of the endpoint",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			endpointID := argetr.GetRequiredString("endpoint-id")
			endpointName := argetr.GetRequiredString("name")
			owner := argetr.GetRequiredString("owner")
			group := argetr.GetString("owner-group")
			desc := argetr.GetString("description")
			location := argetr.GetString("location")
			visibility := argetr.GetString("visibility")
			endpoint := &data.Endpoint{
				EndpointID:  endpointID,
				Name:        endpointName,
				Owner:       owner,
				OwnerGroup:  group,
				Description: desc,
				Location:    location,
				Visibility:  data.EndpointVisiblity(visibility),
			}
			if argetr.Err == nil {
				if isLocalMode(ctx) {
					err = data.GetDataStore().CreateEndpoint(endpoint)
				} else {
					err = orek.RestClient.CreateEndpoint(endpoint)
				}
			}
			if err == nil {
				fmt.Print("Endpoint created")
			}
			return err
		},
	}
	return cmd
}

func updateEndpointCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "update-endpoint",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "endpoint-id",
				Value: "",
				Usage: "The unique id of the endpoint",
			},
			cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: "Name of the endpoint",
			},
			cli.StringFlag{
				Name:  "owner",
				Value: "",
				Usage: "User who owns the endpoint",
			},
			cli.StringFlag{
				Name:  "owner-group",
				Value: "",
				Usage: "Group which owns the endpoint",
			},
			cli.StringFlag{
				Name:  "description",
				Value: "",
				Usage: "Description of the endpoint",
			},
			cli.StringFlag{
				Name:  "location",
				Value: "",
				Usage: "Loacation of the endpoint",
			},
			cli.StringFlag{
				Name:  "visibility",
				Value: "public",
				Usage: "Visibility of the endpoint",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			endpointID := argetr.GetRequiredString("endpoint-id")
			endpointName := argetr.GetRequiredString("name")
			owner := argetr.GetRequiredString("owner")
			group := argetr.GetString("owner-group")
			desc := argetr.GetString("description")
			location := argetr.GetString("location")
			visibility := argetr.GetString("visibility")
			endpoint := &data.Endpoint{
				EndpointID:  endpointID,
				Name:        endpointName,
				Owner:       owner,
				OwnerGroup:  group,
				Description: desc,
				Location:    location,
				Visibility:  data.EndpointVisiblity(visibility),
			}
			if argetr.Err == nil {
				if isLocalMode(ctx) {
					err = data.GetDataStore().UpdateEndpoint(endpoint)
				} else {
					err = orek.RestClient.UpdateEndpoint(endpoint)
				}
			}
			if err == nil {
				fmt.Print("Endpoint updated")
			}
			return err
		},
	}
	return cmd
}

func deleteEndpointCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "delete-endpoint",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "endpoint-id",
				Value: "",
				Usage: "The unique id of the endpoint",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			endpointID := argetr.GetRequiredString("endpoint-id")
			err = argetr.Err
			if err == nil {
				if isLocalMode(ctx) {
					err = data.GetDataStore().DeleteEndpoint(endpointID)
				} else {
					err = orek.RestClient.DeleteEndpoint(endpointID)
				}
			}
			if err == nil {
				fmt.Println("Endpoint Deleted")
			}
			return err
		},
	}
	return cmd
}

func listVariablesCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			err = argetr.Err
			if err == nil {
				if isLocalMode(ctx) {
					// err = data.GetDataStore()
				} else {
					// err = orek.RestClient
				}
			}
			if err == nil {
				// fmt.Println("Endpoint Deleted")
			}
			return err
		},
	}
	return cmd
}

func listVariablesForEndpointCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func showVariableCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func createVariableCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func updateVariableCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func deleteVariableCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func listParametersCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func listParametersForEndpointCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func showParameterCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func createParameterCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func updateParameterCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func deleteParameterCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func listUserGroupsCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func showUserGroupCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func createUserGroupCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func updateUserGroupCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func deleteUserGroupCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func addUserToGroupCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func removeUserFromGroupCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func getUsersInGroupCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func getGroupsForUserCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func clearValuesForVariableCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func getValuesForVariableCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func setPasswordCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}

func updatePasswordCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			return err
		},
	}
	return cmd
}
