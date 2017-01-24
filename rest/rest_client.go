package rest

import "github.com/varunamachi/orekng/data"
import "net"

//RestClient - client for orek service
type RestClient struct {
}

//Connect - login to the server
func (client *RestClient) Connect(ipAddress net.IPAddr, port int,
	userName, password string) (err error) {
	return err
}

//GetAllUsers - Gives all user entries in the server
func (client *RestClient) GetAllUsers() (users []*data.User, err error) {
	users = make([]*data.User, 0, 20)
	return users, err
}

//GetUser - Gives the user with given userName from server
func (client *RestClient) GetUser(userName string) (user *data.User, err error) {
	user = &data.User{}
	return user, err
}

// GetUserWithEmail - Gives the user entry with given EMail id
func (client *RestClient) GetUserWithEmail(
	email string) (user *data.User, err error) {
	user = &data.User{}
	logIfError(err)
	return user, err
}

//CreateUser - creates a user entry in the server with given User object
func (client *RestClient) CreateUser(user *data.User) (err error) {
	logIfError(err)
	return err
}

//UpdateUser - Upadates the user entry in the server with the information
//in the given user object
func (client *RestClient) UpdateUser(user *data.User) (err error) {
	logIfError(err)
	return err
}

//DeleteUser - deletes the user entry with given user name
func (client *RestClient) DeleteUser(userName string) (err error) {
	logIfError(err)
	return err
}

//GetAllEndpoints - Gives all the data endpoints which have entries in server
func (client *RestClient) GetAllEndpoints() (endpoints []*data.Endpoint, err error) {
	endpoints = make([]*data.Endpoint, 0, 100)
	logIfError(err)
	return endpoints, err
}

//GetEndpoint - Gives data endpoint object correspoing to the server entry with
// given name
func (client *RestClient) GetEndpoint(endpointID string) (endpoint *data.Endpoint, err error) {
	endpoint = &data.Endpoint{}
	logIfError(err)
	return endpoint, err
}

//CreateEndpoint - Creates a endpoint entry in server according to the endpoint
//object
func (client *RestClient) CreateEndpoint(endpoint *data.Endpoint) (err error) {
	logIfError(err)
	return err
}

//UpdateEndpoint - Updates the endpoint entry in server with information provided
//in the endpoint object
func (client *RestClient) UpdateEndpoint(endpoint *data.Endpoint) (err error) {
	logIfError(err)
	return err
}

//DeleteEndpoint - deletes an endpoint
func (client *RestClient) DeleteEndpoint(endpointID string) (err error) {
	logIfError(err)
	return err
}

//GetAllVariables - Gives list of all variables
func (client *RestClient) GetAllVariables() (variables []*data.Variable, err error) {
	variables = make([]*data.Variable, 0, 100)
	logIfError(err)
	return variables, err
}

//GetVariablesForEndpoint - Gives all the variables exported by an endpoint
func (client *RestClient) GetVariablesForEndpoint(
	endpointID string) (variables []*data.Variable, err error) {
	variables = make([]*data.Variable, 0, 100)
	logIfError(err)
	return variables, err
}

//GetVariable - Gives the variable with the given ID
func (client *RestClient) GetVariable(variableID string) (variable *data.Variable, err error) {
	variable = &data.Variable{}
	logIfError(err)
	return variable, err
}

//CreateVariable - creates a variable in the server
func (client *RestClient) CreateVariable(variable *data.Variable) (err error) {
	logIfError(err)
	return err
}

//UpdateVariable - updates a variable in the server
func (client *RestClient) UpdateVariable(variable *data.Variable) (err error) {
	logIfError(err)
	return err
}

//DeleteVariable - delete a variable from the server
func (client *RestClient) DeleteVariable(variableID string) (err error) {
	logIfError(err)
	return err
}

//GetAllParameters - Gives list of all parameters
func (client *RestClient) GetAllParameters() (parameters []*data.Parameter, err error) {
	parameters = make([]*data.Parameter, 0, 100)
	logIfError(err)
	return parameters, err
}

//GetParametersForEndpoint - Gives all the parameters exported by an endpoint
func (client *RestClient) GetParametersForEndpoint(
	endpointID string) (parameters []*data.Parameter, err error) {
	parameters = make([]*data.Parameter, 0, 100)
	logIfError(err)
	return parameters, err
}

//GetParameter - Gives the parameter with the given ID
func (client *RestClient) GetParameter(
	parameterID string) (parameter *data.Parameter, err error) {
	parameter = &data.Parameter{}
	logIfError(err)
	return parameter, err
}

//CreateParameter - creates a parameter in the server
func (client *RestClient) CreateParameter(parameter *data.Parameter) (err error) {
	logIfError(err)
	return err
}

//UpdateParameter - updates a parameter in the server
func (client *RestClient) UpdateParameter(parameter *data.Parameter) (err error) {
	logIfError(err)
	return err
}

//DeleteParameter - delete a parameter from the server
func (client *RestClient) DeleteParameter(parameterID string) (err error) {
	logIfError(err)
	return err
}

