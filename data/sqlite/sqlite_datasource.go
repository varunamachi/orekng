package sqlite

import (
	"github.com/varunamachi/orekng/data"
)

func (sqlite *DataStore) GetAllUsers() ([]*data.User, error) {
	return nil, nil
}

func (sqlite *DataStore) GetUser(userName string) (*data.User, error) {
	return nil, nil
}

func (sqlite *DataStore) GetUserWithEmail(email string) (*data.User, error) {
	return nil, nil
}

func (sqlite *DataStore) CreateUser(user *data.User) error {
	return nil
}

func (sqlite *DataStore) UpdateUser(user *data.User) error {
	return nil
}

func (sqlite *DataStore) DeleteUser(userName string) error {
	return nil
}

func (sqlite *DataStore) GetAllSources() ([]*data.Source, error) {
	return nil, nil
}

func (sqlite *DataStore) GetSource(sourceID string) (*data.Source, error) {
	return nil, nil
}

func (sqlite *DataStore) CreateSource(source *data.Source) error {
	return nil
}

func (sqlite *DataStore) UpdateSource(source *data.Source) error {
	return nil
}

func (sqlite *DataStore) DeleteSource(sourceID string) error {
	return nil
}

func (sqlite *DataStore) GetAllVariables() ([]*data.Variable, error) {
	return nil, nil
}

func (sqlite *DataStore) GetVariablesForSource(sourceID string) ([]*data.Variable, error) {
	return nil, nil
}

func (sqlite *DataStore) GetVariable(variableID string) (*data.Variable, error) {
	return nil, nil
}

func (sqlite *DataStore) CreateVariable(variable *data.Variable) error {
	return nil
}

func (sqlite *DataStore) UpdateVariable(variable *data.Variable) error {
	return nil
}

func (sqlite *DataStore) DeleteVariable(variableID string) error {
	return nil
}

func (sqlite *DataStore) GetAllUserGroups() ([]*data.UserGroup, error) {
	return nil, nil
}

func (sqlite *DataStore) GetUserGroup(userGroupName string) (*data.UserGroup, error) {
	return nil, nil
}

func (sqlite *DataStore) CreateUserGroup(userGroup *data.UserGroup) error {
	return nil
}

func (sqlite *DataStore) UpdateUserGroup(userGroup *data.UserGroup) error {
	return nil
}

func (sqlite *DataStore) DeleteUserGroup(userGroupName string) error {
	return nil
}

func (sqlite *DataStore) AddUserToGroup(userName, groupName string) error {
	return nil
}

func (sqlite *DataStore) RemoveUserFromGroup(userName, groupName string) error {
	return nil
}

func (sqlite *DataStore) GetUsersInGroup(groupName string) ([]*data.User, error) {
	return nil, nil
}

func (sqlite *DataStore) GetGroupsForUser(userName string) ([]*data.UserGroup, error) {
	return nil, nil
}

func (sqlite *DataStore) AddVariableValue(variableID, value string) error {
	return nil
}

func (sqlite *DataStore) ClearValuesForVariable(variableID string) error {
	return nil
}

func (sqlite *DataStore) GetValuesForVariable(variableID string) ([]*string, error) {
	return nil, nil
}
