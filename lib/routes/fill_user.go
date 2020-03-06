package routes

import (
	"fmt"
	"github.com/chaseWilliams/family-map/lib/database"
	"github.com/chaseWilliams/family-map/lib/datagen/simulation"
	"github.com/chaseWilliams/family-map/lib/models"
	"github.com/chaseWilliams/family-map/lib/util"
	"net/http"
	"strconv"
	"strings"
)

/*
FillUser will fill out a user's ancestry tree, and remove previous ancestry data
if it exists.
*/
func FillUser(w http.ResponseWriter, r *http.Request, user models.User) (err error) {
	paramString := strings.Replace(r.URL.Path, "/fill/", "", 1)
	params := strings.Split(paramString, "/")
	username := params[0]
	count := 4
	if len(params) > 1 && len(params[1]) > 0 {
		count, err = strconv.Atoi(params[1])
		if err != nil {
			panic(err)
		}
	}

	if count < 1 {
		util.WriteBadResponse(
			w,
			"generation must be a positive number",
		)
		return
	}

	user, err = models.GetUser(username)
	if err != nil {
		util.WriteBadResponse(
			w,
			"could not find user: " + err.Error(),
		)
		return
	}

	// clear database of all persons and events attached to user
	// create user person
	err = database.ClearFamily(username)
	if err != nil {
		util.WriteBadResponse(
			w,
			fmt.Sprintf("could not clean database (%v)", err.Error()),
		)
		return
	}
	person := models.Person{
		PersonID:  util.RandomID(),
		Username:  username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Gender:    user.Gender,
	}
	numPeople, numEvents, err := simulation.CreateFamily(person, count)
	if err != nil {
		util.WriteBadResponse(
			w,
			fmt.Sprintf("could not create family: %v", err.Error()),
		)
	}

	// good response
	util.WriteOKResponse(
		w,
		util.Message{
			Message: fmt.Sprintf(
				"Successfully added %d persons and %d events to the database",
				numPeople,
				numEvents,
			),
			Success: "true",
		},
	)

	return
}
