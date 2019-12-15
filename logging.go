package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type logger struct {
	handler http.Handler
}

func (l *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	log.Printf("[%s] body:\n%s", r.Method, string(body))
	l.handler.ServeHTTP(w, r)
}

func NewLogger(handler http.Handler) *logger {
	return &logger{handler}
}
