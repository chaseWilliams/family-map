package models


/*
Person is the model for all people and their relationships
and other data
*/
type Person struct {
	personID string
	username string
	firstName string
	lastName string
	gender string
	fatherID string
	motherID string
	spouseID string
}

/*
create will take this Person model and create it in the database
*/
func (u *Person) create() {
	panic("not implemented")
}

/*
GetPerson returns a User given a username
*/
func GetPerson(personID string) Person {
	panic("not implemented")
}