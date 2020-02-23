package routes

import (
	"github.com/chaseWilliams/family-map/lib/models"
	"github.com/chaseWilliams/family-map/lib/util"
	"net/http"
)

/*
GetPerson will get the specified person object
*/
func GetPerson(w http.ResponseWriter, r *http.Request) (err error) {
	person, _ := models.GetPerson("abc")
	util.WriteResponse(w, person, http.StatusOK)
	return
}
