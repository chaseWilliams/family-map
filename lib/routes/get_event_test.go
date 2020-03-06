package routes

import (
	"net/http"
	"testing"

	"github.com/chaseWilliams/family-map/lib/database"
	"github.com/chaseWilliams/family-map/lib/models"
)

func TestGetEvent(t *testing.T) {
	database.StartTestingSession(t)
	req, err := http.NewRequest("GET", "/event/GLAPAAEFUQSEHFCV", nil)
	if err != nil {
		t.Fatal(err)
	}

	assertExpectedResponse(
		routeTest{
			req:     req,
			service: GetEvent,
			user: models.User{
				Username: "chasew",
			},
			code:           http.StatusOK,
			responseStruct: models.Event{},
			expectError:    false,
			name:           "get event",
		},
		t,
	)
}

func TestGetEventFailure(t *testing.T) {
	database.StartTestingSession(t)
	missingEventReq, err := http.NewRequest("GET", "/event/asdf", nil)
	fatal(err, t)
	missingEventReq.Header.Set("Authorization", authToken)
	unauthorizedEventReq, err := http.NewRequest("GET", "/event/GLAPAAEFUQSEHFCV", nil)
	fatal(err, t)
	unauthorizedEventReq.Header.Set("Authorization", "litty_auth")
	user := models.User{
		Username: "chasew",
	}

	tests := []routeTest{
		routeTest{
			req:            missingEventReq,
			service:        GetEvent,
			user:           user,
			code:           http.StatusBadRequest,
			responseStruct: models.Event{},
			expectError:    true,
			name:           "nonexistent event request",
		},
		routeTest{
			req:     unauthorizedEventReq,
			service: GetEvent,
			user: models.User{
				Username: "test_user",
			},
			code:           http.StatusBadRequest,
			responseStruct: models.Event{},
			expectError:    true,
			name:           "unauthorized auth request",
		},
	}

	for _, test := range tests {
		assertExpectedResponse(test, t)
	}
}
