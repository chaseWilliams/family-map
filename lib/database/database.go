package database

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"testing"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sqlx.DB
	tx *sqlx.Tx
)

/*
StartSession must be called at the beginning of every request, and appropriately
sets the database connection and transaction object
*/
func StartSession() (err error) {
	db, err = sqlx.Connect("sqlite3", "/Users/chasew/go/src/family_map_server/database.sqlite")
	if err != nil {
		return
	}
	tx, err = db.Beginx()
	return
}

/*
StartTestingSession must be called at the beginning of every test, and appropriately
sets the database connection and transaction object
*/
func StartTestingSession(t *testing.T) {
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		t.Errorf("could not open test database: %v", err)
		return
	}
	fileBytes, err := ioutil.ReadFile("../../database_ddl.sql")
	if err != nil {
		t.Errorf("could not open database DDL script: %v", err)
		return
	}
	_, err = db.Exec(string(fileBytes))
	if err != nil {
		t.Errorf("DDL script failed: %v", err)
		return
	}
	tx, err = db.Beginx()
	return
}

/*
GetTransaction will return the current transaction object
*/
func GetTransaction() *sqlx.Tx {
	return tx
}

/*
ClearDatabase will drop all tables
*/
func ClearDatabase(db *sql.DB) {
	panic("not implemented")
}

/*
CreateTables will create all necessary tables
*/
func CreateTables(db *sql.DB) {
	panic("not implemented")
}
