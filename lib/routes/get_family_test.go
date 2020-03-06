package routes

import (
	"net/http"
	"testing"

	"github.com/chaseWilliams/family-map/lib/database"
	"github.com/chaseWilliams/family-map/lib/models"
)

func TestGetFamily(t *testing.T) {
	database.StartTestingSession(t)
	req, err := http.NewRequest("GET", "/person", nil)
	fatal(err, t)
	req.Header.Set("Authorization", authToken)

	assertExpectedResponse(
		routeTest{
			req:     req,
			service: GetFamily,
			user: models.User{
				Username: "chasew",
			},
			code:           http.StatusOK,
			responseStruct: dataResponse{},
			expectError:    false,
			name:           "get all persons",
		},
		t,
	)
}

func TestGetFamilyFailure(t *testing.T) {
	database.StartTestingSession(t)
	invalidAuthReq, err := http.NewRequest("GET", "/person", nil)
	fatal(err, t)
	assertExpectedResponse(
		routeTest{
			req:     invalidAuthReq,
			service: GetFamily,
			user: models.User{
				Username: "chasew",
			},
			code:           http.StatusOK,
			responseStruct: dataResponse{},
			expectError:    false,
			name:           "invalid auth get all persons",
		},
		t,
	)
}