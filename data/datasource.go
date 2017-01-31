package data

import "github.com/varunamachi/orekng/olog"

var db OrekDataStore

//SetStore sets the application wIDe data store to be used
func SetStore(dbInst OrekDataStore) error {
	db = dbInst
	return nil
}

//GetStore - give the application data store
func GetStore() OrekDataStore {
	if db == nil {
		olog.Fatal("Orek", "DataStore is Nil!!!")
	}
	return db
}

//OrekDataStore - interface declares the operation that will be exposed by a
//application data store
type OrekDataStore interface {
	GetAllUsers() (users []*User, err error)
	GetUser(userName string) (users *User, err error)
	GetUserWithEmail(email string) (user *User, err error)
	UserExists(userName string) (exists bool, err error)
	UserExistsWithEmail(email string) (exists bool, err error)
	CreateUser(user *User) (err error)
	UpdateUser(user *User) (err error)
	DeleteUser(userName string) (err error)
	GetAllEndpoints() (endpoints []*Endpoint, err error)
	GetEndpoint(endpointID string) (endpoint *Endpoint, err error)
	EndpointExists(endpointID string) (exists bool, err error)
	CreateEndpoint(endpoint *Endpoint) (err error)
	UpdateEndpoint(endpoint *Endpoint) (err error)
	DeleteEndpoint(endpointID string) (err error)
	GetAllVariables() (variables []*Variable, err error)
	GetVariablesForEndpoint(endpointID string) (variables []*Variable, err error)
	GetVariable(variableID string) (variable *Variable, err error)
	VariableExists(variableID string) (exists bool, err error)
	VariableExistsInEndpoint(variableID, endpointID string) (exists bool, err error)
	CreateVariable(variable *Variable) (err error)
	UpdateVariable(variable *Variable) (err error)
	DeleteVariable(variableID string) (err error)
	GetAllParameters() (parameters []*Parameter, err error)
	GetParametersForEndpoint(endpointID string) (parameters []*Parameter, err error)
	GetParameter(parameterID string) (parameter *Parameter, err error)
	ParameterExists(parameterID string) (exists bool, err error)
	ParameterExistsInEndpoint(variableID, endpointID string) (exists bool, err error)
	CreateParameter(parameter *Parameter) (err error)
	UpdateParameter(parameter *Parameter) (err error)
	DeleteParameter(parameterID string) (err error)
	GetAllUserGroups() (groups []*UserGroup, err error)
	GetUserGroup(userGroupID string) (group *UserGroup, err error)
	UserGroupExists(userGroupID string) (exists bool, err error)
	CreateUserGroup(userGroup *UserGroup) (err error)
	UpdateUserGroup(userGroup *UserGroup) (err error)
	DeleteUserGroup(usergroupID string) (err error)
	AddUserToGroup(userName, groupID string) (err error)
	UserExistsInGroup(userName, groupID string) (exists bool, err error)
	RemoveUserFromGroup(userName, groupID string) (err error)
	GetUsersInGroup(groupID string) (userInGroup []*User, err error)
	GroupHasUser(groupID, userName string) (has bool, err error)
	GetGroupsForUser(userName string) (groupsForUser []*UserGroup, err error)
	AddVariableValue(variableID, value string) (err error)
	ClearValuesForVariable(variableID string) (err error)
	GetValuesForVariable(variableID string) (values []*string, err error)
	SetPasswordHash(userName, passwordHash string) (err error)
	GetPasswordHash(userName string) (hash string, err error)
	UpdatePasswordHash(userName, passwordHash string) (err error)

	Init() (err error)
	ClearData() (err error)
	DeleteSchema() (err error)
}
