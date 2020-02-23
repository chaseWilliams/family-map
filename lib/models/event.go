package models

/*
Event is a significant event concerning a person
*/
type Event struct {
	eventID   string
	username  string
	personID  string
	latitude  string
	longitude string
	country   string
	city      string
	eventType string
	year      string
}

/*
create will take this Event model and create it in the database
*/
func (u *Event) create() {
	panic("not implemented")
}

/*
GetEvent returns a Event given a eventID
*/
func GetEvent(eventID string) Event {
	panic("not implemented")
}
