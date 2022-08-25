package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golangcollege/sessions"

	"mahmud139/snippetbox/pkg/models"
	"mahmud139/snippetbox/pkg/models/mysql"

	_ "github.com/go-sql-driver/mysql"
)

type contexKey string

const contextKeyIsAuthenticated = contexKey("isAuthenticated")

// func checkErr(err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }
// type Config struct {
// 	Addr string
// 	StaticDir string
// }

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	session  *sessions.Session
	//snippets      *mysql.SnippetModel
	snippets interface {
		Insert(string, string, string) (int, error)
		Get(int) (*models.Snippet, error)
		Latest() ([]*models.Snippet, error)
		Delete(int) error
	}
	templateCache map[string]*template.Template
	//users         *mysql.UserModel
	users interface {
		Insert(string, string, string) error
		Authenticate(string, string) (int, error)
		Get(int) (*models.User, error)
		ChangePassword(int, string, string) error
	}
}

func main() {
	/* using pre-existing variable for command-line flags*/
	// cfg := new(Config)
	// flag.StringVar(&cfg.Addr, "addr", "localhost:8080", "HTTP network address")
	// flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	// flag.Parse()

	/* using Environment variable for command-line flags */
	// addr := os.Getenv("SNIPPETBOX_ADDR")

	addr := flag.String("addr", "localhost:8080", "HTTP network address")
	//Define a new command line flag for MySQL DSN string
	dsn := flag.String("dsn", "web:mahmud@/snippetbox?parseTime=true", "MySQL data source name")
	// Define a new command-line flag for the session secret (a random key which
	// will be used to encrypt and authenticate session cookies). It should be 32 bytes long.
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	//Initialize a new template cache
	templateCache, err := newTemplateCache("M:/code_of_Golang/go_workspace/src/projects/snippetbox/ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	// Use the sessions.New() function to initialize a new session manager,
	// passing in the secret key as the parameter. Then we configure it so
	// sessions always expires after 1 hours.
	session := sessions.New([]byte(*secret))
	session.Lifetime = 1 * time.Hour
	session.Secure = true
	session.SameSite = http.SameSiteStrictMode

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		session:       session,
		snippets:      &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
		users:         &mysql.UserModel{DB: db},
	}

	// Initialize a tls.Config struct to hold the non-default TLS settings we want
	// the server to use.
	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	/* transfer to routes.go file
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fileServer := http.FileServer(http.Dir("M:/code_of_Golang/go_workspace/src/projects/snippetbox/ui/static"))
	mux.Handle("/static/",http.StripPrefix("/static", fileServer)) */

	/* logging to a file */
	// file, err := os.OpenFile("/tmp/infoLog.txt", os.O_RDWR|os.O_CREATE, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// infoLog := log.New(file, "INFO\t", log.Ldate | log.Ltime)

	/* using pre-existing variable for command-line flags*/
	// log.Printf("Starting Server on %v \n", cfg.Addr)
	// err := http.ListenAndServe(cfg.Addr, mux)
	// checkErr(err)

	/* using Environment variable for command-line flags */
	// log.Printf("Starting Server on %v \n", addr)
	// err := http.ListenAndServe(addr, mux)
	// checkErr(err)

	/* using command-line flags and custom logger*/
	//log.Printf("Starting Server on %v \n", *addr)
	// infoLog.Printf("Starting Server on %v \n", *addr)
	// err := http.ListenAndServe(*addr, mux)
	//checkErr(err)
	//errorLog.Fatal(err)

	/* implementing the http.Server Error log using our custom logger*/
	infoLog.Printf("Starting Server on https://%v \n", *addr)
	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err = srv.ListenAndServeTLS("M:/code_of_Golang/go_workspace/src/projects/snippetbox/tls/cert.pem", "M:/code_of_Golang/go_workspace/src/projects/snippetbox/tls/key.pem")
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
