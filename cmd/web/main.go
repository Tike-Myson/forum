package main

import (
	"database/sql"
	"flag"
	"github.com/Tike-Myson/forum/cmd/database"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	Reset  = "\033[0m"
	Red = "\033[31m"
	Green  = "\033[32m"
	db *sql.DB
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "./forum.db", "SQLite 3 data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, Green + "INFO\t" + Reset, log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, Red + "ERROR\t" + Reset, log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	createTables(db)
	defer db.Close()

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
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func createTables(db *sql.DB) {
	database.CreatePostsTable(db)
	database.CreateUsersTable(db)
}

