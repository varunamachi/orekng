package data

import (
	"fmt"
	"time"
)

//EndpointVisiblity - determines the visibility of a variable or parameter
type EndpointVisiblity string

const (
	//Public - the endpoint and its properties are visible to everybody
	Public EndpointVisiblity = "public"

	//GroupPrivate - the endpoint and its properties are visible only to
	// owner's groups
	GroupPrivate EndpointVisiblity = "group_private"

	//Private - the endpoint and its properties are visible only to owner
	Private EndpointVisiblity = "private"
)

//User - struct representing the user of the service
type User struct {
	Name       string `json:"name" db:"user_name"`
	FirstName  string `json:"firstName" db:"first_name"`
	SecondName string `json:"secondName" db:"second_name"`
	Email      string `json:"email" db:"email"`
}

//UserGroup - is a set of users, who can share permissions
type UserGroup struct {
	GroupID     string `json:"userGroupID" db:"group_id"`
	Name        string `json:"userGroupName" db:"name"`
	Owner       string `json:"userGroupOwner" db:"owner"`
	Description string `json:"userGroupDesc" db:"description"`
}

//Endpoint - represents the a data endpoint that exposes variables and
//parameters
type Endpoint struct {
	EndpointID  string            `json:"endpointID" db:"endpoint_id"`
	Name        string            `json:"endpointName" db:"name"`
	Owner       string            `json:"owner" db:"owner"`
	OwnerGroup  string            `json:"group" db:"owner_group"`
	Description string            `json:"description" db:"description"`
	Location    string            `json:"location" db:"location"`
	Visibility  EndpointVisiblity `json:"visibility" db:"visibility"`
}

//Variable - is an entity associated with a endpoint that varies with time
type Variable struct {
	VariableID  string `json:"variableID" db:"variable_id"`
	Name        string `json:"variableName" db:"name"`
	EndpointID  string `json:"endpointID" db:"endpoint_id"`
	Description string `json:"description" db:"description"`
	Unit        string `json:"unit" db:"unit"`
	Type        string `json:"type" db:"type"`
}

//Parameter - is an entity associated with a endpoint that can be changed to
//change its behaviour
type Parameter struct {
	ParameterID string `json:"parameterID" db:"parameter_id"`
	Name        string `json:"parameterName" db:"name"`
	EndpointID  string `json:"endpointID" db:"endpoint_id"`
	Description string `json:"description" db:"description"`
	Unit        string `json:"unit" db:"unit"`
	Type        string `json:"type" db:"type"`
	Permission  string `json:"permission" db:"permission"`
}

//Session - represents a session object
type Session struct {
	SessionID string            `json:"sessionId" db:"session_id"`
	UserID    string            `json:"userId" db:"user"`
	StartTime time.Time         `json:"startTime" db:"start_time"`
	Props     map[string]string `json:"properties"`
}

func (user *User) String() string {
	return "User: " + user.Name + "[" + user.Email + "]"
}

func (endpoint *Endpoint) String() string {
	return "Endpoint: " + endpoint.Name + "[" + endpoint.EndpointID + "]"
}

func (variable *Variable) String() string {
	return "Variable: " +
		variable.Name +
		"[" +
		variable.EndpointID +
		" : " +
		variable.VariableID + "]"
}

func (parameter *Parameter) String() string {
	return "Variable: " +
		parameter.Name +
		"[" +
		parameter.EndpointID +
		" : " +
		parameter.ParameterID + "]"
}

func (userGroup *UserGroup) String() string {
	return fmt.Sprintf("UserGroup: %s [Owner: %s]",
		userGroup.Name, userGroup.Owner)
}
