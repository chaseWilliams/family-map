package routes

import (
	"net/http"
	"strings"
	"testing"
	"math/rand"
	"time"

	"github.com/chaseWilliams/family-map/lib/database"
	"github.com/chaseWilliams/family-map/lib/models"
)

func TestRegister(t *testing.T) {
	database.StartTestingSession(t)
	rand.Seed(time.Now().Unix())
	req, err := http.NewRequest("POST", "/user/register", strings.NewReader(
		`{
			"userName": "susan",
			"password": "password",
			"email": "ok_boomer@yahoo.mail",
			"firstName": "Karen",
			"lastName": "Smith",
			"gender": "f"
		}`,
	))
	fatal(err, t)
	
	assertExpectedResponse(
		routeTest{
			req: req,
			service: Register,
			user: models.User{},
			code: http.StatusOK,
			responseStruct: loginSuccess{},
			expectError: false,
			name: "register",
		},
		t,
	)
}

func TestRegisterFailure(t *testing.T) {
	database.StartTestingSession(t)
	rand.Seed(time.Now().Unix())
	usernameTakenReq, err := http.NewRequest("POST", "/user/register", strings.NewReader(
		`{
			"userName": "chasew",
			"password": "password",
			"email": "ok_boomer@yahoo.mail",
			"firstName": "Karen",
			"lastName": "Smith",
			"gender": "f"
		}`,
	))
	fatal(err, t)
	missingPasswordReq, err := http.NewRequest("POST", "/user/register", strings.NewReader(
		`{
			"userName": "chasew",
			"email": "ok_boomer@yahoo.mail",
			"firstName": "Karen",
			"lastName": "Smith",
			"gender": "f"
		}`,
	))
	fatal(err, t)
	
	tests := []routeTest{
		routeTest{
			req: usernameTakenReq,
			service: Register,
			user: models.User{},
			code: http.StatusBadRequest,
			responseStruct: loginSuccess{},
			expectError: true,
			name: "username taken register",
		},
		routeTest{
			req: missingPasswordReq,
			service: Register,
			user: models.User{},
			code: http.StatusBadRequest,
			responseStruct: loginSuccess{},
			expectError: true,
			name: "missing password register",
		},
	}

	for _, test := range tests {
		assertExpectedResponse(
			test,
			t,
		)
	}
}