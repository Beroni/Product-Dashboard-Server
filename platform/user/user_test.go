package user

import (
	"testing"
)

func TestAdd(t *testing.T) {
	users := New()
	users.Add(User{})

	if len(users.Users) != 1 {
		t.Errorf("Item was not added")
	}

}

func TestGetAll(t *testing.T) {

	users := New()
	users.Add(User{})

	results := users.GetAll()

	if len(results) != 1 {
		t.Errorf("Item was not added")
	}

}
