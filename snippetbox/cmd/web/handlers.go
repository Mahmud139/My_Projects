package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//initialize a slice containing the paths to the two files. Note that the home.page.tmpl 
	//must be the first file in the slice.
	files := []string{
		"M:/code_of_Golang/go_workspace/src/projects/snippetbox/ui/html/home.page.tmpl",
		"M:/code_of_Golang/go_workspace/src/projects/snippetbox/ui/html/base.layout.tmpl",
		"M:/code_of_Golang/go_workspace/src/projects/snippetbox/ui/html/footer.partial.tmpl",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		// http.Error(w, "Internal Server Error", 500)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	//w.Write([]byte("Hello from SnippetBox"))
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	//w.Write([]byte("Display a new snippet"))
	fmt.Fprintf(w,"Display with specific snippet with ID %d...", id)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		// w.WriteHeader(405)
		// w.Write([]byte("Method not Allowed!"))
		// http.Error(w, "Method not Allowed!", 405)
		http.Error(w,"Method not Allowed!", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet"))
}