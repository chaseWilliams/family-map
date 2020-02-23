package routes

import (
	"encoding/json"
	"fmt"
	"github.com/chaseWilliams/family-map/lib/models"
	"github.com/chaseWilliams/family-map/lib/util"
	"net/http"
)

type loginSuccess struct {
	AuthToken string `json:"authToken"`
	Username  string `json:"userName"`
	PersonID  string `json:"personID"`
	Success   string `json:"success"`
}

type loginFailure struct {
	Message string `json:"message"`
	Success string `json:"success"`
}

/*
Login checks for a User that matches the provided credentials and
writes the appropriate response to the Response object
*/
func Login(w http.ResponseWriter, r *http.Request) (err error) {
	loginData := new(models.LoginData)

	// check that the method is POST
	if r.Method != "POST" {
		util.WriteNotFound(w)
		err = fmt.Errorf("/user/login must be a POST request; request was %s", r.Method)
		return
	}

	// convert the request payload from JSON to a loginData struct
	err = json.NewDecoder(r.Body).Decode(loginData)
	if err != nil {
		util.WriteResponse(
			w,
			loginFailure{
				fmt.Sprintf("error: could not decode JSON (%v)", err.Error()),
				"false",
			},
			http.StatusBadRequest,
		)
		return
	}

	// get the User matching the credentials
	user, err := loginData.GetUser()
	if err != nil {
		util.WriteResponse(
			w,
			loginFailure{
				"error: no User found for those credentials",
				"false",
			},
			http.StatusBadRequest,
		)
		return
	}

	// get the User's auth token
	token, err := models.GetAuthToken(user)
	if err != nil {
		util.WriteResponse(
			w,
			loginFailure{
				"error: could not find user's auth tokens.",
				"false",
			},
			http.StatusInternalServerError,
		)
		return
	}

	// return a successful response
	util.WriteResponse(
		w,
		loginSuccess{
			AuthToken: token,
			Username:  user.Username,
			PersonID:  user.PersonID,
			Success:   "true",
		},
		http.StatusOK,
	)
	return
}
