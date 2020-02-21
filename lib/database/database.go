package database

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
	tx *sqlx.Tx
)

/*
StartSession is a factory-style function that returns a new connection
to the database.
*/
func StartSession() (err error) {
	db, err = sqlx.Connect("sqlite3", "/Users/chasew/go/src/family_map_server/database.sqlite")
	if err != nil {
		return
	}
	tx, err = db.Beginx()
	return
}

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
