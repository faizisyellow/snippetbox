package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

const methodNotAllowed = 405

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {

		http.NotFound(w, r)
		return
	}

	// Initialize a slice containing the paths to the two files. Note that the
	// home.page.tmpl file must be the *first* file in the slice.
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())

		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a spesific snippet with ID...%d", id)

}

func (app *application) createSnipper(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", methodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippe..."))
}
