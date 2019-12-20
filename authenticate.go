package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func addSlackAuthenticity(h http.HandlerFunc) http.HandlerFunc {
	signingSecret := os.Getenv("SLACK_SIGNING_SECRET")
	if signingSecret == "" {
		log.Println("WARNING: no signing secret")
	}
	// Use the testing version of the signing secret. This will not work with slack
	signingSecret = "abcd123"

	return func(w http.ResponseWriter, r *http.Request) {
		slackSignedHMAC := r.Header.Get("X-Slack-Signature")
		slackSignedHMAC = slackSignedHMAC[strings.Index(slackSignedHMAC, "=")+1:]

		slackTimestamp := r.Header.Get("X-Slack-Request-Timestamp")

		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, "Could not read slack body")
		}

		body := string(b)

		basestring := "v0:" + slackTimestamp + ":" + body

		mac := hmac.New(sha256.New, []byte(signingSecret))
		mac.Write([]byte(basestring))

		expectedMac := hex.EncodeToString(mac.Sum(nil))
		if !hmac.Equal([]byte(slackSignedHMAC), []byte(expectedMac)) {
			w.WriteHeader(500)
			fmt.Fprintf(w, "this message was not authenticated")
			return
		}

		h(w, r)
	}
}
