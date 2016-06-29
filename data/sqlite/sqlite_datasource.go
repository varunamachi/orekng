package sqlite

func (sqlite *DataStore) GetAllUsers() ([]*User, error) {
	return nil, nil
}

func (sqlite *DataStore) GetUser(userName string) (*User, error) {
	return nil, nil
}

func (sqlite *DataStore) GetUserWithEmail(email string) (*User, error) {
	return nil, nil
}

func (sqlite *DataStore) CreateUser(user *User) error {
	return nil
}

func (sqlite *DataStore) UpdateUser(user *User) error {
	return nil
}

func (sqlite *DataStore) DeleteUser(userName string) error {
	return nil
}

func (sqlite *DataStore) GetAllSources() ([]*Source, error) {
	return nil, nil
}

func (sqlite *DataStore) GetSource(sourceID string) (*Source, error) {
	return nil, nil
}

func (sqlite *DataStore) CreateSource(source *Source) error {
	return nil
}

func (sqlite *DataStore) UpdateSource(source *Source) error {
	return nil
}

func (sqlite *DataStore) DeleteSource(sourceID string) error {
	return nil
}

func (sqlite *DataStore) GetAllVariables() ([]*Variable, error) {
	return nil, nil
}

func (sqlite *DataStore) GetVariablesForSource(sourceID string) ([]*Variable, error) {
	return nil, nil
}

func (sqlite *DataStore) GetVariable(variableID string) (*Variable, error) {
	return nil, nil
}

func (sqlite *DataStore) CreateVariable(variable *Variable) error {
	return nil, nil
}

func (sqlite *DataStore) UpdateVariable(variable *Variable) error {
	return nil
}

func (sqlite *DataStore) DeleteVariable(variableID string) error {
	return nil
}

func (sqlite *DataStore) GetAllUserGroups() ([]*UserGroup, error) {
	return nil, nil
}

func (sqlite *DataStore) GetUserGroup(userGroupName string) (*UserGroup, error) {
	return nil, nil
}

func (sqlite *DataStore) CreateUserGroup(userGroup *UserGroup) error {
	return nil
}

func (sqlite *DataStore) UpdateUserGroup(userGroup *UserGroup) error {
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

func (sqlite *DataStore) GetUsersInGroup(groupName string) ([]*User, error) {
	return nil, nil
}

func (sqlite *DataStore) GetGroupsForUser(userName string) ([]*UserGroup, error) {
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
