package models

import (
	"github.com/chaseWilliams/family-map/lib/database"
	"github.com/chaseWilliams/family-map/lib/util"
)

/*
Auth represents the Auth table, correlating Users to auth tokens
*/
type Auth struct {
	Username  string `db:"username"`
	AuthToken string `db:"auth_token"`
}

/*
Save will save the auth object to the database
*/
func (a Auth) Save() (err error) {
	tx, err := database.GetTransaction()
	if err != nil {
		return
	}
	_, err = tx.Exec(
		`INSERT INTO Auth VALUES (?, ?);`,
		a.Username,
		a.AuthToken,
	)
	return
}

/*
NewAuthToken generates a new auth token for the user and persists
that token
*/
func NewAuthToken(user User) (token string, err error) {
	token = util.RandomID()
	auth := Auth{user.Username, token}
	err = auth.Save()
	return
}

func AssertAuth(username, auth string) bool {
	return false
}
