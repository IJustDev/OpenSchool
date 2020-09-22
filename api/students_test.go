package main

import (
	"github.com/gorilla/mux"
	"github.com/matryer/is"
	models "main/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStudentDetailsWithValidID(t *testing.T) {
	is := is.New(t)
	srv := newServer()

	userRepository := new(MockUserRepository)
	userRepository.InsertUser(models.User{
		ID:       20,
		Username: "userOne",
		Password: "SamplePassword",
	})

	handlerFunc := srv.handleStudentDetails(userRepository)

	r := httptest.NewRequest("GET", "/details/20", nil)

	vars := map[string]string{
		"id": "20",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	handlerFunc.ServeHTTP(w, r)
	is.Equal(w.Result().StatusCode, http.StatusOK)
}

func TestStudentDetailsWithNotFoundValidID(t *testing.T) {
	is := is.New(t)
	srv := newServer()
	userRepository := new(MockUserRepository)

	handlerFunc := srv.handleStudentDetails(userRepository)

	r := httptest.NewRequest("GET", "/details/20", nil)
	vars := map[string]string{
		"id": "20",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	handlerFunc.ServeHTTP(w, r)
	is.Equal(w.Result().StatusCode, http.StatusNotFound)
}

func TestStudentDetailsWithInvalidID(t *testing.T) {
	is := is.New(t)
	srv := newServer()
	userRepository := new(MockUserRepository)

	handlerFunc := srv.handleStudentDetails(userRepository)

	r := httptest.NewRequest("GET", "/details/randomString", nil)

	vars := map[string]string{
		"id": "string",
	}
	r = mux.SetURLVars(r, vars)

	w := httptest.NewRecorder()
	handlerFunc.ServeHTTP(w, r)
	is.Equal(w.Result().StatusCode, http.StatusBadRequest)
}
