package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	srv := newServer()
	srv.routes()
	http.HandleFunc("/", srv.ServeHTTP)
	http.ListenAndServe(":4000", nil)
	return nil
}
