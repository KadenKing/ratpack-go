package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hey")
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":8080"
	}
	return ":" + port
}

func testPointHandler(w http.ResponseWriter, r *http.Request) {
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

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/give", testPointHandler)
	mux.HandleFunc("/", handler)

	withLogging := NewLogger(mux)

	err := http.ListenAndServe(getPort(), withLogging)
	if err != nil {
		log.Fatal(err)
	}
}
