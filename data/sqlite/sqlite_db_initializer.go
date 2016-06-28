package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const EXISTS = `SELECT COUNT (*) FROM sqlite_master 
        WHERE type='table' AND name= ? ;`

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
    FOREIGN KEY( user_name ) REFERENCES orek_user( user_name ),
    FOREIGN KEY( email ) REFERENCES orek_user( email )
    );`,

	`CREATE TABLE orek_source(
    source_id     CHAR( 36 )     NOT NULL,
    name          VARCHAR( 255 ) NOT NULL,
    owner         VARCHAR( 255 ) NOT NULL,
    description   TEXT,
    location      VARCHAR( 255 ) NOT NULL,
    access        CHAR( 20 )     NOT NULL,
    PRIMARY KEY( source_id ),
    UNIQUE(name, owner),
    FOREIGN KEY( owner ) REFERENCES orek_user( user_name )
    );`,

	`CREATE TABLE orek_variable(
    variable_id  CHAR( 36 )     NOT NULL,
    name         VARCHAR( 255 ) NOT NULL,
    source_id    CHAR( 36 )     NOT NULL,
    description  TEXT           NOT NULL,
    unit         CHAR( 30 )     NOT NULL,
    PRIMARY KEY( variable_id ),
    UNIQUE(source_id, name),
    FOREIGN KEY( source_id ) REFERENCES orek_source( source_id )
    );`,

	`CREATE TABLE orek_user_group(
    name        VARCHAR( 256 ) NOT NULL,
    owner       VARCHAR( 256 ) NOT NULL,
    description TEXT NOT NULL,
    PRIMARY KEY( name ),
    FOREIGN KEY( owner ) REFERENCES orek_user( user_name )
    );`,

	`CREATE TABLE orek_user_to_group(
    group_name  VARCHAR( 256 ) NOT NULL,
    user_name   VARCHAR( 256 ) NOT NULL,
    FOREIGN KEY( group_name ) REFERENCES orek_user_group( name ),
    FOREIGN KEY( user_name ) REFERENCES orek_user( user_name ),
    PRIMARY KEY( group_name, user_name )
    );`,

	`CREATE TABLE orek_variable_group(
    group_id    CHAR( 36 )     NOT NULL,
    name        VARCHAR( 256 ) NOT NULL,
    owner       VARCHAR( 256 ) NOT NULL,
    description TEXT           NOT NULL,
    access      CHAR( 20 )     NOT NULL DEFAULT "public",
    PRIMARY KEY( group_id ),
    UNIQUE(name, owner),
    FOREIGN KEY( owner ) REFERENCES orek_user( user_name )
    );`,

	`CREATE TABLE orek_variable_to_group(
    var_group_id        CHAR( 36 ) NOT NULL,
    variable_id         CHAR( 36 ) NOT NULL,
    PRIMARY KEY( var_group_id, variable_id ),
    FOREIGN KEY( var_group_id ) REFERENCES orek_variable_group( group_id ),
    FOREIGN KEY( variable_id ) REFERENCES orek_variable( variable_id )
    );`,

	`CREATE TABLE orek_variable_value(
    variable_id         CHAR( 36 ) NOT NULL,
    value               VARCHAR( 256 ) NOT NULL,
    time                TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`,

	`CREATE INDEX idx_orek_var_value ON orek_variable_value( variable_id );`,

	//	`CREATE TABLE orel_session(
	//		session_id		CHAR( 36 ) NOT NULL,
	//		user_name		VARCHAR( 256 ) NOT NULL,
	//		start_time		TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	//		data			TEXT
	//	);`,
}

type MysqlOptions struct {
	UserName string
	Password string
	Host     string
	Port     string
	DbName   string
}

func MysqlInit(options *MysqlOptions) (*MysqlDb, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		options.UserName,
		options.Password,
		options.Host,
		options.Port)
	db, err := sql.Open("mysql", connStr)
	var mdb *MysqlDb
	if err == nil {
		defer db.Close()
		if err = db.Ping(); err == nil {
			row := db.QueryRow(EXISTS, options.DbName)
			var count int
			err = row.Scan(&count)
			if err == nil {
				if count == 0 {
					mdb, err = mysqlCreate(options, db)
				} else {
					mdb, err = mysqlConnect(options)
				}
			}
		}
	}
	if err != nil {
		log.Print("Could not connect to mysql database", err)
	}
	return mdb, err
}

func mysqlConnect(options *MysqlOptions) (*MysqlDb, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		options.UserName,
		options.Password,
		options.Host,
		options.Port,
		options.DbName)
	mdb, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Print("Could not connect to mysql database")
	} else if err = mdb.Ping(); err != nil {
		log.Print("Could not connect to mysql database")
	} else {
		log.Print("Database opened successfuly")
	}
	return &MysqlDb{mdb}, err
}

func mysqlCreate(options *MysqlOptions, db *sql.DB) (*MysqlDb, error) {
	query := fmt.Sprintf("CREATE DATABASE %s;", options.DbName)
	//	_, err := db.Exec(`CREATE DATABASE ?;`, options.DbName)
	_, err := db.Exec(query)
	var mdb *MysqlDb
	if err == nil {
		mdb, err = mysqlConnect(options)
		if err == nil {
			for index, query := range queries {
				_, err = mdb.Exec(query)
				if err != nil {
					log.Printf(`Failed to create database: query %d failed`,
						index)
					break
				}
			}
		}
	}
	if err == nil {
		log.Printf(`Database %s created successfully`, options.DbName)
	} else {
		log.Printf(`Could not create database %s: %v`, options.DbName, err)
	}
	return mdb, err
}
