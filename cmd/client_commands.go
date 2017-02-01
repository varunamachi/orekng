package cmd

import (
	"errors"

	"github.com/varunamachi/orekng/data"
	"github.com/varunamachi/orekng/olog"
	"gopkg.in/urfave/cli.v1"
)

//ClientCommandProvider - Providers commands for running Orek app as client
type ClientCommandProvider struct {
	Client OrekClient
}

//GetCommand - gives commands for running Orek app as client to Orek Service
func (ccp *ClientCommandProvider) GetCommand() cli.Command {
	subcmds := []cli.Command{
		ccp.listUsersCommand(),
		ccp.showUserCommand(),
		ccp.showUserWithEmailCommand(),
		ccp.createUserCommand(),
		ccp.updateUserCommand(),
		ccp.deleteUserCommand(),
		ccp.listEndpointsCommand(),
		ccp.showEndpointCommand(),
		ccp.createEndpointCommand(),
		ccp.updateEndpointCommand(),
		ccp.deleteEndpointCommand(),
		ccp.listVariablesCommand(),
		ccp.listVariablesForEndpointCommand(),
		ccp.showVariableCommand(),
		ccp.createVariableCommand(),
		ccp.updateVariableCommand(),
		ccp.deleteVariableCommand(),
		ccp.listParametersCommand(),
		ccp.listParametersForEndpointCommand(),
		ccp.showParameterCommand(),
		ccp.createParameterCommand(),
		ccp.updateParameterCommand(),
		ccp.deleteParameterCommand(),
		ccp.listUserGroupsCommand(),
		ccp.showUserGroupCommand(),
		ccp.createUserGroupCommand(),
		ccp.updateUserGroupCommand(),
		ccp.deleteUserGroupCommand(),
		ccp.addUserToGroupCommand(),
		ccp.removeUserFromGroupCommand(),
		ccp.getUsersInGroupCommand(),
		ccp.getGroupsForUserCommand(),
		ccp.clearValuesForVariableCommand(),
		ccp.getValuesForVariableCommand(),
		ccp.setPasswordCommand(),
		ccp.updatePasswordCommand(),
	}
	return cli.Command{
		Name:        "client",
		Subcommands: subcmds,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "type",
				Value: "local",
				Usage: "Client type to use - Local Data Store or REST client",
			},
			cli.StringFlag{
				Name:  "api-host",
				Value: "localhost",
				Usage: "Host where the Orek server is running " +
					"(Only applicable to REST client)",
			},
			cli.IntFlag{
				Name:  "api-port",
				Value: 8000,
				Usage: "Port where the Orek server is running " +
					"(Only applicable to REST client)",
			},
			cli.StringFlag{
				Name:  "api-user",
				Value: "",
				Usage: "User name for using Orek service",
			},
			cli.StringFlag{
				Name:  "api-password",
				Value: "",
				Usage: "Orek password, use this falg only for testing",
			},
		},
		Before: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			clientType := argetr.GetRequiredString("type")
			if clientType == "local" {
				ccp.Client = &LocalClient{data.GetStore()}
			} else {
				olog.Error("Client", "REST client is not yet implemented")
			}
			return err
		},
		Usage: "Use orek as a client to Orek service or to DataStore",
	}
}

