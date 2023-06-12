package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type aplication struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	templateCache map[string]*template.Template
}

func main() {
	//arguments
	addr := flag.String("addr", ":4000", "HTTP network address")
	connStr := flag.String("connStr", "user=dev password=dev dbname=dev sslmode=disable", "Connection string to Postgresql database")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := aplication{
		errorLog:      errorLog,
		infoLog:       infoLog,
		templateCache: nil, // TODO: add cached templates
	}
	_ = app
	//connect to db
	db, err := connectToDB(*connStr)
	if err != nil {
		errorLog.Fatal(err)
	}
	//close db after end of program
	defer db.Close()

	http.ListenAndServe(*addr, app.getRoutes())
}

func connectToDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
