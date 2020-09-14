package main

import (
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
)

type User struct {
  ID int `json:"id"`
  Username string `json:"username"`
  password string
}

type GetAllUsers struct {}

func (u GetAllUsers) Handler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  users := []User {User{ID:1, Username: "asdf"}, User{ID:2, Username: "OjunbamO"}}
  json.NewEncoder(w).Encode(users)
}

func (u GetAllUsers) EndpointURL() string {
  return "/"
}

type GetSingleUser struct {}

func (u GetSingleUser) Handler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  w.Header().Set("Content-Type", "application/json")
  user := User{
    ID: 0,
    Username: vars["id"],
  }
  json.NewEncoder(w).Encode(user)
}

func (u GetSingleUser) EndpointURL() string {
  return "/{id}"
}

// User Component
type UserComponent struct {}

func (u UserComponent) EndpointBaseURL() string {
  return "/user"
}

func (u UserComponent) Endpoints() []Endpoint {
  return []Endpoint {GetSingleUser{}, GetAllUsers{}}
}
