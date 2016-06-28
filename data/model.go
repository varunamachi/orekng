package data

import (
	"fmt"
)

type SourceVisiblity string

const (
	//Public - the source and its properties are visible to everybody
	Public SourceVisiblity = "public"

	//GroupPrivate - the source and its properties are visible only to
	// owner's groups
	GroupPrivate SourceVisiblity = "group_private"

	//Private - the source and its properties are visible only to owner
	Private SourceVisiblity = "private"
)

//User - struct representing the user of the service
type User struct {
	Name       string `json:"name"`
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	Email      string `json:"email"`
}

//Source - represents the a data source that exposes variables and parameters
type Source struct {
	SourceID    string          `json:"sourceID"`
	Name        string          `json:"sourceName"`
	Owner       string          `json:"owner"`
	OwnerGroup  string          `json:"group"`
	Description string          `json:"description"`
	Location    string          `json:"location"`
	Visibility  SourceVisiblity `json:"access"`
}

//Variable - is an entity associated with a source that varies with time
type Variable struct {
	VariableID  string `json:"variableID"`
	Name        string `json:"variableName"`
	SourceID    string `json:"sourceID"`
	Description string `json:"description"`
	Unit        string `json:"unit"`
	Type        string `json:"type"`
}

//Parameter - is an entity associated with a source that can be changed to
//change its behaviour
type Parameter struct {
	ParameterID string `json:"parameterID"`
	Name        string `json:"variableName"`
	SourceID    string `json:"sourceID"`
	Description string `json:"description"`
	Unit        string `json:"unit"`
	Type        string `json:"type"`
	Permission  string `json:"permission"`
}

//UserGroup - is a set of users, who can share permissions
type UserGroup struct {
	Name        string `json:"userGroupName"`
	Owner       string `json:"userGroupOwner"`
	Description string `json:"userGroupDesc"`
}

func (user *User) String() string {
	return "User: " + user.Name + "[" + user.Email + "]"
}

func (source *Source) String() string {
	return "Source: " + source.Name + "[" + source.SourceID + "]"
}

func (variable *Variable) String() string {
	return "Variable: " +
		variable.Name +
		"[" +
		variable.SourceID +
		" : " +
		variable.VariableID + "]"
}

func (userGroup *UserGroup) String() string {
	return fmt.Sprintf("UserGroup: %s [Owner: %s]", userGroup.Name, userGroup.Owner)
}
