package main

import (
	"fmt"
	"log"
	"net/http"
)

func (s *server) handleGivePoints(slackWriterGenerator slackResponseWriterGenerator, giveCommandParserGenerator func() whoDidWhatParser) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		giveCommandParser := giveCommandParserGenerator()

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

		whoDidWhat, err := giveCommandParser.Parse(sr, s.slackAPI)

		if err != nil {
			fmt.Fprint(w, "error: could not figure out who added points")
		}

		err = s.storage.IncrementPoints(sr.UserID, whoDidWhat.points)

		userWhoAddedPoints := whoDidWhat.who
		_, err = fmt.Fprintf(slackWriter, "%s added points", userWhoAddedPoints)

		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(200)
		fmt.Fprint(w, "")
	}
}
