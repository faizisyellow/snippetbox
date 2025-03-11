package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

const methodNotAllowed = 405

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {

		http.NotFound(w, r)
		return
	}
	ts, err := template.ParseFiles("./ui/html/home.page.html")
	if err != nil {
		fmt.Println(err.Error())

		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a spesific snippet with ID...%d", id)

}

func createSnipper(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", methodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippe..."))
}
