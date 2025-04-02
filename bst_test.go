package bst_test

import (
	"fmt"
	"testing"

	bst "github.com/adityarizkyramadhan/binary-search-tree-go"
)

type User struct {
	ID   int
	Name string
}

func TestBSTSearch(t *testing.T) {
	users := []User{
		{ID: 1, Name: "User1"},
		{ID: 2, Name: "User2"},
		{ID: 3, Name: "User3"},
		{ID: 4, Name: "User4"},
		{ID: 5, Name: "User5"},
		{ID: 6, Name: "User6"},
		{ID: 7, Name: "User7"},
		{ID: 8, Name: "User8"},
		{ID: 9, Name: "User9"},
		{ID: 10, Name: "User10"},
	}

	bstUsers := bst.NewBSTArray(users)

	lessByID := func(a, b User) bool {
		return a.ID < b.ID
	}

	equalByID := func(a, b User) bool {
		return a.ID == b.ID
	}

	tests := []struct {
		name   string
		target User
		found  bool
	}{
		{"User Exists", User{ID: 10}, true},
		{"User Exists", User{ID: 1}, true},
		{"User Exists", User{ID: 5}, true},
		{"User Does Not Exist", User{ID: 11}, false},
		{"User Does Not Exist", User{ID: 0}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := bstUsers.Search(tt.target, lessByID, equalByID)
			if (user != nil) != tt.found {
				t.Errorf("Test case failed for target ID %d. Expected found: %v, but got: %v", tt.target.ID, tt.found, user != nil)
			} else if user != nil {
				fmt.Println("User ditemukan:", user)
			}
		})
	}
}
