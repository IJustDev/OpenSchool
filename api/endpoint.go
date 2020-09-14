package main

import "net/http"

type Endpoint interface {
  EndpointURL() string
  Handler(w http.ResponseWriter, r *http.Request)
}

type Component interface {
  EndpointBaseURL() string
  Endpoints() []Endpoint
}
