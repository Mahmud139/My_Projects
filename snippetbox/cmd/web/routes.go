package main

import "net/http"

// Update the signature for the routes() method so that it returns a 
// http.Handler instead of *http.ServeMux.
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fileServer := http.FileServer(http.Dir("M:/code_of_Golang/go_workspace/src/projects/snippetbox/ui/static"))
	mux.Handle("/static/",http.StripPrefix("/static", fileServer))

	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}