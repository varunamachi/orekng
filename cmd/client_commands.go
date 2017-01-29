package cmd

import (
	"errors"
	"log"

	"github.com/varunamachi/orekng/data"
	"gopkg.in/urfave/cli.v1"
)

//ClientCommandProvider - Providers commands for running Orek app as client
type ClientCommandProvider struct{}

//GetCommand - gives commands for running Orek app as client to Orek Service
func (ccp *ClientCommandProvider) GetCommand() cli.Command {
	client := &LocalClient{data.GetStore()}
	subcmds := []cli.Command{
		listUsersCommand(client),
		showUserCommand(client),
		showUserWithEmailCommand(client),
		createUserCommand(client),
		updateUserCommand(client),
		deleteUserCommand(client),
		listEndpointsCommand(client),
		showEndpointCommand(client),
		createEndpointCommand(client),
		updateEndpointCommand(client),
		deleteEndpointCommand(client),
		listVariablesCommand(client),
		listVariablesForEndpointCommand(client),
		showVariableCommand(client),
		createVariableCommand(client),
		updateVariableCommand(client),
		deleteVariableCommand(client),
		listParametersCommand(client),
		listParametersForEndpointCommand(client),
		showParameterCommand(client),
		createParameterCommand(client),
		updateParameterCommand(client),
		deleteParameterCommand(client),
		listUserGroupsCommand(client),
		showUserGroupCommand(client),
		createUserGroupCommand(client),
		updateUserGroupCommand(client),
		deleteUserGroupCommand(client),
		addUserToGroupCommand(client),
		removeUserFromGroupCommand(client),
		getUsersInGroupCommand(client),
		getGroupsForUserCommand(client),
		clearValuesForVariableCommand(client),
		getValuesForVariableCommand(client),
		setPasswordCommand(client),
		updatePasswordCommand(client),
	}
	return cli.Command{
		Name:        "client",
		Subcommands: subcmds,
		Flags:       []cli.Flag{},
	}
}

func listUsersCommand(client OrekClient) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "list-users",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			var users []*data.User
			users, err = client.GetAllUsers()
			//Make it better
			if err == nil {
				for _, user := range users {
					log.Println(user)
				}
			}
			return err
		},
	}
	return cmd
}

func showUserCommand(client OrekClient) (cmd cli.Command) {
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
				user, err = client.GetUser(userName)
				if err == nil {
					log.Println(user)
				}
			}
			return err
		},
	}
	return cmd
}

func showUserWithEmailCommand(client OrekClient) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			email := argetr.GetRequiredString("email")
			err = argetr.Err
			var user *data.User
			if err == nil {
				user, err = client.GetUserWithEmail(email)
				if err == nil {
					log.Println(user)
				}
			}
			return err
		},
	}
	return cmd
}

func createUserCommand(client OrekClient) (cmd cli.Command) {
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
				err = client.CreateUser(user)
				if err == nil {
					log.Println("User created successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func updateUserCommand(client OrekClient) (cmd cli.Command) {
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
				err = client.UpdateUser(user)
				if err == nil {
					log.Println("User updated successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func deleteUserCommand(client OrekClient) (cmd cli.Command) {
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
				err = client.DeleteUser(userName)
				if err == nil {
					log.Println("User deleted successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func listEndpointsCommand(client OrekClient) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "list-endpoint",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			var endpoints []*data.Endpoint
			endpoints, err = client.GetAllEndpoints()
			if err == nil {
				for _, ep := range endpoints {
					log.Println(ep)
				}
			}
			return err
		},
	}
	return cmd
}

func showEndpointCommand(client OrekClient) (cmd cli.Command) {
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
				ep, err = client.GetEndpoint(endpointID)
			}
			if err == nil {
				log.Println(ep)
			}
			return err
		},
	}
	return cmd
}

func createEndpointCommand(client OrekClient) (cmd cli.Command) {
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
				err = client.CreateEndpoint(endpoint)
				if err == nil {
					log.Println("Endpoint created")
				}
			}
			return err
		},
	}
	return cmd
}

func updateEndpointCommand(client OrekClient) (cmd cli.Command) {
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
				err = client.UpdateEndpoint(endpoint)
			}
			if err == nil {
				log.Println("Endpoint updated")
			}
			return err
		},
	}
	return cmd
}

func deleteEndpointCommand(client OrekClient) (cmd cli.Command) {
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
				err = client.DeleteEndpoint(endpointID)
				if err == nil {
					log.Println("Endpoint Deleted")
				}
			}
			return err
		},
	}
	return cmd
}

func listVariablesCommand(client OrekClient) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "list-vars",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			if err = argetr.Err; err == nil {
				var variables []*data.Variable
				variables, err = client.GetAllVariables()
				if err == nil {
					for _, vrb := range variables {
						log.Println(vrb)
					}
				}
			}
			return err
		},
	}
	return cmd
}

func listVariablesForEndpointCommand(client OrekClient) (cmd cli.Command) {
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
				variables, err = client.GetVariablesForEndpoint(endpointID)
				if err == nil {
					for _, vrb := range variables {
						log.Println(vrb)
					}
				}
			}
			return err
		},
	}
	return cmd
}

func showVariableCommand(client OrekClient) (cmd cli.Command) {
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
				variable, err = client.GetVariable(variableID)
				if err == nil {
					log.Println(variable)
				}
			}
			return err
		},
	}
	return cmd
}

