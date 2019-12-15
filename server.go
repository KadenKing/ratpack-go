package main

import "net/http"

type server struct {
	router *http.ServeMux
}

func newServer() *server {
	router := http.NewServeMux()
	server := &server{router}

	server.routes()
	return server
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
