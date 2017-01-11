package sqlite

import "github.com/varunamachi/orekng/data"
import "log"

//GetAllUsers - Gives all user entries in the database
func (sqlite *DataStore) GetAllUsers() ([]*data.User, error) {
	users := make([]*data.User, 0, 20)
	err := sqlite.Select(&users, "SELECT * FROM orek_user ORDER BY user_name")
	if err != nil {
		log.Printf("Error: %s", err)
	}
	return users, err
}

//GetUser - Gives the user with given userName from database
func (sqlite *DataStore) GetUser(userName string) (*data.User, error) {
	user := &data.User{}
	query := `SELECT * FROM orek_user WHERE user_name = $1`
	err := sqlite.Select(user, query, userName)
	return user, err
}

// GetUserWithEmail - Gives the user entry with given EMail id
func (sqlite *DataStore) GetUserWithEmail(email string) (*data.User, error) {
	user := &data.User{}
	query := `SELECT * FROM orek_user WHERE email = $1`
	err := sqlite.Select(user, query, email)
	return user, err
}

//CreateUser - creates a user entry in the database with given User object
func (sqlite *DataStore) CreateUser(user *data.User) error {
	query := `INSERT INTO orek_user( 
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
	_, err := sqlite.NamedExec(query, user)
	if err != nil {
		log.Printf("Error:DB: %v", err)
	}
	return err
}

//UpdateUser - Upadates the user entry in the database with the information
//in the given user object
func (sqlite *DataStore) UpdateUser(user *data.User) error {
	query := `UPDATE orek_user SET
		first_name = :first_name,
		second_name = :second_name,
		email = :email
		WHERE user_name = :user_name
	`
	_, err := sqlite.NamedExec(query, user)
	if err != nil {
		log.Printf("Error:DB: %v", err)
	}
	return nil
}

//DeleteUser - deletes the user entry with given user name
func (sqlite *DataStore) DeleteUser(userName string) error {
	query := `DELETE FROM orek_user WHERE user_id = ?`
	_, err := sqlite.Exec(query, userName)
	if err != nil {
		log.Printf("Error:DB: %v", err)
	}
	return err
}

//GetAllEndpoints - Gives all the data endpoints which have entries in database
func (sqlite *DataStore) GetAllEndpoints() ([]*data.Endpoint, error) {
	query := `SELECT * FROM orek_endpoint ORDER BY endpoint_id`
	endpoints := make([]*data.Endpoint, 0, 100)
	err := sqlite.Select(&endpoints, query)
	return endpoints, err
}

//GetEndpoint - Gives data endpoint object correspoing to the database entry with
// given name
func (sqlite *DataStore) GetEndpoint(endpointID string) (*data.Endpoint, error) {
	endpoint := &data.Endpoint{}
	query := `SELECT * FROM orek_endpoint WHERE endpoint_id = ?`
	err := sqlite.Select(endpoint, query)
	return endpoint, err
}

//CreateEndpoint - Creates a endpoint entry in database according to the endpoint
//object
func (sqlite *DataStore) CreateEndpoint(endpoint *data.Endpoint) error {
	query := `INSERT INTO orek_endpoints(
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
	_, err := sqlite.NamedExec(query, endpoint)
	if err != nil {
		log.Printf("Error:DB: %v", err)
	}
	return err
}

//UpdateEndpoint - Updates the endpoint entry in database with information provided
//in the endpoint object
func (sqlite *DataStore) UpdateEndpoint(endpoint *data.Endpoint) error {
	query := `UPDATE orek_endpoint SET
			name = :name,
			owner = :owner,
			owner_group = :owner_group,
			description = :description,
			location = :location,
			visibility = :visibility
		WHERE endpoint_id = :endpoint_id`
	_, err := sqlite.NamedExec(query, endpoint)
	if err != nil {
		log.Printf("Error:DB: %v", err)
	}
	return err
}

//DeleteEndpoint - deletes an endpoint
func (sqlite *DataStore) DeleteEndpoint(endpointID string) error {
	query := `DELETE FROM orek_endpoit WHERE endpoint_id = ?`
	_, err := sqlite.Exec(query, endpointID)
	if err != nil {
		log.Printf("Error:DB: %v", err)
	}
	return err
}

//GetAllVariables - Gives list of all variables
func (sqlite *DataStore) GetAllVariables() ([]*data.Variable, error) {
	query := `SELECT * FROM orek_variable ORDER BY variable_id`
	variables := make([]*data.Variable, 0, 100)
	err := sqlite.Select(&variables, query)
	if err != nil {
		log.Printf("Error:DB: %v", err)
	}
	return variables, err
}

//GetVariablesForEndpoint - Gives all the variables exported by an endpoint
func (sqlite *DataStore) GetVariablesForEndpoint(
	endpointID string) ([]*data.Variable, error) {
	// query := `SELECT * FROM orek_variable WHERE endpoint_id = ?`

	return nil, nil
}

func (sqlite *DataStore) GetVariable(variableID string) (*data.Variable, error) {
	return nil, nil
}

func (sqlite *DataStore) CreateVariable(variable *data.Variable) error {
	return nil
}

func (sqlite *DataStore) UpdateVariable(variable *data.Variable) error {
	return nil
}

func (sqlite *DataStore) DeleteVariable(variableID string) error {
	return nil
}

func (sqlite *DataStore) GetAllUserGroups() ([]*data.UserGroup, error) {
	return nil, nil
}

func (sqlite *DataStore) GetUserGroup(userGroupName string) (*data.UserGroup, error) {
	return nil, nil
}

func (sqlite *DataStore) CreateUserGroup(userGroup *data.UserGroup) error {
	return nil
}

func (sqlite *DataStore) UpdateUserGroup(userGroup *data.UserGroup) error {
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

func (sqlite *DataStore) GetUsersInGroup(groupName string) ([]*data.User, error) {
	return nil, nil
}

func (sqlite *DataStore) GetGroupsForUser(userName string) ([]*data.UserGroup, error) {
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
