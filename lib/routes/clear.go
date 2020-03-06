package routes

import (
	"fmt"
	"net/http"

	"github.com/chaseWilliams/family-map/lib/database"
	"github.com/chaseWilliams/family-map/lib/models"
	"github.com/chaseWilliams/family-map/lib/util"
)

/*
Clear will wipe all data from the database
*/
func Clear(w http.ResponseWriter, r *http.Request, user models.User) (err error) {
	err = database.ClearDatabase()
	if err != nil {
		util.WriteBadResponse(
			w,
			fmt.Sprintf("could not clear the database (%v)", err.Error()),
		)
		return
	}
	util.WriteOKResponse(
		w,
		map[string]string{
			"message": "Clear succeeded.",
			"success": "true",
		},
	)
	return
}
