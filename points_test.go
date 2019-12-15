package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
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

func newTestServer() *server {
	return &server{}
}

func getTestSlackParameters() []byte {
	r := slackRequest{
		Token: "123",
	}
	s, _ := json.Marshal(r)
	return s
}

func TestHandleGivePoints(t *testing.T) {

	slackParams := getTestSlackParameters()
	req, err := http.NewRequest("POST", "/api/give", bytes.NewReader(slackParams))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	testServer := newTestServer()
	handler := http.HandlerFunc(testServer.handleGivePoints())
	handler.ServeHTTP(rr, req)
	expected := "token:123"
	if rr.Body.String() != expected {
		t.Errorf("\nexpected %s\ngot %s", expected, rr.Body.String())
	}
}
