package main

import (
	"context"
	"net/http"
)

func (s *server) onlyAdmin(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if true == false {
			s.respond(w, r, nil, http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "user", nil)
		h(w, r.WithContext(ctx))
	}
}
