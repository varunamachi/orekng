package rest

import "github.com/varunamachi/orekng/data"
import "net"

//Client - client for orek service
type Client struct {
}

//Connect - login to the server
func (client *Client) Connect(ipAddress net.IPAddr, port int,
	userName, password string) (err error) {
	return err
}

//GetAllUsers - Gives all user entries in the server
func (client *Client) GetAllUsers() (users []*data.User, err error) {
	users = make([]*data.User, 0, 20)
	return users, err
}

//GetUser - Gives the user with given userName from server
func (client *Client) GetUser(userName string) (user *data.User, err error) {
	user = &data.User{}
	return user, err
}

// GetUserWithEmail - Gives the user entry with given EMail id
func (client *Client) GetUserWithEmail(
	email string) (user *data.User, err error) {
	user = &data.User{}
	logIfError(err)
	return user, err
}

//CreateUser - creates a user entry in the server with given User object
func (client *Client) CreateUser(user *data.User) (err error) {
	logIfError(err)
	return err
}

//UpdateUser - Upadates the user entry in the server with the information
//in the given user object
func (client *Client) UpdateUser(user *data.User) (err error) {
	logIfError(err)
	return err
}

//DeleteUser - deletes the user entry with given user name
func (client *Client) DeleteUser(userName string) (err error) {
	logIfError(err)
	return err
}

//GetAllEndpoints - Gives all the data endpoints which have entries in server
func (client *Client) GetAllEndpoints() (endpoints []*data.Endpoint, err error) {
	endpoints = make([]*data.Endpoint, 0, 100)
	logIfError(err)
	return endpoints, err
}

//GetEndpoint - Gives data endpoint object correspoing to the server entry with
// given name
func (client *Client) GetEndpoint(endpointID string) (endpoint *data.Endpoint, err error) {
	endpoint = &data.Endpoint{}
	logIfError(err)
	return endpoint, err
}

//CreateEndpoint - Creates a endpoint entry in server according to the endpoint
//object
func (client *Client) CreateEndpoint(endpoint *data.Endpoint) (err error) {
	logIfError(err)
	return err
}

//UpdateEndpoint - Updates the endpoint entry in server with information provided
//in the endpoint object
func (client *Client) UpdateEndpoint(endpoint *data.Endpoint) (err error) {
	logIfError(err)
	return err
}

//DeleteEndpoint - deletes an endpoint
func (client *Client) DeleteEndpoint(endpointID string) (err error) {
	logIfError(err)
	return err
}

//GetAllVariables - Gives list of all variables
func (client *Client) GetAllVariables() (variables []*data.Variable, err error) {
	variables = make([]*data.Variable, 0, 100)
	logIfError(err)
	return variables, err
}

//GetVariablesForEndpoint - Gives all the variables exported by an endpoint
func (client *Client) GetVariablesForEndpoint(
	endpointID string) (variables []*data.Variable, err error) {
	variables = make([]*data.Variable, 0, 100)
	logIfError(err)
	return variables, err
}

//GetVariable - Gives the variable with the given ID
func (client *Client) GetVariable(variableID string) (variable *data.Variable, err error) {
	variable = &data.Variable{}
	logIfError(err)
	return variable, err
}

//CreateVariable - creates a variable in the server
func (client *Client) CreateVariable(variable *data.Variable) (err error) {
	logIfError(err)
	return err
}

//UpdateVariable - updates a variable in the server
func (client *Client) UpdateVariable(variable *data.Variable) (err error) {
	logIfError(err)
	return err
}

//DeleteVariable - delete a variable from the server
func (client *Client) DeleteVariable(variableID string) (err error) {
	logIfError(err)
	return err
}

//GetAllParameters - Gives list of all parameters
func (client *Client) GetAllParameters() (parameters []*data.Parameter, err error) {
	parameters = make([]*data.Parameter, 0, 100)
	logIfError(err)
	return parameters, err
}

//GetParametersForEndpoint - Gives all the parameters exported by an endpoint
func (client *Client) GetParametersForEndpoint(
	endpointID string) (parameters []*data.Parameter, err error) {
	parameters = make([]*data.Parameter, 0, 100)
	logIfError(err)
	return parameters, err
}

//GetParameter - Gives the parameter with the given ID
func (client *Client) GetParameter(
	parameterID string) (parameter *data.Parameter, err error) {
	parameter = &data.Parameter{}
	logIfError(err)
	return parameter, err
}

//CreateParameter - creates a parameter in the server
func (client *Client) CreateParameter(parameter *data.Parameter) (err error) {
	logIfError(err)
	return err
}

//UpdateParameter - updates a parameter in the server
func (client *Client) UpdateParameter(parameter *data.Parameter) (err error) {
	logIfError(err)
	return err
}

//DeleteParameter - delete a parameter from the server
func (client *Client) DeleteParameter(parameterID string) (err error) {
	logIfError(err)
	return err
}

