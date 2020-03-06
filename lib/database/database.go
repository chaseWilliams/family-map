package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"path"
	"runtime"
	"testing"
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
func StartTestingSession(t testing.TB) {
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		t.Errorf("could not open test database: %v", err)
		return
	}
	_, filename, _, _ := runtime.Caller(0)
	sqlFiles := []string{
		"database_ddl.sql",
		"cities.sql",
		"users.sql",
		"persons.sql",
		"events.sql",
		"auth.sql",
	}
	for _, name := range sqlFiles {
		filepath := path.Join(path.Dir(filename), "../../test/data/"+name)
		fileBytes, err := ioutil.ReadFile(filepath)
		if err != nil {
			t.Errorf("could not open sql script %v: %v", name, err)
			return
		}
		_, err = db.Exec(string(fileBytes))
		if err != nil {
			t.Errorf("SQL script failed %v: %v", name, err)
			return
		}
	}
	tx, _ = db.Beginx()
	return
}

/*
GetTransaction will return the current transaction object
*/
func GetTransaction() (*sqlx.Tx, error) {
	if tx == nil {
		return nil, fmt.Errorf("database session hasn't been properly started")
	}
	return tx, nil
}

/*
ClearDatabase will drop all rows from tables
Persons, Events, Users, and Auth
*/
func ClearDatabase() (err error) {
	_, err = tx.Exec(
		`DELETE FROM Persons;
		DELETE FROM Events;
		DELETE FROM Users;
		DELETE FROM Auth;`,
	)
	return
}

/*
ClearFamily will clear all family data for the user
(this also delete the user's corresponding person)
*/
func ClearFamily(username string) (err error) {
	_, err = tx.Exec(
		`DELETE FROM Persons WHERE username = ?;
		DELETE FROM Events WHERE username = ?;`,
		username,
		username,
	)
	return
}
