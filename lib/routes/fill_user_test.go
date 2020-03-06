package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chaseWilliams/family-map/lib/database"
	"github.com/chaseWilliams/family-map/lib/models"
	"github.com/chaseWilliams/family-map/lib/util"
)

func TestFillUser(t *testing.T) {
	database.StartTestingSession(t)
	req, err := http.NewRequest("POST", "/fill/chasew/2", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	err = FillUser(rr, req, models.User{})
	if err != nil {
		t.Error(err)
	}

	if rr.Code != http.StatusOK {
		t.Errorf("non 200 response: %v", rr.Code)
	}

	message := util.Message{}
	err = json.Unmarshal(rr.Body.Bytes(), &message)
	if err != nil {
		t.Errorf("could not unmarshal response: %v", err)
	}
}

func TestFillUserFailure(t *testing.T) {
	database.StartTestingSession(t)
	badNumReq, err := http.NewRequest("POST", "/fill/chasew/-1", nil)
	if err != nil {
		t.Fatal(err)
	}
	noUserReq, err := http.NewRequest("POST", "/fill/bob/3", nil)
	if err != nil {
		t.Fatal(err)
	}
	tests := []routeTest{
		routeTest{
			req:            badNumReq,
			service:        FillUser,
			user:           models.User{},
			code:           http.StatusBadRequest,
			responseStruct: util.Message{},
			expectError:    true,
			name:           "bad generation number",
		},
		routeTest{
			req:            noUserReq,
			service:        FillUser,
			user:           models.User{},
			code:           http.StatusBadRequest,
			responseStruct: util.Message{},
			expectError:    true,
			name:           "bad generation number",
		},
	}

	for _, test := range tests {
		assertExpectedResponse(test, t)
	}
}