//GetAllUserGroups - gets the list of user group from the server
func (client *Client) GetAllUserGroups() (userGroups []*data.UserGroup, err error) {
	userGroups = make([]*data.UserGroup, 0, 100)
	logIfError(err)
	return userGroups, err
}

//GetUserGroup - get an instance of user group for give group name
func (client *Client) GetUserGroup(
	userGroupName string) (userGroup *data.UserGroup, err error) {
	userGroup = &data.UserGroup{}
	logIfError(err)
	return nil, err
}

//CreateUserGroup - creates an user group with give details
func (client *Client) CreateUserGroup(userGroup *data.UserGroup) (err error) {
	logIfError(err)
	return err
}

//UpdateUserGroup - Updates an existing user group with details from the
//given object
func (client *Client) UpdateUserGroup(userGroup *data.UserGroup) (err error) {
	logIfError(err)
	return err
}

//DeleteUserGroup - deletes an user group with the given group name
func (client *Client) DeleteUserGroup(userGroupName string) (err error) {
	logIfError(err)
	return err
}

//AddUserToGroup - adds user with given user name to a group with given group
//name
func (client *Client) AddUserToGroup(userName, groupID string) (err error) {
	logIfError(err)
	return err
}

//RemoveUserFromGroup - disassociates user with given user name from group with
//given group name
func (client *Client) RemoveUserFromGroup(userName, groupID string) (err error) {
	logIfError(err)
	return err
}

//GetUsersInGroup - gives a list of users who are associated with the group
//with given group name
func (client *Client) GetUsersInGroup(
	groupID string) (users []*data.User, err error) {
	logIfError(err)
	return users, err
}

//GetGroupsForUser - Gives a list of groups with which the user with given user
//name is associated
func (client *Client) GetGroupsForUser(
	userName string) (groups []*data.UserGroup, err error) {
	groups = make([]*data.UserGroup, 0, 100)
	logIfError(err)
	return groups, err
}

//AddVariableValue - Adds value to list of values of a variable
func (client *Client) AddVariableValue(variableID, value string) (err error) {
	logIfError(err)
	return err
}

//ClearValuesForVariable - clears values from the list of values associated with
//the variable with given variable id
func (client *Client) ClearValuesForVariable(variableID string) (err error) {
	logIfError(err)
	return err
}

//GetValuesForVariable - Gives list of values associated with a variable with
//given variable id
func (client *Client) GetValuesForVariable(
	variableID string) (values []*string, err error) {
	values = make([]*string, 0, 100)
	logIfError(err)
	return values, err
}

//SetPasswordHash - stores password hash for an user in the server
func (client *Client) SetPasswordHash(userName, passwordHash string) (err error) {
	logIfError(err)
	return err
}

//GetPasswordHash - Retrieves password hash for an user from the server
func (client *Client) GetPasswordHash(userName string) (hash string, err error) {
	logIfError(err)
	return hash, err
}

//UpdatePasswordHash - updates password hash for a user in the server
func (client *Client) UpdatePasswordHash(userName, passwordHash string) (err error) {
	logIfError(err)
	return err
}

//UserExists - Checks if an user record exists for given user bane
func (client *Client) UserExists(userName string) (exists bool, err error) {
	return exists, err
}

//UserExistsWithEmail - checks if an user record exists with given email
func (client *Client) UserExistsWithEmail(email string) (exists bool, err error) {
	return exists, err
}

//EndpointExists - checks if an endpoint exists with given ID
func (client *Client) EndpointExists(endpointID string) (exists bool, err error) {
	return exists, err
}

//VariableExists - checks if a variable exists with given variable ID
func (client *Client) VariableExists(variableID string) (exists bool, err error) {
	return exists, err
}

//VariableExistsInEndpoint - checks if a variable with given variableID in an
//endpoint given by the endpointID
func (client *Client) VariableExistsInEndpoint(
	variableID, endpointID string) (exists bool, err error) {
	return exists, err
}

//ParameterExists - checks if a parameter exists with given parameter ID
func (client *Client) ParameterExists(parameterID string) (exists bool, err error) {
	return exists, err
}

//ParameterExistsInEndpoint - checks if a parameter with given parameterID in an
//endpoint given by the endpointID
func (client *Client) ParameterExistsInEndpoint(
	parameterID, endpointID string) (exists bool, err error) {
	return exists, err
}

//UserGroupExists - checks if an User group exists with given ID
func (client *Client) UserGroupExists(
	userGroupID string) (exists bool, err error) {
	return exists, err
}

//UserExistsInGroup - checks if user with given user ID is associated with the
//group with given groupID
func (client *Client) UserExistsInGroup(
	userName, groupID string) (exists bool, err error) {
	return exists, err
}

//GroupHasUser - checks if group with given ID has a user with given userName
//associated with it
func (client *Client) GroupHasUser(
	groupID, userName string) (has bool, err error) {
	return has, err
}
