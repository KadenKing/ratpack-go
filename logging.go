package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func addLogger(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		log.Printf("[%s] body:\n%s", r.Method, string(body))
		f(w, r)
	}
}
