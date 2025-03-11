package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnipper)

	log.Println("Starting server on localhost:4000")
	err := http.ListenAndServe("localhost:4000", mux)
	log.Fatal(err)
}
