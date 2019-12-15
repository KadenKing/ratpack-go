package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (s *server) handleGivePoints() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("point handler:")
		r.ParseForm()
		slackRequest, _ := json.Marshal(map[string]string{
			"text": "howdy partner",
		})

		responseUrl := r.Form.Get("response_url")
		fmt.Printf("url: %s", responseUrl)

		_, err := http.Post(r.Form.Get("response_url"), "application/json", bytes.NewBuffer(slackRequest))
		if err != nil {
			log.Fatal((err))
		}
	}
}
