package main

import "net/http"

type server struct {
	router  *http.ServeMux
	storage storage
}

func newServer() *server {
	router := http.NewServeMux()
	postgres := newPostgres()

	server := &server{router: router, storage: postgres}

	server.routes()
	return server
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
