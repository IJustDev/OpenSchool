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

func TestAuthLoginValidationWithShortPassword(t *testing.T) {
	is := is.New(t)
	srv := newServer()

	userRepository := new(MockUserRepository)

	reqBody, _ := json.Marshal(map[string]string{
		"username": "userOne",
		"password": "pass",
	})

	handlerFunc := srv.handleAuthLogin(userRepository)

	r := httptest.NewRequest("POST", "/login", bytes.NewBuffer(reqBody))
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

func TestAuthRegisterWithValidData(t *testing.T) {
	is := is.New(t)
	srv := newServer()
	userRepository := new(MockUserRepository)

	reqBody, _ := json.Marshal(map[string]string{
		"username": "userOne",
		"password": "password",
		"email":    "Email",
	})

	handlerFunc := srv.handleAuthRegister(userRepository)

	r := httptest.NewRequest("POST", "/register", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	handlerFunc.ServeHTTP(w, r)
	is.Equal(w.Result().StatusCode, http.StatusCreated)
}

func TestAuthRegisterWithOutPassword(t *testing.T) {
	is := is.New(t)
	srv := newServer()
	userRepository := new(MockUserRepository)

	reqBody, _ := json.Marshal(map[string]string{
		"username": "userOne",
		"email":    "Email",
	})

	handlerFunc := srv.handleAuthRegister(userRepository)

	r := httptest.NewRequest("POST", "/register", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	handlerFunc.ServeHTTP(w, r)
	is.Equal(w.Result().StatusCode, http.StatusBadRequest)
}

func TestAuthRegisterWithOutEmail(t *testing.T) {
	is := is.New(t)
	srv := newServer()
	userRepository := new(MockUserRepository)

	reqBody, _ := json.Marshal(map[string]string{
		"username": "userOne",
		"password": "password",
	})

	handlerFunc := srv.handleAuthRegister(userRepository)

	r := httptest.NewRequest("POST", "/register", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	handlerFunc.ServeHTTP(w, r)
	is.Equal(w.Result().StatusCode, http.StatusBadRequest)
}

func TestAuthRegisterWasUserCreated(t *testing.T) {
	is := is.New(t)
	srv := newServer()
	userRepository := new(MockUserRepository)

	is.Equal(len(userRepository.GetAllUsers()), 0)

	reqBody, _ := json.Marshal(map[string]string{
		"username": "userOne",
		"password": "password",
		"email":    "email@test.de",
	})

	handlerFunc := srv.handleAuthRegister(userRepository)

	r := httptest.NewRequest("POST", "/register", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	handlerFunc.ServeHTTP(w, r)
	is.Equal(len(userRepository.GetAllUsers()), 1)
}

func TestAuthRegisterWithAllocatedUsername(t *testing.T) {
	is := is.New(t)
	srv := newServer()
	userRepository := new(MockUserRepository)
	userRepository.InsertUser(models.User{
		Username: "userOne",
		Password: "password",
		Email:    "email@test.de",
	})

	reqBody, _ := json.Marshal(map[string]string{
		"username": "userOne",
		"password": "password",
		"email":    "email@test.de",
	})

	handlerFunc := srv.handleAuthRegister(userRepository)

	r := httptest.NewRequest("POST", "/register", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	handlerFunc.ServeHTTP(w, r)
	is.Equal(w.Result().StatusCode, http.StatusConflict)
	is.Equal(len(userRepository.GetAllUsers()), 1)
}
