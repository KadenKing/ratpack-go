package main

import "net/http"

type router struct {
	mux *http.ServeMux
}

func newRouter() router {
	return router{http.NewServeMux()}
}

func (s *server) routes() {
	s.router.HandleFunc("/api/give", s.addSlackAuthenticity(s.handleGivePoints(newSlackWriterGenerator(), func() whoDidWhatParser {
		return giveCommandParser{}
	})))
}
