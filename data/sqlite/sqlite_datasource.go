package sqlite

import (
	"log"

	"time"

	"github.com/jmoiron/sqlx"
	"github.com/varunamachi/orekng/data"
)

func logIfError(err error) {
	if err != nil {
		log.Printf("Error: %s", err)
	}
}

//GetAllUsers - Gives all user entries in the database
func (sqlite *DataStore) GetAllUsers() (users []*data.User, err error) {
	users = make([]*data.User, 0, 20)
	err = sqlite.Select(&users, "SELECT * FROM orek_user ORDER BY user_name")
	return users, err
}

//GetUser - Gives the user with given userName from database
func (sqlite *DataStore) GetUser(userName string) (user *data.User, err error) {
	user = &data.User{}
	queryStr := `SELECT * FROM orek_user WHERE user_name = ?`
	err = sqlite.Select(user, queryStr, userName)
	return user, err
}

// GetUserWithEmail - Gives the user entry with given EMail id
func (sqlite *DataStore) GetUserWithEmail(
	email string) (user *data.User, err error) {
	user = &data.User{}
	queryStr := `SELECT * FROM orek_user WHERE email = ?`
	err = sqlite.Select(user, queryStr, email)
	logIfError(err)
	return user, err
}

//CreateUser - creates a user entry in the database with given User object
func (sqlite *DataStore) CreateUser(user *data.User) (err error) {
	queryStr := `INSERT INTO orek_user( 
		user_name,  
		first_name, 
		second_name,
		email,      
	) VALUES (
		:user_name,
		:first_name,
		:second_name,
		:email
	)`
	_, err = sqlite.NamedExec(queryStr, user)
	logIfError(err)
	return err
}

//UpdateUser - Upadates the user entry in the database with the information
//in the given user object
func (sqlite *DataStore) UpdateUser(user *data.User) (err error) {
	queryStr := `UPDATE orek_user SET
		first_name = :first_name,
		second_name = :second_name,
		email = :email
		WHERE user_name = :user_name
	`
	_, err = sqlite.NamedExec(queryStr, user)
	logIfError(err)
	return err
}

//DeleteUser - deletes the user entry with given user name
func (sqlite *DataStore) DeleteUser(userName string) (err error) {
	queryStr := `DELETE FROM orek_user WHERE user_id = ?`
	_, err = sqlite.Exec(queryStr, userName)
	logIfError(err)
	return err
}

//GetAllEndpoints - Gives all the data endpoints which have entries in database
func (sqlite *DataStore) GetAllEndpoints() (endpoints []*data.Endpoint, err error) {
	queryStr := `SELECT * FROM orek_endpoint ORDER BY endpoint_id`
	endpoints = make([]*data.Endpoint, 0, 100)
	err = sqlite.Select(&endpoints, queryStr)
	logIfError(err)
	return endpoints, err
}

//GetEndpoint - Gives data endpoint object correspoing to the database entry with
// given name
func (sqlite *DataStore) GetEndpoint(endpointID string) (endpoint *data.Endpoint, err error) {
	endpoint = &data.Endpoint{}
	queryStr := `SELECT * FROM orek_endpoint WHERE endpoint_id = ?`
	err = sqlite.Select(endpoint, queryStr, endpointID)
	logIfError(err)
	return endpoint, err
}

//CreateEndpoint - Creates a endpoint entry in database according to the endpoint
//object
func (sqlite *DataStore) CreateEndpoint(endpoint *data.Endpoint) (err error) {
	queryStr := `INSERT INTO orek_endpoints(
		endpoint_id,
		name ,
		owner,
		owner_group,
		description,
		location,
		visibility,
	) VALUES (
		:endpoint_id,
		:name,
		:owner,
		:owner_group,
		:description,
		:location,
		:visibility
	)`
	_, err = sqlite.NamedExec(queryStr, endpoint)
	logIfError(err)
	return err
}

//UpdateEndpoint - Updates the endpoint entry in database with information provided
//in the endpoint object
func (sqlite *DataStore) UpdateEndpoint(endpoint *data.Endpoint) (err error) {
	queryStr := `UPDATE orek_endpoint SET
			name = :name,
			owner = :owner,
			owner_group = :owner_group,
			description = :description,
			location = :location,
			visibility = :visibility
		WHERE endpoint_id = :endpoint_id`
	_, err = sqlite.NamedExec(queryStr, endpoint)
	logIfError(err)
	return err
}

