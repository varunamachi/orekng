package data

import (
	"log"
	"time"

	"github.com/varunamachi/orekng/rest"
)

var db OrekDataStore

//SetDataStore sets the application wIDe data store to be used
func SetDataStore(dbInst OrekDataStore) error {
	db = dbInst
	return nil
}

//GetDataStore - give the application data store
func GetDataStore() OrekDataStore {
	if db == nil {
		log.Fatal("DataStore is Nil!!!")
	}
	return db
}

//OrekDataStore - interface declares the operation that will be exposed by a
//application data store
type OrekDataStore interface {
	GetAllUsers() ([]*User, error)
	GetUser(userName string) (*User, error)
	GetUserWithEmail(email string) (*User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(userName string) error

	GetAllEndpoints() ([]*Endpoint, error)
	GetEndpoint(endpointID string) (*Endpoint, error)
	CreateEndpoint(endpoint *Endpoint) error
	UpdateEndpoint(endpoint *Endpoint) error
	DeleteEndpoint(endpointID string) error

	GetAllVariables() ([]*Variable, error)
	GetVariablesForEndpoint(endpointID string) ([]*Variable, error)
	GetVariable(variableID string) (*Variable, error)
	CreateVariable(variable *Variable) error
	UpdateVariable(variable *Variable) error
	DeleteVariable(variableID string) error

	GetAllParameters() ([]*Parameter, error)
	GetParametersForEndpoint(endpointID string) ([]*Parameter, error)
	GetParameter(parameterID string) (*Parameter, error)
	CreateParameter(parameter *Parameter) error
	UpdateParameter(parameter *Parameter) error
	DeleteParameter(parameterID string) error

	GetAllUserGroups() ([]*UserGroup, error)
	GetUserGroup(userGroupName string) (*UserGroup, error)
	CreateUserGroup(userGroup *UserGroup) error
	UpdateUserGroup(userGroup *UserGroup) error
	DeleteUserGroup(userGroupName string) error

	AddUserToGroup(userName, groupName string) error
	RemoveUserFromGroup(userName, groupName string) error
	GetUsersInGroup(groupName string) ([]*User, error)
	GetGroupsForUser(userName string) ([]*UserGroup, error)

	AddVariableValue(variableID, value string) error
	ClearValuesForVariable(variableID string) error
	GetValuesForVariable(variableID string) ([]*string, error)

	CreateUserSession(session *rest.Session) error
	GetUserSession(sessionID string) (rest.Session, error)
	RemoveUserSession(sessionID string) error
	ClearExpiredSessions(expiryTime time.Time) error
}
