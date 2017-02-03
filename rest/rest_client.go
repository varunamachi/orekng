package rest

import (
	"bytes"
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
	BaseURL    string
}

//NewRestClient - creates a new rest client
func NewRestClient(address, versionStr string) *Client {
	return &Client{
		Client: http.Client{
			Timeout: 0,
		},
		Address:    address,
		VersionStr: versionStr,
		BaseURL:    fmt.Sprintf("%s/%s", address, versionStr),
	}
}

func handleStatusCode(statusCode int, decoder *json.Decoder) {
	if statusCode == http.StatusInternalServerError ||
		statusCode == http.StatusBadRequest ||
		statusCode == http.StatusUnauthorized ||
		statusCode == http.StatusOK {
		var res Result
		err := decoder.Decode(&res)
		if err == nil && len(res.Error) != 0 {
			olog.Error("REST", "%s : %s - %s", res.Operation,
				res.Message,
				res.Error)
		} else if err != nil {
			olog.Error("RESTClient", "Result decode failed: ", err)
		} else {
			olog.Info("REST", "%s : %s", res.Operation, res.Message)
		}
	} else {
		olog.Error("REST", "Unexpected response status %s",
			http.StatusText(statusCode))
	}
}

func (client *Client) orekDo(req *http.Request) (resp *http.Response, err error) {
	authHeader := fmt.Sprintf("Bearer %s", client.Token)
	req.Header.Add("Authorization", authHeader)
	resp, err = client.Do(req)
	return resp, err
}

func (client *Client) getURL(args ...string) (str string) {
	var buffer bytes.Buffer
	buffer.WriteString(client.BaseURL)
	buffer.WriteString("/in/")
	for i := 0; i < len(args); i++ {
		buffer.WriteString(args[i])
		if i < len(args)-1 {
			buffer.WriteString("/")
		}
	}
	str = buffer.String()
	return str
}

func (client *Client) orekPutOrPost(
	method string,
	content interface{},
	urlArgs ...string) (err error) {

	var data []byte
	var resp *http.Response
	data, err = json.Marshal(bytes.Buffer{})
	apiURL := client.getURL(urlArgs...)
	req, err := http.NewRequest(method, apiURL, bytes.NewBuffer(data))
	authHeader := fmt.Sprintf("Bearer %s", client.Token)
	req.Header.Add("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		decoder := json.NewDecoder(resp.Body)
		handleStatusCode(resp.StatusCode, decoder)
	}
	return err
}

func (client *Client) orekGet(
	content interface{},
	urlArgs ...string) (err error) {

	var req *http.Request
	var resp *http.Response
	apiURL := client.getURL(urlArgs...)
	req, err = http.NewRequest("GET", apiURL, nil)
	authHeader := fmt.Sprintf("Bearer %s", client.Token)
	req.Header.Add("Authorization", authHeader)
	resp, err = client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		decoder := json.NewDecoder(resp.Body)
		if resp.StatusCode == http.StatusOK {
			err = decoder.Decode(content)
		} else {
			handleStatusCode(resp.StatusCode, decoder)
		}
	}
	return err
}

func (client *Client) orekDelete(
	urlArgs ...string) (err error) {
	var req *http.Request
	var resp *http.Response
	apiURL := client.getURL(urlArgs...)
	req, err = http.NewRequest("GET", apiURL, nil)
	authHeader := fmt.Sprintf("Bearer %s", client.Token)
	req.Header.Add("Authorization", authHeader)
	resp, err = client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		decoder := json.NewDecoder(resp.Body)
		handleStatusCode(resp.StatusCode, decoder)
	}
	return err
}

func (client *Client) orekPost(content interface{},
	urlArgs ...string) (err error) {
	return client.orekPutOrPost("POST", content, urlArgs...)
}

func (client *Client) orekPut(content interface{},
	urlArgs ...string) (err error) {
	return client.orekPutOrPost("PUT", content, urlArgs...)
}

//GetAllUsers - Gives all user entries in the server
func (client *Client) GetAllUsers() (users []*data.User, err error) {
	// apiURL := fmt.Sprintf("%s/%s/in/users", client.Address, client.VersionStr)
	users = make([]*data.User, 0, 20)
	err = client.orekGet(&users, "users")
	if err != nil {
		olog.PrintError("RESTClient", err)
	}
	return users, err
}

//GetUser - Gives the user with given userName from server
func (client *Client) GetUser(userName string) (user *data.User, err error) {
	user = &data.User{}
	err = client.orekGet(&user, "users", userName)
	if err != nil {
		olog.PrintError("RESTClient", err)
	}
	return user, err
}