//GetAllUserGroups - gets the list of user group from the server
func (client *RestClient) GetAllUserGroups() (userGroups []*data.UserGroup, err error) {
	userGroups = make([]*data.UserGroup, 0, 100)
	logIfError(err)
	return userGroups, err
}

//GetUserGroup - get an instance of user group for give group name
func (client *RestClient) GetUserGroup(
	userGroupName string) (userGroup *data.UserGroup, err error) {
	userGroup = &data.UserGroup{}
	logIfError(err)
	return nil, err
}

//CreateUserGroup - creates an user group with give details
func (client *RestClient) CreateUserGroup(userGroup *data.UserGroup) (err error) {
	logIfError(err)
	return err
}

//UpdateUserGroup - Updates an existing user group with details from the
//given object
func (client *RestClient) UpdateUserGroup(userGroup *data.UserGroup) (err error) {
	logIfError(err)
	return err
}

//DeleteUserGroup - deletes an user group with the given group name
func (client *RestClient) DeleteUserGroup(userGroupName string) (err error) {
	logIfError(err)
	return err
}

//AddUserToGroup - adds user with given user name to a group with given group
//name
func (client *RestClient) AddUserToGroup(userName, groupID string) (err error) {
	logIfError(err)
	return err
}

//RemoveUserFromGroup - disassociates user with given user name from group with
//given group name
func (client *RestClient) RemoveUserFromGroup(userName, groupID string) (err error) {
	logIfError(err)
	return err
}

//GetUsersInGroup - gives a list of users who are associated with the group
//with given group name
func (client *RestClient) GetUsersInGroup(
	groupID string) (users []*data.User, err error) {
	logIfError(err)
	return users, err
}

//GetGroupsForUser - Gives a list of groups with which the user with given user
//name is associated
func (client *RestClient) GetGroupsForUser(
	userName string) (groups []*data.UserGroup, err error) {
	groups = make([]*data.UserGroup, 0, 100)
	logIfError(err)
	return groups, err
}

//AddVariableValue - Adds value to list of values of a variable
func (client *RestClient) AddVariableValue(variableID, value string) (err error) {
	logIfError(err)
	return err
}

//ClearValuesForVariable - clears values from the list of values associated with
//the variable with given variable id
func (client *RestClient) ClearValuesForVariable(variableID string) (err error) {
	logIfError(err)
	return err
}

//GetValuesForVariable - Gives list of values associated with a variable with
//given variable id
func (client *RestClient) GetValuesForVariable(
	variableID string) (values []*string, err error) {
	values = make([]*string, 0, 100)
	logIfError(err)
	return values, err
}

//SetPasswordHash - stores password hash for an user in the server
func (client *RestClient) SetPasswordHash(userName, passwordHash string) (err error) {
	logIfError(err)
	return err
}

//GetPasswordHash - Retrieves password hash for an user from the server
func (client *RestClient) GetPasswordHash(userName string) (hash string, err error) {
	logIfError(err)
	return hash, err
}

//UpdatePasswordHash - updates password hash for a user in the server
func (client *RestClient) UpdatePasswordHash(userName, passwordHash string) (err error) {
	logIfError(err)
	return err
}

//UserExists - Checks if an user record exists for given user bane
func (client *RestClient) UserExists(userName string) (exists bool, err error) {
	return exists, err
}

//UserExistsWithEmail - checks if an user record exists with given email
func (client *RestClient) UserExistsWithEmail(email string) (exists bool, err error) {
	return exists, err
}

//EndpointExists - checks if an endpoint exists with given ID
func (client *RestClient) EndpointExists(endpointID string) (exists bool, err error) {
	return exists, err
}

//VariableExists - checks if a variable exists with given variable ID
func (client *RestClient) VariableExists(variableID string) (exists bool, err error) {
	return exists, err
}

//VariableExistsInEndpoint - checks if a variable with given variableID in an
//endpoint given by the endpointID
func (client *RestClient) VariableExistsInEndpoint(
	variableID, endpointID string) (exists bool, err error) {
	return exists, err
}

//ParameterExists - checks if a parameter exists with given parameter ID
func (client *RestClient) ParameterExists(parameterID string) (exists bool, err error) {
	return exists, err
}

//ParameterExistsInEndpoint - checks if a parameter with given parameterID in an
//endpoint given by the endpointID
func (client *RestClient) ParameterExistsInEndpoint(
	parameterID, endpointID string) (exists bool, err error) {
	return exists, err
}

//UserGroupExists - checks if an User group exists with given ID
func (client *RestClient) UserGroupExists(
	userGroupID string) (exists bool, err error) {
	return exists, err
}

//UserExistsInGroup - checks if user with given user ID is associated with the
//group with given groupID
func (client *RestClient) UserExistsInGroup(
	userName, groupID string) (exists bool, err error) {
	return exists, err
}

//GroupHasUser - checks if group with given ID has a user with given userName
//associated with it
func (client *RestClient) GroupHasUser(
	groupID, userName string) (has bool, err error) {
	return has, err
}
