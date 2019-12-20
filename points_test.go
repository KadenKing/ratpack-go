package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
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

func getTestSlackParameters() url.Values {
	// r := slackRequest{
	// 	Token:       "123",
	// 	ResponseURL: "abc123",
	// }
	data := url.Values{}
	data.Add("token", "123")
	data.Add("response_url", "abc")
	return data

}

type testSlackWriter struct {
	buf *bytes.Buffer
}

func newTestSlackWriter() *testSlackWriter {
	return &testSlackWriter{buf: bytes.NewBuffer(nil)}
}

func (b *testSlackWriter) SetDestination(dest string) {

}

func (b *testSlackWriter) Write(p []byte) (n int, err error) {
	b.buf.Write(p)
	return len(p), nil
}

func TestHandleGivePoints(t *testing.T) {
	slackParams := getTestSlackParameters()
	req, err := http.NewRequest("POST", "/api/give", strings.NewReader(slackParams.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	slackWriter := newTestSlackWriter()

	rr := httptest.NewRecorder()
	testServer := newTestServer()
	handler := http.HandlerFunc(testServer.handleGivePoints(slackWriter))

	handler.ServeHTTP(rr, req)

	if slackWriter.buf.String() != "you added points" {
		t.Errorf("\nexpected: %s\ngot: %s", "you added points", slackWriter.buf.String())
	}
}
