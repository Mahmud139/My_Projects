package main

import (
	"log"
	"net/http"
	"html/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("view.gohtml")
	check(err)
	err = tmpl.Execute(writer, copy)
	check(err)
}
var copy string
func pasteHandler(writer http.ResponseWriter, request *http.Request) {
	copy = request.FormValue("paste")
	// _, err := writer.Write([]byte(copy))
	// check(err)
	http.Redirect(writer, request, "/home", http.StatusFound)
}

func main() {
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/paste", pasteHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	check(err)
}