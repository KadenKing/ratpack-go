package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
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
	Token       string `json:"token"`
	TeamDomain  string `json:"team_domain"`
	ChannelID   string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	UserID      string `json:"user_id"`
	UserName    string `json:"user_name"`
	Command     string `json:"command"`
	ResponseURL string `json:"response_url"`
	TriggerID   string `json:"trigger_id"`
}

type slackWriter struct {
	destination string
}

type slackResponseWriter interface {
	SetDestination(dest string)
	Destination() string
	Write(p []byte) (n int, err error)
}

func newSlackWriter() *slackWriter {
	return &slackWriter{}
}

func (sw slackWriter) Write(p []byte) (n int, err error) {
	if sw.destination == "" {
		return 0, errors.New("no destination to respond to slack specified")
	}

	text := string(p)
	payload := struct{ text string }{
		text: text,
	}
	serialized, err := json.Marshal(payload)
	if err != nil {
		return 0, err
	}

	r, err := http.Post(sw.destination, "application/json", bytes.NewReader(serialized))
	if err != nil {
		return 0, err
	}

	if r.StatusCode != 200 {
		return 0, errors.New("failed to send slack message")
	}

	return len(p), nil
}

func unmarshalSlackRequest(body io.Reader) slackRequest {
	b, _ := ioutil.ReadAll(body)

	var res slackRequest
	json.Unmarshal(b, &res)
	return res
}
