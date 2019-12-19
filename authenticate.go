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
		log.Fatal("no signing secret")
	}

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

		fmt.Printf("as bytes: %s\n", b)
		body := string(b)

		basestring := "v0:" + slackTimestamp + ":" + body

		fmt.Printf("Basestring: %s\n", basestring)
		fmt.Printf("Slack hmac: %s\n", slackSignedHMAC)

		mac := hmac.New(sha256.New, []byte(signingSecret))

		mac.Write([]byte(basestring))
		expectedMac := hex.EncodeToString(mac.Sum(nil))

		fmt.Printf("expected mac: %s\n", string(expectedMac))

		if hmac.Equal([]byte(slackSignedHMAC), []byte(expectedMac)) {
			log.Println("Correct!")
		} else {
			log.Fatal("Incorrect!")
		}

	}
}
