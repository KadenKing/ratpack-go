// +build integration

package main

import "testing"

func TestIncrement(t *testing.T) {
	server := newServer()
	server.storage.IncrementPoints("kaden", 1)
}
