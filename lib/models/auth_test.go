package models

import (
	"testing"

	"github.com/chaseWilliams/family-map/lib/database"
	"github.com/chaseWilliams/family-map/lib/util"
)

func TestSaveAuth(t *testing.T) {
	database.StartTestingSession(t)
	auth := Auth {
		AuthToken: util.RandomID(),
		Username: "chasew",
	}

	/*
	badAuth := Auth {
		AuthToken: util.RandomID(),
		Username: "idontexist",
	}
	*/

	err := auth.Save()
	if err != nil {
		t.Errorf("could not save auth: %v", err)
	}

	/*
	err = badAuth.Save()
	if err == nil {
		t.Error("bad auth should have failed, but it saved")
	}
	*/
}

func TestAssertAuth(t *testing.T) {
	_, ok := AssertAuth("VEZQWVDNYNIDSVTA")
	if !ok {
		t.Error("auth failed for a valid token")
	}

	_, ok = AssertAuth("blahblahblah")
	if ok {
		t.Error("auth did not fail for invalid token")
	}
}