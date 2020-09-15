package main

import (
	"net/http"
)

func (s *server) handleAuthLogin() http.HandlerFunc {
	type response struct {
		Token *Token
	}
	token, _ := newToken(30, false)
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, response{
			Token: token,
		}, http.StatusGone)
	}
}
