package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
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

func unmarshalSlackRequest(body io.Reader) slackRequest {
	b, _ := ioutil.ReadAll(body)

	var res slackRequest
	json.Unmarshal(b, &res)
	return res
}
