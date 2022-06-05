package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id")) 
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	//w.Write([]byte("Display a new snippet"))
	fmt.Fprintf(w, "Display with specific snippet with ID %d...", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		// w.WriteHeader(405)
		// w.Write([]byte("Method not Allowed"))
		//http.Error(w, "Method not Allowed!", 405)
		http.Error(w, "Method not Allowed!", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on localhost:8080")
	err := http.ListenAndServe("localhost:8080", mux)
	checkErr(err)
}