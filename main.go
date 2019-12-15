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
		log.Fatal("could not find a port")
	}
	return ":" + port
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(getPort(), nil)
	if err != nil {
		log.Fatal(err)
	}
}
