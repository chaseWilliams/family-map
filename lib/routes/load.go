package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.WriteBadResponse(
			w,
			err.Error(),
		)
		return
	}
	err = json.Unmarshal(reqBytes, data)
	if err != nil {
		util.WriteBadResponse(
			w,
			fmt.Sprintf("could not parse JSON (%v)", err.Error()),
		)
		return
	}

	if data.Users == nil || data.Events == nil || data.Persons == nil {
		util.WriteBadResponse(
			w,
			"missing or malformed fields",
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
