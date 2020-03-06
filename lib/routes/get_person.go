package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/chaseWilliams/family-map/lib/models"
	"github.com/chaseWilliams/family-map/lib/util"
)

/*
GetPerson will get the specified person object
*/
func GetPerson(w http.ResponseWriter, r *http.Request, user models.User) (err error) {
	personID := strings.TrimPrefix(r.URL.Path, "/person/")
	person, err := models.GetPerson(personID)
	if err != nil {
		util.WriteBadResponse(
			w,
			fmt.Sprintf("could not get person: %v", err.Error()),
		)
		return
	}
	// check auth
	if person.Username != user.Username {
		util.WriteBadResponse(
			w,
			"unauthorized resource access",
		)
		return
	}
	util.WriteResponse(w, person.ToJSON(), http.StatusOK)
	return
}
