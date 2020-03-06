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
