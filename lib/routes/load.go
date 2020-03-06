package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chaseWilliams/family-map/lib/database"
	"github.com/chaseWilliams/family-map/lib/models"
	"github.com/chaseWilliams/family-map/lib/util"
)

type loadData struct {
	Users   []models.User       `json:"users"`
	Persons []models.PersonJSON `json:"persons"`
	Events  []models.Event      `json:"events"`
}

/*
Load will clear the database, and then load all data
*/
func Load(w http.ResponseWriter, r *http.Request, user models.User) (err error) {
	defer r.Body.Close()
	data := new(loadData)
	err = json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		util.WriteBadResponse(
			w,
			fmt.Sprintf("could not parse JSON (%v)", err.Error()),
		)
		return
	}

	err = database.ClearDatabase()
	if err != nil {
		util.WriteBadResponse(
			w,
			fmt.Sprintf("could not clear database: %v", err.Error()),
		)
		return
	}

	// save Users
	for _, user := range data.Users {
		err = user.Save()
		if err != nil {
			util.WriteBadResponse(
				w,
				fmt.Sprintf("could not save User (%v)", err.Error()),
			)
			return
		}
	}

	// save Persons
	for _, personJSON := range data.Persons {
		person := personJSON.ToPerson()
		err = person.Save()
		if err != nil {
			util.WriteBadResponse(
				w,
				fmt.Sprintf("could not save Person (%v)", err.Error()),
			)
			return
		}
	}

	// save Events
	for _, event := range data.Events {
		err = event.Save()
		if err != nil {
			util.WriteBadResponse(
				w,
				fmt.Sprintf("could not save Event (%v)", err.Error()),
			)
			return
		}
	}

	// good response
	util.WriteOKResponse(
		w,
		util.Message{
			Message: fmt.Sprintf(
				"Successfully added %d users, %d persons, and %d events to the database.",
				len(data.Users),
				len(data.Persons),
				len(data.Events),
			),
			Success: "true",
		},
	)
	return
}