//DeleteEndpoint - deletes an endpoint
func (sqlite *DataStore) DeleteEndpoint(endpointID string) error {
	queryStr := `DELETE FROM orek_endpoit WHERE endpoint_id = ?`
	_, err := sqlite.Exec(queryStr, endpointID)
	logIfError(err)
	return err
}

//GetAllVariables - Gives list of all variables
func (sqlite *DataStore) GetAllVariables() (variables []*data.Variable, err error) {
	queryStr := `SELECT * FROM orek_variable ORDER BY variable_id`
	variables = make([]*data.Variable, 0, 100)
	err = sqlite.Select(&variables, queryStr)
	logIfError(err)
	return variables, err
}

//GetVariablesForEndpoint - Gives all the variables exported by an endpoint
func (sqlite *DataStore) GetVariablesForEndpoint(
	endpointID string) (variables []*data.Variable, err error) {
	queryStr := `SELECT * FROM orek_variable WHERE endpoint_id = ?
		ORDER BY variable_id`
	variables = make([]*data.Variable, 0, 100)
	err = sqlite.Select(&variables, queryStr, endpointID)
	logIfError(err)
	return variables, err
}

//GetVariable - Gives the variable with the given ID
func (sqlite *DataStore) GetVariable(variableID string) (variable *data.Variable, err error) {
	queryStr := `SELECT * FROM orek_variable WHERE variable_id = ?`
	variable = &data.Variable{}
	err = sqlite.Select(variable, queryStr, variableID)
	logIfError(err)
	return variable, err
}

//CreateVariable - creates a variable in the datasource
func (sqlite *DataStore) CreateVariable(variable *data.Variable) (err error) {
	queryStr := `INSERT INTO orek_variable(
		variable_id,
    	name,
    	endpoint_id,
    	description,
    	unit,
		type
	) VALUES (
		:variable_id,
    	:name,
    	:endpoint_id,
    	:description,
    	:unit,
		:type		
	)`
	_, err = sqlite.NamedExec(queryStr, variable)
	logIfError(err)
	return err
}

//UpdateVariable - updates a variable in the datasource
func (sqlite *DataStore) UpdateVariable(variable *data.Variable) (err error) {
	queryStr := `UPDATE orek_variable SET
    		name = :name,
    		endpoint_id = :endpoint_id,
    		description = :description,
    		unit = :unit ,
			type = :type
		WHERE variable_id = :variable_id
	`
	_, err = sqlite.NamedExec(queryStr, variable)
	logIfError(err)
	return err
}

//DeleteVariable - delete a variable from the datasource
func (sqlite *DataStore) DeleteVariable(variableID string) (err error) {
	queryStr := `DELETE FROM orek_variable WHERE variable_id = ?`
	_, err = sqlite.Exec(queryStr, variableID)
	logIfError(err)
	return err
}

//GetAllParameters - Gives list of all parameters
func (sqlite *DataStore) GetAllParameters() (parameters []*data.Parameter, err error) {
	queryStr := `SELECT * FROM orek_parameter ORDER BY parameter_id`
	parameters = make([]*data.Parameter, 0, 100)
	err = sqlite.Select(&parameters, queryStr)
	logIfError(err)
	return parameters, err
}

//GetParametersForEndpoint - Gives all the parameters exported by an endpoint
func (sqlite *DataStore) GetParametersForEndpoint(
	endpointID string) (parameters []*data.Parameter, err error) {
	queryStr := `SELECT * FROM orek_parameter WHERE endpoint_id = ?
		ORDER BY parameter_id`
	parameters = make([]*data.Parameter, 0, 100)
	err = sqlite.Select(&parameters, queryStr, endpointID)
	logIfError(err)
	return parameters, err
}

//GetParameter - Gives the parameter with the given ID
func (sqlite *DataStore) GetParameter(parameterID string) (parameter *data.Parameter, err error) {
	queryStr := `SELECT * FROM orek_parameter WHERE parameter_id = ?`
	parameter = &data.Parameter{}
	err = sqlite.Select(parameter, queryStr, parameterID)
	logIfError(err)
	return parameter, err
}

