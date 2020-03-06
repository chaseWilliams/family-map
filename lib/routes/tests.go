package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chaseWilliams/family-map/lib/models"
)

const (
	authToken string = "VEZQWVDNYNIDSVTA"
)

type routeTest struct {
	req            *http.Request
	service        Route
	user           models.User
	code           int
	responseStruct interface{}
	expectError    bool
	name           string
}

func assertExpectedResponse(rt routeTest, t *testing.T) {
	rr := httptest.NewRecorder()
	err := rt.service(rr, rt.req, rt.user)
	if err != nil && !rt.expectError {
		t.Errorf("service for %s returned an error: %v", rt.name, err)
	}

	if rr.Code != rt.code {
		t.Errorf("%s: non %d code received: %d", rt.name, rt.code, rr.Code)
	}

	err = json.Unmarshal(rr.Body.Bytes(), &rt.responseStruct)
	if err != nil {
		t.Errorf("%s: could not unmarshal response: %v", rt.name, err)
	}
}

func fatal(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}