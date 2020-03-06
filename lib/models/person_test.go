package models

import (
	"database/sql"
	"github.com/chaseWilliams/family-map/lib/database"
	"testing"
)

func TestGetPerson(t *testing.T) {
	database.StartTestingSession(t)
	person := Person{
		PersonID:  "abc",
		Username:  sql.NullString{"chasew", true},
		FirstName: "chase",
		LastName:  "williams",
		Gender:    "m",
		FatherID:  sql.NullString{"", false},
		MotherID:  sql.NullString{"", false},
		SpouseID:  sql.NullString{"", false},
	}

	p, err := GetPerson("abc")
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
		FirstName: "Chase",
		LastName:  "Williams",
		Gender:    "m",
	}
	fullPerson := Person{
		PersonID:  "12345",
		Username:  sql.NullString{"chasew", true},
		FirstName: "Chase",
		LastName:  "Williams",
		Gender:    "m",
		FatherID:  sql.NullString{"abc", true},
		MotherID:  sql.NullString{"abc", true},
		SpouseID:  sql.NullString{"abc", true},
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
	duplicatePerson := Person{
		PersonID:  "abc",
		Username:  sql.NullString{"chasew", true},
		FirstName: "Chase",
		LastName:  "Williams",
		Gender:    "m",
	}
	badPersonData := Person{
		Username:  sql.NullString{"chasew", true},
		FirstName: "Chase",
		LastName:  "Williams",
		Gender:    "m",
	}
	err := duplicatePerson.Save()
	if err == nil {
		t.Error("person.Save() didn't return an error for duplicate personID")
	}
	err = badPersonData.Save()
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
