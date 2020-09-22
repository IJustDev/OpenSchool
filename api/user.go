package main

import (
	models "main/models"
	"net/http"
)

func (s *server) handleSelfDetails(userRepository BaseUserRepository) http.HandlerFunc {
	type response struct {
		User *models.User
	}

	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, nil, http.StatusOK)
	}
}
