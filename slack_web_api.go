package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
)

type slackProfile struct {
	RealNameNormalized    string `json:"real_name_normalized"`
	DisplayNameNormalized string `json:"display_name_normalized"`
}

type slackProfileResponse struct {
	Ok      bool         `json:"ok"`
	Profile slackProfile `json:"profile"`
	Error   string       `json:"error"`
}

type slackMembersResponse struct {
	ID       string       `json:"id"`
	Name     string       `json:"name"`
	RealName string       `json:"real_name"`
	Profile  slackProfile `json:"profile"`
}

type slackUsersListResponse struct {
	Ok      bool                   `json:"ok"`
	Members []slackMembersResponse `json:"members"`
}

type slackAPI interface {
	slackIDTranslater
}

type slackIDTranslater interface {
	GetProfileByID(id string) (string, error)
	GetProfileByUsername(username string) (slackMembersResponse, error)
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

func searchUsersForUsernameID(username string, members []slackMembersResponse) (slackMembersResponse, bool) {
	for _, member := range members {
		if member.Name == username {
			return member, true
		}
	}

	return slackMembersResponse{}, false
}

func (s *slackConnection) GetProfileByUsername(username string) (slackMembersResponse, error) {
	reqURL, err := url.Parse("https://slack.com/api/users.list")
	if err != nil {
		return slackMembersResponse{}, err
	}

	query := url.Values{}
	query.Add("token", s.token)
	reqURL.RawQuery = query.Encode()

	resp, err := http.Get(reqURL.String())

	if err != nil {
		return slackMembersResponse{}, err
	}

	var response slackUsersListResponse
	json.NewDecoder(resp.Body).Decode(&response)

	if !response.Ok {
		return slackMembersResponse{}, errors.New("response not ok")
	}

	member, ok := searchUsersForUsernameID(username[1:], response.Members)

	if !ok {
		return slackMembersResponse{}, errors.New("could not find a username with that name")
	}

	return member, nil
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
	return response.Profile.DisplayNameNormalized, nil
}