func createVariableCommand(client OrekClient) (cmd cli.Command) {
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
				err = client.CreateVariable(variable)
				if err == nil {
					log.Println("Variable created successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func updateVariableCommand(client OrekClient) (cmd cli.Command) {
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
				err = client.UpdateVariable(variable)
				if err == nil {
					log.Println("Variable updated successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func deleteVariableCommand(client OrekClient) (cmd cli.Command) {
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
				err = client.DeleteVariable(variableID)
				if err == nil {
					log.Println("Variable deleted successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func listUserGroupsCommand(client OrekClient) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "list-groups",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			if err = argetr.Err; err == nil {
				var groups []*data.UserGroup
				groups, err = client.GetAllUserGroups()
				if err == nil {
					for _, group := range groups {
						log.Println(group)
					}
				}
			}
			return err
		},
	}
	return cmd
}

func showUserGroupCommand(client OrekClient) (cmd cli.Command) {
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
				group, err = client.GetUserGroup(groupID)
				if err == nil {
					log.Println(group)
				}
			}
			return err
		},
	}
	return cmd
}

func createUserGroupCommand(client OrekClient) (cmd cli.Command) {
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
				err = client.CreateUserGroup(userGroup)
				if err == nil {
					log.Println("User group created successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func updateUserGroupCommand(client OrekClient) (cmd cli.Command) {
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
				err = client.UpdateUserGroup(userGroup)
				if err == nil {
					log.Println("User group updated successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func deleteUserGroupCommand(client OrekClient) (cmd cli.Command) {
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
				err = client.DeleteUserGroup(groupID)
				if err == nil {
					log.Println("User group deleted successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func addUserToGroupCommand(client OrekClient) (cmd cli.Command) {
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
				err = client.AddUserToGroup(userName, groupID)
				if err == nil {
					log.Println("User added to group")
				}
			}
			return err
		},
	}
	return cmd
}

func removeUserFromGroupCommand(client OrekClient) (cmd cli.Command) {
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
				err = client.RemoveUserFromGroup(userName, groupID)
				if err == nil {
					log.Println("User removed from group")
				}
			}
			return err
		},
	}
	return cmd
}

func getUsersInGroupCommand(client OrekClient) (cmd cli.Command) {
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
				users, err = client.GetUsersInGroup(groupID)
				if err == nil {
					for _, user := range users {
						log.Println(user.Name)
					}
				}
			}
			return err
		},
	}
	return cmd
}

func getGroupsForUserCommand(client OrekClient) (cmd cli.Command) {
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
				groups, err = client.GetGroupsForUser(userName)
				if err == nil {
					for _, group := range groups {
						log.Println(group.GroupID)
					}
				}
			}
			return err
		},
	}
	return cmd
}

func clearValuesForVariableCommand(client OrekClient) (cmd cli.Command) {
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
				err = client.ClearValuesForVariable(variableID)
				if err == nil {
					log.Println("All variable values are cleared")
				}
			}
			return err
		},
	}
	return cmd
}

func getValuesForVariableCommand(client OrekClient) (cmd cli.Command) {
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
				values, err = client.GetValuesForVariable(variableID)
				if err == nil {
					for _, val := range values {
						log.Println(val)
					}
				}
			}
			return err
		},
	}
	return cmd
}

func setPasswordCommand(client OrekClient) (cmd cli.Command) {
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
					err = client.SetPassword(userName, password)
					if err == nil {
						log.Println("Password set successfully")
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

func updatePasswordCommand(client OrekClient) (cmd cli.Command) {
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
					err = client.UpdatePassword(userName,
						currentPassword, newPassword)
					if err == nil {
						log.Println("Password set successfully")
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
func listParametersCommand(client OrekClient) (cmd cli.Command) {
	cmd = cli.Command{
		Name:  "list-params",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) (err error) {
			argetr := ArgGetter{Ctx: ctx}
			if err = argetr.Err; err == nil {
				var parameters []*data.Parameter
				parameters, err = client.GetAllParameters()
				if err == nil {
					for _, vrb := range parameters {
						log.Println(vrb)
					}
				}
			}
			return err
		},
	}
	return cmd
}

func listParametersForEndpointCommand(client OrekClient) (cmd cli.Command) {
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
				parameters, err = client.GetParametersForEndpoint(endpointID)
				if err == nil {
					for _, vrb := range parameters {
						log.Println(vrb)
					}
				}
			}
			return err
		},
	}
	return cmd
}

func showParameterCommand(client OrekClient) (cmd cli.Command) {
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
				parameter, err = client.GetParameter(parameterID)
				if err == nil {
					log.Println(parameter)
				}
			}
			return err
		},
	}
	return cmd
}

func createParameterCommand(client OrekClient) (cmd cli.Command) {
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
				err = client.CreateParameter(parameter)
				if err == nil {
					log.Println("Parameter created successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func updateParameterCommand(client OrekClient) (cmd cli.Command) {
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
				err = client.UpdateParameter(parameter)
				if err == nil {
					log.Println("Parameter updated successfully")
				}
			}
			return err
		},
	}
	return cmd
}

func deleteParameterCommand(client OrekClient) (cmd cli.Command) {
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
				err = client.DeleteParameter(parameterID)
				if err == nil {
					log.Println("Parameter deleted successfully")
				}
			}
			return err
		},
	}
	return cmd
}
