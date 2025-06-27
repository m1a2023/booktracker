package main

import (
	"book-tracker/internal/books/handlers"
	"book-tracker/internal/db"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var ( // server vars
	HOST 		string
	PORT		string
)

func main() {
	// Setup flags 
	setupFlags()

	// Database connection 
	con := db.GetConnection()
	if err := con.Ping(); err != nil {
		log.Fatalln(err)
	}
	// Defer connection close  
	defer con.Close()

	// Setting up endpoints
	setupRoutes()
}	


func setupRoutes() {

	// Creating new server
	mux := http.NewServeMux()

	// Handlers
	mux.HandleFunc("/books", handlers.GetBooks)
	mux.HandleFunc("/book", handlers.PostBook)

	// Listening
	adr := []string{HOST, PORT}
	err := http.ListenAndServe(strings.Join(adr, ":"), mux)
	log.Fatal(err)
}


func setupFlags() {
	// host 
	flag.StringVar(&HOST, "host", "0.0.0.0", "IP address to serve on")
	flag.StringVar(&HOST, "h", "0.0.0.0", "IP address shortcut")
	// port 
	flag.StringVar(&PORT, "port", "8080", "Port to listen on")
	flag.StringVar(&PORT, "p", "8080", "Port shortcut")


	// Setting up 
	flag.Usage = printHelp
	flag.Parse() 

	if len(flag.Args()) == 0 {
		printUsage()
	}

	fmt.Printf("\nServer started at http://%s:%s\nUsing database file: %s\n", HOST, PORT, db.DSN)
}


func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage: cmd [OPTIONS]\n\n")
	fmt.Fprintf(os.Stderr, "Options:\n")
	fmt.Fprintf(os.Stderr, "  -h, --host string      IP address to serve on (default \"%s\")\n", HOST)
	fmt.Fprintf(os.Stderr, "  -p, --port string      Port to listen on (default \"%s\")\n", PORT)
	fmt.Fprintf(os.Stderr, "  -d, --dsn string       Path to SQLite database file (default \"%s\")\n", db.DSN)
	fmt.Fprintf(os.Stderr, "  --help                 Show this help message\n")
}


func printHelp() {
	fmt.Fprintf(os.Stderr, "cmd - Serve your app with configurable options\n\n")
	fmt.Fprintf(os.Stderr, "Options:\n")
	fmt.Fprintf(os.Stderr, "  -h, --host string\n")
	fmt.Fprintf(os.Stderr, "        IP address to serve on.\n")
	fmt.Fprintf(os.Stderr, "        Default is \"%s\".\n\n", HOST)

	fmt.Fprintf(os.Stderr, "  -p, --port string\n")
	fmt.Fprintf(os.Stderr, "        Port to listen on.\n")
	fmt.Fprintf(os.Stderr, "        Default is \"%s\".\n\n", PORT)

	fmt.Fprintf(os.Stderr, "  -d, --dsn string\n")
	fmt.Fprintf(os.Stderr, "        Path to the SQLite database file.\n")
	fmt.Fprintf(os.Stderr, "        Default is \"%s\".\n\n", db.DSN)

	fmt.Fprintf(os.Stderr, "Example:\n")
	fmt.Fprintf(os.Stderr, "  cmd --host=0.0.0.0 --port=8080 --dsn=store.db\n")
}

