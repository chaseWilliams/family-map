package routes

import (
	"encoding/json"
	"fmt"
	"github.com/chaseWilliams/family-map/lib/datagen/simulation"
	"github.com/chaseWilliams/family-map/lib/models"
	"github.com/chaseWilliams/family-map/lib/util"
	"net/http"
)

/*
Register creates a new user account, generates 4 generations of ancestor data for the new
user, logs the user in, and returns an auth token.
*/
func Register(w http.ResponseWriter, r *http.Request) (err error) {
	user := new(models.User)
	json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		util.WriteBadResponse(
			w,
			fmt.Sprintf("error: could not decode JSON (%v)", err.Error()),
		)
		return
	}
	person := models.Person{
		PersonID:  util.RandomID(),
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Gender:    user.Gender,
	}
	user.PersonID = person.PersonID
	err = person.Save()
	if err != nil {
		util.WriteBadResponse(
			w,
			err.Error(),
		)
		return
	}
	err = user.Save()
	if err != nil {
		util.WriteBadResponse(
			w,
			err.Error(),
		)
		return
	}

	simulation.CreateFamily(person, 4)

	token, err := models.NewAuthToken(*user)
	if err != nil {
		util.WriteBadResponse(
			w,
			err.Error(),
		)
		return
	}

	resp := loginSuccess{
		AuthToken: token,
		Username:  user.Username,
		PersonID:  user.PersonID,
		Success:   "true",
	}
	util.WriteOKResponse(w, resp)

	return
}
