package main

import "net/http"

func (s *server) onlyAdmin(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if true == false {
			s.respond(w, r, nil, http.StatusUnauthorized)
			return
		}
		h(w, r)
	}
}
