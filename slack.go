package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	schema "github.com/gorilla/schema"
)

/**
token=KYFf0k2jd7qroRnIxWt4sqDg&team_id=TJP5Y3A3V&
team_domain=peskypockers&
channel_id=DK0JX3HQ8&
channel_name=directmessage&
user_id=UK2DV0U2K&
user_name=kaden.king.king&
command=%2Fgive&text=kaden+25&
response_url=https%3A%2F%2Fhooks.slack.com%2Fcommands%2FTJP5Y3A3V%2F862889225923%2F1YyKGINXB3ci1kHzko7QvQSh&
trigger_id=876217331607.635202112131.776d240216a617135974615d3ace76e8
**/

type slackRequest struct {
	Token       string `schema:"token"`
	TeamDomain  string `schema:"team_domain"`
	TeamID      string `schema:"team_id"`
	Text        string `schema:"text"`
	ChannelID   string `schema:"channel_id"`
	ChannelName string `schema:"channel_name"`
	UserID      string `schema:"user_id"`
	UserName    string `schema:"user_name"`
	Command     string `schema:"command"`
	ResponseURL string `schema:"response_url"`
	TriggerID   string `schema:"trigger_id"`
}

type slackWriter struct {
	destination string
}

type slackResponseWriter interface {
	SetDestination(dest string)
	Write(p []byte) (n int, err error)
}

func newSlackWriter() *slackWriter {
	return &slackWriter{}
}

func (sw *slackWriter) SetDestination(dest string) {
	sw.destination = dest
}

func (sw slackWriter) Write(p []byte) (n int, err error) {
	if sw.destination == "" {
		return 0, errors.New("no destination to respond to slack specified")
	}

	text := string(p)
	payload := &struct {
		Text string `json:"text"`
	}{
		Text: text,
	}
	serialized, err := json.Marshal(payload)
	if err != nil {
		return 0, err
	}
	fmt.Println(string(serialized))

	r, err := http.Post(sw.destination, "application/json", strings.NewReader(string(serialized)))
	if err != nil {
		return 0, err
	}

	if r.StatusCode != 200 {
		body, _ := ioutil.ReadAll(r.Body)
		return 0, errors.New(string(body))
	}

	return len(p), nil
}

func unmarshalSlackRequest(r *http.Request) (slackRequest, error) {
	err := r.ParseForm()

	if err != nil {
		return slackRequest{}, err
	}

	var res slackRequest
	decoder := schema.NewDecoder()

	err = decoder.Decode(&res, r.PostForm)
	if err != nil {
		return slackRequest{}, err
	}

	return res, nil
}
