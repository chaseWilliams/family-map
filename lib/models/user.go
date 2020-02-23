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

/*
LoginData is what data needs to be passed to login a User
*/
type LoginData struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

/*
GetUser will perform the SQL query to get the appropriate User
object for the login data
*/
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
Save will take this User model and create it in the database
*/
func (u *User) Save() (err error) {
	tx := database.GetTransaction()
	_, err = tx.Exec(
		`INSERT INTO USERS
		VALUES (?, ?, ?, ?, ?, ?, ?);`,
		u.Username,
		u.Password,
		u.Email,
		u.FirstName,
		u.LastName,
		u.Gender,
		u.PersonID,
	)
	return
}

/*
DeleteUser will delete the row in Users with the username
*/
func DeleteUser(username string) (err error) {
	tx := database.GetTransaction()
	_, err = tx.Exec(
		"DELETE FROM Users WHERE username = ?",
		username,
	)
	return
}
