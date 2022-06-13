package main

import (
	"flag"
	"log"
	"net/http"
	//"os"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// addr := os.Getenv("SNIPPETBOX_ADDR")
	addr := flag.String("addr", "localhost:8080",  "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("M:/code_of_Golang/go_workspace/src/projects/snippetbox/ui/static"))
	mux.Handle("/static/",http.StripPrefix("/static", fileServer))

	// log.Printf("Starting Server on %v \n", addr)
	// err := http.ListenAndServe(addr, mux)
	// checkErr(err)
	log.Printf("Starting Server on %v \n", *addr)
	err := http.ListenAndServe(*addr, mux)
	checkErr(err)
}