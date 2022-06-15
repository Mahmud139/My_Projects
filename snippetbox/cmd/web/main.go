package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// func checkErr(err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
// type Config struct {
// 	Addr string
// 	StaticDir string
// }

func main() {
	// cfg := new(Config)
	// flag.StringVar(&cfg.Addr, "addr", "localhost:8080", "HTTP network address")
	// flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	// flag.Parse()
	// addr := os.Getenv("SNIPPETBOX_ADDR")
	addr := flag.String("addr", "localhost:8080",  "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate | log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate | log.Ltime | log.Lshortfile)
	
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("M:/code_of_Golang/go_workspace/src/projects/snippetbox/ui/static"))
	mux.Handle("/static/",http.StripPrefix("/static", fileServer))

	//logging to a file
	// file, err := os.OpenFile("/tmp/infoLog.txt", os.O_RDWR|os.O_CREATE, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// infoLog := log.New(file, "INFO\t", log.Ldate | log.Ltime)

	// log.Printf("Starting Server on %v \n", cfg.Addr)
	// err := http.ListenAndServe(cfg.Addr, mux)
	// checkErr(err)

	// log.Printf("Starting Server on %v \n", addr)
	// err := http.ListenAndServe(addr, mux)
	// checkErr(err)

	//log.Printf("Starting Server on %v \n", *addr)
	// infoLog.Printf("Starting Server on %v \n", *addr)
	// err := http.ListenAndServe(*addr, mux)
	//checkErr(err)
	//errorLog.Fatal(err)
	infoLog.Printf("Starting Server on %v \n", *addr)
	srv := &http.Server {
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: mux,
	}
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}