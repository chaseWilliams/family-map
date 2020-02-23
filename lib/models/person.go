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
Save will take this Person model and create it in the database
*/
func (p *Person) Save() (err error) {
	tx := database.GetTransaction()
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

/*
GetPerson returns a User given a username
*/
func GetPerson(personID string) (person *Person, err error) {
	tx := database.GetTransaction()
	person = new(Person)
	err = tx.QueryRowx(
		"SELECT * FROM Persons WHERE person_id = ?",
		personID,
	).StructScan(person)
	return
}

/*
DeletePerson deletes the person from the table
*/
func DeletePerson(personID string) (err error) {
	tx := database.GetTransaction()
	_, err = tx.Exec(
		"DELETE FROM Persons WHERE person_id = ?",
		personID,
	)
	return
}
