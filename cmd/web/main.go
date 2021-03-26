package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var (
	Reset  = "\033[0m"
	Red = "\033[31m"
	Green  = "\033[32m"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, Green + "INFO\t" + Reset, log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, Red + "ERROR\t" + Reset, log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/login", app.login)
	mux.HandleFunc("/signup", app.signup)
	mux.HandleFunc("/logout", app.logout)
	mux.HandleFunc("/comment", app.comment)
	mux.HandleFunc("/like", app.like)
	mux.HandleFunc("/dislike", app.dislike)
	mux.HandleFunc("/profile", app.profile)
	mux.HandleFunc("/likedPosts", app.likedPosts)
	mux.HandleFunc("/category", app.filterByCategory)

	fileServer := http.FileServer(http.Dir("../ui/static/"))
	mux.Handle("/ui/", http.StripPrefix("/static", fileServer))

	srv := &http.Server {
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: mux,
	}

	infoLog.Printf("Server run on http://127.0.0.1%s\n", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
