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
	Name       string `json:"name" db:"user_name"`
	FirstName  string `json:"firstName" db:"first_name"`
	SecondName string `json:"secondName" db:"second_name"`
	Email      string `json:"email" db:"email"`
}

//UserGroup - is a set of users, who can share permissions
type UserGroup struct {
	GroupID     string `json:"userGroupID" db:"id"`
	Name        string `json:"userGroupName" db:"name"`
	Owner       string `json:"userGroupOwner" db:"owner"`
	Description string `json:"userGroupDesc" db:"description"`
}

//Source - represents the a data source that exposes variables and parameters
type Source struct {
	SourceID    string          `json:"sourceID" db:"source_id"`
	Name        string          `json:"sourceName" db:"name"`
	Owner       string          `json:"owner" db:"owner"`
	OwnerGroup  string          `json:"group" db:"owner_group"`
	Description string          `json:"description" db:"description"`
	Location    string          `json:"location" db:"location"`
	Visibility  SourceVisiblity `json:"visibility" db:"visibility"`
}

//Variable - is an entity associated with a source that varies with time
type Variable struct {
	VariableID  string `json:"variableID" db:"variable_id"`
	Name        string `json:"variableName" db:"name"`
	SourceID    string `json:"sourceID" db:"source_id"`
	Description string `json:"description" db:"description"`
	Unit        string `json:"unit" db:"unit"`
	Type        string `json:"type" db:"type"`
}

//Parameter - is an entity associated with a source that can be changed to
//change its behaviour
type Parameter struct {
	ParameterID string `json:"parameterID" db:"parameter_id"`
	Name        string `json:"parameterName" db:"name"`
	SourceID    string `json:"sourceID" db:"source_id"`
	Description string `json:"description" db:"description"`
	Unit        string `json:"unit" db:"unit"`
	Type        string `json:"type" db:"type"`
	Permission  string `json:"permission" db:"permission"`
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
