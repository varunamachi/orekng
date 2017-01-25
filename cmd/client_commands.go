package cmd

import (
	"errors"
	"fmt"

	"github.com/varunamachi/orekng/data"
	"gopkg.in/urfave/cli.v1"
)

//ClientCommandProvider - Providers commands for running Orek app as client
type ClientCommandProvider struct{}

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
			users, err = orek.Client.GetAllUsers()
			//Make it better
			if err == nil {
				for _, user := range users {
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
			var user *data.User
			if err = argetr.Err; err == nil {
				user, err = orek.Client.GetUser(userName)
				if err == nil {
					fmt.Println(user)
				}
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
				user, err = orek.Client.GetUserWithEmail(email)
				if err == nil {
					fmt.Println(user)
				}
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
			if err = argetr.Err; err == nil {
				user := &data.User{
					Name:       userName,
					FirstName:  firstName,
					SecondName: secondName,
					Email:      email,
				}
				err = orek.Client.CreateUser(user)
				if err == nil {
					fmt.Println("User created successfully")
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
			if err = argetr.Err; err == nil {
				user := &data.User{
					Name:       userName,
					FirstName:  firstName,
					SecondName: secondName,
					Email:      email,
				}
				err = orek.Client.UpdateUser(user)
				if err == nil {
					fmt.Println("User updated successfully")
				}
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
			if err = argetr.Err; err == nil {
				err = orek.Client.DeleteUser(userName)
				if err == nil {
					fmt.Println("User deleted successfully")
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
			endpoints, err = orek.Client.GetAllEndpoints()
			if err == nil {
				for _, ep := range endpoints {
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
				ep, err = orek.Client.GetEndpoint(endpointID)
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
			if err = argetr.Err; err == nil {
				err = orek.Client.CreateEndpoint(endpoint)
				if err == nil {
					fmt.Print("Endpoint created")
				}
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
				err = orek.Client.UpdateEndpoint(endpoint)
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
			if err = argetr.Err; err == nil {
				err = orek.Client.DeleteEndpoint(endpointID)
				if err == nil {
					fmt.Println("Endpoint Deleted")
				}
			}
			return err
		},
	}
	return cmd
}

func listVariablesCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "list-vars",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			if err = argetr.Err; err == nil {
				var variables []*data.Variable
				variables, err = orek.Client.GetAllVariables()
				if err == nil {
					for _, vrb := range variables {
						fmt.Println(vrb)
					}
				}
			}
			return err
		},
	}
	return cmd
}

func listVariablesForEndpointCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "list-ep-vars",
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
				variables, err = orek.Client.GetVariablesForEndpoint(endpointID)
				if err == nil {
					for _, vrb := range variables {
						fmt.Println(vrb)
					}
				}
			}
			return err
		},
	}
	return cmd
}

func showVariableCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "show-var",
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
				variable, err = orek.Client.GetVariable(variableID)
				if err == nil {
					fmt.Println(variable)
				}
			}
			return err
		},
	}
	return cmd
}

func createVariableCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "create-var",
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
				err = orek.Client.CreateVariable(variable)
				if err == nil {
					fmt.Println("Variable created successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func updateVariableCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "update-var",
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
				err = orek.Client.UpdateVariable(variable)
				if err == nil {
					fmt.Println("Variable updated successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func deleteVariableCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "delete-var",
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
				err = orek.Client.DeleteVariable(variableID)
				if err == nil {
					fmt.Println("Variable deleted successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func listUserGroupsCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "list-groups",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			if err = argetr.Err; err == nil {
				var groups []*data.UserGroup
				groups, err = orek.Client.GetAllUserGroups()
				if err == nil {
					for _, group := range groups {
						fmt.Println(group)
					}
				}
			}
			return err
		},
	}
	return cmd
}

func showUserGroupCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "show-group",
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
				group, err = orek.Client.GetUserGroup(groupID)
				if err == nil {
					fmt.Println(group)
				}
			}
			return err
		},
	}
	return cmd
}

func createUserGroupCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "create-group",
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
				err = orek.Client.CreateUserGroup(userGroup)
				if err == nil {
					fmt.Println("User group created successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func updateUserGroupCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "update-group",
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
				err = orek.Client.UpdateUserGroup(userGroup)
				if err == nil {
					fmt.Println("User group updated successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func deleteUserGroupCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "delete-group",
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
				err = orek.Client.DeleteUserGroup(groupID)
				if err == nil {
					fmt.Println("User group deleted successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func addUserToGroupCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "group-user-link",
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
				err = orek.Client.AddUserToGroup(userName, groupID)
				if err == nil {
					fmt.Println("User added to group")
				}
			}
			return err
		},
	}
	return cmd
}

func removeUserFromGroupCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "group-user-unlink",
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
				err = orek.Client.RemoveUserFromGroup(userName, groupID)
				if err == nil {
					fmt.Println("User removed from group")
				}
			}
			return err
		},
	}
	return cmd
}

func getUsersInGroupCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "users-in-group",
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
				users, err = orek.Client.GetUsersInGroup(groupID)
				if err == nil {
					for _, user := range users {
						fmt.Println(user.Name)
					}
				}
			}
			return err
		},
	}
	return cmd
}

func getGroupsForUserCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "groups-of-user",
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
				groups, err = orek.Client.GetGroupsForUser(userName)
				if err == nil {
					for _, group := range groups {
						fmt.Println(group.GroupID)
					}
				}
			}
			return err
		},
	}
	return cmd
}

func clearValuesForVariableCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "clear-vars",
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
				err = orek.Client.ClearValuesForVariable(variableID)
				if err == nil {
					fmt.Println("All variable values are cleared")
				}
			}
			return err
		},
	}
	return cmd
}

func getValuesForVariableCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "var-values",
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
				values, err = orek.Client.GetValuesForVariable(variableID)
				if err == nil {
					for _, val := range values {
						fmt.Println(val)
					}
				}
			}
			return err
		},
	}
	return cmd
}

func setPasswordCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "set-password",
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
					err = orek.Client.SetPassword(userName, password)
					if err == nil {
						fmt.Println("Password set successfully")
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

func updatePasswordCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "set-password",
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
					err = orek.Client.UpdatePassword(userName,
						currentPassword, newPassword)
					if err == nil {
						fmt.Println("Password set successfully")
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
func listParametersCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "list-params",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			if err = argetr.Err; err == nil {
				var parameters []*data.Parameter
				parameters, err = orek.Client.GetAllParameters()
				if err == nil {
					for _, vrb := range parameters {
						fmt.Println(vrb)
					}
				}
			}
			return err
		},
	}
	return cmd
}

func listParametersForEndpointCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "list-ep-params",
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
				parameters, err = orek.Client.GetParametersForEndpoint(endpointID)
				if err == nil {
					for _, vrb := range parameters {
						fmt.Println(vrb)
					}
				}
			}
			return err
		},
	}
	return cmd
}

func showParameterCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "show-param",
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
				parameter, err = orek.Client.GetParameter(parameterID)
				if err == nil {
					fmt.Println(parameter)
				}
			}
			return err
		},
	}
	return cmd
}

func createParameterCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "create-param",
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
				err = orek.Client.CreateParameter(parameter)
				if err == nil {
					fmt.Println("Parameter created successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func updateParameterCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "update-param",
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
				err = orek.Client.UpdateParameter(parameter)
				if err == nil {
					fmt.Println("Parameter updated successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func deleteParameterCommand(orek *OrekApp) (cmd cli.Command) {
	cmd = cli.Command{
		Name: "delete-param",
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
				err = orek.Client.DeleteParameter(parameterID)
				if err == nil {
					fmt.Println("Parameter deleted successfully")
				}
			}
			return err
		},
	}
	return cmd
}
