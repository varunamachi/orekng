package data

import (
	"log"
)

var db OrekDb = nil

func SetDataSource(dbInst OrekDb) error {
	db = dbInst
	return nil
}

func DataSource() OrekDb {
	if db == nil {
		log.Fatal("DataSource is Nil!!!")
	}
	return db
}

type OrekDb interface {
	GetAllUsers() ([]*User, error)
	GetUser(userName string) (*User, error)
	GetUserWithEmail(email string) (*User, error)
	CreateUser(user *User) error
    UpdateUser(user *User) error
	DeleteUser(userName string) error

	GetAllSources() ([]*Source, error)
	GetSource(sourceId string) (*Source, error)
	CreateSource(source *Source) error
    UpdateSource(source *Source) error
	DeleteSource(sourceId string) error

	GetAllVariables() ([]*Variable, error)
    GetVariablesForSource(sourceId string) ([]*Variable, error)
	GetVariable(variableId string) (*Variable, error)
	CreateVariable(variable *Variable) error
    UpdateVariable(variable *Variable) error
	DeleteVariable(variableId string) error

	GetAllUserGroups() ([]*UserGroup, error)
	GetUserGroup(userGroupName string) (*UserGroup, error)
	CreateUserGroup(userGroup *UserGroup) error
    UpdateUserGroup(userGroup *UserGroup) error
	DeleteUserGroup(userGroupName string) error

	AddUserToGroup(userName, groupName string) error
	RemoveUserFromGroup(userName, groupName string) error
	GetUsersInGroup(groupName string) ([]*User, error)
	GetGroupsForUser(userName string) ([]*UserGroup, error)

	AddVariableValue(variableId, value string) error
	ClearValuesForVariable(variableId string) error
	GetValuesForVariable(variableId string) ([]*string, error)
}
