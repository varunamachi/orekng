package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"

	"strconv"

	"path/filepath"

	"github.com/varunamachi/orekng/data"
	"github.com/varunamachi/orekng/data/sqlite"
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
func (ds LocalClient) SetPassword(userName, password string) (err error) {
	return err
}

//UpdatePassword - updates the password for the user
func (ds LocalClient) UpdatePassword(userName,
	currentPassword, newPassword string) (err error) {
	return err
}

//CliCommandProvider - gives commands supported by the application
type CliCommandProvider interface {
	GetCommand() cli.Command
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

func fromOrekDir(relative string) (path string) {
	home := os.Getenv("HOME")
	if runtime.GOOS == "windows" {
		home = os.Getenv("APPDATA")
	}
	return filepath.Join(home, ".orek", relative)
}

//Run - runs the application
func (orek *OrekApp) Run(args []string) (err error) {
	if runtime.GOOS != "windows" {

	}
	app := cli.NewApp()
	app.Name = "Orek"
	app.Version = "0.0.1"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Varun Amachi",
			Email: "varunamachi@github.com",
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "ds",
			Value: "sqlite",
			Usage: "Datasource name, sqlite|postgres",
		},
		cli.StringFlag{
			Name:  "db-path",
			Value: fromOrekDir("orek.db"),
			Usage: "Path to SQLite database [Only applicable for SQLite]",
		},
		cli.StringFlag{
			Name:  "ds-host",
			Value: "localhost",
			Usage: "DataBase host name [Not applicable for SqliteDataSource]",
		},
		cli.IntFlag{
			Name:  "ds-port",
			Value: 5432,
			Usage: "DataBase port [Not applicable for SqliteDataSource]",
		},
		cli.StringFlag{
			Name:  "db-name",
			Value: "orek",
			Usage: "DataBase name [Not applicable for SqliteDataSource]",
		},
		cli.StringFlag{
			Name:  "db-user",
			Value: "",
			Usage: "DataBase username [Not applicable for SqliteDataSource]",
		},
		cli.StringFlag{
			Name:  "db-password",
			Value: "",
			Usage: "Option db password for testing " +
				"[Not applicable for SqliteDataSource]",
		},
	}
	app.Before = func(ctx *cli.Context) (err error) {
		argetr := ArgGetter{Ctx: ctx}
		ds := argetr.GetRequiredString("ds")
		var store data.OrekDataStore
		if ds == "sqlite" {
			path := argetr.GetRequiredString("db-path")
			dirPath := filepath.Dir(path)
			fmt.Println(dirPath)
			if _, err := os.Stat(dirPath); os.IsNotExist(err) {
				err = os.Mkdir(dirPath, 0755)
				fmt.Println(err)
			}
			store, err = sqlite.Init(&sqlite.Options{
				Path: path,
			})
			if err == nil {
				data.SetStore(store)
				// err = data.GetStore().Init()
				if err != nil {
					log.Fatalf("Data Store initialization failed: %v", err)
				} //else {
				// 	log.Printf("Data Store initialized")
				// }
			}
		} else if ds == "postgres" {
			host := argetr.GetRequiredString("db-host")
			port := argetr.GetRequiredInt("db-port")
			dbName := argetr.GetRequiredString("db-name")
			user := argetr.GetRequiredString("db-user")
			pswd := argetr.GetRequiredString("db-password")
			if len(pswd) == 0 {
				fmt.Printf("Password for %s: ", user)
				var pbyte []byte
				pbyte, err = terminal.ReadPassword(int(syscall.Stdin))
				if err != nil {
					log.Fatalf("Could not retrieve DB password: %v", err)
				} else {
					pswd = string(pbyte)
				}
			}
			log.Printf(`Postgres isnt supported yet. Here are the args
				Host: %s,
				Port: %d,
				DbName: %s,
				User: %s`, host, port, dbName, user)
		} else {
			log.Fatalf("Unknown datasource %s requested", ds)
		}
		return err
	}
	// app.Action = func(ctx *cli.Context) error {
	// 	return errors.New("Please provide arguments")
	// }
	app.Commands = make([]cli.Command, 0, 30)
	for _, cmdp := range orek.CommandProviders {
		app.Commands = append(app.Commands, cmdp.GetCommand())
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

func readInput(text *string) (err error) {
	scanner := bufio.NewScanner(os.Stdin)
	*text = ""
	for scanner.Scan() {
		*text = scanner.Text()
		break
	}
	err = scanner.Err()
	return err
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
		err := readInput(&val)
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
		err := readInput(&val)
		if err != nil || len(val) == 0 {
			val = ""
			retriever.Err = fmt.Errorf("Required argument %s not provided", key)
		}
	}
	return val
}

//GetRequiredInt - gives a Integer argument either from commandline or from
//blocking user input, this method sets the error if required arg-val is empty
func (retriever *ArgGetter) GetRequiredInt(key string) (val int) {
	if retriever.Err != nil {
		return val
	}
	val = retriever.Ctx.Int(key)
	if !retriever.Ctx.IsSet(key) {
		fmt.Print(key + ": ")
		var strval string
		err := readInput(&strval)
		if err != nil || len(strval) == 0 {
			val = 0
			retriever.Err = fmt.Errorf("Required argument %s not provided", key)
		} else {
			val, err = strconv.Atoi(strval)
			if err != nil {
				retriever.Err = fmt.Errorf("Invalid value for %s given", key)
				val = 0
			}
		}
	}
	return val
}

//GetInt - gives a Integer argument either from commandline or from blocking
//user input, this method doesnt complain even if the arg-value is empty
func (retriever *ArgGetter) GetInt(key string) (val int) {
	if retriever.Err != nil {
		return val
	}
	val = retriever.Ctx.Int(key)
	if !retriever.Ctx.IsSet(key) {
		fmt.Print(key + ": ")
		var strval string
		err := readInput(&strval)
		if err != nil || len(strval) == 0 {
			val = 0
		} else {
			val, err = strconv.Atoi(strval)
			if err != nil {
				val = 0
			}
		}
	}
	return val
}