//CreateParameter - creates a parameter in the datasource
func (sqlite *DataStore) CreateParameter(parameter *data.Parameter) (err error) {
	queryStr := `INSERT INTO orek_parameter(
		parameter_id,
    	name,
    	endpoint_id,
    	description,
    	unit,
		type
	) VALUES (
		:parameter_id,
    	:name,
    	:endpoint_id,
    	:description,
    	:unit,
		:type		
	)`
	_, err = sqlite.NamedExec(queryStr, parameter)
	logIfError(err)
	return err
}

//UpdateParameter - updates a parameter in the datasource
func (sqlite *DataStore) UpdateParameter(parameter *data.Parameter) (err error) {
	queryStr := `UPDATE orek_parameter SET
    		name = :name,
    		endpoint_id = :endpoint_id,
    		description = :description,
    		unit = :unit ,
			type = :type
		WHERE parameter_id = :parameter_id
	`
	_, err = sqlite.NamedExec(queryStr, parameter)
	logIfError(err)
	return err
}

//DeleteParameter - delete a parameter from the datasource
func (sqlite *DataStore) DeleteParameter(parameterID string) (err error) {
	queryStr := `DELETE FROM orek_parameter WHERE parameter_id = ?`
	_, err = sqlite.Exec(queryStr, parameterID)
	logIfError(err)
	return err
}

//GetAllUserGroups - gets the list of user group from the database
func (sqlite *DataStore) GetAllUserGroups() (userGroups []*data.UserGroup, err error) {
	queryStr := `SELECT * FROM orek_user_group ORDER BY group_id`
	userGroups = make([]*data.UserGroup, 0, 100)
	err = sqlite.Select(&userGroups, queryStr)
	logIfError(err)
	return userGroups, err
}

//GetUserGroup - get an instance of user group for give group name
func (sqlite *DataStore) GetUserGroup(
	userGroupName string) (userGroup *data.UserGroup, err error) {
	queryStr := `SELECT * FROM orek_user_group WHERE group_id = ?`
	userGroup = &data.UserGroup{}
	err = sqlite.Select(userGroup, queryStr, userGroupName)
	logIfError(err)
	return nil, err
}

//CreateUserGroup - creates an user group with give details
func (sqlite *DataStore) CreateUserGroup(userGroup *data.UserGroup) (err error) {
	queryStr := `INSERT INTO orek_user_group(
		group_id,
		name,
		owner,
		description 
	) VALUES (
		:group_id,
		:name,
		:owner,
		:description 
	)`
	_, err = sqlite.NamedExec(queryStr, userGroup)
	logIfError(err)
	return err
}

//UpdateUserGroup - Updates an existing user group with details from the
//given object
func (sqlite *DataStore) UpdateUserGroup(userGroup *data.UserGroup) (err error) {
	queryStr := `UPDATE orek_user_group SET
			name = :name,
			owner = :owner,
			description = :description
		WHERE group_id = :group_id`
	_, err = sqlite.NamedExec(queryStr, userGroup)
	logIfError(err)
	return err
}

//DeleteUserGroup - deletes an user group with the given group name
func (sqlite *DataStore) DeleteUserGroup(userGroupName string) (err error) {
	queryStr := `DELETE FROM orek_user_group WHERE group_id = ?`
	_, err = sqlite.Exec(queryStr, userGroupName)
	logIfError(err)
	return err
}

//AddUserToGroup - adds user with given user name to a group with given group
//name
func (sqlite *DataStore) AddUserToGroup(userName, groupID string) (err error) {
	queryStr := `INSERT INTO orek_user_to_group( 
		group_id,
		user_name
	) VALUES (
		?,
		?
	)`
	_, err = sqlite.Exec(queryStr, groupID, userName)
	logIfError(err)
	return err
}

//RemoveUserFromGroup - disassociates user with given user name from group with
//given group name
func (sqlite *DataStore) RemoveUserFromGroup(userName, groupID string) (err error) {
	queryStr := `DELETE FROM orek_user_to_group 
		WHERE group_id = ? AND user_name = ?`
	_, err = sqlite.Exec(queryStr, groupID, userName)
	logIfError(err)
	return err
}

