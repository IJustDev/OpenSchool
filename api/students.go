package main

import (
	"github.com/gorilla/mux"
	models "main/models"
	"net/http"
	"strconv"
)

func (s *server) handleStudentDetails(userRepository BaseUserRepository) http.HandlerFunc {
	type response struct {
		User *models.User
	}
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.respond(w, r, nil, http.StatusBadRequest)
		}
		user := userRepository.GetUserById(id)
		if user != nil {
			s.respond(w, r, response{
				User: user,
			}, http.StatusOK)
		} else {
			s.respond(w, r, nil, http.StatusNotFound)
		}
	}
}
