package main

import (
	"fmt"
	"log"
	"net/http"
)

func (s *server) handleGivePoints(slackWriter slackResponseWriter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sr, err := unmarshalSlackRequest(r)

		if err != nil {
			log.Fatal(err)
		}

		if sr.ResponseURL == "" {
			log.Fatal("no response url given")
		}

		slackWriter.SetDestination(sr.ResponseURL)
		_, err = fmt.Fprintf(slackWriter, "you added points")

		if err != nil {
			log.Fatal(err)
		}
	}
}
