package main

import "net/http"

func (s *server) onlyAdmin(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if true == false {
			s.respond(w, r, nil, 200)
			return
		}
		h(w, r)
	}
}
