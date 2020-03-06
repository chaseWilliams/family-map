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

func TestClear(t *testing.T) {
	database.StartTestingSession(t)
	req, err := http.NewRequest("GET", "/clear", nil)
	recorder := httptest.NewRecorder()
	if err != nil {
		t.Fatal(err)
	}

	Clear(recorder, req, models.User{})

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("wrong status code. wanted %d, got %d", http.StatusOK, recorder.Code)
	}

	response := util.Message{}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("response was not expected structure. %v", err)
	}
}
