// +build integration

package main

import "testing"

import "fmt"

func TestGetProfileById(t *testing.T) {

	env := newEnv()
	slackApi := newSlackAPI(env)
	name, err := slackApi.GetProfileByID("UK2DV0U2K")

	if err != nil {
		t.Errorf("error: %s", err)
	}

	fmt.Printf("name: %s\n", name)
}

func TestGetIDByProfileName(t *testing.T) {
	env := newEnv()
	slackAPI := newSlackAPI(env)

	id, err := slackAPI.GetIDByUsername("kaden.king.king")
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("\n id: %s\n", id)
}
