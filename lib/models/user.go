package models


/*
User is the model for all accounts
*/
type User struct {
	userName string
	password string
	email string
	firstName string
	lastName string
	gender string
	personID string
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

