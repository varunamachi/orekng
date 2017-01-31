package sqlite

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/varunamachi/orekng/olog"
)

//CreateQuery - query for creating table with table name
type CreateQuery struct {
	TableName   string
	QueryString string
}

const exists = `SELECT EXISTS (SELECT 1 FROM sqlite_master 
        WHERE type = 'table' AND name = '%s')`

var queries = [...]CreateQuery{
	CreateQuery{
		TableName: "orek_user",
		QueryString: `CREATE TABLE orek_user(
			user_name     VARCHAR( 255 ) NOT NULL,
			first_name    VARCHAR( 255 ),
			second_name   VARCHAR( 255 ),
			email         VARCHAR( 255 ) NOT NULL,
			PRIMARY KEY( user_name ),
			UNIQUE(email)
    	);`,
	},
	CreateQuery{
		TableName: "orek_user_password",
		QueryString: `CREATE TABLE orek_user_password(
    		user_name   VARCHAR( 255 ) NOT NULL,
    		hash        VARCHAR( 255 ) NOT NULL,
    		PRIMARY KEY( user_name ),
    		FOREIGN KEY( user_name ) REFERENCES orek_user( user_name ) 
				ON DELETE CASCADE
    	);`,
	},
	CreateQuery{
		TableName: "orek_user_group",
		QueryString: `CREATE TABLE orek_user_group(
			group_id	VARCHAR( 256 ) NOT NULL,
    		name        VARCHAR( 256 ) NOT NULL,
    		owner       VARCHAR( 256 ) NOT NULL,
    		description TEXT NOT NULL,
    		PRIMARY KEY( group_id ),
    		FOREIGN KEY( owner ) REFERENCES orek_user( user_name )
    	);`,
	},
	CreateQuery{
		TableName: "orek_user_to_group",
		QueryString: `CREATE TABLE orek_user_to_group(
    		group_id    VARCHAR( 256 ) NOT NULL,
    		user_name   VARCHAR( 256 ) NOT NULL,
    		PRIMARY KEY( group_id, user_name ) 
    		FOREIGN KEY( group_id ) REFERENCES orek_user_group( group_id )
				ON DELETE CASCADE,
    		FOREIGN KEY( user_name ) REFERENCES orek_user( user_name )
				ON DELETE CASCADE
    	);`,
	},
	CreateQuery{
		TableName: "orek_endpoint",
		QueryString: `CREATE TABLE orek_endpoint(
    		endpoint_id		CHAR( 36 )     NOT NULL,
    		name       		VARCHAR( 255 ) NOT NULL,
    		owner      		VARCHAR( 255 ) NOT NULL,
			owner_group		VARCHAR( 255 ) NOT NULL,
    		description		TEXT NOT NULL,
    		location   		VARCHAR( 255 ) NOT NULL,
    		visibility 		CHAR( 20 )     NOT NULL,
    		PRIMARY KEY( endpoint_id ),
    		UNIQUE(name, owner),
    		FOREIGN KEY( owner ) REFERENCES orek_user( user_name )
				ON DELETE CASCADE,
			FOREIGN KEY( owner_group ) REFERENCES orek_user_group( group_id )
				ON DELETE CASCADE
    	);`,
	},
	CreateQuery{
		TableName: "orek_variable",
		QueryString: `CREATE TABLE orek_variable(
    		variable_id  CHAR( 36 )     NOT NULL,
    		name         VARCHAR( 255 ) NOT NULL,
    		endpoint_id  CHAR( 36 )     NOT NULL,
    		description  TEXT           NOT NULL,
    		unit         CHAR( 30 )     NOT NULL,
			type		 CHAR( 30 )		NOT NULL,
    		PRIMARY KEY( variable_id ),
    		UNIQUE(endpoint_id, name),
    		FOREIGN KEY( endpoint_id ) REFERENCES orek_endpoint( endpoint_id )
				ON DELETE CASCADE
    	);`,
	},
	CreateQuery{
		TableName: "orek_parameter",
		QueryString: `CREATE TABLE orek_parameter(
    			parameter_id  	CHAR( 36 )     	NOT NULL,
    			name         	VARCHAR( 255 ) 	NOT NULL,
    			endpoint_id    	CHAR( 36 )     	NOT NULL,
    			description  	TEXT           	NOT NULL,
    			unit         	CHAR( 30 )     	NOT NULL,
				type		 	CHAR( 30 )		NOT NULL,
				permission   	CHAR( 20 )		NOT NULL,
    			PRIMARY KEY( parameter_id ),
    			UNIQUE(endpoint_id, name),
    			FOREIGN KEY( endpoint_id ) REFERENCES orek_endpoint( endpoint_id )
					ON DELETE CASCADE
    		);`,
	},
	CreateQuery{
		TableName: "orek_variable_value",
		QueryString: `CREATE TABLE orek_variable_value(
    		variable_id         CHAR( 36 ) NOT NULL,
    		value               VARCHAR( 256 ) NOT NULL,
    		time                TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    	);`,
	},
	CreateQuery{
		TableName: "orek_user_session",
		QueryString: `CREATE TABLE orek_user_session(
    		session_id          CHAR( 36 ) NOT NULL,
    		user_name			VARCHAR( 256 ) NOT NULL,
    		time                TIMESTAMP NOT NULL,
			PRIMARY KEY( session_id ),
			UNIQUE( user_name ),
			FOREIGN KEY( user_name ) REFERENCES orek_user( user_name )
				ON DELETE CASCADE
    	);`,
	},
}

