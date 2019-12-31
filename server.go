package main

import "net/http"

type server struct {
	router   *http.ServeMux
	storage  storage
	env      environment
	slackAPI slackAPI
}

func newServer() *server {
	env := newEnv()
	router := http.NewServeMux()
	mongodb := newMongodb(env)
	slackAPI := newSlackAPI(env)

	server := &server{router: router, env: env}
	server.storage = mongodb
	server.slackAPI = slackAPI

	server.routes()
	return server
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
