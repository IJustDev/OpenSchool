package main

import (
	"bytes"
	"encoding/json"
	models "main/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
)

func TestAuthLoginWithValidData(t *testing.T) {
	is := is.New(t)
	srv := newServer()

	userRepository := new(MockUserRepository)
	userRepository.InsertUser(models.User{
		ID:       20,
		Username: "userOne",
		Password: "password",
	})

	reqBody, _ := json.Marshal(map[string]string{
		"username": "userOne",
		"password": "password",
	})

	handlerFunc := srv.handleAuthLogin(userRepository)

	r := httptest.NewRequest("POST", "/login", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	handlerFunc.ServeHTTP(w, r)
	is.Equal(w.Result().StatusCode, http.StatusOK)
}

func TestAuthLoginWithoutPostBody(t *testing.T) {
	is := is.New(t)
	srv := newServer()
	userRepository := new(MockUserRepository)
	handlerFunc := srv.handleAuthLogin(userRepository)
	r := httptest.NewRequest("POST", "/login", nil)
	w := httptest.NewRecorder()
	handlerFunc.ServeHTTP(w, r)
	is.Equal(w.Result().StatusCode, http.StatusBadRequest)
}

func TestAuthLoginWithInvalidUsernameAndPassword(t *testing.T) {
	is := is.New(t)
	srv := newServer()

	userRepository := new(MockUserRepository)

	reqBody, _ := json.Marshal(map[string]string{
		"username": "user",
		"password": "password",
	})
	handlerFunc := srv.handleAuthLogin(userRepository)

	r := httptest.NewRequest("POST", "/login", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	handlerFunc.ServeHTTP(w, r)
	is.Equal(w.Result().StatusCode, http.StatusNotFound)
}

func TestAuthLoginWithCorrectUsernameButInvalidPassword(t *testing.T) {
	is := is.New(t)
	srv := newServer()

	userRepository := new(MockUserRepository)
	userRepository.InsertUser(models.User{
		ID:       20,
		Username: "userOne",
		Password: "password",
	})

	reqBody, _ := json.Marshal(map[string]string{
		"username": "userOne",
		"password": "passwordTwo",
	})
	handlerFunc := srv.handleAuthLogin(userRepository)

	r := httptest.NewRequest("POST", "/login", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	handlerFunc.ServeHTTP(w, r)
	is.Equal(w.Result().StatusCode, http.StatusNotFound)
}

func TestAuthRegister(t *testing.T) {
	is := is.New(t)
	srv := newServer()
	userRepository := new(MockUserRepository)
	handlerFunc := srv.handleAuthRegister(userRepository)
	r := httptest.NewRequest("POST", "/register", nil)
	w := httptest.NewRecorder()
	handlerFunc.ServeHTTP(w, r)
	is.Equal(w.Result().StatusCode, http.StatusOK)
}
