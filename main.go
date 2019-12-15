package main

import (
	"fmt"
	"io/ioutil"
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
	reqBody, _ := ioutil.ReadAll(r.Body)

	fmt.Fprint(w, string(reqBody))
}

func main() {
	http.HandleFunc("/api/give", testPointHandler)
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(getPort(), nil)
	if err != nil {
		log.Fatal(err)
	}
}
