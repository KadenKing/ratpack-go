package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
)

type slackProfile struct {
	RealNameNormalized string `json:"real_name_normalized"`
}

type slackProfileResponse struct {
	Ok      bool         `json:"ok"`
	Profile slackProfile `json:"profile"`
	Error   string       `json:"error"`
}

type slackAPI interface {
	GetProfileByID(id string) (string, error)
}

func newSlackAPI(e environment) *slackConnection {
	token := e.Get("SLACK_TOKEN")
	if token == "" {
		log.Fatal("No slack token found")
	}

	return &slackConnection{token}
}

type slackConnection struct {
	token string
}

func (s *slackConnection) GetProfileByID(id string) (string, error) {
	reqURL, err := url.Parse("https://slack.com/api/users.profile.get")
	if err != nil {
		return "", err
	}

	query := url.Values{}
	query.Add("token", s.token)
	query.Add("user", id)

	reqURL.RawQuery = query.Encode()
	resp, err := http.Get(reqURL.String())

	if err != nil {
		return "", err
	}
	var response slackProfileResponse

	json.NewDecoder(resp.Body).Decode(&response)

	if !response.Ok {
		return "", errors.New(response.Error)
	}
	return response.Profile.RealNameNormalized, nil
}
