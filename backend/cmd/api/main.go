package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	//arguments
	addr := flag.String("addr", ":4000", "HTTP network address")
	connStr := flag.String("connStr", "zatimNic", "Connection string to Postgresql database")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
}
