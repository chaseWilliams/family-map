package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chaseWilliams/family-map/lib/database"
	"github.com/chaseWilliams/family-map/lib/models"
)

func TestAllEvents(t *testing.T) {
	database.StartTestingSession(t)
	req, err := http.NewRequest("GET", "/event", nil)
	rr := httptest.NewRecorder()
	req.Header.Set("Authorization", "VEZQWVDNYNIDSVTA")
	if err != nil {
		t.Fatal(err)
	}

	err = AllEvents(rr, req, models.User{
		Username: "chasew",
	})
	if err != nil {
		t.Error(err)
	}

	if rr.Code != http.StatusOK {
		t.Errorf("non 200 status code: %v", rr.Code)
	}

	events := dataResponse{}
	err = json.Unmarshal(rr.Body.Bytes(), &events)
	if err != nil {
		t.Errorf("could not unmarshal response: %v", err)
	}
}

func TestAllEventsFailure(t *testing.T) {
	database.StartTestingSession(t)

	// missing auth token
	req, err := http.NewRequest("GET", "/event", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	defer func() {
		if p := recover(); p == nil {
			t.Error("no panic")
		}
	}()
	err = AllEvents(rr, req, models.User{})
}
