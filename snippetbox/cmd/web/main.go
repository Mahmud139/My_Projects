package main

import (
	"log"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("M:/code_of_Golang/go_workspace/src/projects/snippetbox/ui/static"))
	mux.Handle("/static/",http.StripPrefix("/static", fileServer))

	log.Println("Starting Server on localhost:8080")
	err := http.ListenAndServe("localhost:8080", mux)
	checkErr(err)
}