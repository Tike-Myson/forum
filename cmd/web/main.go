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

	srv := &http.Server {
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	infoLog.Printf("Server run on http://127.0.0.1%s\n", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

