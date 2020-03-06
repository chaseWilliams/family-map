package routes

import (
	"net/http"
	"strings"
	"testing"

	"github.com/chaseWilliams/family-map/lib/database"
	"github.com/chaseWilliams/family-map/lib/models"
)

func TestLogin(t *testing.T) {
	database.StartTestingSession(t)
	req, err := http.NewRequest("POST", "/user/login", strings.NewReader(
		`{"userName": "chasew", "password": "password"}`,
	))
	fatal(err, t)
	
	assertExpectedResponse(
		routeTest{
			req: req,
			service: Login,
			user: models.User{},
			code: http.StatusOK,
			responseStruct: loginSuccess{},
			expectError: false,
			name: "login",
		},
		t,
	)
}

func TestLoginFailure(t *testing.T) {
	database.StartTestingSession(t)
	req, err := http.NewRequest("POST", "/user/login", strings.NewReader(
		`{"userName": "chasew", "password": "bad_password"}`,
	))
	fatal(err, t)
	
	assertExpectedResponse(
		routeTest{
			req: req,
			service: Login,
			user: models.User{},
			code: http.StatusBadRequest,
			responseStruct: loginSuccess{},
			expectError: true,
			name: "bad password login",
		},
		t,
	)
}