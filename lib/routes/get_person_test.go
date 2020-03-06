package routes

import (
	"net/http"
	"testing"

	"github.com/chaseWilliams/family-map/lib/database"
	"github.com/chaseWilliams/family-map/lib/models"
)

func TestGetPerson(t *testing.T) {
	database.StartTestingSession(t)
	req, err := http.NewRequest("GET", "/person/XVLBZGBAICMRAJWW", nil)
	fatal(err, t)
	req.Header.Set("Authorization", authToken)

	assertExpectedResponse(
		routeTest{
			req:     req,
			service: GetPerson,
			user: models.User{
				Username: "chasew",
			},
			code:           http.StatusOK,
			responseStruct: models.PersonJSON{},
			expectError:    false,
			name:           "get person",
		},
		t,
	)
}

func TestGetPersonFailure(t *testing.T) {
	database.StartTestingSession(t)
	req, err := http.NewRequest("GET", "/person/LRPJPSGVMPYMXBUE", nil)
	fatal(err, t)
	req.Header.Set("Authorization", "litty_auth")

	assertExpectedResponse(
		routeTest{
			req:     req,
			service: GetPerson,
			user: models.User{
				Username: "test_user",
			},
			code:           http.StatusBadRequest,
			responseStruct: models.PersonJSON{},
			expectError:    true,
			name:           "invalid get person",
		},
		t,
	)
}
