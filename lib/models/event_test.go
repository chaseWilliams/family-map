package models

import (
	"testing"
	"github.com/chaseWilliams/family-map/lib/database"
)

func TestGetEvent(t *testing.T) {
	database.StartTestingSession(t)
	event := Event {
		EventID: "chase_dies",
		Username: "chasew",
		PersonID: "abc",
		Latitude: 37,
		Longitude: 38,
		Country: "USA",
		City: "Roswell",
		EventType: "Death",
		Year: 2078,
	}
	eventCopy, err := GetEvent("chase_dies")
	if err != nil {
		t.Errorf("Couldn't get Event object: %v", err)
	}
	if *eventCopy != event {
		t.Error("Event doesn't match expected values")
	}
}

func TestGetEventFailure(t *testing.T) {
	database.StartTestingSession(t)
	_, err := GetEvent("blahblahblah")
	if err == nil {
		t.Error("nonexistent Event should have thrown an error")
	}
}

func TestSaveEvent(t *testing.T) {
	database.StartTestingSession(t)
	event := Event {
		EventID: "chase_born",
		Username: "chasew",
		PersonID: "abc",
		Latitude: 37,
		Longitude: 38,
		Country: "USA",
		City: "Roswell",
		EventType: "Death",
		Year: 1999,
	}
	err := event.Save()
	if err != nil {
		t.Errorf("could not save Event: %v", err)
	}
}

func TestSaveEventFailure(t *testing.T) {
	database.StartTestingSession(t)
	malformedEvent := Event {
		Username: "chasew",
		PersonID: "abc",
		Latitude: 37,
		Longitude: 38,
		Country: "USA",
		City: "Roswell",
		EventType: "Death",
		Year: 2078,
	}

	nonexistentPersonEvent := Event {
		EventID: "chase_dies",
		Username: "chasewww",
		PersonID: "abc",
		Latitude: 37,
		Longitude: 38,
		Country: "USA",
		City: "Roswell",
		EventType: "Death",
		Year: 2078,
	}

	err := malformedEvent.Save()
	if err == nil {
		t.Error("malformed data did not throw an error when saving")
	}
	err = nonexistentPersonEvent.Save()
	if err == nil {
		t.Error("event data had nonexistent person, did not throw error")
	}
}