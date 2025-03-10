package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a spesific snippet..."))
}
func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on localhost:4000")
	err := http.ListenAndServe("localhost:4000", mux)
	log.Fatal(err)
}
