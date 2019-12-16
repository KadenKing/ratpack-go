package main

import "net/http"

type middlewareRouter interface {
	HandleFunc(route string, h http.HandlerFunc)
}

type router struct {
	mux *http.ServeMux
}

func newRouter() router {
	return router{http.NewServeMux()}
}

func (s *server) routes() {
	s.router.HandleFunc("/api/give", s.handleGivePoints(newSlackWriter()))
}
