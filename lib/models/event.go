package models

import (
	"fmt"
	"github.com/chaseWilliams/family-map/lib/database"
)

/*
Event is a significant event concerning a person
*/
type Event struct {
	EventID   string  `json:"eventID" db:"event_id"`
	Username  string  `json:"associatedUsername" db:"username"`
	PersonID  string  `json:"personID" db:"person_id"`
	Latitude  float64 `json:"latitude" db:"latitude"`
	Longitude float64 `json:"longitude" db:"longitude"`
	Country   string  `json:"country" db:"country"`
	City      string  `json:"city" db:"city"`
	EventType string  `json:"eventType" db:"event_type"`
	Year      int     `json:"year" db:"year"`
}

/*
Save will take this Event model and create it in the database
*/
func (e *Event) Save() (err error) {
	tx, err := database.GetTransaction()
	if err != nil {
		return err
	}
	if e.EventID == "" {
		return fmt.Errorf("e.EventID must not be an empty string")
	}
	_, err = tx.Exec(
		`INSERT INTO Events
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);`,
		e.EventID,
		e.Username,
		e.PersonID,
		e.Latitude,
		e.Longitude,
		e.Country,
		e.City,
		e.EventType,
		e.Year,
	)
	return
}

/*
GetEvent returns a Event given a eventID
*/
func GetEvent(eventID string) (event *Event, err error) {
	tx, err := database.GetTransaction()
	if err != nil {
		return
	}
	event = new(Event)
	err = tx.QueryRowx(
		"SELECT * FROM Events WHERE event_id = ?",
		eventID,
	).StructScan(event)
	return
}