func (ccp *ClientCommandProvider) listUsersCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "list-users",
		Usage:    "Lists the user presetnt in the data store",
		Category: "User",
		Flags:    []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			var users []*data.User
			users, err = ccp.Client.GetAllUsers()
			//Make it better
			if err == nil {
				for _, user := range users {
					olog.Print("Client", "%v", user)
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) showUserCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "show-user",
		Usage:    "Shows details of a user identified by user-name ",
		Category: "User",
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
			var user *data.User
			if err = argetr.Err; err == nil {
				user, err = ccp.Client.GetUser(userName)
				if err == nil {
					olog.Print("Client", "%v", user)
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) showUserWithEmailCommand() (
	cmd cli.Command) {
	cmd = cli.Command{
		Name:     "user-with-email",
		Usage:    "Shows details of a user identified by user email ",
		Category: "User",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "email",
				Value: "",
				Usage: "User's unique email address",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			email := argetr.GetRequiredString("email")
			err = argetr.Err
			var user *data.User
			if err == nil {
				user, err = ccp.Client.GetUserWithEmail(email)
				if err == nil {
					olog.Print("Client", "%v", user)
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) createUserCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "create-user",
		Usage:    "Creates user with given details ",
		Category: "User",
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
			if err = argetr.Err; err == nil {
				user := &data.User{
					Name:       userName,
					FirstName:  firstName,
					SecondName: secondName,
					Email:      email,
				}
				err = ccp.Client.CreateUser(user)
				if err == nil {
					olog.Print("Client", "User created successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) updateUserCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "update-user",
		Usage:    "Updates a existing user record with new information",
		Category: "User",
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
			if err = argetr.Err; err == nil {
				user := &data.User{
					Name:       userName,
					FirstName:  firstName,
					SecondName: secondName,
					Email:      email,
				}
				err = ccp.Client.UpdateUser(user)
				if err == nil {
					olog.Print("Client", "User updated successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) deleteUserCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "delete-user",
		Usage:    "Removes a user record identified by user-name ",
		Category: "User",
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
			if err = argetr.Err; err == nil {
				err = ccp.Client.DeleteUser(userName)
				if err == nil {
					olog.Print("Client", "User deleted successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) listEndpointsCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "list-endpoints",
		Usage:    "Lists all the endpoints registered",
		Category: "Endpoint",
		Flags:    []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			var endpoints []*data.Endpoint
			endpoints, err = ccp.Client.GetAllEndpoints()
			if err == nil {
				for _, ep := range endpoints {
					olog.Print("Client", "%v", ep)
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) showEndpointCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "show-endpoint",
		Usage:    "Shows the details of endpoint identified by its ID",
		Category: "Endpoint",
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
				ep, err = ccp.Client.GetEndpoint(endpointID)
			}
			if err == nil {
				olog.Print("Client", "%v", ep)
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) createEndpointCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "create-endpoint",
		Usage:    "Creates an endpoint with given information",
		Category: "Endpoint",
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
				Visibility:  visibility,
			}
			if err = argetr.Err; err == nil {
				err = ccp.Client.CreateEndpoint(endpoint)
				if err == nil {
					olog.Print("Client", "Endpoint created")
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) updateEndpointCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "update-endpoint",
		Usage:    "Updates existing endpoints with new information",
		Category: "Endpoint",
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
				Visibility:  visibility,
			}
			if argetr.Err == nil {
				err = ccp.Client.UpdateEndpoint(endpoint)
			}
			if err == nil {
				olog.Print("Client", "Endpoint updated")
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) deleteEndpointCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "delete-endpoint",
		Usage:    "Deletes an endpoint identified by its ID",
		Category: "Endpoint",
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
			if err = argetr.Err; err == nil {
				err = ccp.Client.DeleteEndpoint(endpointID)
				if err == nil {
					olog.Print("Client", "Endpoint Deleted")
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) listVariablesCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "list-vars",
		Usage:    "Lists all the registered variables",
		Category: "Variable",
		Flags:    []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			if err = argetr.Err; err == nil {
				var variables []*data.Variable
				variables, err = ccp.Client.GetAllVariables()
				if err == nil {
					for _, vrb := range variables {
						olog.Print("Client", "%v", vrb)
					}
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) listVariablesForEndpointCommand() (
	cmd cli.Command) {
	cmd = cli.Command{
		Name:     "list-ep-vars",
		Usage:    "Lists variables registered with the given endpoint",
		Category: "Variable",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "endpoint-id",
				Value: "",
				Usage: "The unique ID of the endpoint",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			endpointID := argetr.GetRequiredString("endpoint-id")
			if err = argetr.Err; err == nil {
				var variables []*data.Variable
				variables, err = ccp.Client.GetVariablesForEndpoint(endpointID)
				if err == nil {
					for _, vrb := range variables {
						olog.Print("Client", "%v", vrb)
					}
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) showVariableCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "show-var",
		Usage:    "Displays information about a variable with given ID",
		Category: "Variable",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "variable-id",
				Value: "",
				Usage: "The unique ID of the variable",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			variableID := argetr.GetRequiredString("variable-id")
			if err = argetr.Err; err == nil {
				var variable *data.Variable
				variable, err = ccp.Client.GetVariable(variableID)
				if err == nil {
					olog.Print("Client", "%v", variable)
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) createVariableCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "create-var",
		Usage:    "Creates variable from given information",
		Category: "Variable",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "variable-id",
				Value: "",
				Usage: "Unique ID of the variable",
			},
			cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: "Name of the variable",
			},
			cli.StringFlag{
				Name:  "endpoint-id",
				Value: "",
				Usage: "Unique ID of the endpoint to which variable belongs to",
			},
			cli.StringFlag{
				Name:  "description",
				Value: "",
				Usage: "Description of the variable",
			},
			cli.StringFlag{
				Name:  "unit",
				Value: "",
				Usage: "Unit in which variable is measured",
			},
			cli.StringFlag{
				Name:  "type",
				Value: "",
				Usage: "Data type of the variable",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			variableID := argetr.GetRequiredString("variable-id")
			name := argetr.GetRequiredString("name")
			endpointID := argetr.GetRequiredString("endpoint-id")
			description := argetr.GetRequiredString("description")
			unit := argetr.GetRequiredString("unit")
			variableType := argetr.GetRequiredString("type")
			if err = argetr.Err; err == nil {
				variable := &data.Variable{
					VariableID:  variableID,
					Name:        name,
					EndpointID:  endpointID,
					Description: description,
					Unit:        unit,
					Type:        variableType,
				}
				err = ccp.Client.CreateVariable(variable)
				if err == nil {
					olog.Print("Client", "%v", "Variable created successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) updateVariableCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "update-var",
		Usage:    "Updates the existing variable with new information",
		Category: "Variable",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "variable-id",
				Value: "",
				Usage: "Unique ID of the variable",
			},
			cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: "Name of the variable",
			},
			cli.StringFlag{
				Name:  "endpoint-id",
				Value: "",
				Usage: "Unique ID of the endpoint to which variable belongs to",
			},
			cli.StringFlag{
				Name:  "description",
				Value: "",
				Usage: "Description of the variable",
			},
			cli.StringFlag{
				Name:  "unit",
				Value: "",
				Usage: "Unit in which variable is measured",
			},
			cli.StringFlag{
				Name:  "type",
				Value: "",
				Usage: "Data type of the variable",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			variableID := argetr.GetRequiredString("variable-id")
			name := argetr.GetRequiredString("name")
			endpointID := argetr.GetRequiredString("endpoint-id")
			description := argetr.GetRequiredString("description")
			unit := argetr.GetRequiredString("unit")
			variableType := argetr.GetRequiredString("type")
			if err = argetr.Err; err == nil {
				variable := &data.Variable{
					VariableID:  variableID,
					Name:        name,
					EndpointID:  endpointID,
					Description: description,
					Unit:        unit,
					Type:        variableType,
				}
				err = ccp.Client.UpdateVariable(variable)
				if err == nil {
					olog.Print("Client", "%v", "Variable updated successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) deleteVariableCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "delete-var",
		Usage:    "Deletes a variable with given ID",
		Category: "Variable",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "variable-id",
				Value: "",
				Usage: "The unique ID of the variable",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			variableID := argetr.GetRequiredString("variable-id")
			if err = argetr.Err; err == nil {
				err = ccp.Client.DeleteVariable(variableID)
				if err == nil {
					olog.Print("Client", "Variable deleted successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) listUserGroupsCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "list-groups",
		Usage:    "Lists all the user groups present in the data store",
		Category: "Group",
		Flags:    []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			if err = argetr.Err; err == nil {
				var groups []*data.UserGroup
				groups, err = ccp.Client.GetAllUserGroups()
				if err == nil {
					for _, group := range groups {
						olog.Print("Client", "%v", group)
					}
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) showUserGroupCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "show-group",
		Usage:    "Shows a group identified by the given ID",
		Category: "Group",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "group-id",
				Value: "",
				Usage: "The unique ID of the user group",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			groupID := argetr.GetRequiredString("group-id")
			if err = argetr.Err; err == nil {
				var group *data.UserGroup
				group, err = ccp.Client.GetUserGroup(groupID)
				if err == nil {
					olog.Print("Client", "%v", group)
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) createUserGroupCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "create-group",
		Usage:    "Creates a user group with given information",
		Category: "Group",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "group-id",
				Value: "",
				Usage: "Unique ID of the group",
			},
			cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: "Name of the group",
			},
			cli.StringFlag{
				Name:  "owner",
				Value: "",
				Usage: "Owner of the group",
			},
			cli.StringFlag{
				Name:  "description",
				Value: "",
				Usage: "Description of the group",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			groupID := argetr.GetRequiredString("group-id")
			name := argetr.GetRequiredString("name")
			owner := argetr.GetRequiredString("owner")
			description := argetr.GetString("description")
			err = argetr.Err
			if err == nil {
				userGroup := &data.UserGroup{
					GroupID:     groupID,
					Name:        name,
					Owner:       owner,
					Description: description,
				}
				err = ccp.Client.CreateUserGroup(userGroup)
				if err == nil {
					olog.Print("Client", "User group created successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) updateUserGroupCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "update-group",
		Usage:    "Updates a existing group with new information",
		Category: "Group",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "group-id",
				Value: "",
				Usage: "Unique ID of the group",
			},
			cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: "Name of the group",
			},
			cli.StringFlag{
				Name:  "owner",
				Value: "",
				Usage: "Owner of the group",
			},
			cli.StringFlag{
				Name:  "description",
				Value: "",
				Usage: "Description of the group",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			groupID := argetr.GetRequiredString("group-id")
			name := argetr.GetRequiredString("name")
			owner := argetr.GetRequiredString("owner")
			description := argetr.GetString("description")
			err = argetr.Err
			if err == nil {
				userGroup := &data.UserGroup{
					GroupID:     groupID,
					Name:        name,
					Owner:       owner,
					Description: description,
				}
				err = ccp.Client.UpdateUserGroup(userGroup)
				if err == nil {
					olog.Print("Client", "User group updated successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) deleteUserGroupCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "delete-group",
		Usage:    "Deletes a user group identified by the given ID",
		Category: "Group",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "group-id",
				Value: "",
				Usage: "The unique ID of the user group",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			groupID := argetr.GetRequiredString("group-id")
			if err = argetr.Err; err == nil {
				err = ccp.Client.DeleteUserGroup(groupID)
				if err == nil {
					olog.Print("Client", "User group deleted successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) addUserToGroupCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name: "link-group-user",
		Usage: "Adds the user identified by user-name to group " +
			"identified by groupID",
		Category: "Group",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "user-name",
				Value: "",
				Usage: "Unique name of an existing user",
			},
			cli.StringFlag{
				Name:  "group-id",
				Value: "",
				Usage: "Unique ID of an existing user group",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			userName := argetr.GetRequiredString("user-name")
			groupID := argetr.GetRequiredString("group-id")
			if err = argetr.Err; err == nil {
				err = ccp.Client.AddUserToGroup(userName, groupID)
				if err == nil {
					olog.Print("Client", "User added to group")
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) removeUserFromGroupCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name: "unlink-group-user",
		Usage: "Removes a user identified by user-name to group " +
			"identified by groupID",
		Category: "Group",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "user-name",
				Value: "",
				Usage: "Unique name of an existing user",
			},
			cli.StringFlag{
				Name:  "group-id",
				Value: "",
				Usage: "Unique ID of an existing user group",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			userName := argetr.GetRequiredString("user-name")
			groupID := argetr.GetRequiredString("group-id")
			if err = argetr.Err; err == nil {
				err = ccp.Client.RemoveUserFromGroup(userName, groupID)
				if err == nil {
					olog.Print("Client", "User removed from group")
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) getUsersInGroupCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "users-in-group",
		Usage:    "Lists the users present in the group with given group ID",
		Category: "Group",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "group-id",
				Value: "",
				Usage: "Unique ID of an existing user group",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			groupID := argetr.GetRequiredString("group-id")
			if err = argetr.Err; err == nil {
				var users []*data.User
				users, err = ccp.Client.GetUsersInGroup(groupID)
				if err == nil {
					if len(users) == 0 {
						olog.Print("Client", "No users found")
					}
					for _, user := range users {
						olog.Print("Client", "%v", user.Name)
					}
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) getGroupsForUserCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "groups-of-user",
		Usage:    "Lists the groups with which the user with given ID is associated",
		Category: "Group",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "user-name",
				Value: "",
				Usage: "Unique name of an existing user",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			userName := argetr.GetRequiredString("user-name")
			if err = argetr.Err; err == nil {
				var groups []*data.UserGroup
				groups, err = ccp.Client.GetGroupsForUser(userName)
				if err == nil {
					for _, group := range groups {
						olog.Print("Client", "%v", group.GroupID)
					}
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) clearValuesForVariableCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "clear-vars",
		Usage:    "Clears the values stored for the variable identified by the ID",
		Category: "Variable",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "variable-id",
				Value: "",
				Usage: "Unique ID of the variable",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			variableID := argetr.GetRequiredString("variable-id")
			err = argetr.Err
			if err == nil {
				err = ccp.Client.ClearValuesForVariable(variableID)
				if err == nil {
					olog.Print("Client", "All variable values are cleared")
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) getValuesForVariableCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "var-values",
		Usage:    "Lists the values stored for the variable identified by the ID",
		Category: "Variable",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "variable-id",
				Value: "",
				Usage: "Unique ID of the variable",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			variableID := argetr.GetRequiredString("variable-id")
			err = argetr.Err
			if err == nil {
				var values []*string
				values, err = ccp.Client.GetValuesForVariable(variableID)
				if err == nil {
					for _, val := range values {
						olog.Print("Client", "%v", val)
					}
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) setPasswordCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "set-password",
		Usage:    "Sets the password for an user with given ID",
		Category: "User",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "user-name",
				Value: "",
				Usage: "Unique name of an existing user",
			},
			cli.StringFlag{
				Name:  "password",
				Value: "",
				Usage: "Password for the user",
			},
			cli.StringFlag{
				Name:  "confirm-password",
				Value: "",
				Usage: "New Password for the user",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			userName := argetr.GetRequiredString("user-name")
			password := argetr.GetRequiredString("password")
			confirmPassword := argetr.GetRequiredString("confirm-password")
			if err = argetr.Err; err == nil {
				if password == confirmPassword {
					err = ccp.Client.SetPassword(userName, password)
					if err == nil {
						olog.Print("Client", "Password set successfully")
					}
				} else {
					err = errors.New("Password dont match")
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) updatePasswordCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "update-password",
		Usage:    "Updates the password for an user with given ID",
		Category: "User",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "user-name",
				Value: "",
				Usage: "Unique name of an existing user",
			},
			cli.StringFlag{
				Name:  "current-password",
				Value: "",
				Usage: "Current Password for the user",
			},
			cli.StringFlag{
				Name:  "new-password",
				Value: "",
				Usage: "New Password for the user",
			},
			cli.StringFlag{
				Name:  "confirm-password",
				Value: "",
				Usage: "New Password for the user",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			userName := argetr.GetRequiredString("user-name")
			currentPassword := argetr.GetRequiredString("current-password")
			newPassword := argetr.GetRequiredString("new-password")
			confirmPassword := argetr.GetRequiredString("confirm-password")
			if err = argetr.Err; err == nil {
				if confirmPassword == newPassword {
					err = ccp.Client.UpdatePassword(userName,
						currentPassword, newPassword)
					if err == nil {
						olog.Print("Client", "Password set successfully")
					}
				} else {
					err = errors.New("Password dont match")
				}
			}
			return err
		},
	}
	return cmd
}

//////////////////////////////////
func (ccp *ClientCommandProvider) listParametersCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "list-params",
		Usage:    "Lists all the registered parameters",
		Category: "Parameter",
		Flags:    []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			if err = argetr.Err; err == nil {
				var parameters []*data.Parameter
				parameters, err = ccp.Client.GetAllParameters()
				if err == nil {
					for _, vrb := range parameters {
						olog.Print("Client", "%v", vrb)
					}
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) listParametersForEndpointCommand() (
	cmd cli.Command) {
	cmd = cli.Command{
		Name:     "list-ep-params",
		Usage:    "Lists all the parameters registered with endpoint of given ID",
		Category: "Parameter",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "endpoint-id",
				Value: "",
				Usage: "The unique ID of the endpoint",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			endpointID := argetr.GetRequiredString("endpoint-id")
			if err = argetr.Err; err == nil {
				var parameters []*data.Parameter
				parameters, err = ccp.Client.GetParametersForEndpoint(endpointID)
				if err == nil {
					for _, vrb := range parameters {
						olog.Print("Client", "%v", vrb)
					}
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) showParameterCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "show-param",
		Usage:    "Shows the details of the parameter with given ID",
		Category: "Parameter",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "parameter-id",
				Value: "",
				Usage: "The unique ID of the parameter",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			parameterID := argetr.GetRequiredString("parameter-id")
			if err = argetr.Err; err == nil {
				var parameter *data.Parameter
				parameter, err = ccp.Client.GetParameter(parameterID)
				if err == nil {
					olog.Print("Client", "%v", parameter)
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) createParameterCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "create-param",
		Usage:    "Creates parameter with given information",
		Category: "Parameter",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "parameter-id",
				Value: "",
				Usage: "Unique ID of the parameter",
			},
			cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: "Name of the parameter",
			},
			cli.StringFlag{
				Name:  "endpoint-id",
				Value: "",
				Usage: "Unique ID of the endpoint to which parameter belongs to",
			},
			cli.StringFlag{
				Name:  "description",
				Value: "",
				Usage: "Description of the parameter",
			},
			cli.StringFlag{
				Name:  "unit",
				Value: "",
				Usage: "Unit in which parameter is measured",
			},
			cli.StringFlag{
				Name:  "type",
				Value: "",
				Usage: "Data type of the parameter",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			parameterID := argetr.GetRequiredString("parameter-id")
			name := argetr.GetRequiredString("name")
			endpointID := argetr.GetRequiredString("endpoint-id")
			description := argetr.GetRequiredString("description")
			unit := argetr.GetRequiredString("unit")
			parameterType := argetr.GetRequiredString("type")
			if err = argetr.Err; err == nil {
				parameter := &data.Parameter{
					ParameterID: parameterID,
					Name:        name,
					EndpointID:  endpointID,
					Description: description,
					Unit:        unit,
					Type:        parameterType,
				}
				err = ccp.Client.CreateParameter(parameter)
				if err == nil {
					olog.Print("Client", "Parameter created successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) updateParameterCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "update-param",
		Usage:    "Updates an existing parameter with new information",
		Category: "Parameter",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "parameter-id",
				Value: "",
				Usage: "Unique ID of the parameter",
			},
			cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: "Name of the parameter",
			},
			cli.StringFlag{
				Name:  "endpoint-id",
				Value: "",
				Usage: "Unique ID of the endpoint to which parameter belongs to",
			},
			cli.StringFlag{
				Name:  "description",
				Value: "",
				Usage: "Description of the parameter",
			},
			cli.StringFlag{
				Name:  "unit",
				Value: "",
				Usage: "Unit in which parameter is measured",
			},
			cli.StringFlag{
				Name:  "type",
				Value: "",
				Usage: "Data type of the parameter",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			parameterID := argetr.GetRequiredString("parameter-id")
			name := argetr.GetRequiredString("name")
			endpointID := argetr.GetRequiredString("endpoint-id")
			description := argetr.GetRequiredString("description")
			unit := argetr.GetRequiredString("unit")
			parameterType := argetr.GetRequiredString("type")
			if err = argetr.Err; err == nil {
				parameter := &data.Parameter{
					ParameterID: parameterID,
					Name:        name,
					EndpointID:  endpointID,
					Description: description,
					Unit:        unit,
					Type:        parameterType,
				}
				err = ccp.Client.UpdateParameter(parameter)
				if err == nil {
					olog.Print("Client", "Parameter updated successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func (ccp *ClientCommandProvider) deleteParameterCommand() (cmd cli.Command) {
	cmd = cli.Command{
		Name:     "delete-param",
		Usage:    "Deletes a parameter identified by the given ID",
		Category: "Parameter",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "parameter-id",
				Value: "",
				Usage: "The unique ID of the parameter",
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			parameterID := argetr.GetRequiredString("parameter-id")
			if err = argetr.Err; err == nil {
				err = ccp.Client.DeleteParameter(parameterID)
				if err == nil {
					olog.Print("Client", "Parameter deleted successfully")
				}
			}
			return err
		},
	}
	return cmd
}
