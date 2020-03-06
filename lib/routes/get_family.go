package routes

import (
	"net/http"

	"github.com/chaseWilliams/family-map/lib/models"
	"github.com/chaseWilliams/family-map/lib/util"
)

/*
GetFamily returns all related Persons to specified Person
*/
func GetFamily(w http.ResponseWriter, r *http.Request, user models.User) (err error) {
	family, err := models.GetFamily(user.Username)
	if err != nil {
		util.WriteBadResponse(
			w,
			err.Error(),
		)
		return
	}

	familyJSON := make([]models.PersonJSON, len(family))
	for i, person := range family {
		familyJSON[i] = person.ToJSON()
	}

	util.WriteOKResponse(w, dataResponse{
		Data:    familyJSON,
		Success: "true",
	})
	return
}
