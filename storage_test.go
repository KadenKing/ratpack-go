// +build integration

package main

import "testing"

func TestIncrement(t *testing.T) {
	server := newServer()
	pc := pointChange{
		User:         "kaden",
		Points:       1,
		UserChanging: "test",
		Reason:       "being a good boy",
	}
	err := server.storage.IncrementPoints(pc)

	if err != nil {
		t.Error(err)
	}
}
