package models

import (
	"github.com/chaseWilliams/family-map/lib/database"
)

type Auth struct {
	Username string `db:"username"`
	AuthToken string `db:"auth_token"`
}

func GetAuthToken(user *User) (token string, err error) {
	tx := database.GetTransaction()
	auth := new(Auth)

	err = tx.QueryRowx("select * from Auth where username = ?", user.Username).StructScan(auth)
	token = auth.AuthToken
	return
}