package main

import (
	"fmt"
	"net/http"
	"strconv"
	"html/template"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles("M:/code_of_Golang/go_workspace/src/projects/snippetbox/ui/html/home.page.tmpl")
	checkErr(err)
	err = tmpl.Execute(w, nil)
	checkErr(err)
	//w.Write([]byte("Hello from SnippetBox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	//w.Write([]byte("Display a new snippet"))
	fmt.Fprintf(w,"Display with specific snippet with ID %d...", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
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