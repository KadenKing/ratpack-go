package main

import (
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

}

func main() {
	server := newServer()

	err := http.ListenAndServe(getPort(), server)
	if err != nil {
		log.Fatal(err)
	}
}
