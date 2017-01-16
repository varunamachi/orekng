package sqlite

import (
	"log"

	"github.com/jmoiron/sqlx"
)

const exists = `SELECT COUNT (*) FROM sqlite_master 
        WHERE type = 'table' AND name = orek_user;`

var queries = [...]string{

	`CREATE TABLE orek_user(
		user_name     VARCHAR( 255 ) NOT NULL,
		first_name    VARCHAR( 255 ),
		second_name   VARCHAR( 255 ),
		email         VARCHAR( 255 ) NOT NULL,
		PRIMARY KEY( user_name ),
		UNIQUE(email)
    );`,

	`CREATE TABLE orek_user_identity(
    	user_name   VARCHAR( 255 ) NOT NULL,
    	email       VARCHAR( 255 ) NOT NULL,
    	digest      VARCHAR( 255 ) NOT NULL,
    	PRIMARY KEY( user_name ),
    	UNIQUE(email),
    	FOREIGN KEY( user_name ) REFERENCES orek_user( user_name ) 
			ON DELETE CASCADE,
    	FOREIGN KEY( email ) REFERENCES orek_user( email )
			ON DELETE CASCADE
    );`,

	`CREATE TABLE orek_user_group(
		group_id	VARCHAR( 256 ) NOT NULL
    	name        VARCHAR( 256 ) NOT NULL,
    	owner       VARCHAR( 256 ) NOT NULL,
    	description TEXT NOT NULL,
    	PRIMARY KEY( group_id ),
    	FOREIGN KEY( owner ) REFERENCES orek_user( user_name )
    );`,

	`CREATE TABLE orek_user_to_group(
    	group_id    VARCHAR( 256 ) NOT NULL,
    	user_name   VARCHAR( 256 ) NOT NULL,
    	FOREIGN KEY( group_id ) REFERENCES orek_user_group( group_id )
			ON DELETE CASCADE,
    	FOREIGN KEY( user_name ) REFERENCES orek_user( user_name ),
    	PRIMARY KEY( group_id, user_name )
			ON DELETE CASCADE
    );`,

	`CREATE TABLE orek_endpoint(
    	endpoint_id		CHAR( 36 )     NOT NULL,
    	name       		VARCHAR( 255 ) NOT NULL,
    	owner      		VARCHAR( 255 ) NOT NULL,
		owner_group		VARCHAR( 255 ) NOT NULL
    	description		TEXT,
    	location   		VARCHAR( 255 ) NOT NULL,
    	visibility 		CHAR( 20 )     NOT NULL,
    	PRIMARY KEY( endpoint_id ),
    	UNIQUE(name, owner),
    	FOREIGN KEY( owner ) REFERENCES orek_user( user_name )
			ON DELETE CASCADE
		FOREIGN KEY( owner_group ) REFERENCES orek_user_group( group_id )
			ON DELETE CASCADE
    );`,

	`CREATE TABLE orek_variable(
    	variable_id  CHAR( 36 )     NOT NULL,
    	name         VARCHAR( 255 ) NOT NULL,
    	endpoint_id    CHAR( 36 )     NOT NULL,
    	description  TEXT           NOT NULL,
    	unit         CHAR( 30 )     NOT NULL,
		type		 CHAR( 30 )		NOT NULL,
    	PRIMARY KEY( variable_id ),
    	UNIQUE(endpoint_id, name),
    	FOREIGN KEY( endpoint_id ) REFERENCES orek_endpoint( endpoint_id )
			ON DELETE CASCADE
    );`,

	`CREATE TABLE orek_parameter(
    	parameter_id  CHAR( 36 )     NOT NULL,
    	name         VARCHAR( 255 ) NOT NULL,
    	endpoint_id    CHAR( 36 )     NOT NULL,
    	description  TEXT           NOT NULL,
    	unit         CHAR( 30 )     NOT NULL,
		type		 CHAR( 30 )		NOT NULL,
		permission   CHAR( 20 )		NOT NULL,
    	PRIMARY KEY( parameter_id ),
    	UNIQUE(endpoint_id, name),
    	FOREIGN KEY( endpoint_id ) REFERENCES orek_endpoint( endpoint_id )
			ON DELETE CASCADE
    );`,

	`CREATE TABLE orek_variable_value(
    variable_id         CHAR( 36 ) NOT NULL,
    value               VARCHAR( 256 ) NOT NULL,
    time                TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`,

	`CREATE TABLE orek_session(
    session_id          CHAR( 36 ) NOT NULL,
    user_id				VARCHAR( 256 ) NOT NULL,
    time                TIMESTAMP NOT NULL
    );`,
	`CREATE INDEX idx_orek_session ON orek_session( session_id );`,
}

//DataStore - represents the orek datastore
type DataStore struct {
	*sqlx.DB
}

//Options - options for connecting to sqlite database
type Options struct {

	//Path - path of the sqlite database file
	Path string
}

//Init - initializes the orek datastore
func Init(options *Options) (*DataStore, error) {
	db, err := sqlx.Connect("sqlite3", options.Path)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer db.Close()
	var sqliteDB *DataStore
	if err = db.Ping(); err == nil {
		row := db.QueryRow(exists)
		var count int
		err = row.Scan(&count)
		if err == nil {
			if count == 0 {
				sqliteDB, err = create(options, db)
			} else {
				sqliteDB, err = connect(options)
			}
		} else {
			log.Print(err)
			return nil, err
		}
	}
	return sqliteDB, err
}

//connect - connects to a sqlite database file
func connect(options *Options) (*DataStore, error) {
	mdb, err := sqlx.Open("sqlite3", options.Path)
	if err != nil {
		log.Print(err)
	} else if err = mdb.Ping(); err != nil {
		log.Printf("Could not connect to mysql database: %s", err)
	} else {
		log.Print("Database opened successfuly")
	}
	return &DataStore{mdb}, err
}

//create - connects to a sqlite database file and creates Orek schema
func create(options *Options, db *sqlx.DB) (*DataStore, error) {
	mdb, err := connect(options)
	if err == nil {
		for index, query := range queries {
			_, err = mdb.Exec(query)
			if err != nil {
				log.Printf(`Failed to create database, query %d failed: %s`,
					index, err)
				break
			}
		}
	}
	return mdb, err
}
