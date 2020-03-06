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
func Register(w http.ResponseWriter, r *http.Request, user models.User) (err error) {
	json.NewDecoder(r.Body).Decode(&user)
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
	err = user.Save()
	if err != nil {
		util.WriteBadResponse(
			w,
			err.Error(),
		)
		return
	}

	_, _, err = simulation.CreateFamily(person, 5)
	if err != nil {
		util.WriteBadResponse(
			w,
			"could not simulate family: "+err.Error(),
		)
		return
	}

	token, err := models.NewAuthToken(user)
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
