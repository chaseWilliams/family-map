package routes

import (
	"fmt"
	"github.com/chaseWilliams/family-map/lib/database"
	"github.com/chaseWilliams/family-map/lib/util"
	"net/http"
)

/*
Clear will wipe all data from the database
*/
func Clear(w http.ResponseWriter, r *http.Request) (err error) {
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
