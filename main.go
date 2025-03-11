package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const NotAllowedCode = 405

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a spesific snippet with ID: %d...", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")

		// Behind the scene it will call:
		// - w.WriteHeader(code)
		// - w.Write([]byte(response))
		http.Error(w, "Method Not Allowed", NotAllowedCode)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	/*@Page 47 */
	log.Println("Starting server on localhost:4000")
	err := http.ListenAndServe("localhost:4000", mux)
	log.Fatal(err)
}
