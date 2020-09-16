package main

import (
	"net/http"
)

func (s *server) handleAuthLogin(userRepo BaseUserRepository) http.HandlerFunc {
	type request struct {
		Username string
		Password string
	}
	type response struct {
		Token *Token
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var data request
		err := s.decode(r, &data)
		if err != nil {
			s.respond(w, r, err, http.StatusBadRequest)
		}

		user := userRepo.GetUserByUsername(data.Username)
		if user == nil {
			s.respond(w, r, nil, http.StatusNotFound)
			return
		}
		success := user.CheckPassword(data.Password)

		if success {
			token, _ := newToken(user.ID, false)
			s.respond(w, r, response{
				Token: token,
			}, http.StatusOK)
			return
		}
		s.respond(w, r, nil, http.StatusNotFound)
	}
}

func (s *server) handleAuthRegister(userRepo BaseUserRepository) http.HandlerFunc {
	type request struct {
		Username string
		Password string
		Email    string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, nil, 200)
	}
}
