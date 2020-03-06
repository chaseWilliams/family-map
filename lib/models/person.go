package models

import (
	"database/sql"
	"fmt"
	"github.com/chaseWilliams/family-map/lib/database"
)

/*
Person is the model for all people and their relationships
and other data
*/
type Person struct {
	PersonID  string         `db:"person_id"`
	Username  string         `db:"username"`
	FirstName string         `db:"first_name"`
	LastName  string         `db:"last_name"`
	Gender    string         `db:"gender"`
	FatherID  sql.NullString `db:"father_id"`
	MotherID  sql.NullString `db:"mother_id"`
	SpouseID  sql.NullString `db:"spouse_id"`
}

/*
PersonJSON is a stand-in struct that represents the exact
JSON reponse needed for the API
*/
type PersonJSON struct {
	PersonID  string `json:"personID"`
	Username  string `json:"associatedUsername"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Gender    string `json:"gender"`
	FatherID  string `json:"fatherID,omitempty"`
	MotherID  string `json:"motherID,omitempty"`
	SpouseID  string `json:"spouseID,omitempty"`
}

/*
ToPerson returns a Person struct based on 
this struct's data
*/
func (p PersonJSON) ToPerson() Person {
	person := Person{
		PersonID: p.PersonID,
		Username: p.Username,
		FirstName: p.FirstName,
		LastName: p.LastName,
		Gender: p.Gender,
		FatherID: sql.NullString{p.FatherID, true},
		MotherID: sql.NullString{p.MotherID, true},
		SpouseID: sql.NullString{p.SpouseID, true},
	}
	if len(p.FatherID) == 0 {
		person.FatherID.Valid = false
	}
	if len(p.MotherID) == 0 {
		person.MotherID.Valid = false
	}
	if len(p.SpouseID) == 0 {
		person.SpouseID.Valid = false
	}
	return person
}

/*
ToJSON returns a compliant PersonJSON
based on this person's data
*/
func (p Person) ToJSON() PersonJSON {
	return PersonJSON{
		PersonID: p.PersonID,
		Username: p.Username,
		FirstName: p.FirstName,
		LastName: p.LastName,
		Gender: p.Gender,
		FatherID: p.FatherID.String,
		MotherID: p.MotherID.String,
		SpouseID: p.SpouseID.String,
	}
}

/*
Save will take this Person model and create it in the database
*/
func (p *Person) Save() (err error) {
	tx, err := database.GetTransaction()
	if err != nil {
		return
	}
	if !p.validate() {
		return fmt.Errorf("person is malformed or missing properties")
	}
	if p.PersonID == "" {
		return fmt.Errorf("p.PersonID must not be an empty string")
	}
	_, err = tx.Exec(
		`INSERT INTO Persons
		VALUES (?, ?, ?, ?, ?, ?, ?, ?);`,
		p.PersonID,
		p.Username,
		p.FirstName,
		p.LastName,
		p.Gender,
		p.FatherID,
		p.MotherID,
		p.SpouseID,
	)
	return
}

func (p Person) validate() bool {
	return !(len(p.Gender) != 1 ||
		len(p.Username) == 0 ||
		len(p.FirstName) == 0 ||
		len(p.LastName) == 0)
}

/*
GetPerson returns a Person given a personID
*/
func GetPerson(personID string) (person *Person, err error) {
	tx, err := database.GetTransaction()
	if err != nil {
		return
	}
	person = new(Person)
	err = tx.QueryRowx(
		"SELECT * FROM Persons WHERE person_id = ?",
		personID,
	).StructScan(person)
	return
}

/*
GetFamily returns all people related to the username
*/
func GetFamily(username string) (family []Person, err error) {
	tx, err := database.GetTransaction()
	family = []Person{}
	if err != nil {
		return
	}
	
	err = tx.Select(&family, "SELECT * FROM Persons WHERE username = ?", username)
	return
}

/*
DeletePerson deletes the person from the table
*/
func DeletePerson(personID string) (err error) {
	tx, err := database.GetTransaction()
	if err != nil {
		return
	}
	_, err = tx.Exec(
		"DELETE FROM Persons WHERE person_id = ?",
		personID,
	)
	return
}
