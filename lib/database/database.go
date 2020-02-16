package database

import (
	"database/sql"
	"family-map-server/lib/models"
)

/*
CreateConnection is a factory-style function that returns a new connection
to the database.
*/
func CreateConnection() *sql.DB {
	panic("not implemented")
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

/*
LoginUser will return User associated with auth token
*/
func LoginUser(auth string) models.User {
	panic("not implemented")
}