// GetUserWithEmail - Gives the user entry with given EMail id
func (client *Client) GetUserWithEmail(
	email string) (user *data.User, err error) {
	user = &data.User{}
	err = client.orekGet(&user, "users/email", email)
	return user, err
}

//CreateUser - creates a user entry in the server with given User object
func (client *Client) CreateUser(user *data.User) (err error) {
	err = client.orekPost(user, "users")
	return err
}

//UpdateUser - Upadates the user entry in the server with the information
//in the given user object
func (client *Client) UpdateUser(user *data.User) (err error) {
	err = client.orekPut(user, "users")
	return err
}

//DeleteUser - deletes the user entry with given user name
func (client *Client) DeleteUser(userName string) (err error) {
	err = client.orekDelete("users", userName)
	return err
}

//GetAllEndpoints - Gives all the data endpoints which have entries in server
func (client *Client) GetAllEndpoints() (endpoints []*data.Endpoint, err error) {
	endpoints = make([]*data.Endpoint, 0, 100)
	err = client.orekGet(&endpoints, "endpoints")
	return endpoints, err
}

//GetEndpoint - Gives data endpoint object correspoing to the server entry with
// given name
func (client *Client) GetEndpoint(endpointID string) (endpoint *data.Endpoint, err error) {
	endpoint = &data.Endpoint{}
	err = client.orekGet(endpoint, "endpoints", endpointID)
	return endpoint, err
}

//CreateEndpoint - Creates a endpoint entry in server according to the endpoint
//object
func (client *Client) CreateEndpoint(endpoint *data.Endpoint) (err error) {
	err = client.orekPost(endpoint, "endpoints")
	return err
}

//UpdateEndpoint - Updates the endpoint entry in server with information provided
//in the endpoint object
func (client *Client) UpdateEndpoint(endpoint *data.Endpoint) (err error) {
	err = client.orekPut(endpoint, "endpoints")
	return err
}

//DeleteEndpoint - deletes an endpoint
func (client *Client) DeleteEndpoint(endpointID string) (err error) {
	err = client.orekDelete("endpoints", endpointID)
	return err
}

//GetAllVariables - Gives list of all variables
func (client *Client) GetAllVariables() (variables []*data.Variable, err error) {
	variables = make([]*data.Variable, 0, 100)
	err = client.orekGet(&variables, "variables")
	return variables, err
}

//GetVariablesForEndpoint - Gives all the variables exported by an endpoint
func (client *Client) GetVariablesForEndpoint(
	endpointID string) (variables []*data.Variable, err error) {
	variables = make([]*data.Variable, 0, 100)
	err = client.orekGet(&variables, "endpoints", endpointID, "variables")
	return variables, err
}

//GetVariable - Gives the variable with the given ID
func (client *Client) GetVariable(variableID string) (variable *data.Variable, err error) {
	variable = &data.Variable{}
	err = client.orekGet(variable, "variables", variableID)
	return variable, err
}

//CreateVariable - creates a variable in the server
func (client *Client) CreateVariable(variable *data.Variable) (err error) {
	err = client.orekPost(variable, "variables")
	return err
}

//UpdateVariable - updates a variable in the server
func (client *Client) UpdateVariable(variable *data.Variable) (err error) {
	err = client.orekPut(variable, "variables")
	return err
}

//DeleteVariable - delete a variable from the server
func (client *Client) DeleteVariable(variableID string) (err error) {
	err = client.orekDelete("variables", variableID)
	return err
}

//GetAllParameters - Gives list of all parameters
func (client *Client) GetAllParameters() (parameters []*data.Parameter, err error) {
	parameters = make([]*data.Parameter, 0, 100)
	err = client.orekGet(&parameters, "parameters")
	return parameters, err
}

//GetParametersForEndpoint - Gives all the parameters exported by an endpoint
func (client *Client) GetParametersForEndpoint(
	endpointID string) (parameters []*data.Parameter, err error) {
	parameters = make([]*data.Parameter, 0, 100)
	err = client.orekGet(&parameters, "endpoints", endpointID, "parameters")
	return parameters, err
}

//GetParameter - Gives the parameter with the given ID
func (client *Client) GetParameter(
	parameterID string) (parameter *data.Parameter, err error) {
	parameter = &data.Parameter{}
	err = client.orekGet(parameter, "parameters", parameterID)
	return parameter, err
}

//CreateParameter - creates a parameter in the server
func (client *Client) CreateParameter(parameter *data.Parameter) (err error) {
	err = client.orekPost(parameter, "parameters")
	return err
}

//UpdateParameter - updates a parameter in the server
func (client *Client) UpdateParameter(parameter *data.Parameter) (err error) {
	err = client.orekPut(parameter, "parameters")
	return err
}

//DeleteParameter - delete a parameter from the server
func (client *Client) DeleteParameter(parameterID string) (err error) {
	err = client.orekDelete("parameters", parameterID)
	return err
}

