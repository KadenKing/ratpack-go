package main

import (
	"bytes"
	"io/ioutil"
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

func getTestSlackParameters(vals ...[2]string) url.Values {
	data := url.Values{}
	for _, val := range vals {
		data.Add(val[0], val[1])
	}
	return data

}

type testSlackWriter struct {
	buf *bytes.Buffer
}

func newTestSlackWriter() *testSlackWriter {
	return &testSlackWriter{buf: bytes.NewBuffer(nil)}
}

func (b *testSlackWriter) Write(p []byte) (n int, err error) {
	b.buf.Write(p)
	return len(p), nil
}

func TestHandleGivePoints(t *testing.T) {
	type test struct {
		params          url.Values
		expectedWritten string
		expectedStatus  int
		expectedBody    string
	}

	tests := []test{
		{
			params: getTestSlackParameters(
				[2]string{"response_url", "abc123"},
			),
			expectedWritten: "you added points",
			expectedStatus:  200,
			expectedBody:    "",
		},
		{
			params:          getTestSlackParameters(),
			expectedWritten: "",
			expectedStatus:  500,
			expectedBody:    "no response url given",
		},
	}

	for _, currentTest := range tests {
		slackParams := currentTest.params
		req, err := http.NewRequest("POST", "/api/give", strings.NewReader(slackParams.Encode()))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		slackWriter := newTestSlackWriter()

		slackWriterGenerator := func(url string) slackResponseWriter {
			return slackWriter
		}

		rr := httptest.NewRecorder()
		testServer := newTestServer()
		handler := http.HandlerFunc(testServer.handleGivePoints(slackWriterGenerator))

		handler.ServeHTTP(rr, req)

		if slackWriter.buf.String() != currentTest.expectedWritten {
			t.Errorf("\nexpected to be written to slackwriter: %s\ngot: %s", currentTest.expectedWritten, slackWriter.buf.String())
		}
		res := rr.Result()
		if res.StatusCode != currentTest.expectedStatus {
			t.Errorf("\nexpected status code: %d\ngot: %d\n", currentTest.expectedStatus, res.StatusCode)
		}
		body, _ := ioutil.ReadAll(res.Body)
		if string(body) != currentTest.expectedBody {
			t.Errorf("\nexpected response body: %s\n got: %s\n", currentTest.expectedBody, body)
		}
	}
}
