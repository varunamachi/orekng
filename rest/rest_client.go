package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"strings"

	"github.com/varunamachi/orekng/data"
	"github.com/varunamachi/orekng/olog"
)

//Client - client for orek service
type Client struct {
	http.Client
	Address    string
	VersionStr string
	Token      string
}

//NewRestClient - creates a new rest client
func NewRestClient(address, versionStr string) *Client {
	return &Client{
		Client: http.Client{
			Timeout: 0,
		},
		Address:    address,
		VersionStr: versionStr,
	}
}

//Login - login to the server
func (client *Client) Login(userName, password string) (err error) {
	apiURL := fmt.Sprintf("%s/%s/login", client.Address, client.VersionStr)
	form := url.Values{}
	form.Add("username", userName)
	form.Add("password", password)
	var req *http.Request
	req, err = http.NewRequest("POST", apiURL, strings.NewReader(form.Encode()))
	if err == nil {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		var resp *http.Response
		resp, err = client.Do(req)
		if err == nil {
			if resp.StatusCode == http.StatusOK {
				defer resp.Body.Close()
				decoder := json.NewDecoder(resp.Body)
				tmap := make(map[string]string)
				err = decoder.Decode(&tmap)
				client.Token = tmap["token"]
			} else {
				olog.Error("REST", "Unexpected response status %s", resp.Status)
			}
		}
	}
	if err != nil {
		olog.PrintError("RESTClient", err)
	}
	return err
}

//GetAllUsers - Gives all user entries in the server
func (client *Client) GetAllUsers() (users []*data.User, err error) {
	apiURL := fmt.Sprintf("%s/%s/in/users", client.Address, client.VersionStr)
	var req *http.Request
	req, err = http.NewRequest("GET", apiURL, nil)
	if err == nil {
		// req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		authHeader := fmt.Sprintf("Bearer %s", client.Token)
		req.Header.Add("Authorization", authHeader)
		var resp *http.Response
		resp, err = client.Do(req)
		if err == nil {
			defer resp.Body.Close()
			decoder := json.NewDecoder(resp.Body)
			if resp.StatusCode == http.StatusOK {
				users = make([]*data.User, 0, 20)
				err = decoder.Decode(&users)
			} else if resp.StatusCode == http.StatusInternalServerError {
				var res Result
				err = decoder.Decode(&res)
				olog.Error("REST", "%s : %s - %s", res.Operation,
					res.Message,
					res.Error)
			} else {
				olog.Print("REST", "Unexpected response status %s", resp.Status)
			}
		}
	}
	if err != nil {
		olog.PrintError("RESTClient", err)
	}
	return users, err
}

//GetUser - Gives the user with given userName from server
func (client *Client) GetUser(userName string) (user *data.User, err error) {
	apiURL := fmt.Sprintf("%s/%s/in/users/%s", client.Address,
		client.VersionStr,
		userName)
	var req *http.Request
	req, err = http.NewRequest("GET", apiURL, nil)
	if err == nil {
		// req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		authHeader := fmt.Sprintf("Bearer %s", client.Token)
		req.Header.Add("Authorization", authHeader)
		var resp *http.Response
		resp, err = client.Do(req)
		if err == nil {
			defer resp.Body.Close()
			decoder := json.NewDecoder(resp.Body)
			if resp.StatusCode == http.StatusOK {
				user = &data.User{}
				err = decoder.Decode(&user)
			} else if resp.StatusCode == http.StatusInternalServerError {
				var res Result
				err = decoder.Decode(&res)
				olog.Error("REST", "%s : %s - %s", res.Operation,
					res.Message,
					res.Error)
			} else {
				olog.Print("REST", "Unexpected response status %s", resp.Status)
			}
		}
	}
	if err != nil {
		olog.PrintError("RESTClient", err)
	}
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

//SetPassword - stores password hash for an user in the server
func (client *Client) SetPassword(userName, password string) (err error) {
	logIfError(err)
	return err
}

//UpdatePassword - updates password hash for a user in the server
func (client *Client) UpdatePassword(userName,
	currentPassword, newPassword string) (err error) {
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
