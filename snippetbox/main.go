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

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting server on localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	checkErr(err)
}