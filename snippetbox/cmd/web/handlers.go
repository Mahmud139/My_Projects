package main

import (
	"errors"
	"fmt"
	//"html/template"
	"net/http"
	"strconv"
	"strings" 
	"unicode/utf8"

	"mahmud139/snippetbox/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	//http.NotFound(w, r)
	// 	app.notFound(w)
	// 	return
	// }

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{
		Snippets: s,
	})

	/*
	data := &templateData{Snippets: s}
	// for _, snippet := range s {
	// 	fmt.Fprintf(w, "%v\n", snippet)
	// }

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
	err = tmpl.Execute(w, data)
	if err != nil {
		// app.errorLog.Println(err.Error())
		// http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		app.serverError(w, err)
	}
	//w.Write([]byte("Hello from SnippetBox")) */
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
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

	app.render(w, r, "show.page.tmpl", &templateData{
		Snippet: s,
	})

	/*
	data := &templateData{Snippet: s}

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

	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
	}

	//w.Write([]byte("Display a new snippet"))
	//fmt.Fprintf(w,"Display with specific snippet with ID %d...", id)
	//fmt.Fprintf(w, "%v", s) */
}

func (app *application) deleteSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Println("deleteSnippet is running", id)

	err = app.snippets.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Create a new snippet..."))
	app.render(w, r, "create.page.tmpl", nil)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST" {
	// 	// w.Header().Set("Allow", "POST")
	// 	w.Header().Set("Allow", http.MethodPost)
	// 	// w.WriteHeader(405)
	// 	// w.Write([]byte("Method not Allowed!"))
	// 	// http.Error(w, "Method not Allowed!", 405)
	// 	// http.Error(w,"Method not Allowed!", http.StatusMethodNotAllowed)
	// 	app.clientError(w, http.StatusMethodNotAllowed)
	// 	return
	// }

	// First we call r.ParseForm() which adds any data in POST request bodies 
	// to the r.PostForm map. This also works in the same way for PUT and PATCH 
	// requests. If there are any errors, we use our app.ClientError helper to send 
	// a 400 Bad Request response to the user.
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Use the r.PostForm.Get() method to retrieve the relevant data fields 
	// from the r.PostForm map.
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires := r.PostForm.Get("expires")

	// Initialize a map to hold any validation errors.
	errors := make(map[string]string)
	// Check that the title field is not blank and is not more than 100 characters
	// long. If it fails either of those checks, add a message to the errors 
	// map using the field name as the key.
	if strings.TrimSpace(title) == "" {
		errors["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(title) > 100 {
		errors["title"] = "This field is too long (maximum is 100 characters"
	}

	if strings.TrimSpace(content) == "" {
		errors["content"] = "This field cannot be blank"
	}

	if strings.TrimSpace(expires) == "" {
		errors["expires"] = "This field cannot be blank"
	} else if expires != "365" && expires != "30" && expires != "7" {
		errors["expires"] = "This field is invalid"
	}

	if len(errors) > 0 {
		fmt.Fprint(w, errors)
		return
	}

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	//w.Write([]byte("Create a new snippet"))
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d",id), http.StatusSeeOther)
	/*The HTTP response status code 303 See Other is a way to redirect web applications 
	to a new URI, particularly after a HTTP POST has been performed*/
}