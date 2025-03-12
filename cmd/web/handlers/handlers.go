package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"faizisyellow.com/snippetbox/cmd/web/config"
)

const methodNotAllowed = 405

func Home(app *config.Aplication) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

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
			app.ErrorLog.Println(err.Error())

			http.Error(w, "Internal Server Error", 500)
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	}
}

func ShowSnippet(app *config.Aplication) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			http.NotFound(w, r)
			return
		}

		fmt.Fprintf(w, "Display a spesific snippet with ID...%d", id)

	}
}

func CreateSnipper(app *config.Aplication) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			w.Header().Set("Allow", "POST")
			http.Error(w, "Method Not Allowed", methodNotAllowed)
			return
		}

		w.Write([]byte("Create a new snippe..."))
	}
}
