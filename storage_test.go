package main

import "testing"

func TestIncrement(t *testing.T) {
	server := newServer()
	whoDidWhat := whoDidWhat{
		Who:    "thrifty watermelon",
		Did:    "gave",
		Points: -65,
	}
	err := server.storage.IncrementPoints("1234", whoDidWhat.Points)

	if err != nil {
		t.Error(err)
	}

	err = server.storage.LogChange("1234", whoDidWhat)

	if err != nil {
		t.Error(err)
	}
}
