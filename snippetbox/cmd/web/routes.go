package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

// Update the signature for the routes() method so that it returns a 
// http.Handler instead of *http.ServeMux.
func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	dynamicMiddleware := alice.New(app.session.Enable)
	
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", app.home)
	// mux.HandleFunc("/snippet", app.showSnippet)
	// mux.HandleFunc("/snippet/create", app.createSnippet)

	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippetForm))
	mux.Post("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippet))
	mux.Get("/snippet/delete/:id", dynamicMiddleware.ThenFunc(app.deleteSnippet))
	mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(app.showSnippet))

	fileServer := http.FileServer(http.Dir("M:/code_of_Golang/go_workspace/src/projects/snippetbox/ui/static"))
	//mux.Handle("/static/",http.StripPrefix("/static", fileServer))
	mux.Get("/static/",http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}