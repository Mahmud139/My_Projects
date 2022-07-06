package main

import (
	"errors"
	"fmt"
	"html/template"
	"mahmud139/snippetbox/pkg/models"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		//http.NotFound(w, r)
		app.notFound(w)
		return
	}

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, snippet := range s {
		fmt.Fprintf(w, "%v\n", snippet)
	}

	/* commented out for testing our multiple query rows function from snippets.go

	//initialize a slice containing the paths to the two files. Note that the home.page.tmpl 
	//must be the first file in the slice.
	files := []string{
		"M:/code_of_Golang/go_workspace/src/projects/snippetbox/ui/html/home.page.tmpl",
		"M:/code_of_Golang/go_workspace/src/projects/snippetbox/ui/html/base.layout.tmpl",
		"M:/code_of_Golang/go_workspace/src/projects/snippetbox/ui/html/footer.partial.tmpl",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		//app.errorLog.Println(err.Error())
		// http.Error(w, "Internal Server Error", 500)
		//http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		app.serverError(w, err)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		// app.errorLog.Println(err.Error())
		// http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		app.serverError(w, err)
	}
	//w.Write([]byte("Hello from SnippetBox")) */
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		// http.NotFound(w, r)
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	files := []string {
		"M:/code_of_Golang/go_workspace/src/projects/snippetbox/ui/html/show.page.tmpl",
		"M:/code_of_Golang/go_workspace/src/projects/snippetbox/ui/html/base.layout.tmpl",
		"M:/code_of_Golang/go_workspace/src/projects/snippetbox/ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, s)
	if err != nil {
		app.serverError(w, err)
	}

	//w.Write([]byte("Display a new snippet"))
	//fmt.Fprintf(w,"Display with specific snippet with ID %d...", id)
	//fmt.Fprintf(w, "%v", s)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// w.Header().Set("Allow", "POST")
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(405)
		// w.Write([]byte("Method not Allowed!"))
		// http.Error(w, "Method not Allowed!", 405)
		// http.Error(w,"Method not Allowed!", http.StatusMethodNotAllowed)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// Create some variables holding dummy data. We'll remove these later on 
	// during the build.
	title := "O snail" 
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa" 
	expires := "7"

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	//w.Write([]byte("Create a new snippet"))
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d",id), http.StatusSeeOther)
	/*The HTTP response status code 303 See Other is a way to redirect web applications 
	to a new URI, particularly after a HTTP POST has been performed*/
}