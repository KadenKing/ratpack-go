package main

import (
	"fmt"
	"net/http"
)

func (s *server) handleGivePoints() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sr := unmarshalSlackRequest(r.Body)

		fmt.Fprintf(w, "token:%s", sr.Token)
	}
}