//GetAllUserGroups - gets the list of user group from the server
func (client *Client) GetAllUserGroups() (userGroups []*data.UserGroup, err error) {
	userGroups = make([]*data.UserGroup, 0, 100)
	err = client.orekGet(&userGroups, "groups")
	return userGroups, err
}

//GetUserGroup - get an instance of user group for give group name
func (client *Client) GetUserGroup(
	groupID string) (userGroup *data.UserGroup, err error) {
	userGroup = &data.UserGroup{}
	err = client.orekGet(userGroup, "groups", groupID)
	return nil, err
}

//CreateUserGroup - creates an user group with give details
func (client *Client) CreateUserGroup(userGroup *data.UserGroup) (err error) {
	err = client.orekPost(userGroup, "groups")
	return err
}

//UpdateUserGroup - Updates an existing user group with details from the
//given object
func (client *Client) UpdateUserGroup(userGroup *data.UserGroup) (err error) {
	err = client.orekPut(userGroup, "groups")
	return err
}

//DeleteUserGroup - deletes an user group with the given group name
func (client *Client) DeleteUserGroup(groupID string) (err error) {
	err = client.orekDelete("groups", groupID)
	return err
}

//AddUserToGroup - adds user with given user name to a group with given group
//name
func (client *Client) AddUserToGroup(userName, groupID string) (err error) {
	err = client.orekPut("", "groups", groupID, "users", userName)
	return err
}

//RemoveUserFromGroup - disassociates user with given user name from group with
//given group name
func (client *Client) RemoveUserFromGroup(userName, groupID string) (err error) {
	err = client.orekDelete("", "groups", groupID, "users", userName)
	return err
}

//GetUsersInGroup - gives a list of users who are associated with the group
//with given group name
func (client *Client) GetUsersInGroup(
	groupID string) (users []*data.User, err error) {
	users = make([]*data.User, 0, 20)
	err = client.orekGet(&users, "groups", groupID, "users")
	return users, err
}

//GetGroupsForUser - Gives a list of groups with which the user with given user
//name is associated
func (client *Client) GetGroupsForUser(
	userName string) (groups []*data.UserGroup, err error) {
	groups = make([]*data.UserGroup, 0, 100)
	err = client.orekGet(&groups, "users", userName, "groups")
	return groups, err
}

//AddVariableValue - Adds value to list of values of a variable
func (client *Client) AddVariableValue(variableID, value string) (err error) {
	err = client.orekPost(&value, "variables", variableID, "values")
	return err
}

//ClearValuesForVariable - clears values from the list of values associated with
//the variable with given variable id
func (client *Client) ClearValuesForVariable(variableID string) (err error) {
	err = client.orekDelete("variables", variableID, "values")
	return err
}

//GetValuesForVariable - Gives list of values associated with a variable with
//given variable id
func (client *Client) GetValuesForVariable(
	variableID string) (values []*string, err error) {
	values = make([]*string, 0, 100)
	err = client.orekGet(&values, "variables", variableID, "values")
	return values, err
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
			defer resp.Body.Close()
			decoder := json.NewDecoder(resp.Body)
			if resp.StatusCode == http.StatusOK {
				tmap := make(map[string]string)
				err = decoder.Decode(&tmap)
				client.Token = tmap["token"]
			} else {
				handleStatusCode(resp.StatusCode, decoder)
			}
		}
	}
	if err != nil {
		olog.PrintError("RESTClient", err)
	}
	return err
}

//SetPassword - stores password hash for an user in the server
func (client *Client) SetPassword(userName, password string) (err error) {
	apiURL := fmt.Sprintf("%s/%s/in/manageAuth", client.Address, client.VersionStr)
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
			defer resp.Body.Close()
			decoder := json.NewDecoder(resp.Body)
			handleStatusCode(resp.StatusCode, decoder)
		}
	}
	return err
}

//UpdatePassword - updates password hash for a user in the server
func (client *Client) UpdatePassword(userName,
	currentPassword, newPassword string) (err error) {
	apiURL := fmt.Sprintf("%s/%s/in/manageAuth", client.Address, client.VersionStr)
	form := url.Values{}
	form.Add("username", userName)
	form.Add("oldPassword", currentPassword)
	form.Add("password", newPassword)
	var req *http.Request
	req, err = http.NewRequest("PUT", apiURL, strings.NewReader(form.Encode()))
	if err == nil {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		var resp *http.Response
		resp, err = client.Do(req)
		if err == nil {
			defer resp.Body.Close()
			decoder := json.NewDecoder(resp.Body)
			handleStatusCode(resp.StatusCode, decoder)
		}
	}
	return err
}
