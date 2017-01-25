package cmd

import (
	"fmt"

	"github.com/varunamachi/orekng/data"
	cli "gopkg.in/urfave/cli.v1"
)

//OrekClient - Defines a client impl used to execute commands from cli
type OrekClient interface {
	GetAllUsers() (users []*data.User, err error)
	GetUser(userName string) (users *data.User, err error)
	GetUserWithEmail(email string) (user *data.User, err error)
	CreateUser(user *data.User) (err error)
	UpdateUser(user *data.User) (err error)
	DeleteUser(userName string) (err error)
	GetAllEndpoints() (endpoints []*data.Endpoint, err error)
	GetEndpoint(endpointID string) (endpoint *data.Endpoint, err error)
	CreateEndpoint(endpoint *data.Endpoint) (err error)
	UpdateEndpoint(endpoint *data.Endpoint) (err error)
	DeleteEndpoint(endpointID string) (err error)
	GetAllVariables() (variables []*data.Variable, err error)
	GetVariablesForEndpoint(endpointID string) (variables []*data.Variable, err error)
	GetVariable(variableID string) (variable *data.Variable, err error)
	CreateVariable(variable *data.Variable) (err error)
	UpdateVariable(variable *data.Variable) (err error)
	DeleteVariable(variableID string) (err error)
	GetAllParameters() (parameters []*data.Parameter, err error)
	GetParametersForEndpoint(endpointID string) (parameters []*data.Parameter, err error)
	GetParameter(parameterID string) (parameter *data.Parameter, err error)
	CreateParameter(parameter *data.Parameter) (err error)
	UpdateParameter(parameter *data.Parameter) (err error)
	DeleteParameter(parameterID string) (err error)
	GetAllUserGroups() (groups []*data.UserGroup, err error)
	GetUserGroup(userGroupID string) (group *data.UserGroup, err error)
	CreateUserGroup(userGroup *data.UserGroup) (err error)
	UpdateUserGroup(userGroup *data.UserGroup) (err error)
	DeleteUserGroup(usergroupID string) (err error)
	AddUserToGroup(userName, groupID string) (err error)
	RemoveUserFromGroup(userName, groupID string) (err error)
	GetUsersInGroup(groupID string) (userInGroup []*data.User, err error)
	GetGroupsForUser(userName string) (groupsForUser []*data.UserGroup, err error)
	AddVariableValue(variableID, value string) (err error)
	ClearValuesForVariable(variableID string) (err error)
	GetValuesForVariable(variableID string) (values []*string, err error)
	SetPassword(userName, password string) (err error)
	UpdatePassword(userName, currentPassword, newPassword string) (err error)
}

//LocalClient - This is local client which executes the command in the current
//process, also uses data source directly
type LocalClient struct {
	data.OrekDataStore
}

//SetPassword - sets the password for the user
func (ds *LocalClient) SetPassword(userName, password string) (err error) {
	return err
}

//UpdatePassword - updates the password for the user
func (ds *LocalClient) UpdatePassword(userName,
	currentPassword, newPassword string) (err error) {
	return err
}

//CliCommandProvider - gives commands supported by the application
type CliCommandProvider interface {
	GetCommands() cli.Command
}

//OrekApp - contains command providers and runs the app
type OrekApp struct {
	Client           OrekClient
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
