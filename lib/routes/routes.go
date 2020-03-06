package routes

import (
	"net/http"

	"github.com/chaseWilliams/family-map/lib/models"
)

/*
Route is a function header that acts as a common way to deal with all requests
*/
type Route func(http.ResponseWriter, *http.Request, models.User) error
