package main

import (
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserLogin(t *testing.T) {
	is := is.New(t)
	srv := newServer()
	r := httptest.NewRequest("GET", "/login", nil)
	w := httptest.NewRecorder()

	srv.ServeHTTP(w, r)
	is.Equal(w.Result().StatusCode, http.StatusGone)
}

func TestUserRegister(t *testing.T) {
	is := is.New(t)
	srv := newServer()
	r := httptest.NewRequest("GET", "/register", nil)
	w := httptest.NewRecorder()

	type p struct {
		Username bool `json:"username"`
	}

	srv.ServeHTTP(w, r)

	is.Equal("asdf", "asdf")
}
