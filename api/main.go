package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

type Student struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var db *sql.DB
var err error

func connectToSQL() {
	db, err = sql.Open("mysql", "test")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func registerComponent(router *mux.Router, component Component) {
	endpoints := component.Endpoints()
	for i, _ := range endpoints {
		endpoint := endpoints[i]
		router.HandleFunc(component.EndpointBaseURL()+endpoint.EndpointURL(), endpoint.Handler)
	}
}

func main() {
	router := mux.NewRouter()
	components := []Component{
		UserComponent{},
	}

	for i, _ := range components {
		registerComponent(router, components[i])
	}
	http.ListenAndServe(":4000", router)
}
