package routes

import (
	"encoding/json"
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
Login will login the user and return an auth token
*/
func Login(w http.ResponseWriter, r *http.Request) (err error) {
	loginData := new(models.LoginData)

	err = json.NewDecoder(r.Body).Decode(loginData)
	if err != nil {
		util.WriteResponse(w, loginFailure{err.Error(), "false"}, http.StatusBadRequest)
		return
	}

	user, err := loginData.GetUser()
	if err != nil {
		response := loginFailure{
			Message: err.Error(),
			Success: "false",
		}
		json.NewEncoder(w).Encode(response)
	}

	token, err := models.GetAuthToken(user)
	if err != nil {
		response := loginFailure{
			Message: err.Error(),
			Success: "false",
		}
		json.NewEncoder(w).Encode(response)
	}

	response := loginSuccess{
		AuthToken: token,
		Username:  user.Username,
		PersonID:  user.PersonID,
		Success:   "true",
	}

	json.NewEncoder(w).Encode(response)
	return
}
