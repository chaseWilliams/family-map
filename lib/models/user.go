package models

import (
	"github.com/chaseWilliams/family-map/lib/database"
)

/*
User is the model for all accounts
*/
type User struct {
	Username  string `db:"username"`
	Password  string `db:"password"`
	Email     string `db:"email"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Gender    string `db:"gender"`
	PersonID  string `db:"person_id"`
}

type LoginData struct {
	Username string `json:userName`
	Password string `json:password`
}

func (d *LoginData) GetUser() (user *User, err error) {
	tx := database.GetTransaction()
	user = new(User)
	err = tx.QueryRowx(
		"select * from Users where username = ? AND password = ?",
		d.Username,
		d.Password,
	).StructScan(user)
	return
}

/*
create will take this User model and create it in the database
*/
func (u *User) create() {
	panic("not implemented")
}

/*
GetUser returns a User given a username
*/
func GetUser(userName string) User {
	panic("not implemented")
}