//GetUsersInGroup - gives a list of users who are associated with the group
//with given group name
func (sqlite *DataStore) GetUsersInGroup(
	groupID string) (users []*data.User, err error) {
	queryStr := `SELECT * FROM orek_user WHERE user_name IN(
		SELECT user_name FROM orek_user_to_group WHERE group_id = ?
	)`
	users = make([]*data.User, 0, 100)
	err = sqlite.Select(users, queryStr, groupID)
	logIfError(err)
	return users, err
}

//GetGroupsForUser - Gives a list of groups with which the user with given user
//name is associated
func (sqlite *DataStore) GetGroupsForUser(
	userName string) (groups []*data.UserGroup, err error) {
	queryStr := `SELECT * FROM orek_user_group WHERE group_id IN (
		SELECT group_id FROM orek_user_to_group WHERE user_name = ?
	)`
	groups = make([]*data.UserGroup, 0, 100)
	err = sqlite.Select(&groups, queryStr, userName)
	logIfError(err)
	return groups, err
}

//AddVariableValue - Adds value to list of values of a variable
func (sqlite *DataStore) AddVariableValue(variableID, value string) (err error) {
	queryStr := `INSERT INTO orek_variable_value(
		variable_id,
		value      
	) VALUES (
		?,
		?
	)`
	_, err = sqlite.Exec(queryStr, variableID, value)
	logIfError(err)
	return err
}

//ClearValuesForVariable - clears values from the list of values associated with
//the variable with given variable id
func (sqlite *DataStore) ClearValuesForVariable(variableID string) (err error) {
	queryStr := `DELETE * FROM orek_variable_value`
	_, err = sqlite.Exec(queryStr)
	logIfError(err)
	return err
}

//GetValuesForVariable - Gives list of values associated with a variable with
//given variable id
func (sqlite *DataStore) GetValuesForVariable(
	variableID string) (values []*string, err error) {
	queryStr := `SELECT value FROM orek_variable_value WHERE variable_id = ?`
	values = make([]*string, 0, 100)
	err = sqlite.Select(values, queryStr, variableID)
	logIfError(err)
	return values, err
}

//CreateUserSession - creates a orek user session
func (sqlite *DataStore) CreateUserSession(session *data.Session) (err error) {
	queryStr := `INSERT INTO orek_user_session(
		session_id,
		user_name,
		time
	) VALUES (
		:session_id,
		:user_name,
		:time
	)`
	_, err = sqlite.NamedExec(queryStr, session)
	logIfError(err)
	return err
}

//GetUserSession - gives a session with given session ID
func (sqlite *DataStore) GetUserSession(
	sessionID string) (session *data.Session, err error) {
	queryStr := `SELECT * FROM orek_user_session WHERE session_id = ?`
	session = &data.Session{}
	err = sqlite.Select(session, queryStr, sessionID)
	logIfError(err)
	return session, err
}

//RemoveUserSession - removes a user session with given ID
func (sqlite *DataStore) RemoveUserSession(sessionID string) (err error) {
	queryStr := `DELETE FROM orek_user_session WHERE session_id = ?`
	_, err = sqlite.Exec(queryStr, sessionID)
	logIfError(err)
	return err
}

//ClearExpiredSessions - clears sessions that have exceeded expiry time i.e.
//sessionStartSize - currentSize > givenExperyTime
func (sqlite *DataStore) ClearExpiredSessions(expiryTimeMillis int64) (err error) {
	queryStrOne := `SELECT * FROM orek_user_session`
	sessions := make([]*data.Session, 0, 100)
	expired := make([]string, 0, 100)
	err = sqlite.Select(&sessions, queryStrOne)
	if err == nil {
		for _, sess := range sessions {
			if sess.StartTime.Unix()+expiryTimeMillis > time.Now().Unix() {
				expired = append(expired, sess.SessionID)
			}
		}
		var query string
		var args []interface{}
		query, args, err = sqlx.In(
			"DELETE FROM orek_user_session WHERE session_id IN (?)", expired)
		if err == nil {
			query = sqlite.Rebind(query)
			_, err = sqlite.Exec(query, args...)
		}
	}
	logIfError(err)
	return err
}

//SetPasswordHash - stores password hash for an user in the database
func SetPasswordHash(userName, passwordHash string) (err error) {

	return err
}

//GetPasswordHash - Retrieves password hash for an user from the database
func GetPasswordHash(userName string) (hash string, err error) {

	return hash, err
}
