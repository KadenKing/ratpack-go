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

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal(err)
	}
}
