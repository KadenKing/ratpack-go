package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/**
expected: v0=553dc5596cf01dd35c4388d9e4d3aca1461c165bac80b8b195c890b3656fe0cf
time: 1576788841
body:
token=KYFf0k2jd7qroRnIxWt4sqDg&team_id=TJP5Y3A3V&team_domain=peskypockers&channel_id=DK0JX3HQ8&channel_name=directmessage&user_id=UK2DV0U2K&user_name=kaden.king.king&command=%2Fgive&text=kaden+250&response_url=https%3A%2F%2Fhooks.slack.com%2Fcommands%2FTJP5Y3A3V%2F881047631104%2FZo0shoWQHRGzY21OeT3BWr1p&trigger_id=880590990212.635202112131.e64522e9f0ed0a991af428716d6cdd6a
Basestring: v0:1576788841:token=KYFf0k2jd7qroRnIxWt4sqDg&team_id=TJP5Y3A3V&team_domain=peskypockers&channel_id=DK0JX3HQ8&channel_name=directmessage&user_id=UK2DV0U2K&user_name=kaden.king.king&command=%2Fgive&text=kaden+250&response_url=https%3A%2F%2Fhooks.slack.com%2Fcommands%2FTJP5Y3A3V%2F881047631104%2FZo0shoWQHRGzY21OeT3BWr1p&trigger_id=880590990212.635202112131.e64522e9f0ed0a991af428716d6cdd6a
**/

func createHMAC(value, key string) string {
	hmac := hmac.New(sha256.New, []byte(key))
	hmac.Write([]byte(value))
	return hex.EncodeToString(hmac.Sum(nil))
}

func TestAddSlackAuthentication(t *testing.T) {
	type test struct {
		body     string
		key      string
		expected bool
	}

	tests := []test{
		{
			body:     "token=KYFf0k2jd7qroRnIxWt4sqDg&team_id=TJP5Y3A3V&team_domain=peskypockers&channel_id=DK0JX3HQ8&channel_name=directmessage&user_id=UK2DV0U2K&user_name=kaden.king.king&command=%2Fgive&text=kaden+250&response_url=https%3A%2F%2Fhooks.slack.com%2Fcommands%2FTJP5Y3A3V%2F881047631104%2FZo0shoWQHRGzY21OeT3BWr1p&trigger_id=880590990212.635202112131.e64522e9f0ed0a991af428716d6cdd6aBasestring: v0:1576788841:token=KYFf0k2jd7qroRnIxWt4sqDg&team_id=TJP5Y3A3V&team_domain=peskypockers&channel_id=DK0JX3HQ8&channel_name=directmessage&user_id=UK2DV0U2K&user_name=kaden.king.king&command=%2Fgive&text=kaden+250&response_url=https%3A%2F%2Fhooks.slack.com%2Fcommands%2FTJP5Y3A3V%2F881047631104%2FZo0shoWQHRGzY21OeT3BWr1p&trigger_id=880590990212.635202112131.e64522e9f0ed0a991af428716d6cdd6a",
			expected: true,
			key:      "abcd123",
		},
		{
			body:     "token=KYFf0k2jd7qroRnIxWt4sqDg&team_id=TJP5Y3A3V&team_domain=peskypockers&channel_id=DK0JX3HQ8&channel_name=directmessage&user_id=UK2DV0U2K&user_name=kaden.king.king&command=%2Fgive&text=kaden+250&response_url=https%3A%2F%2Fhooks.slack.com%2Fcommands%2FTJP5Y3A3V%2F881047631104%2FZo0shoWQHRGzY21OeT3BWr1p&trigger_id=880590990212.635202112131.e64522e9f0ed0a991af428716d6cdd6aBasestring: v0:1576788841:token=KYFf0k2jd7qroRnIxWt4sqDg&team_id=TJP5Y3A3V&team_domain=peskypockers&channel_id=DK0JX3HQ8&channel_name=directmessage&user_id=UK2DV0U2K&user_name=kaden.king.king&command=%2Fgive&text=kaden+250&response_url=https%3A%2F%2Fhooks.slack.com%2Fcommands%2FTJP5Y3A3V%2F881047631104%2FZo0shoWQHRGzY21OeT3BWr1p&trigger_id=880590990212.635202112131.e64522e9f0ed0a991af428716d6cdd6a",
			expected: false,
			key:      "notthekey",
		},
	}

	for _, currentTest := range tests {
		spyHandlerActiviated := false
		body := currentTest.body

		req, err := http.NewRequest("POST", "nowhere", strings.NewReader(body))
		if err != nil {
			t.Error(err)
		}
		time := "1576788841"
		req.Header.Add("X-Slack-Request-Timestamp", time)
		signedSecret := "v0=" + createHMAC("v0:"+time+":"+body, currentTest.key)
		req.Header.Add("X-Slack-Signature", signedSecret)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		var spyMiddleware http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
			spyHandlerActiviated = true
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(addSlackAuthenticity(spyMiddleware))
		handler.ServeHTTP(rr, req)

		if spyHandlerActiviated != currentTest.expected {
			t.Errorf("\nExpected handler to be run: %t\nDid handler run: %t\n", currentTest.expected, spyHandlerActiviated)
		}
	}
}