//Store - represents the orek datastore
type Store struct {
	*sqlx.DB
	Path string
}

//Options - options for connecting to sqlite database
type Options struct {

	//Path - path of the sqlite database file
	Path string
}

//Init - initializes the orek datastore
func Init(options *Options) (*Store, error) {
	mdb, err := sqlx.Open("sqlite3", options.Path)
	if err == nil {
		_, err = mdb.Exec("PRAGMA foreign_keys = ON;")
	}
	if err != nil {
		olog.PrintError("SQLiteDS", err)
	} else if err = mdb.Ping(); err != nil {
		olog.Error("SQLiteDS",
			"Error: Could not connect to SQLite database: %s", err)
	} else {
		olog.Info("SQLiteDS", "Database opened successfuly")
	}
	return &Store{mdb, options.Path}, err
}

//Init - Initialize the database, creates the tables taht aren't yet created
func (sqlite *Store) Init() (err error) {
	for _, query := range queries {
		if !sqlite.tableExists(query.TableName) {
			// _, err = sqlite.Exec("PRAGMA foreign_keys = ON;")
			_, err = sqlite.Exec(query.QueryString)
			if err != nil {
				olog.Error("SQLiteDS", `Failed to create table %s: %s`,
					query.TableName, err)
			} else {
				olog.Info("SQLiteDS",
					"Table %s created successfuly", query.TableName)
			}
		} else {
			olog.Info("SQLiteDS",
				"Table %s exists, nothing to do", query.TableName)
		}
	}
	return err
}

//ClearData - clears data from all the tables
func (sqlite *Store) ClearData() (err error) {
	qtemplate := "DELETE FROM %s;"
	for _, q := range queries {
		query := fmt.Sprintf(qtemplate, q.TableName)
		_, err = sqlite.Exec(query)
		logIfError(err)
	}
	return err
}

//DeleteSchema - Drops all the tables
func (sqlite *Store) DeleteSchema() (err error) {
	qtemplate := "DROP TABLE %s;"
	// for _, q := range queries {
	for i := len(queries) - 1; i >= 0; i-- {
		tableName := queries[i].TableName
		query := fmt.Sprintf(qtemplate, tableName)
		_, err = sqlite.Exec(query)
	}
	err = sqlite.Close()
	if err == nil {
		err = os.Remove(sqlite.Path)
	}
	return err
}

func (sqlite *Store) tableExists(tableName string) (has bool) {
	query := fmt.Sprintf(exists, tableName)
	rows := sqlite.QueryRowx(query)
	err := rows.Err()
	if err == nil {
		err = rows.Scan(&has)
	}
	if err == sql.ErrNoRows {
		has = false
	} else if err != nil {
		olog.Fatal("SQLiteDS", "Could not initialize database: %s", err)
		has = false
	}
	return has
}
