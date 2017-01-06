package sqlite

import (
	"github.com/varunamachi/orekng/data"
)

//GetAllUsers - Gives all user entries in the database
func (sqlite *DataStore) GetAllUsers() ([]*data.User, error) {
	return nil, nil
}

//GetUser - Gives the user with given userName from database
func (sqlite *DataStore) GetUser(userName string) (*data.User, error) {
	return nil, nil
}

// GetUserWithEmail - Gives the user entry with given EMail id
func (sqlite *DataStore) GetUserWithEmail(email string) (*data.User, error) {
	return nil, nil
}

//CreateUser - creates a user entry in the database with given User object
func (sqlite *DataStore) CreateUser(user *data.User) error {
	return nil
}

//UpdateUser - Upadates the user entry in the database with the information
//in the given user object
func (sqlite *DataStore) UpdateUser(user *data.User) error {
	return nil
}

//DeleteUser - deletes the user entry with given user name
func (sqlite *DataStore) DeleteUser(userName string) error {
	return nil
}

//GetAllSources - Gives all the data sources which have entries in database
func (sqlite *DataStore) GetAllSources() ([]*data.Endpoint, error) {
	return nil, nil
}

//GetSource - Gives data source object correspoing to the database entry with
// given name
func (sqlite *DataStore) GetSource(sourceID string) (*data.Endpoint, error) {
	return nil, nil
}

//CreateSource - Creates a source entry in database according to the source
//object
func (sqlite *DataStore) CreateSource(source *data.Endpoint) error {
	return nil
}

//UpdateSource - Updates the source entry in database with information provided
//in the source object
func (sqlite *DataStore) UpdateSource(source *data.Endpoint) error {
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
