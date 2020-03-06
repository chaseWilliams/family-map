package routes

import (
	"strings"
	"net/http"

	"github.com/chaseWilliams/family-map/lib/models"
	"github.com/chaseWilliams/family-map/lib/util"
)

/*
GetEvent returns specified event object
*/
func GetEvent(w http.ResponseWriter, r *http.Request, user models.User) (err error) {
	eventID := strings.TrimPrefix(r.URL.Path, "/event/")
	event, err := models.GetEvent(eventID)
	if err != nil {
		util.WriteBadResponse(
			w,
			"could not get event: " + err.Error(),
		)
		return
	}

	// check auth
	if event.Username != user.Username {
		util.WriteBadResponse(
			w,
			"unauthorized resource access",
		)
		return
	}

	util.WriteOKResponse(w, event)
	return
}
