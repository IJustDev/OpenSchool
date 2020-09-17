package main

import (
	"gopkg.in/validator.v2"
	models "main/models"
	"net/http"
)

func (s *server) handleAuthLogin(userRepo BaseUserRepository) http.HandlerFunc {
	type request struct {
		Username string `validate:"min=3"`
		Password string `validate:"min=8"`
	}
	type response struct {
		Token *Token
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var data request
		if err := s.decode(r, &data); err != nil {
			s.respond(w, r, err, http.StatusBadRequest)
			return
		}
		if err := validator.Validate(data); err != nil {
			s.respond(w, r, err, http.StatusBadRequest)
			return
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
		Username string `validate:"min=3"`
		Password string `validate:"min=8"`
		Email    string `validate:"min=1"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var data request
		if err := s.decode(r, &data); err != nil {
			s.respond(w, r, err, http.StatusBadRequest)
			return
		}
		if err := validator.Validate(data); err != nil {
			s.respond(w, r, err, http.StatusBadRequest)
			return
		}
		isUsernameAllocated := userRepo.GetUserByUsername(
			data.Username,
		) != nil
		if !isUsernameAllocated {
			userRepo.InsertUser(models.User{
				Username: data.Username,
				Password: data.Password,
				Email:    data.Email,
			})
			s.respond(w, r, nil, 201)
		} else {
			s.respond(w, r, nil, http.StatusConflict)
		}
	}
}
