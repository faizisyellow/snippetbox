package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	/* @page 126 */

	// Define a new command-line flag with the name 'addr'. The value of
	// flag will be stored in the addr variable at runtime.
	addr := flag.String("addr", "localhost:4000", "HTTP network address")

	// Importantly, we use the flag.Parse() function to parse the command-line
	// This reads in the command-line flag value and assigns it to the addr
	// variable.
	flag.Parse()

	// Use log.New() to create a logger for writing information messages. This
	// three parameters: the destination to write the logs to (os.Stdout), a st
	// prefix for message (INFO followed by a tab), and flags to indicate what
	// additional information to include (local date and time). Note that the fl
	// are joined using the bitwise OR operator |.
	infoLog := log.New(os.Stdout, "\033[32mINFO\t\033[0m", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "\033[31mERROR\t\033[0m", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize a new instance of application containing the dependencies.
	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	// Initialize a new http.Server struct. We set the Addr and Handler fields
	// that the server uses the same network address and routes as before, and
	// the ErrorLog field so that the server now uses the custom errorLog logger
	// the event of any problems.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// Write messages using the two new loggers, instead of the standard logger
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
