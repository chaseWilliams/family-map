package models

import (
	"github.com/chaseWilliams/family-map/lib/database"
	"testing"
)

func TestLoginUser(t *testing.T) {
	database.StartTestingSession(t)
	d := LoginData{
		Username: "chasew",
		Password: "password",
	}
	user := User{
		Username:  "chasew",
		Password:  "password",
		Email:     "lol@gmail.com",
		FirstName: "Chase",
		LastName:  "Williams",
		Gender:    "m",
		PersonID:  "XVLBZGBAICMRAJWW",
	}
	userResult, err := d.GetUser()
	if err != nil {
		t.Errorf("failed to get user for login credentials: %v", err)
	}

	if user != *userResult {
		t.Errorf("user did not match expected user struct")
	}
}

func TestLoginFailure(t *testing.T) {
	database.StartTestingSession(t)
	d := LoginData{
		Username: "chaseww",
		Password: "password",
	}
	if _, err := d.GetUser(); err == nil {
		t.Errorf("Login did not properly fail for %v", d)
	}
}

func TestSaveUser(t *testing.T) {
	database.StartTestingSession(t)
	user := User{
		Username:  "chasetheman",
		Password:  "password",
		Email:     "email@m.org",
		FirstName: "chase",
		LastName:  "williams",
		Gender:    "m",
		PersonID:  "abc",
	}
	err := user.Save()
	if err != nil {
		t.Errorf("Could not save user: %v", err)
	}
}

func TestSaveUserFailure(t *testing.T) {
	database.StartTestingSession(t)
	user := User{
		Username:  "chasew",
		Password:  "password",
		Email:     "email@m.org",
		FirstName: "chase",
		LastName:  "williams",
		Gender:    "m",
		PersonID:  "abc",
	}
	err := user.Save()
	if err == nil {
		t.Error("user.Save() should have returned an error but didn't")
	}
}

func TestDeleteUser(t *testing.T) {
	database.StartTestingSession(t)
	err := DeleteUser("chasew")
	if err != nil {
		t.Errorf("could not delete user: %v", err)
	}
}
