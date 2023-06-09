package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	//arguments
	addr := flag.String("addr", ":4000", "HTTP network address")
	connStr := flag.String("connStr", "zatimNic", "Connection string to Postgresql database")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	
	fmt.Println(addr, infoLog, errorLog) //unused variable.... TODO: delete

	//connect to db
	db, err := connectToDB(*connStr)
	if err != nil {
		errorLog.Fatal(err)
	}
	//close db after end of program
	defer db.Close()

	

}

func connectToDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil{
		return nil, err
	}
	return db, nil
}