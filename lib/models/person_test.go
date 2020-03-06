package models

import (
	"database/sql"
	"github.com/chaseWilliams/family-map/lib/database"
	"testing"
)

func TestGetPerson(t *testing.T) {
	database.StartTestingSession(t)
	person := Person{
		PersonID:  "XVLBZGBAICMRAJWW",
		Username:  "chasew",
		FirstName: "Chase",
		LastName:  "Williams",
		Gender:    "m",
		FatherID:  sql.NullString{"BQAMCZRGGUMOIAOY", true},
		MotherID:  sql.NullString{"TCBOTOQSPYHBOZUO", true},
		SpouseID:  sql.NullString{"", false},
	}

	p, err := GetPerson("XVLBZGBAICMRAJWW")
	if err != nil {
		t.Errorf("could not get person: %v", err)
	}
	if *p != person {
		t.Error("returned Person object did not match what was expected")
	}
}

func TestGetPersonFailure(t *testing.T) {
	database.StartTestingSession(t)
	_, err := GetPerson("lalala")
	if err == nil {
		t.Error("no error thrown for a nonexistent person")
	}
}

func TestSavePerson(t *testing.T) {
	database.StartTestingSession(t)
	minimalPerson := Person{
		PersonID:  "123",
		Username:  "chasew",
		FirstName: "Chase",
		LastName:  "Williams",
		Gender:    "m",
	}
	fullPerson := Person{
		PersonID:  "12345",
		Username:  "chasew",
		FirstName: "Chase",
		LastName:  "Williams",
		Gender:    "m",
		FatherID:  sql.NullString{"TCBOTOQSPYHBOZUO", true},
		MotherID:  sql.NullString{"TCBOTOQSPYHBOZUO", true},
		SpouseID:  sql.NullString{"TCBOTOQSPYHBOZUO", true},
	}
	err := minimalPerson.Save()
	if err != nil {
		t.Errorf("could not save minimal person: %v", err)
	}
	err = fullPerson.Save()
	if err != nil {
		t.Errorf("could not save full person: %v", err)
	}
}

func TestSavePersonFailure(t *testing.T) {
	database.StartTestingSession(t)
	badPersonData := Person{
		Username:  "chasew",
		FirstName: "Chase",
		LastName:  "Williams",
		Gender:    "m",
	}
	err := badPersonData.Save()
	if err == nil {
		t.Error("person.Save() didn't return an error for missing data")
	}
}

func TestDeletePerson(t *testing.T) {
	database.StartTestingSession(t)
	err := DeletePerson("abc")
	if err != nil {
		t.Errorf("could not delete person: %v", err)
	}
}
