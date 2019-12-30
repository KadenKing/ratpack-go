package main

import "net/http"

type server struct {
	router  *http.ServeMux
	storage storage
	env     environment
}

func newServer() *server {
	env := newEnv()
	router := http.NewServeMux()
	mongodb := newMongodb(env)

	server := &server{router: router}
	server.storage = mongodb

	server.routes()
	return server
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
