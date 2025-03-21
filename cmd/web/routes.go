package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
	mux := pat.New()
	mux.Get("/", app.session.Enable(noSurf(http.HandlerFunc(app.home))))
	mux.Get("/snippet/create", app.session.Enable(noSurf(app.requireAuthenticatedUser(http.HandlerFunc(app.createSnippetForm)))))
	mux.Post("/snippet/create", app.session.Enable(noSurf(app.requireAuthenticatedUser(http.HandlerFunc(app.createSnippet)))))
	mux.Get("/snippet/:id", app.session.Enable(noSurf(http.HandlerFunc(app.showSnippet))))

	// users routes
	mux.Get("/user/signup", app.session.Enable(noSurf(http.HandlerFunc(app.signupUserForm))))
	mux.Post("/user/signup", app.session.Enable(noSurf(http.HandlerFunc(app.signupUser))))
	mux.Get("/user/login", app.session.Enable(noSurf(http.HandlerFunc(app.loginUserForm))))
	mux.Post("/user/login", app.session.Enable(noSurf(http.HandlerFunc(app.loginUser))))
	mux.Post("/user/logout", app.session.Enable(noSurf(http.HandlerFunc(app.logoutUser))))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
