package routes

import (
	"net/http"

	"github.com/chaseWilliams/family-map/lib/models"
	"github.com/chaseWilliams/family-map/lib/util"
)

type dataResponse struct {
	Data    interface{} `json:"data"`
	Success string      `json:"success"`
}

/*
AllEvents Returns all events relevant to current user
*/
func AllEvents(w http.ResponseWriter, r *http.Request, user models.User) (err error) {
	events, err := models.GetAllEvents(user.Username)
	if err != nil {
		util.WriteBadResponse(
			w,
			err.Error(),
		)
		return
	}

	util.WriteOKResponse(w, dataResponse{
		Data: events,
		Success: "true",
	})
	return
}
