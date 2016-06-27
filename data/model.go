package data

import (
    "fmt"
)

type User struct {
	Name       string `json:"name"`
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	Email      string `json:"email"`
}

type Source struct {
	SourceId    string `json:"sourceId"`
	Name        string `json:"sourceName"`
	Owner       string `json:"owner"`
    OwnerGroup  string `json:"group"`
	Description string `json:"description"`
	Location    string `json:"location"`
    Access      string `json:"access"`
}

type Variable struct {
	VariableId  string `json:"variableId"`
	Name        string `json:"variableName"`
	SourceId    string `json:"sourceId"`
	Description string `json:"description"`
	Unit        string `json:"unit"`
    Type        string `json:"type"`
}

type Parameter struct {
    Variable    
}

type UserGroup struct {
	Name        string `json:"userGroupName"`
	Owner       string `json:"userGroupOwner"`
	Description string `json:"userGroupDesc"`
}

func (user *User) String() string {
	return "User: " + user.Name + "[" + user.Email + "]"
}

func (source *Source) String() string {
	return "Source: " + source.Name + "[" + source.SourceId + "]"
}

func (variable *Variable) String() string {
	return "Variable: " + 
        variable.Name + 
        "[" + 
        variable.SourceId + 
        " : " + 
        variable.VariableId + 
        "]"
}

func (userGroup *UserGroup) String() string {
	return fmt.Sprintf("UserGroup: %s [Owner: %s]", userGroup.Name, userGroup.Owner)
}
