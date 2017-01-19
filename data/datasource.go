package data

import (
	"log"
	"time"
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
	GetAllUsers() (users []*User, err error)
	GetUser(userName string) (users *User, err error)
	GetUserWithEmail(email string) (user *User, err error)
	CreateUser(user *User) (err error)
	UpdateUser(user *User) (err error)
	DeleteUser(userName string) (err error)

	GetAllEndpoints() (endpoints []*Endpoint, err error)
	GetEndpoint(endpointID string) (endpoint *Endpoint, err error)
	CreateEndpoint(endpoint *Endpoint) (err error)
	UpdateEndpoint(endpoint *Endpoint) (err error)
	DeleteEndpoint(endpointID string) (err error)

	GetAllVariables() (variables []*Variable, err error)
	GetVariablesForEndpoint(endpointID string) (variables []*Variable, err error)
	GetVariable(variableID string) (variable *Variable, err error)
	CreateVariable(variable *Variable) (err error)
	UpdateVariable(variable *Variable) (err error)
	DeleteVariable(variableID string) (err error)

	GetAllParameters() (parameters []*Parameter, err error)
	GetParametersForEndpoint(endpointID string) (parameters []*Parameter, err error)
	GetParameter(parameterID string) (parameter *Parameter, err error)
	CreateParameter(parameter *Parameter) (err error)
	UpdateParameter(parameter *Parameter) (err error)
	DeleteParameter(parameterID string) (err error)

	GetAllUserGroups() (groups []*UserGroup, err error)
	GetUserGroup(userGroupName string) (group *UserGroup, err error)
	CreateUserGroup(userGroup *UserGroup) (err error)
	UpdateUserGroup(userGroup *UserGroup) (err error)
	DeleteUserGroup(userGroupName string) (err error)

	AddUserToGroup(userName, groupName string) (err error)
	RemoveUserFromGroup(userName, groupName string) (err error)
	GetUsersInGroup(groupName string) (userInGroup []*User, err error)
	GetGroupsForUser(userName string) (groupsForUser []*UserGroup, err error)

	AddVariableValue(variableID, value string) (err error)
	ClearValuesForVariable(variableID string) (err error)
	GetValuesForVariable(variableID string) (values []*string, err error)

	CreateUserSession(session *Session) (err error)
	GetUserSession(sessionID string) (session *Session, err error)
	RemoveUserSession(sessionID string) (err error)
	ClearExpiredSessions(expiryTime time.Time) (err error)

	SetPasswordHash(userName, passwordHash string) (err error)
	GetPasswordHash(userName string) (hash string, err error)
}
