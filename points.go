package main

import (
	"fmt"
	"log"
	"net/http"
)

func (s *server) handleGivePoints(slackWriterGenerator slackResponseWriterGenerator, pointCommandGenerator pointsCommandGenerator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sr, err := unmarshalSlackRequest(r)
		slackWriter := slackWriterGenerator(sr.ResponseURL)

		if err != nil {
			log.Fatal(err)
		}

		if sr.ResponseURL == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "no response url given")
			return
		}
		fmt.Printf("\n%s\n", sr.Text)

		command := pointCommandGenerator(GIVE, s.storage)
		if err != nil {
			fmt.Fprintf(w, "error: %s", err.Error())
			return
		}
		pd := pointData{user: sr.UserID, arguments: sr.Text}
		command(pd)

		userWhoAddedPoints, err := s.slackAPI.GetProfileByID(sr.UserID)
		if err != nil {
			fmt.Fprint(w, "error: could not figure out who added points")
		}

		_, err = fmt.Fprintf(slackWriter, "%s added points", userWhoAddedPoints)

		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(200)
		fmt.Fprint(w, "")
	}
}
