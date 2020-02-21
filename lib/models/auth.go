package models

import (
	"github.com/chaseWilliams/family-map/lib/database"
)

/*
Auth represents the Auth table, correlating Users to auth tokens
*/
type Auth struct {
	Username string `db:"username"`
	AuthToken string `db:"auth_token"`
}

/*
GetAuthToken will get the auth token of the provided User
*/
func GetAuthToken(user *User) (token string, err error) {
	tx := database.GetTransaction()
	auth := new(Auth)

	err = tx.QueryRowx("select * from Auth where username = ?", user.Username).StructScan(auth)
	token = auth.AuthToken
	return
}