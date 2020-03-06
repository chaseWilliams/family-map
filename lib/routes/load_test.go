package routes

import (
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/chaseWilliams/family-map/lib/database"
	"github.com/chaseWilliams/family-map/lib/models"
	"github.com/chaseWilliams/family-map/lib/util"
)

func TestLoad(t *testing.T) {
	database.StartTestingSession(t)
	f, err := os.Open("load_data.json")
	fatal(err, t)
	req, err := http.NewRequest("POST", "/load", f)
	fatal(err, t)

	assertExpectedResponse(
		routeTest{
			req: req,
			service: Load,
			user: models.User{},
			code: http.StatusOK,
			responseStruct: util.Message{},
			expectError: false,
			name: "load",
		},
		t,
	)
}

func TestLoadFailure(t *testing.T) {
	database.StartTestingSession(t)
	req, err := http.NewRequest("POST", "/load", strings.NewReader(`{"a": "foo"}`))
	fatal(err, t)

	assertExpectedResponse(
		routeTest{
			req: req,
			service: Load,
			user: models.User{},
			code: http.StatusBadRequest,
			responseStruct: util.Message{},
			expectError: true,
			name: "malformed data load",
		},
		t,
	)
}