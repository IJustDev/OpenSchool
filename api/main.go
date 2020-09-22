package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func establishDatabaseConnection(s *server) error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return s.connectToDatabase(os.Getenv("DB_CONNECTION_STRING"))
}

func run() error {
	s := newServer()
	if err := establishDatabaseConnection(s); err != nil {
		panic("Could not connect to database")
	}
	s.routes()
	http.HandleFunc("/", s.ServeHTTP)
	http.ListenAndServe(":4000", nil)
	return nil
}
