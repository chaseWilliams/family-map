package models

import (
	"fmt"
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
	tx, err := database.GetTransaction()
	if err != nil {
		return
	}
	user = new(User)
	err = tx.QueryRowx(
		"select * from Users where username = ? AND password = ?",
		d.Username,
		d.Password,
	).StructScan(user)
	return
}

/*
GetUser will return a User struct for given username
*/
func GetUser(username string) (user User, err error) {
	tx, err := database.GetTransaction()
	if err != nil {
		return
	}
	user = User{}
	err = tx.QueryRowx(
		"SELECT * FROM Users WHERE username = ?",
		username,
	).StructScan(&user)
	return
}

/*
Save will take this User model and create it in the database
*/
func (u *User) Save() (err error) {
	tx, err := database.GetTransaction()
	if !u.validate() {
		return fmt.Errorf("request property missing or malformed")
	}
	if err != nil {
		return
	}
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

func (u User) validate() bool {
	return !(len(u.Username) == 0 ||
		len(u.Password) == 0 ||
		len(u.Email) == 0 ||
		len(u.FirstName) == 0 ||
		len(u.LastName) == 0 ||
		len(u.Gender) != 1)

}

/*
DeleteUser will delete the row in Users with the username
*/
func DeleteUser(username string) (err error) {
	tx, err := database.GetTransaction()
	if err != nil {
		return
	}
	_, err = tx.Exec(
		"DELETE FROM Users WHERE username = ?",
		username,
	)
	return
}
